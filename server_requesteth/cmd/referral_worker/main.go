// 铸造分佣 Worker：监听铸造盲盒事件(nft_minted)，计算 B 0.3% + D 0.2% 待发记录
// 二级市场购买盲盒不分佣
// 运行：go run . -conf=../../configs
// 环境变量：ETH_RPC_URL, NFT_ADDR, STAKE_ADDR（可选，默认从 config.yaml eth.bsc 读取）
package main

import (
	"context"
	"flag"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gopkg.in/yaml.v3"
)

var (
	flagConf   = flag.String("conf", "../../configs", "config dir")
	cfgRPC     string
	cfgNFT     string
	cfgStake   string
	cfgDB      string
	referralEnabled = false // 铸造分佣开关，false=关闭（默认）
	pollIntervalSec = 30
	// 铸造分佣比例（可 config 修改）
	bBps     = 30  // B 推广者 0.3%
	dBps     = 20  // D 团队长 0.2%
	holdMinU = 300 // B 合格阈值 300 USDT
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
	Referral struct {
		Enabled  bool `yaml:"enabled"`       // 铸造分佣开关，false 则不再处理
		IntervalSec int `yaml:"interval_sec"`
		BBps     int  `yaml:"b_bps"`        // B 0.3%=30
		DBps     int  `yaml:"d_bps"`        // D 0.2%=20
		HoldMinU int  `yaml:"hold_min_usd"` // B 合格阈值 USDT
	} `yaml:"referral"`
}

func loadConfig() {
	rpc := os.Getenv("ETH_RPC_URL")
	nft := os.Getenv("NFT_ADDR")
	stake := os.Getenv("STAKE_ADDR")
	if rpc != "" {
		cfgRPC = rpc
	}
	if nft != "" {
		cfgNFT = nft
	}
	if stake != "" {
		cfgStake = stake
	}
	if v := os.Getenv("REFERRAL_INTERVAL_SEC"); v != "" {
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
	if ds := os.Getenv("DATABASE_SOURCE"); ds != "" {
		cfgDB = ds
	}
	if nf := os.Getenv("NFT_ADDR"); nf != "" {
		cfgNFT = nf
	}
	if st := os.Getenv("STAKE_ADDR"); st != "" {
		cfgStake = st
	}
	if c.Referral.BBps >= 0 {
		bBps = c.Referral.BBps
	}
	if c.Referral.IntervalSec > 0 {
		pollIntervalSec = c.Referral.IntervalSec
	}
	if c.Referral.DBps >= 0 {
		dBps = c.Referral.DBps
	}
	if c.Referral.HoldMinU > 0 {
		holdMinU = c.Referral.HoldMinU
	}
	referralEnabled = c.Referral.Enabled
	// 环境变量覆盖
	if v := os.Getenv("REFERRAL_B_BPS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			bBps = n
		}
	}
	if v := os.Getenv("REFERRAL_D_BPS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			dBps = n
		}
	}
	if v := os.Getenv("REFERRAL_HOLD_MIN_USD"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			holdMinU = n
		}
	}
	if v := os.Getenv("REFERRAL_ENABLED"); v != "" {
		referralEnabled = (v == "1" || v == "true" || v == "yes")
	}
}

func main() {
	flag.Parse()
	loadConfig()

	if cfgRPC == "" {
		log.Fatal("need ETH_RPC_URL or config eth.bsc.addr")
	}
	if cfgNFT == "" || cfgStake == "" {
		log.Fatal("need NFT_ADDR and STAKE_ADDR env")
	}
	if cfgDB == "" {
		log.Fatal("need config data.database.source")
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

	// 自动迁移（含 minted_id 新列）
	if err := db.AutoMigrate(&referralPendingModel{}); err != nil {
		log.Printf("migrate: %v", err)
	}

	runner := &referralRunner{
		cli:       cli,
		db:        db,
		nftAddr:   common.HexToAddress(cfgNFT),
		stakeAddr: common.HexToAddress(cfgStake),
	}
	if err := runner.initBindings(); err != nil {
		log.Fatalf("bindings: %v", err)
	}

	ctx := context.Background()
	for {
		loadReferralRuntimeOverrides(db)
		if referralEnabled {
			if err := runner.process(ctx); err != nil {
				log.Printf("process: %v", err)
			}
		}
		time.Sleep(time.Duration(pollIntervalSec) * time.Second)
	}
}

type runtimeConfigRow struct {
	ConfigKey   string `gorm:"column:config_key"`
	ConfigValue string `gorm:"column:config_value"`
}

func (runtimeConfigRow) TableName() string { return "runtime_config" }

func loadReferralRuntimeOverrides(db *gorm.DB) {
	keys := []string{
		"worker.referral.interval_sec",
		"referral.enabled",
		"referral.b_bps",
		"referral.d_bps",
		"referral.hold_min_usd",
	}
	var rows []runtimeConfigRow
	if err := db.Table("runtime_config").Where("config_key IN ?", keys).Find(&rows).Error; err != nil {
		return
	}
	for _, row := range rows {
		switch row.ConfigKey {
		case "referral.enabled":
			referralEnabled = row.ConfigValue == "1" || row.ConfigValue == "true" || row.ConfigValue == "yes"
		case "worker.referral.interval_sec":
			if n, err := strconv.Atoi(row.ConfigValue); err == nil && n > 0 {
				pollIntervalSec = n
			}
		case "referral.b_bps":
			if n, err := strconv.Atoi(row.ConfigValue); err == nil && n >= 0 {
				bBps = n
			}
		case "referral.d_bps":
			if n, err := strconv.Atoi(row.ConfigValue); err == nil && n >= 0 {
				dBps = n
			}
		case "referral.hold_min_usd":
			if n, err := strconv.Atoi(row.ConfigValue); err == nil && n > 0 {
				holdMinU = n
			}
		}
	}
}

type referralPendingModel struct {
	ID              uint64     `gorm:"primaryKey;autoIncrement"`
	MintedID        uint64     `gorm:"index:idx_minted;default:0;comment:关联 nft_minted.id（铸造分佣）"`
	PurchaseID      uint64     `gorm:"default:0;comment:历史二级购买 id，现仅铸造分佣"`
	BlockNumber     uint64
	BlockTime       uint64
	TokenID        uint64
	Minter         string
	Seller         string
	Buyer          string
	Price          string
	AddrB          string
	AddrC          string
	AmountToB      string
	AmountToC      string
	AmountToProject string
	Reason          uint8
	PaidAt         *time.Time
	PaidTxHash     string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (referralPendingModel) TableName() string { return "referral_pending" }

type referralRunner struct {
	cli       *ethclient.Client
	db        *gorm.DB
	nftAddr   common.Address
	stakeAddr common.Address

	nftABI   abi.ABI
	stakeABI abi.ABI
}

func (r *referralRunner) initBindings() error {
	nftABI := `[{"inputs":[{"name":"user","type":"address"}],"name":"holdingValueOf","outputs":[{"name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
	stakeABI := `[{"inputs":[{"name":"","type":"address"}],"name":"parentOf","outputs":[{"name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"name":"","type":"address"}],"name":"topOf","outputs":[{"name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"name":"top","type":"address"}],"name":"isTopEligible","outputs":[{"name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
	if err := r.nftABI.UnmarshalJSON([]byte(nftABI)); err != nil {
		return err
	}
	if err := r.stakeABI.UnmarshalJSON([]byte(stakeABI)); err != nil {
		return err
	}
	return nil
}

func (r *referralRunner) process(ctx context.Context) error {
	if !referralEnabled {
		return nil // 铸造分佣已禁用
	}
	// 查 nft_minted 中尚未生成分佣的记录
	var minted []struct {
		ID          uint64
		BlockNumber uint64
		BlockTime   uint64
		TokenID     uint64
		MintAddr    string
		UsdtPaid    float64
	}
	if err := r.db.Table("nft_minted").FindInBatches(&minted, 50, func(tx *gorm.DB, batch int) error {
		for _, m := range minted {
			var exists int64
			r.db.Table("referral_pending").Where("minted_id = ?", m.ID).Count(&exists)
			if exists > 0 {
				continue
			}
			if err := r.handleMint(ctx, &m); err != nil {
				log.Printf("handle mint id=%d tokenId=%d err=%v", m.ID, m.TokenID, err)
			}
		}
		return nil
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *referralRunner) handleMint(ctx context.Context, m *struct {
	ID          uint64
	BlockNumber uint64
	BlockTime   uint64
	TokenID     uint64
	MintAddr    string
	UsdtPaid    float64
}) error {
	minter := common.HexToAddress(m.MintAddr)

	// 2. B = parentOf(minter) 推广者
	data, err := r.stakeABI.Pack("parentOf", minter)
	if err != nil {
		return err
	}
	msg := ethereum.CallMsg{To: &r.stakeAddr, Data: data}
	out, err := r.cli.CallContract(ctx, msg, nil)
	if err != nil {
		return err
	}
	var B common.Address
	r.stakeABI.UnpackIntoInterface(&B, "parentOf", out)

	// 3. D = topOf(minter) 团队长（无上级时 D=minter）
	data, _ = r.stakeABI.Pack("topOf", minter)
	out, _ = r.cli.CallContract(ctx, ethereum.CallMsg{To: &r.stakeAddr, Data: data}, nil)
	var D common.Address
	r.stakeABI.UnpackIntoInterface(&D, "topOf", out)

	priceWei := priceToWei(m.UsdtPaid)
	holdMinVal := new(big.Int).Mul(big.NewInt(int64(holdMinU)), big.NewInt(1e18))

	// 4. isTopEligible(D) — D 合格则 D 拿 d_bps%，不合格则归项目方
	data, _ = r.stakeABI.Pack("isTopEligible", D)
	out, _ = r.cli.CallContract(ctx, ethereum.CallMsg{To: &r.stakeAddr, Data: data}, nil)
	var dEligible bool
	r.stakeABI.UnpackIntoInterface(&dEligible, "isTopEligible", out)

	// 5. holdingValueOf(B) >= holdMinU — B 合格则 B 拿 b_bps%，不合格则归项目方
	data, _ = r.nftABI.Pack("holdingValueOf", B)
	out, _ = r.cli.CallContract(ctx, ethereum.CallMsg{To: &r.nftAddr, Data: data}, nil)
	var bVal big.Int
	r.nftABI.UnpackIntoInterface(&bVal, "holdingValueOf", out)
	bEligible := bVal.Cmp(holdMinVal) >= 0

	// 分别计算：B b_bps%、D d_bps%，各自不满足时其份额归项目方
	toB := big.NewInt(0)
	toD := big.NewInt(0)
	toProj := big.NewInt(0)
	amtB := new(big.Int).Mul(priceWei, big.NewInt(int64(bBps)))
	amtB.Div(amtB, big.NewInt(10000))
	amtD := new(big.Int).Mul(priceWei, big.NewInt(int64(dBps)))
	amtD.Div(amtD, big.NewInt(10000))
	if bEligible {
		toB = amtB
	} else {
		toProj.Add(toProj, amtB)
	}
	if dEligible {
		toD = amtD
	} else {
		toProj.Add(toProj, amtD)
	}

	// reason: 0=正常 1=仅D不合格 2=仅B不合格 3=无上级(B=0) 4=B和D都不合格
	reason := uint8(0)
	if B == (common.Address{}) {
		reason = 3
	} else if !bEligible && !dEligible {
		reason = 4
	} else if !dEligible {
		reason = 1
	} else if !bEligible {
		reason = 2
	}
	return r.savePending(m, minter.String(), B.String(), D.String(), toB.String(), toD.String(), toProj.String(), reason)
}

func (r *referralRunner) savePending(m *struct {
	ID          uint64
	BlockNumber uint64
	BlockTime   uint64
	TokenID     uint64
	MintAddr    string
	UsdtPaid    float64
}, minter, addrB, addrC, amtB, amtC, amtProj string, reason uint8) error {
	priceWei := priceToWei(m.UsdtPaid)
	model := &referralPendingModel{
		MintedID:        m.ID,
		PurchaseID:      0,
		BlockNumber:    m.BlockNumber,
		BlockTime:      m.BlockTime,
		TokenID:        m.TokenID,
		Minter:         minter,
		Seller:         "0x0",
		Buyer:          "0x0",
		Price:          priceWei.String(),
		AddrB:          addrB,
		AddrC:          addrC,
		AmountToB:      amtB,
		AmountToC:      amtC,
		AmountToProject: amtProj,
		Reason:         reason,
	}
	return r.db.Create(model).Error
}

func priceToWei(f float64) *big.Int {
	x := new(big.Float).SetFloat64(f)
	e18 := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	x.Mul(x, e18)
	i, _ := x.Int(nil)
	return i
}
