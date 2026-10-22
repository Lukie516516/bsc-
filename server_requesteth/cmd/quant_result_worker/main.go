package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"requestEth/internal/quantfund"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	fs := flag.NewFlagSet("quant_result_worker", flag.ExitOnError)
	confPath := quantfund.AddCommonFlags(fs, "../../configs")
	runOnce := fs.Bool("once", false, "run one batch and exit")
	_ = fs.Parse(os.Args[1:])

	cfg, err := quantfund.LoadConfig(*confPath)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}
	db, err := quantfund.OpenDB(cfg.DatabaseSource)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	if err := quantfund.AutoMigrate(db); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	worker := &resultWorker{db: db, cfg: cfg}
	ctx := context.Background()
	for {
		created, err := worker.run(ctx)
		if err != nil {
			log.Printf("quant result: %v", err)
		} else {
			log.Printf("quant result batch done, created=%d", created)
		}
		if *runOnce {
			return
		}
		time.Sleep(time.Duration(cfg.IntervalMin) * time.Minute)
	}
}

type resultWorker struct {
	db  *gorm.DB
	cfg quantfund.Config
}

func (w *resultWorker) run(ctx context.Context) (int, error) {
	rows := make([]quantfund.QuantFundLedger, 0)
	if err := w.db.WithContext(ctx).
		Where("status = ?", quantfund.LedgerStatusTransferred).
		Order("id ASC").
		Limit(w.cfg.ResultBatchSize).
		Find(&rows).Error; err != nil {
		return 0, err
	}
	if len(rows) == 0 {
		return 0, nil
	}

	results := make([]quantfund.QuantProfitResult, 0, len(rows))
	ids := make([]uint64, 0, len(rows))
	for _, row := range rows {
		taskNo := quantfund.NewTaskNo()
		profitBps := quantfund.PickProfitBps(
			w.cfg.ResultMode,
			w.cfg.FixedProfitBps,
			w.cfg.RandomProfitMinBps,
			w.cfg.RandomProfitMaxBps,
		)
		profitAmount, returnAmount, err := quantfund.CalcProfit(row.QuantAmount, profitBps)
		if err != nil {
			return 0, err
		}
		results = append(results, quantfund.QuantProfitResult{
			TaskNo:           taskNo,
			LedgerID:         row.ID,
			TokenID:          row.TokenID,
			UserAddr:         row.UserAddr,
			MintAddr:         row.MintAddr,
			QuantAmount:      row.QuantAmount,
			ProfitBps:        profitBps,
			ProfitAmount:     profitAmount,
			ReturnAmount:     returnAmount,
			SettlementReward: profitAmount,
			ResultMode:       quantfund.NormalizeResultMode(w.cfg.ResultMode),
			Status:           quantfund.ResultStatusPendingSettle,
		})
		ids = append(ids, row.ID)
	}

	err := w.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "ledger_id"}},
			DoNothing: true,
		}).Create(&results)
		if result.Error != nil {
			return result.Error
		}
		if int(result.RowsAffected) != len(results) {
			return gorm.ErrDuplicatedKey
		}

		taskByLedger := make(map[uint64]string, len(results))
		for _, item := range results {
			taskByLedger[item.LedgerID] = item.TaskNo
		}
		for _, id := range ids {
			updates := map[string]interface{}{
				"status":         quantfund.LedgerStatusResultReady,
				"result_task_no": taskByLedger[id],
				"remark":         "",
			}
			if err := tx.Model(&quantfund.QuantFundLedger{}).
				Where("id = ? AND status = ?", id, quantfund.LedgerStatusTransferred).
				Updates(updates).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return len(results), nil
}
