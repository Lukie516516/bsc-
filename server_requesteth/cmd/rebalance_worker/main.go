// 调仓脚本：定时将 accountA 的 USDT 余额按比例转到交易所地址
// 运行：go run . -conf=../../configs
// 环境变量：ETH_RPC_URL, NFT_ADDR, USDT_ADDR, ACCOUNT_A_ADDR, EXCHANGE_ADDR, REBALANCE_PRIVATE_KEY
package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/yaml.v3"
)

var (
	flagConf = flag.String("conf", "../../configs", "config dir")
	cfgRPC   string
	cfgNFT   string
	cfgUSDT  string
	cfgAcctA string
	cfgExch  string
	cfgKey   string

	rebalanceBps  = 9500  // 95%（与铸造分账 accountAFeeBps 一致，5% 留 staking）
	minAmountWei  = new(big.Int).Mul(big.NewInt(100), big.NewInt(1e18)) // 100 USDT 起转
	chainID       = int64(56)
	intervalMins  = 60
)

type configStruct struct {
	Eth struct {
		Bsc struct {
			Addr string `yaml:"addr"`
		} `yaml:"bsc"`
	} `yaml:"eth"`
	Rebalance struct {
		RebalanceBps int   `yaml:"rebalance_bps"`  // 9500=95%
		MinAmountU   int64 `yaml:"min_amount_usd"` // 起转阈值 USDT
		ChainID      int   `yaml:"chain_id"`
		IntervalMin  int   `yaml:"interval_min"`
	} `yaml:"rebalance"`
}

func loadConfig() {
	cfgRPC = os.Getenv("ETH_RPC_URL")
	cfgNFT = os.Getenv("NFT_ADDR")
	cfgUSDT = os.Getenv("USDT_ADDR")
	cfgAcctA = os.Getenv("ACCOUNT_A_ADDR")
	cfgExch = os.Getenv("EXCHANGE_ADDR")
	cfgKey = os.Getenv("REBALANCE_PRIVATE_KEY")

	if cfgUSDT == "" {
		cfgUSDT = "0x55d398326f99059fF775485246999027B3197955" // BSC USDT
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
	if c.Rebalance.RebalanceBps > 0 {
		rebalanceBps = c.Rebalance.RebalanceBps
	}
	if c.Rebalance.MinAmountU > 0 {
		minAmountWei = new(big.Int).Mul(big.NewInt(c.Rebalance.MinAmountU), big.NewInt(1e18))
	}
	if c.Rebalance.ChainID > 0 {
		chainID = int64(c.Rebalance.ChainID)
	}
	if c.Rebalance.IntervalMin > 0 {
		intervalMins = c.Rebalance.IntervalMin
	}
}

// ERC20 最小 ABI
const erc20ABI = `[
	{"inputs":[{"name":"account","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"stateMutability":"view","type":"function"},
	{"inputs":[{"name":"to","type":"address"},{"name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"}
]`

func main() {
	flag.Parse()
	loadConfig()

	if cfgRPC == "" {
		log.Fatal("need ETH_RPC_URL or config eth.bsc.addr")
	}
	if cfgKey == "" {
		log.Fatal("need REBALANCE_PRIVATE_KEY env (持有 USDT 的热钱包私钥，hex 无 0x)")
	}
	if cfgExch == "" {
		log.Fatal("need EXCHANGE_ADDR env (交易所充值地址)")
	}

	// 转出地址 = 私钥对应地址（即持有 USDT 的热钱包，通常为 accountA）
	fromAddr := crypto.PubkeyToAddress(mustParseKey(cfgKey).PublicKey)
	cfgAcctA = fromAddr.Hex()

	cli, err := ethclient.Dial(cfgRPC)
	if err != nil {
		log.Fatalf("eth client: %v", err)
	}
	defer cli.Close()

	parsed, _ := abi.JSON(strings.NewReader(erc20ABI))
	auth, _ := bind.NewKeyedTransactorWithChainID(mustParseKey(cfgKey), big.NewInt(chainID))

	runner := &rebalanceRunner{
		cli:    cli,
		usdt:   common.HexToAddress(cfgUSDT),
		abi:    parsed,
		auth:   auth,
		from:   common.HexToAddress(cfgAcctA),
		to:     common.HexToAddress(cfgExch),
	}
	ctx := context.Background()
	for {
		if err := runner.run(ctx); err != nil {
			log.Printf("rebalance: %v", err)
		}
		time.Sleep(time.Duration(intervalMins) * time.Minute)
	}
}

func mustParseKey(hex string) *ecdsa.PrivateKey {
	if strings.HasPrefix(hex, "0x") {
		hex = hex[2:]
	}
	k, err := crypto.HexToECDSA(hex)
	if err != nil {
		log.Fatalf("parse key: %v", err)
	}
	return k
}

type rebalanceRunner struct {
	cli  *ethclient.Client
	usdt common.Address
	abi  abi.ABI
	auth *bind.TransactOpts
	from common.Address
	to   common.Address
}

func (r *rebalanceRunner) run(ctx context.Context) error {
	bal, err := r.balanceOf(ctx, r.from)
	if err != nil {
		return err
	}
	if bal.Cmp(minAmountWei) < 0 {
		log.Printf("balance %s < min %s, skip", bal.String(), minAmountWei.String())
		return nil
	}
	amount := new(big.Int).Mul(bal, big.NewInt(int64(rebalanceBps)))
	amount.Div(amount, big.NewInt(10000))
	if amount.Sign() <= 0 {
		return nil
	}
	txHash, err := r.transfer(ctx, r.to, amount)
	if err != nil {
		return err
	}
	log.Printf("rebalance %s USDT -> %s tx=%s", weiToUsd(amount), r.to.Hex(), txHash)
	return nil
}

func (r *rebalanceRunner) balanceOf(ctx context.Context, addr common.Address) (*big.Int, error) {
	data, _ := r.abi.Pack("balanceOf", addr)
	msg := ethereum.CallMsg{To: &r.usdt, Data: data}
	out, err := r.cli.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}
	if len(out) < 32 {
		return big.NewInt(0), nil
	}
	return new(big.Int).SetBytes(out[:32]), nil
}

func (r *rebalanceRunner) transfer(ctx context.Context, to common.Address, amount *big.Int) (string, error) {
	data, _ := r.abi.Pack("transfer", to, amount)
	nonce, err := r.cli.PendingNonceAt(ctx, r.auth.From)
	if err != nil {
		return "", err
	}
	gasPrice, err := r.cli.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}
	tx := types.NewTransaction(
		nonce,
		r.usdt,
		big.NewInt(0),
		100000,
		gasPrice,
		data,
	)
	signed, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID)), mustParseKey(cfgKey))
	if err != nil {
		return "", err
	}
	err = r.cli.SendTransaction(ctx, signed)
	if err != nil {
		return "", err
	}
	return signed.Hash().Hex(), nil
}

func weiToUsd(w *big.Int) string {
	if w == nil || w.Sign() == 0 {
		return "0"
	}
	f := new(big.Float).SetInt(w)
	f.Quo(f, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)))
	s, _ := f.Float64()
	return fmt.Sprintf("%.2f", s)
}

