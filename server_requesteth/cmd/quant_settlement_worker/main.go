package main

import (
	"context"
	"flag"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"

	"requestEth/internal/quantfund"
	"requestEth/internal/service"
)

type nftOpenedRow struct {
	ID          uint64 `gorm:"column:id"`
	TokenID     uint64 `gorm:"column:token_id"`
	BlockTime   uint64 `gorm:"column:block_time"`
	OpenedAt    uint64 `gorm:"column:opened_at"`
	CheckStatus uint64 `gorm:"column:check_status"`
}

func (nftOpenedRow) TableName() string { return "nft_opened" }

func main() {
	fs := flag.NewFlagSet("quant_settlement_worker", flag.ExitOnError)
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

	worker := &settlementWorker{db: db, cfg: cfg}
	if !cfg.SettleDryRun {
		keyHex := os.Getenv("SET_REWARD_ADMIN_PRIVATE_KEY")
		if keyHex == "" {
			log.Fatalf("need SET_REWARD_ADMIN_PRIVATE_KEY when QUANT_SETTLE_DRY_RUN is not enabled")
		}
		key, err := crypto.HexToECDSA(trimHex(keyHex))
		if err != nil {
			log.Fatalf("invalid SET_REWARD_ADMIN_PRIVATE_KEY: %v", err)
		}
		client, err := ethclient.Dial(cfg.RPCURL)
		if err != nil {
			log.Fatalf("eth client: %v", err)
		}
		auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(cfg.ChainID))
		if err != nil {
			log.Fatalf("transactor: %v", err)
		}
		nftAddr := common.HexToAddress(os.Getenv("NFT_ADDR"))
		if (nftAddr == common.Address{}) {
			log.Fatalf("need NFT_ADDR when QUANT_SETTLE_DRY_RUN is not enabled")
		}
		instance, err := service.NewNft(nftAddr, client)
		if err != nil {
			log.Fatalf("nft binding: %v", err)
		}
		worker.client = client
		worker.auth = auth
		worker.nft = instance
		defer client.Close()
	}

	ctx := context.Background()
	for {
		settled, err := worker.run(ctx)
		if err != nil {
			log.Printf("quant settlement: %v", err)
		} else {
			log.Printf("quant settlement batch done, settled=%d", settled)
		}
		if *runOnce {
			return
		}
		time.Sleep(time.Duration(cfg.IntervalMin) * time.Minute)
	}
}

type settlementWorker struct {
	db     *gorm.DB
	cfg    quantfund.Config
	client *ethclient.Client
	auth   *bind.TransactOpts
	nft    *service.Nft
}

func (w *settlementWorker) run(ctx context.Context) (int, error) {
	results := make([]quantfund.QuantProfitResult, 0)
	if err := w.db.WithContext(ctx).
		Where("status = ?", quantfund.ResultStatusPendingSettle).
		Order("id ASC").
		Limit(w.cfg.SettleBatchSize).
		Find(&results).Error; err != nil {
		return 0, err
	}
	if len(results) == 0 {
		return 0, nil
	}

	settled := 0
	for _, item := range results {
		ok, err := w.settleOne(ctx, &item)
		if err != nil {
			log.Printf("settle token %d err=%v", item.TokenID, err)
			continue
		}
		if ok {
			settled++
		}
	}
	return settled, nil
}

func (w *settlementWorker) settleOne(ctx context.Context, result *quantfund.QuantProfitResult) (bool, error) {
	var opened nftOpenedRow
	if err := w.db.WithContext(ctx).Where("token_id = ?", result.TokenID).Order("id DESC").First(&opened).Error; err != nil {
		_ = w.markFailed(ctx, result.ID, result.LedgerID, "nft_opened not found")
		return false, nil
	}

	rewardWei, err := quantfund.DecimalToWei(result.SettlementReward, 18)
	if err != nil {
		return false, err
	}
	txHash := quantfund.BuildDryRunTxHash("dry-settle", result.TaskNo)
	if !w.cfg.SettleDryRun {
		tx, err := w.nft.SetRewardOnce(&bind.TransactOpts{
			From:     w.auth.From,
			Signer:   w.auth.Signer,
			GasLimit: 500000,
			Context:  ctx,
		}, new(big.Int).SetUint64(result.TokenID), rewardWei)
		if err != nil {
			_ = w.markFailed(ctx, result.ID, result.LedgerID, err.Error())
			return false, err
		}
		txHash = tx.Hash().Hex()
	}

	err = w.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&quantfund.QuantProfitResult{}).
			Where("id = ? AND status = ?", result.ID, quantfund.ResultStatusPendingSettle).
			Updates(map[string]interface{}{
				"status":             quantfund.ResultStatusSettled,
				"settlement_tx_hash": txHash,
				"remark":             "",
			}).Error; err != nil {
			return err
		}
		if err := tx.Model(&quantfund.QuantFundLedger{}).
			Where("id = ?", result.LedgerID).
			Updates(map[string]interface{}{
				"status":             quantfund.LedgerStatusSettled,
				"settlement_tx_hash": txHash,
				"remark":             "",
			}).Error; err != nil {
			return err
		}
		if err := tx.Table("nft_opened").Where("id = ?", opened.ID).
			Updates(map[string]interface{}{
				"reward":       result.SettlementReward,
				"check_status": 1,
			}).Error; err != nil {
			return err
		}
		return tx.Table("nft_minted").Where("token_id = ?", result.TokenID).
			Updates(map[string]interface{}{
				"open_status": 1,
				"opened_at":   opened.BlockTime,
			}).Error
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (w *settlementWorker) markFailed(ctx context.Context, resultID, ledgerID uint64, remark string) error {
	return w.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&quantfund.QuantProfitResult{}).
			Where("id = ?", resultID).
			Updates(map[string]interface{}{
				"status": quantfund.ResultStatusSettleFailed,
				"remark": remark,
			}).Error; err != nil {
			return err
		}
		return tx.Model(&quantfund.QuantFundLedger{}).
			Where("id = ?", ledgerID).
			Updates(map[string]interface{}{
				"status": quantfund.LedgerStatusSettlementFailed,
				"remark": remark,
			}).Error
	})
}

func trimHex(v string) string {
	if len(v) >= 2 && v[:2] == "0x" {
		return v[2:]
	}
	return v
}
