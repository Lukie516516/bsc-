package quantfund

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type fileConfig struct {
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
	Rebalance struct {
		MinAmountU  int64 `yaml:"min_amount_usd"`
		ChainID     int64 `yaml:"chain_id"`
		IntervalMin int   `yaml:"interval_min"`
	} `yaml:"rebalance"`
}

type Config struct {
	DatabaseSource     string
	RPCURL             string
	USDTAddr           string
	ExchangeAddr       string
	PrivateKey         string
	ChainID            int64
	QuantBps           int
	LedgerBatchSize    int
	TransferBatchSize  int
	ResultBatchSize    int
	SettleBatchSize    int
	MinTransferUSD     int64
	IntervalMin        int
	DryRun             bool
	SettleDryRun       bool
	ResultMode         string
	FixedProfitBps     int
	RandomProfitMinBps int
	RandomProfitMaxBps int
}

func AddCommonFlags(fs *flag.FlagSet, defaultConf string) *string {
	return fs.String("conf", defaultConf, "config dir")
}

func LoadConfig(confDir string) (Config, error) {
	cfg := Config{
		USDTAddr:           "0x55d398326f99059fF775485246999027B3197955",
		ChainID:            56,
		QuantBps:           9500,
		LedgerBatchSize:    500,
		TransferBatchSize:  500,
		ResultBatchSize:    200,
		SettleBatchSize:    50,
		MinTransferUSD:     100,
		IntervalMin:        60,
		ResultMode:         "random_range",
		FixedProfitBps:     1500,
		RandomProfitMinBps: -200,
		RandomProfitMaxBps: 200,
	}

	var fc fileConfig
	confPath := resolveConfigPath(confDir)
	if confPath != "" {
		if data, err := os.ReadFile(confPath); err == nil {
			_ = yaml.Unmarshal(data, &fc)
		}
	}

	cfg.DatabaseSource = firstNonEmpty(
		os.Getenv("REQUESTETH_DB_SOURCE"),
		fc.Data.Database.Source,
	)
	cfg.RPCURL = firstNonEmpty(
		os.Getenv("ETH_RPC_URL"),
		os.Getenv("ETH_BSC_RPC_URL"),
		firstCSV(os.Getenv("ETH_BSC_RPC_URLS")),
		fc.Eth.Bsc.Addr,
	)
	cfg.USDTAddr = firstNonEmpty(os.Getenv("USDT_ADDR"), cfg.USDTAddr)
	cfg.ExchangeAddr = firstNonEmpty(os.Getenv("EXCHANGE_ADDR"))
	cfg.PrivateKey = firstNonEmpty(os.Getenv("REBALANCE_PRIVATE_KEY"))
	cfg.DryRun = isTrue(os.Getenv("QUANT_DRY_RUN"))
	cfg.SettleDryRun = isTrue(os.Getenv("QUANT_SETTLE_DRY_RUN"))
	cfg.ResultMode = firstNonEmpty(os.Getenv("QUANT_RESULT_MODE"), cfg.ResultMode)

	cfg.ChainID = firstInt64(
		os.Getenv("QUANT_CHAIN_ID"),
		cfg.ChainID,
		fc.Rebalance.ChainID,
	)
	cfg.QuantBps = firstInt(
		os.Getenv("QUANT_BPS"),
		cfg.QuantBps,
	)
	cfg.LedgerBatchSize = firstInt(
		os.Getenv("QUANT_LEDGER_BATCH"),
		cfg.LedgerBatchSize,
	)
	cfg.TransferBatchSize = firstInt(
		os.Getenv("QUANT_TRANSFER_BATCH"),
		cfg.TransferBatchSize,
	)
	cfg.ResultBatchSize = firstInt(
		os.Getenv("QUANT_RESULT_BATCH"),
		cfg.ResultBatchSize,
	)
	cfg.SettleBatchSize = firstInt(
		os.Getenv("QUANT_SETTLE_BATCH"),
		cfg.SettleBatchSize,
	)
	cfg.MinTransferUSD = firstInt64(
		os.Getenv("QUANT_MIN_TRANSFER_USDT"),
		cfg.MinTransferUSD,
		fc.Rebalance.MinAmountU,
	)
	cfg.FixedProfitBps = firstSignedInt(
		os.Getenv("QUANT_RESULT_FIXED_BPS"),
		cfg.FixedProfitBps,
	)
	cfg.RandomProfitMinBps = firstSignedInt(
		os.Getenv("QUANT_RESULT_MIN_BPS"),
		cfg.RandomProfitMinBps,
	)
	cfg.RandomProfitMaxBps = firstSignedInt(
		os.Getenv("QUANT_RESULT_MAX_BPS"),
		cfg.RandomProfitMaxBps,
	)
	cfg.IntervalMin = firstInt(
		os.Getenv("QUANT_INTERVAL_MIN"),
		cfg.IntervalMin,
		int64(fc.Rebalance.IntervalMin),
	)

	if cfg.DatabaseSource == "" {
		return cfg, errors.New("missing database source")
	}
	if cfg.QuantBps <= 0 || cfg.QuantBps >= 10000 {
		return cfg, errors.New("QUANT_BPS must be between 1 and 9999")
	}
	if cfg.ResultBatchSize <= 0 {
		cfg.ResultBatchSize = 200
	}
	if cfg.SettleBatchSize <= 0 {
		cfg.SettleBatchSize = 50
	}
	if cfg.RandomProfitMinBps > cfg.RandomProfitMaxBps {
		cfg.RandomProfitMinBps, cfg.RandomProfitMaxBps = cfg.RandomProfitMaxBps, cfg.RandomProfitMinBps
	}

	return cfg, nil
}

func resolveConfigPath(conf string) string {
	conf = strings.TrimSpace(conf)
	if conf == "" {
		return ""
	}
	if strings.EqualFold(filepath.Ext(conf), ".yaml") || strings.EqualFold(filepath.Ext(conf), ".yml") {
		return conf
	}
	return filepath.Join(conf, "config.yaml")
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}

func firstCSV(value string) string {
	for _, part := range strings.Split(value, ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			return part
		}
	}
	return ""
}

func firstInt(value string, fallback int, more ...int64) int {
	if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil && v > 0 {
		return v
	}
	for _, item := range more {
		if item > 0 {
			return int(item)
		}
	}
	return fallback
}

func firstInt64(value string, fallback int64, more ...int64) int64 {
	if v, err := strconv.ParseInt(strings.TrimSpace(value), 10, 64); err == nil && v > 0 {
		return v
	}
	for _, item := range more {
		if item > 0 {
			return item
		}
	}
	return fallback
}

func firstSignedInt(value string, fallback int) int {
	value = strings.TrimSpace(value)
	if value == "" {
		return fallback
	}
	if v, err := strconv.Atoi(value); err == nil {
		return v
	}
	return fallback
}

func isTrue(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "1", "true", "yes", "y", "on":
		return true
	default:
		return false
	}
}
