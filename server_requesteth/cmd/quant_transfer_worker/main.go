package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"requestEth/internal/quantfund"

	"gorm.io/gorm"
)

func main() {
	fs := flag.NewFlagSet("quant_transfer_worker", flag.ExitOnError)
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
	var runner *quantfund.ERC20Runner
	if !cfg.DryRun {
		runner, err = quantfund.NewERC20Runner(cfg)
		if err != nil {
			log.Fatalf("erc20 runner: %v", err)
		}
		defer runner.Close()
	}

	worker := &transferWorker{db: db, cfg: cfg, runner: runner}
	ctx := context.Background()
	for {
		if err := worker.run(ctx); err != nil {
			log.Printf("quant transfer: %v", err)
		}
		if *runOnce {
			return
		}
		time.Sleep(time.Duration(cfg.IntervalMin) * time.Minute)
	}
}

type transferWorker struct {
	db     *gorm.DB
	cfg    quantfund.Config
	runner *quantfund.ERC20Runner
}

func (w *transferWorker) run(ctx context.Context) error {
	ledgers, task, totalWei, err := w.lockBatch(ctx)
	if err != nil {
		return err
	}
	if len(ledgers) == 0 {
		log.Printf("quant transfer: no pending ledger")
		return nil
	}

	minWei, err := quantfund.DecimalToWei(fmt.Sprintf("%d", w.cfg.MinTransferUSD), 18)
	if err != nil {
		return err
	}
	if totalWei.Cmp(minWei) < 0 {
		remark := fmt.Sprintf("pending amount %s below min %d", quantfund.WeiToDecimal(totalWei, 18), w.cfg.MinTransferUSD)
		return w.unlockBatch(ctx, task.TaskNo, remark)
	}

	if w.cfg.DryRun {
		txHash := "dry-run-" + task.TaskNo
		if err := w.finishBatch(ctx, task.TaskNo, txHash); err != nil {
			return err
		}
		log.Printf("quant transfer dry-run success: ledgers=%d amount=%s tx=%s", len(ledgers), quantfund.WeiToDecimal(totalWei, 18), txHash)
		return nil
	}

	bal, err := w.runner.BalanceOf(ctx, w.runner.From)
	if err != nil {
		_ = w.failBatch(ctx, task.TaskNo, "balance check failed: "+err.Error())
		return err
	}
	if bal.Cmp(totalWei) < 0 {
		remark := fmt.Sprintf("hot wallet balance %s < needed %s", quantfund.WeiToDecimal(bal, 18), quantfund.WeiToDecimal(totalWei, 18))
		return w.unlockBatch(ctx, task.TaskNo, remark)
	}

	txHash, err := w.runner.Transfer(ctx, totalWei)
	if err != nil {
		_ = w.failBatch(ctx, task.TaskNo, "transfer failed: "+err.Error())
		return err
	}

	if err := w.finishBatch(ctx, task.TaskNo, txHash); err != nil {
		return err
	}
	log.Printf("quant transfer success: ledgers=%d amount=%s tx=%s", len(ledgers), quantfund.WeiToDecimal(totalWei, 18), txHash)
	return nil
}

func (w *transferWorker) lockBatch(ctx context.Context) ([]quantfund.QuantFundLedger, *quantfund.QuantTransferTask, *big.Int, error) {
	rows := make([]quantfund.QuantFundLedger, 0)
	totalWei := big.NewInt(0)
	taskNo := quantfund.NewTaskNo()

	err := w.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("status = ?", 0).
			Order("id ASC").
			Limit(w.cfg.TransferBatchSize).
			Find(&rows).Error; err != nil {
			return err
		}
		if len(rows) == 0 {
			return nil
		}

		ids := make([]uint64, 0, len(rows))
		for _, row := range rows {
			ids = append(ids, row.ID)
			amountWei, err := quantfund.DecimalToWei(row.QuantAmount, 18)
			if err != nil {
				return err
			}
			totalWei.Add(totalWei, amountWei)
		}

		result := tx.Model(&quantfund.QuantFundLedger{}).
			Where("id IN ? AND status = ?", ids, 0).
			Updates(map[string]interface{}{
				"status":           1,
				"transfer_task_no": taskNo,
				"remark":           "",
			})
		if result.Error != nil {
			return result.Error
		}
		if int(result.RowsAffected) != len(ids) {
			return errors.New("pending ledgers changed during lock")
		}

		task := &quantfund.QuantTransferTask{
			TaskNo:      taskNo,
			LedgerCount: len(rows),
			QuantBps:    w.cfg.QuantBps,
			TotalAmount: quantfund.WeiToDecimal(totalWei, 18),
			FromAddr:    fromAddr(w.runner, w.cfg),
			ToAddr:      toAddr(w.runner, w.cfg),
			Status:      0,
		}
		return tx.Create(task).Error
	})
	if err != nil {
		return nil, nil, nil, err
	}
	if len(rows) == 0 {
		return nil, nil, big.NewInt(0), nil
	}

	return rows, &quantfund.QuantTransferTask{
		TaskNo:      taskNo,
		LedgerCount: len(rows),
		QuantBps:    w.cfg.QuantBps,
		TotalAmount: quantfund.WeiToDecimal(totalWei, 18),
		FromAddr:    fromAddr(w.runner, w.cfg),
		ToAddr:      toAddr(w.runner, w.cfg),
	}, totalWei, nil
}

func (w *transferWorker) unlockBatch(ctx context.Context, taskNo, remark string) error {
	return w.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&quantfund.QuantFundLedger{}).
			Where("transfer_task_no = ? AND status = ?", taskNo, 1).
			Updates(map[string]interface{}{
				"status":           0,
				"transfer_task_no": "",
				"remark":           remark,
			}).Error; err != nil {
			return err
		}
		return tx.Model(&quantfund.QuantTransferTask{}).
			Where("task_no = ?", taskNo).
			Updates(map[string]interface{}{
				"status": 1,
				"remark": remark,
			}).Error
	})
}

func (w *transferWorker) failBatch(ctx context.Context, taskNo, remark string) error {
	return w.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&quantfund.QuantFundLedger{}).
			Where("transfer_task_no = ? AND status = ?", taskNo, 1).
			Updates(map[string]interface{}{
				"status": 3,
				"remark": remark,
			}).Error; err != nil {
			return err
		}
		return tx.Model(&quantfund.QuantTransferTask{}).
			Where("task_no = ?", taskNo).
			Updates(map[string]interface{}{
				"status": 3,
				"remark": remark,
			}).Error
	})
}

func (w *transferWorker) finishBatch(ctx context.Context, taskNo, txHash string) error {
	return w.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&quantfund.QuantFundLedger{}).
			Where("transfer_task_no = ? AND status = ?", taskNo, 1).
			Updates(map[string]interface{}{
				"status":           2,
				"transfer_tx_hash": txHash,
				"remark":           "",
			}).Error; err != nil {
			return err
		}
		return tx.Model(&quantfund.QuantTransferTask{}).
			Where("task_no = ?", taskNo).
			Updates(map[string]interface{}{
				"status":  2,
				"tx_hash": txHash,
				"remark":  "",
			}).Error
	})
}

func fromAddr(r *quantfund.ERC20Runner, cfg quantfund.Config) string {
	if r != nil {
		return r.From.Hex()
	}
	return "dry-run-from"
}

func toAddr(r *quantfund.ERC20Runner, cfg quantfund.Config) string {
	if r != nil {
		return r.To.Hex()
	}
	if cfg.ExchangeAddr != "" {
		return cfg.ExchangeAddr
	}
	return "dry-run-to"
}
