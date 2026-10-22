package quantfund

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const erc20ABI = `[
	{"inputs":[{"name":"account","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"stateMutability":"view","type":"function"},
	{"inputs":[{"name":"to","type":"address"},{"name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"}
]`

type ERC20Runner struct {
	Client  *ethclient.Client
	Token   common.Address
	To      common.Address
	From    common.Address
	Auth    *bind.TransactOpts
	ABI     abi.ABI
	ChainID int64
	Key     *ecdsa.PrivateKey
}

func NewERC20Runner(cfg Config) (*ERC20Runner, error) {
	if strings.TrimSpace(cfg.RPCURL) == "" {
		return nil, fmt.Errorf("missing ETH_RPC_URL / eth.bsc.addr")
	}
	if strings.TrimSpace(cfg.PrivateKey) == "" {
		return nil, fmt.Errorf("missing REBALANCE_PRIVATE_KEY")
	}
	if strings.TrimSpace(cfg.ExchangeAddr) == "" {
		return nil, fmt.Errorf("missing EXCHANGE_ADDR")
	}

	key, err := ParsePrivateKey(cfg.PrivateKey)
	if err != nil {
		return nil, err
	}
	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		client.Close()
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(cfg.ChainID))
	if err != nil {
		client.Close()
		return nil, err
	}

	from := crypto.PubkeyToAddress(key.PublicKey)
	return &ERC20Runner{
		Client:  client,
		Token:   common.HexToAddress(cfg.USDTAddr),
		To:      common.HexToAddress(cfg.ExchangeAddr),
		From:    from,
		Auth:    auth,
		ABI:     parsed,
		ChainID: cfg.ChainID,
		Key:     key,
	}, nil
}

func (r *ERC20Runner) Close() {
	if r.Client != nil {
		r.Client.Close()
	}
}

func ParsePrivateKey(hexKey string) (*ecdsa.PrivateKey, error) {
	hexKey = strings.TrimSpace(hexKey)
	hexKey = strings.TrimPrefix(hexKey, "0x")
	return crypto.HexToECDSA(hexKey)
}

func (r *ERC20Runner) BalanceOf(ctx context.Context, addr common.Address) (*big.Int, error) {
	data, err := r.ABI.Pack("balanceOf", addr)
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{To: &r.Token, Data: data}
	out, err := r.Client.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}
	if len(out) < 32 {
		return big.NewInt(0), nil
	}
	return new(big.Int).SetBytes(out[:32]), nil
}

func (r *ERC20Runner) Transfer(ctx context.Context, amount *big.Int) (string, error) {
	data, err := r.ABI.Pack("transfer", r.To, amount)
	if err != nil {
		return "", err
	}
	nonce, err := r.Client.PendingNonceAt(ctx, r.Auth.From)
	if err != nil {
		return "", err
	}
	gasPrice, err := r.Client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}

	tx := types.NewTransaction(
		nonce,
		r.Token,
		big.NewInt(0),
		100000,
		gasPrice,
		data,
	)
	signed, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(r.ChainID)), r.Key)
	if err != nil {
		return "", err
	}
	if err := r.Client.SendTransaction(ctx, signed); err != nil {
		return "", err
	}
	return signed.Hash().Hex(), nil
}
