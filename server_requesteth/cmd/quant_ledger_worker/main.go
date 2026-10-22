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
	fs := flag.NewFlagSet("quant_ledger_worker", flag.ExitOnError)
	confDir := quantfund.AddCommonFlags(fs, "../../configs")
	runOnce := fs.Bool("once", false, "run one batch and exit")
	_ = fs.Parse(os.Args[1:])

	cfg, err := quantfund.LoadConfig(*confDir)
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

	worker := &ledgerWorker{db: db, cfg: cfg}
	ctx := context.Background()
	for {
		inserted, err := worker.run(ctx)
		if err != nil {
			log.Printf("quant ledger: %v", err)
		} else {
			log.Printf("quant ledger batch done, inserted=%d", inserted)
		}
		if *runOnce {
			return
		}
		time.Sleep(time.Duration(cfg.IntervalMin) * time.Minute)
	}
}

type ledgerWorker struct {
	db  *gorm.DB
	cfg quantfund.Config
}

func (w *ledgerWorker) run(ctx context.Context) (int, error) {
	sub := w.db.WithContext(ctx).
		Table("nft_minted AS t").
		Select("MAX(t.id)").
		Where("t.token_id > 0").
		Group("t.token_id")

	rows := make([]quantfund.MintedRow, 0)
	err := w.db.WithContext(ctx).
		Table("nft_minted AS m").
		Select("m.id, m.token_id, m.to_addr, m.mint_addr, CAST(m.usdt_paid AS CHAR) AS usdt_paid, m.block_number, m.block_time").
		Where("m.id IN (?)", sub).
		Joins("LEFT JOIN quant_fund_ledger q ON q.token_id = m.token_id").
		Where("q.id IS NULL").
		Where("m.usdt_paid > 0").
		Order("m.id ASC").
		Limit(w.cfg.LedgerBatchSize).
		Scan(&rows).Error
	if err != nil {
		return 0, err
	}
	if len(rows) == 0 {
		return 0, nil
	}

	ledgers := make([]quantfund.QuantFundLedger, 0, len(rows))
	for _, row := range rows {
		totalWei, err := quantfund.DecimalToWei(row.UsdtPaid, 18)
		if err != nil {
			return 0, err
		}
		quantWei, reserveWei := quantfund.SplitByBps(totalWei, w.cfg.QuantBps)
		if quantWei.Sign() <= 0 {
			continue
		}

		ledgers = append(ledgers, quantfund.QuantFundLedger{
			MintedID:      row.ID,
			TokenID:       row.TokenID,
			UserAddr:      row.ToAddr,
			MintAddr:      row.MintAddr,
			UsdtPaid:      quantfund.WeiToDecimal(totalWei, 18),
			QuantBps:      w.cfg.QuantBps,
			QuantAmount:   quantfund.WeiToDecimal(quantWei, 18),
			ReserveAmount: quantfund.WeiToDecimal(reserveWei, 18),
			BlockNumber:   row.BlockNumber,
			BlockTime:     row.BlockTime,
			Status:        0,
		})
	}
	if len(ledgers) == 0 {
		return 0, nil
	}

	result := w.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "token_id"}},
			DoNothing: true,
		}).
		Create(&ledgers)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}
