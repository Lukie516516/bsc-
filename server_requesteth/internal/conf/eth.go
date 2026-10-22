// Package conf 链上配置：合约地址、RPC 等，优先从环境变量读取
package conf

import (
	"os"
	"strconv"
	"strings"
)

// 环境变量 key
const (
	EnvBscRpcUrls       = "ETH_BSC_RPC_URLS"         // 逗号分隔、前者优先；读链时与 POLLING 合并，本列表先尝试（自购节点放最前）
	EnvBscRpcUrlsPoll   = "ETH_BSC_RPC_URLS_POLLING" // 事件扫链与读链回退用；不填则自动等于 ETH_BSC_RPC_URLS（只配一处即可）
	EnvNftContract      = "ETH_NFT_CONTRACT"
	EnvStakeContract    = "ETH_STAKE_CONTRACT"
	EnvAicatAdmin       = "ETH_AICAT_ADMIN_CONTRACT"
	EnvAicatAdminNew    = "ETH_AICAT_ADMIN_NEW_CONTRACT"
	EnvUsdtContract     = "ETH_USDT_CONTRACT"
	EnvDividendContract = "ETH_DIVIDEND_CONTRACT"
	EnvSwapContract     = "ETH_SWAP_CONTRACT"
	EnvPrimaryMarket    = "ETH_PRIMARY_MARKET_CONTRACT"
	EnvRegisterContract = "ETH_REGISTER_CONTRACT"
	EnvVaultFactory        = "ETH_VAULT_FACTORY_CONTRACT"
	EnvAicatToken          = "ETH_AICAT_TOKEN_CONTRACT" // AICAT 代币地址（GetLpByOrderId 等用）
	EnvDeployBlockNftMint     = "ETH_DEPLOY_BLOCK_NFT_MINT"                  // 铸造事件起始块（BlindBoxNFT 部署块，BscScan 合约页可查）
	EnvDeployBlockNftListed   = "ETH_DEPLOY_BLOCK_NFT_LISTED"                // Listed 事件起始块；不填则与 ETH_DEPLOY_BLOCK_NFT_MINT 相同
	EnvDeployBlockNftTransfer = "ETH_DEPLOY_BLOCK_NFT_TRANSFER"              // Transfer 事件起始块；不填则与 ETH_DEPLOY_BLOCK_NFT_MINT 相同（同合约）
)

// 默认 BSC RPC（公开节点，用于交易等）
var defaultBscRpcUrls = []string{
	"https://bsc-dataseed4.binance.org/",
	"https://binance.llamarpc.com/",
	"https://bscrpc.com/",
	"https://bsc-pokt.nodies.app/",
	"https://data-seed-prebsc-1-s3.binance.org:8545/",
}

// 默认合约地址（测试网/示例，部署时需替换）
var defaultContracts = map[string]string{
	"nft":             "0xE447b0391d3F03befeC0dC09E25c049132618fd9",
	"stake":           "0x47CC8586f34ecdbbcE595faf7A1402139A37eb23",
	"aicat_admin":     "0x18Ac4F491F4A365B5a66F5c4e44C2b311e516bC7",
	"aicat_admin_new": "0x4984BE11B25b22615f20B96d997AD64fbC7a28D3",
	"usdt":            "0x55d398326f99059fF775485246999027B3197955",
	"dividend":        "0x135d05e0d1a8083525D64B9C9831579D4B6811d2",
	"swap":            "0x3ff0acac62d1c5f74581d7a45b974bcce5f054e1",
	"primary_market":   "0xe817f45dcd271dde3cac7d6ad9cb91de820c688a",
	"register":        "0x135d05e0d1a8083525D64B9C9831579D4B6811d2",
	"vault_factory":    "0x36e79a79829f3a0dad3f928d1ea8c28965b68237",
	"aicat_token":      "0xdA08bA041b901cd6F4744C924b6A111A2F57904e",
}

func splitBscRpcList(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

// GetBscRpcUrls 返回 BSC RPC 列表，优先从 ETH_BSC_RPC_URLS 读取
func GetBscRpcUrls() []string {
	s := os.Getenv(EnvBscRpcUrls)
	if s == "" {
		return defaultBscRpcUrls
	}
	out := splitBscRpcList(s)
	if len(out) == 0 {
		return defaultBscRpcUrls
	}
	return out
}

// 默认事件轮询 RPC（本地未配置任何环境变量时使用）
var defaultBscRpcUrlsPoll = []string{
	"https://bsc-dataseed4.binance.org/",
	"https://bscrpc.com/",
}

// GetBscRpcUrlsForPolling 返回事件扫链用 RPC。
// 若设置了 ETH_BSC_RPC_URLS_POLLING 则用之；否则与 ETH_BSC_RPC_URLS 相同（只填 Allnodes 一处即可全链路只用该节点）。
func GetBscRpcUrlsForPolling() []string {
	s := os.Getenv(EnvBscRpcUrlsPoll)
	if s != "" {
		out := splitBscRpcList(s)
		if len(out) == 0 {
			return defaultBscRpcUrlsPoll
		}
		return out
	}
	if main := os.Getenv(EnvBscRpcUrls); main != "" {
		out := splitBscRpcList(main)
		if len(out) > 0 {
			return out
		}
	}
	return defaultBscRpcUrlsPoll
}

// GetContract 返回合约地址，优先从环境变量读取
func GetContract(key string) string {
	switch key {
	case "nft":
		if v := os.Getenv(EnvNftContract); v != "" {
			return v
		}
	case "stake":
		if v := os.Getenv(EnvStakeContract); v != "" {
			return v
		}
	case "aicat_admin":
		if v := os.Getenv(EnvAicatAdmin); v != "" {
			return v
		}
	case "aicat_admin_new":
		if v := os.Getenv(EnvAicatAdminNew); v != "" {
			return v
		}
	case "usdt":
		if v := os.Getenv(EnvUsdtContract); v != "" {
			return v
		}
	case "dividend":
		if v := os.Getenv(EnvDividendContract); v != "" {
			return v
		}
	case "swap":
		if v := os.Getenv(EnvSwapContract); v != "" {
			return v
		}
	case "primary_market":
		if v := os.Getenv(EnvPrimaryMarket); v != "" {
			return v
		}
	case "register":
		if v := os.Getenv(EnvRegisterContract); v != "" {
			return v
		}
	case "vault_factory":
		if v := os.Getenv(EnvVaultFactory); v != "" {
			return v
		}
	case "aicat_token":
		if v := os.Getenv(EnvAicatToken); v != "" {
			return v
		}
	}
	return defaultContracts[key]
}

// GetDeployBlockNftMint 铸造事件扫描起始块，0 表示用代码默认值
func GetDeployBlockNftMint(defaultBlock uint64) uint64 {
	s := os.Getenv(EnvDeployBlockNftMint)
	if s == "" {
		return defaultBlock
	}
	v, err := strconv.ParseUint(strings.TrimSpace(s), 10, 64)
	if err != nil {
		return defaultBlock
	}
	return v
}

// GetDeployBlockNftListed Listed 事件扫描起始块；未设置时与铸造起始块一致（同一 NFT 合约）
func GetDeployBlockNftListed(defaultBlock uint64) uint64 {
	s := os.Getenv(EnvDeployBlockNftListed)
	if s == "" {
		return GetDeployBlockNftMint(defaultBlock)
	}
	v, err := strconv.ParseUint(strings.TrimSpace(s), 10, 64)
	if err != nil {
		return GetDeployBlockNftMint(defaultBlock)
	}
	return v
}

// GetDeployBlockNftTransfer ERC721 Transfer 扫块起始块。
// 若设置 ETH_DEPLOY_BLOCK_NFT_TRANSFER 则用之；否则与 ETH_DEPLOY_BLOCK_NFT_MINT 一致（避免 Mint 已换新部署块而 Transfer 仍从旧常量爬导致永远追不上）。
func GetDeployBlockNftTransfer(codeDefaultWhenMintUnset uint64) uint64 {
	s := os.Getenv(EnvDeployBlockNftTransfer)
	if s != "" {
		v, err := strconv.ParseUint(strings.TrimSpace(s), 10, 64)
		if err == nil {
			return v
		}
	}
	return GetDeployBlockNftMint(codeDefaultWhenMintUnset)
}
