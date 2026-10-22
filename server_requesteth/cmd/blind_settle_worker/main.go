// 盲盒结算 Worker：监听 nft_opened（开盒事件），生成随机盈亏，调用 setRewardOnce
// 运行：go run . -conf=../../configs
// 环境变量：ETH_RPC_URL, NFT_ADDR, BLIND_SETTLE_PRIVATE_KEY, DATABASE_SOURCE
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"requestEth/internal/service"
)

var (
	flagConf = flag.String("conf", "../../configs", "config dir")
	cfgRPC   string
	cfgNFT   string
	cfgDB    string
	cfgKey   string

	// 结算随机规则（可配置）
	winRateBps      = 5000  // 胜率 50% (5000/10000)
	profitMinBps    = 5000  // 盈利下限：usdtPaid 的 50%
	profitMaxBps    = 20000 // 盈利上限：usdtPaid 的 200%
	lossMinBps      = 5000  // 亏损下限：usdtPaid 的 50%
	lossMaxBps      = 10000 // 亏损上限：usdtPaid 的 100%
	chainID         = int64(56)
	pollIntervalSec = 15
)

type configStruct struct {
	Data struct {
		Database struct {
			Source string `yaml:"source"`
		} `yaml:"database"`
	} `yaml:"data"`
	Eth struct {
		Bsc struct {
			Addr string `yaml:"addr"`
		} `yaml:"bsc"`
	} `yaml:"eth"`
	BlindSettle struct {
		IntervalSec  int `yaml:"interval_sec"`
		WinRateBps   int `yaml:"win_rate_bps"`   // 胜率 5000=50%
		ProfitMinBps int `yaml:"profit_min_bps"` // 盈利下限 5000=50% of usdtPaid
		ProfitMaxBps int `yaml:"profit_max_bps"` // 盈利上限 20000=200%
		LossMinBps   int `yaml:"loss_min_bps"`   // 亏损下限 5000=50%
		LossMaxBps   int `yaml:"loss_max_bps"`   // 亏损上限 10000=100%
		ChainID      int `yaml:"chain_id"`       // 56=BSC mainnet, 97=testnet
	} `yaml:"blind_settle"`
}

func loadConfig() {
	cfgRPC = os.Getenv("ETH_RPC_URL")
	cfgNFT = os.Getenv("NFT_ADDR")
	cfgDB = os.Getenv("DATABASE_SOURCE")
	cfgKey = os.Getenv("BLIND_SETTLE_PRIVATE_KEY")
	if v := os.Getenv("BLIND_SETTLE_INTERVAL_SEC"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			pollIntervalSec = n
		}
	}

	f, err := os.Open(*flagConf + "/config.yaml")
	if err != nil {
		log.Printf("warning: open config: %v", err)
		return
	}
	defer f.Close()
	var c configStruct
	if err := yaml.NewDecoder(f).Decode(&c); err != nil {
		log.Printf("warning: parse config: %v", err)
		return
	}
	if cfgRPC == "" && c.Eth.Bsc.Addr != "" {
		cfgRPC = c.Eth.Bsc.Addr
	}
	if cfgDB == "" && c.Data.Database.Source != "" {
		cfgDB = c.Data.Database.Source
	}
	if cfgNFT == "" {
		cfgNFT = os.Getenv("NFT_ADDR")
	}
	if c.BlindSettle.IntervalSec > 0 {
		pollIntervalSec = c.BlindSettle.IntervalSec
	}
	if c.BlindSettle.WinRateBps > 0 {
		winRateBps = c.BlindSettle.WinRateBps
	}
	if c.BlindSettle.ProfitMinBps > 0 {
		profitMinBps = c.BlindSettle.ProfitMinBps
	}
	if c.BlindSettle.ProfitMaxBps > 0 {
		profitMaxBps = c.BlindSettle.ProfitMaxBps
	}
	if c.BlindSettle.LossMinBps > 0 {
		lossMinBps = c.BlindSettle.LossMinBps
	}
	if c.BlindSettle.LossMaxBps > 0 {
		lossMaxBps = c.BlindSettle.LossMaxBps
	}
	if c.BlindSettle.ChainID > 0 {
		chainID = int64(c.BlindSettle.ChainID)
	}
}

type nftOpenedRow struct {
	ID          uint64  `gorm:"column:id"`
	BlockNumber uint64  `gorm:"column:block_number"`
	BlockTime   uint64  `gorm:"column:block_time"`
	UserAddr    string  `gorm:"column:user_addr"`
	TokenID     uint64  `gorm:"column:token_id"`
	OpenedAt    uint64  `gorm:"column:opened_at"`
	Reward      float64 `gorm:"column:reward"`
	CheckStatus uint64  `gorm:"column:check_status"`
}

func (nftOpenedRow) TableName() string { return "nft_opened" }

type nftMintedRow struct {
	ID        uint64  `gorm:"column:id"`
	TokenID   uint64  `gorm:"column:token_id"`
	UsdtPaid  float64 `gorm:"column:usdt_paid"`
	BlockTime uint64  `gorm:"column:block_time"`
}

func (nftMintedRow) TableName() string { return "nft_minted" }

func main() {
	flag.Parse()
	loadConfig()

	if cfgRPC == "" {
		log.Fatal("need ETH_RPC_URL or config eth.bsc.addr")
	}
	if cfgNFT == "" {
		log.Fatal("need NFT_ADDR env")
	}
	if cfgDB == "" {
		log.Fatal("need DATABASE_SOURCE or config data.database.source")
	}
	if cfgKey == "" {
		log.Fatal("need BLIND_SETTLE_PRIVATE_KEY env (hex, no 0x prefix)")
	}

	cli, err := ethclient.Dial(cfgRPC)
	if err != nil {
		log.Fatalf("eth client: %v", err)
	}
	defer cli.Close()

	db, err := gorm.Open(mysql.Open(cfgDB), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(cfgKey)
	if err != nil {
		log.Fatalf("private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		log.Fatalf("transactor: %v", err)
	}

	nftAddr := common.HexToAddress(cfgNFT)
	instance, err := service.NewNft(nftAddr, cli)
	if err != nil {
		log.Fatalf("nft binding: %v", err)
	}

	runner := &settleRunner{
		cli:     cli,
		db:      db,
		nft:     instance,
		auth:    auth,
		nftAddr: nftAddr,
	}

	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	for {
		loadBlindSettleRuntimeOverrides(db)
		if err := runner.process(ctx); err != nil {
			log.Printf("process: %v", err)
		}
		time.Sleep(time.Duration(pollIntervalSec) * time.Second)
	}
}

type runtimeConfigRow struct {
	ConfigKey   string `gorm:"column:config_key"`
	ConfigValue string `gorm:"column:config_value"`
}

func (runtimeConfigRow) TableName() string { return "runtime_config" }

func loadBlindSettleRuntimeOverrides(db *gorm.DB) {
	keys := []string{
		"worker.blind_settle.interval_sec",
		"blind_settle.win_rate_bps",
		"blind_settle.profit_min_bps",
		"blind_settle.profit_max_bps",
		"blind_settle.loss_min_bps",
		"blind_settle.loss_max_bps",
	}
	var rows []runtimeConfigRow
	if err := db.Table("runtime_config").Where("config_key IN ?", keys).Find(&rows).Error; err != nil {
		return
	}
	for _, row := range rows {
		n, err := strconv.Atoi(row.ConfigValue)
		if err != nil {
			continue
		}
		switch row.ConfigKey {
		case "worker.blind_settle.interval_sec":
			if n > 0 {
				pollIntervalSec = n
			}
		case "blind_settle.win_rate_bps":
			if n >= 0 {
				winRateBps = n
			}
		case "blind_settle.profit_min_bps":
			profitMinBps = n
		case "blind_settle.profit_max_bps":
			profitMaxBps = n
		case "blind_settle.loss_min_bps":
			lossMinBps = n
		case "blind_settle.loss_max_bps":
			lossMaxBps = n
		}
	}
}

type settleRunner struct {
	cli     *ethclient.Client
	db      *gorm.DB
	nft     *service.Nft
	auth    *bind.TransactOpts
	nftAddr common.Address
}

func (r *settleRunner) process(ctx context.Context) error {
	var opened []nftOpenedRow
	if err := r.db.WithContext(ctx).Where("check_status = ?", 0).Order("id asc").Limit(10).Find(&opened).Error; err != nil {
		return err
	}
	if len(opened) == 0 {
		return nil
	}

	tokenIds := make([]uint64, len(opened))
	for i, o := range opened {
		tokenIds[i] = o.TokenID
	}

	var minted []nftMintedRow
	if err := r.db.WithContext(ctx).Table("nft_minted").Where("token_id IN ?", tokenIds).
		Select("id, token_id, usdt_paid, block_time").Find(&minted).Error; err != nil {
		return err
	}
	mintedMap := make(map[uint64]*nftMintedRow)
	for i := range minted {
		mintedMap[minted[i].TokenID] = &minted[i]
	}

	for _, o := range opened {
		m, ok := mintedMap[o.TokenID]
		if !ok {
			log.Printf("skip token %d: no nft_minted", o.TokenID)
			continue
		}
		if o.BlockTime < m.BlockTime {
			continue
		}
		if err := r.settleOne(ctx, &o, m); err != nil {
			log.Printf("settle token %d err=%v", o.TokenID, err)
		}
	}
	return nil
}

// genRandomReward 根据 usdtPaid 生成随机盈亏（USDT 数值），返回链上 int256 格式
func genRandomReward(usdtPaid float64) *big.Int {
	paidWei := usdToWei(usdtPaid)
	if paidWei.Sign() <= 0 {
		return big.NewInt(0)
	}
	win := rand.Intn(10000) < winRateBps
	if win {
		// 盈利：profitMinBps ~ profitMaxBps 的 usdtPaid
		ratioBps := profitMinBps
		if profitMaxBps > profitMinBps {
			ratioBps = profitMinBps + rand.Intn(profitMaxBps-profitMinBps)
		}
		profitWei := new(big.Int).Mul(paidWei, big.NewInt(int64(ratioBps)))
		profitWei.Div(profitWei, big.NewInt(10000))
		return profitWei
	}
	// 亏损：lossMinBps ~ lossMaxBps 的 usdtPaid，取负
	ratioBps := lossMinBps
	if lossMaxBps > lossMinBps {
		ratioBps = lossMinBps + rand.Intn(lossMaxBps-lossMinBps)
	}
	lossWei := new(big.Int).Mul(paidWei, big.NewInt(int64(ratioBps)))
	lossWei.Div(lossWei, big.NewInt(10000))
	return new(big.Int).Neg(lossWei)
}

func usdToWei(f float64) *big.Int {
	x := new(big.Float).SetFloat64(f)
	e18f := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	x.Mul(x, e18f)
	i, _ := x.Int(nil)
	return i
}

func (r *settleRunner) settleOne(ctx context.Context, opened *nftOpenedRow, minted *nftMintedRow) error {
	rewardWei := genRandomReward(minted.UsdtPaid)
	rewardFloat := weiToUsd(rewardWei)

	tokenIdBig := new(big.Int).SetUint64(opened.TokenID)
	tx, err := r.nft.SetRewardOnce(&bind.TransactOpts{
		From:     r.auth.From,
		Signer:   r.auth.Signer,
		GasLimit: 500000,
		Context:  ctx,
	}, tokenIdBig, rewardWei)
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(ctx, r.cli, tx)
	if err != nil {
		return fmt.Errorf("wait tx %s: %w", tx.Hash().Hex(), err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("tx %s reverted, receipt status=%d", tx.Hash().Hex(), receipt.Status)
	}
	log.Printf("settle token %d reward=%v tx=%s", opened.TokenID, rewardFloat, tx.Hash().Hex())

	// 更新 DB
	if err := r.db.WithContext(ctx).Table("nft_opened").Where("id = ?", opened.ID).
		Updates(map[string]interface{}{
			"check_status": 1,
			"reward":       rewardFloat,
		}).Error; err != nil {
		return err
	}
	if err := r.db.WithContext(ctx).Table("nft_minted").Where("id = ?", minted.ID).
		Updates(map[string]interface{}{
			"open_status": 1,
			"opened_at":   opened.BlockTime,
		}).Error; err != nil {
		return err
	}
	return nil
}

func weiToUsd(w *big.Int) float64 {
	if w == nil || w.Sign() == 0 {
		return 0
	}
	neg := w.Sign() < 0
	if neg {
		w = new(big.Int).Neg(w)
	}
	f := new(big.Float).SetInt(w)
	e18 := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	f.Quo(f, e18)
	val, _ := f.Float64()
	if neg {
		return -val
	}
	return val
}
