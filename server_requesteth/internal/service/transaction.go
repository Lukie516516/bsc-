package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"

	"requestEth/internal/biz"
	"requestEth/internal/conf"
	"requestEth/internal/data"

	pb "requestEth/api/requestEth/v1"
)

func formatReward4(value float64) string {
	return strconv.FormatFloat(value, 'f', 4, 64)
}

// 铸造分佣打款环境变量
const (
	envReferralPayoutKey = "REFERRAL_PAYOUT_PRIVATE_KEY"
	envReferralUsdtAddr  = "REFERRAL_USDT_ADDR"
	envAICATAdminKey     = "AICAT_ADMIN_PRIVATE_KEY"      // AICAT 相关接口需配置
	envSetRewardAdminKey = "SET_REWARD_ADMIN_PRIVATE_KEY" // SetReward/SetRewardTwo 需配置
)

const (
	addressBoxEthReadTimeout     = 1200 * time.Millisecond
	addressBoxVaultResolveWall   = 1500 * time.Millisecond
	addressBoxVaultCacheTTL      = 5 * time.Minute
	addressBoxEmptyVaultCacheTTL = 90 * time.Second
	addressBoxRepairSyncTimeout  = 8 * time.Second
	addressBoxPageSize           = 10
	boxRangeEthReadTimeout       = 4 * time.Second
	boxRangeWorkerCount          = 4
	// maxBoxTokenRangeSpan：get_box_new / get_box_open 的 start、end 为链上 tokenId。过大则 RPC/缓冲与切片峰值过高易 OOM。
	maxBoxTokenRangeSpan         uint64 = 2000
	allInfoChainCacheTTL                = 15 * time.Second
	allInfoChainReadTimeout             = 20 * time.Second
	allInfoMintPageTimeout              = 6 * time.Second
	allInfoMintBatchSize         uint64 = 200
	marketPriceCacheTTL                 = 10 * time.Second
	userRelationEthReadTimeout          = 1200 * time.Millisecond
	userRelationResolveWall             = 1500 * time.Millisecond
	userRelationCacheTTL                = 5 * time.Minute
	userRelationNegativeCacheTTL        = 90 * time.Second
	zeroAddressHex                      = "0x0000000000000000000000000000000000000000"
)

type vaultCacheEntry struct {
	vault     string
	expiresAt time.Time
}

type relationCacheEntry struct {
	summary   biz.UserRelationSummary
	expiresAt time.Time
}

type priceCacheEntry struct {
	price     float64
	expiresAt time.Time
}

type allInfoChainCacheEntry struct {
	summary   allInfoMintSummary
	expiresAt time.Time
}

type allInfoMintSummary struct {
	totalCount      uint64
	totalSum        *big.Int
	rangeCount      uint64
	rangeSum        *big.Int
	notOpenCount    uint64
	notOpenSum      *big.Int
	tierCounts      [4]uint64
	tierSums        [4]*big.Int
	totalOpenReward *big.Int
}

type TransactionService struct {
	pb.UnimplementedTransactionServer
	log             *log.Helper
	ac              *biz.AppUsecase
	refRepo         *data.ReferralPendingRepo
	vaultCacheMu    sync.RWMutex
	vaultCache      map[string]vaultCacheEntry
	relationCacheMu sync.RWMutex
	relationCache   map[string]relationCacheEntry
	priceCacheMu    sync.RWMutex
	priceCache      map[uint64]priceCacheEntry
	allInfoCacheMu  sync.RWMutex
	allInfoCache    map[string]allInfoChainCacheEntry
}

func NewTransactionService(ac *biz.AppUsecase, refRepo *data.ReferralPendingRepo, logger log.Logger) *TransactionService {
	return &TransactionService{
		ac:            ac,
		refRepo:       refRepo,
		log:           log.NewHelper(logger),
		vaultCache:    make(map[string]vaultCacheEntry),
		relationCache: make(map[string]relationCacheEntry),
		priceCache:    make(map[uint64]priceCacheEntry),
		allInfoCache:  make(map[string]allInfoChainCacheEntry),
	}
}

var (
	stakeRelationABIOnce sync.Once
	stakeRelationABIInst abi.ABI
	stakeRelationABIErr  error
)

const stakeRelationABIJSON = `[
  {"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"registered","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},
  {"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"parentOf","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},
  {"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"topOf","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}
]`

// bscReadURLs 供读链（eth_call、FilterLogs 等）按顺序回退的 URL 列表。
// 先 ETH_BSC_RPC_URLS，再 ETH_BSC_RPC_URLS_POLLING，去重保序；主配置中的节点（如自购 Allnodes）应写在 ETH_BSC_RPC_URLS 最前，避免被轮询长列表里同 URL 抢先。
func bscReadURLs() []string {
	deduped := make([]string, 0)
	seen := make(map[string]struct{})
	appendURLs := func(urls []string) {
		for _, raw := range urls {
			raw = strings.TrimSpace(raw)
			if raw == "" {
				continue
			}
			if _, ok := seen[raw]; ok {
				continue
			}
			seen[raw] = struct{}{}
			deduped = append(deduped, raw)
		}
	}
	appendURLs(conf.GetBscRpcUrls())
	appendURLs(conf.GetBscRpcUrlsForPolling())
	return deduped
}

func currentNftContract() common.Address {
	return common.HexToAddress(conf.GetContract("nft"))
}

func currentSwapContract() common.Address {
	return common.HexToAddress(conf.GetContract("swap"))
}

func normalizeAddress(raw string) string {
	return strings.ToLower(strings.TrimSpace(raw))
}

func parseDecimalOrderID(raw string) (*big.Int, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, false
	}
	id, ok := new(big.Int).SetString(raw, 10)
	if !ok || id == nil || id.Sign() < 0 {
		return nil, false
	}
	return id, true
}

func parseRequiredBigInt(raw string, field string) (*big.Int, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, fmt.Errorf("%s required", field)
	}
	v, ok := new(big.Int).SetString(raw, 10)
	if !ok || v == nil || v.Sign() < 0 {
		return nil, fmt.Errorf("invalid %s", field)
	}
	return v, nil
}

// parseRequiredSignedBigInt 用于 int256 参数（如 setRewardOnce 的盈亏，可为负）
func parseRequiredSignedBigInt(raw string, field string) (*big.Int, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, fmt.Errorf("%s required", field)
	}
	v, ok := new(big.Int).SetString(raw, 10)
	if !ok || v == nil {
		return nil, fmt.Errorf("invalid %s", field)
	}
	return v, nil
}

func parseBoxTokenRange(startRaw, endRaw string) (uint64, uint64, bool) {
	start, err := strconv.ParseUint(strings.TrimSpace(startRaw), 10, 64)
	if err != nil {
		return 0, 0, false
	}
	end, err := strconv.ParseUint(strings.TrimSpace(endRaw), 10, 64)
	if err != nil || end < start {
		return 0, 0, false
	}
	return start, end, true
}

func clampBoxTokenRange(startTokenID, endTokenID uint64) (uint64, uint64, bool) {
	if endTokenID < startTokenID {
		return startTokenID, startTokenID, false
	}
	span := endTokenID - startTokenID + 1
	if span <= maxBoxTokenRangeSpan {
		return startTokenID, endTokenID, false
	}
	return startTokenID, startTokenID + maxBoxTokenRangeSpan - 1, true
}

func filterBoxInfosByTokenRange(infos []BlindBoxNFTBoxInfo, startTokenID, endTokenID uint64, wantOpened bool) []BlindBoxNFTBoxInfo {
	res := make([]BlindBoxNFTBoxInfo, 0, len(infos))
	for _, info := range infos {
		if info.TokenId == nil {
			continue
		}
		tokenID := info.TokenId.Uint64()
		if tokenID < startTokenID || tokenID > endTokenID {
			continue
		}
		if wantOpened {
			if info.OpenedAt == 0 {
				continue
			}
		} else if info.OpenedAt != 0 {
			continue
		}
		res = append(res, info)
	}
	return res
}

func normalizeBoxInfo(tokenID uint64, raw struct {
	Tier        uint8
	UsdtPaid    *big.Int
	MintedAt    uint64
	OpenedAt    uint64
	Reward      *big.Int
	RewardSetAt uint64
	TokenId     *big.Int
}) (BlindBoxNFTBoxInfo, bool) {
	info := BlindBoxNFTBoxInfo{
		Tier:        raw.Tier,
		UsdtPaid:    raw.UsdtPaid,
		MintedAt:    raw.MintedAt,
		OpenedAt:    raw.OpenedAt,
		Reward:      raw.Reward,
		RewardSetAt: raw.RewardSetAt,
		TokenId:     raw.TokenId,
	}
	if info.UsdtPaid == nil {
		info.UsdtPaid = new(big.Int)
	}
	if info.Reward == nil {
		info.Reward = new(big.Int)
	}
	if info.TokenId == nil || info.TokenId.Sign() == 0 {
		if info.MintedAt == 0 && info.UsdtPaid.Sign() == 0 && info.OpenedAt == 0 && info.Reward.Sign() == 0 && info.RewardSetAt == 0 {
			return BlindBoxNFTBoxInfo{}, false
		}
		info.TokenId = new(big.Int).SetUint64(tokenID)
	}
	return info, true
}

func loadBoxInfosByTokenRange(
	ctx context.Context,
	instance *Nft,
	startTokenID uint64,
	endTokenID uint64,
	wantOpened bool,
) ([]BlindBoxNFTBoxInfo, error) {
	if endTokenID < startTokenID {
		return make([]BlindBoxNFTBoxInfo, 0), nil
	}

	type task struct {
		index   int
		tokenID uint64
	}
	type result struct {
		index int
		info  BlindBoxNFTBoxInfo
		ok    bool
		err   error
	}

	total := int(endTokenID-startTokenID) + 1
	tasks := make(chan task)
	workers := boxRangeWorkerCount
	if total < workers {
		workers = total
	}
	// 勿用 make(chan result, total)：大 total 时 channel 缓冲会按元素大小占用巨量内存，易 OOM。
	results := make(chan result, workers*2)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range tasks {
				callCtx, cancel := context.WithTimeout(ctx, boxRangeEthReadTimeout)
				raw, err := instance.BoxInfo(&bind.CallOpts{Context: callCtx}, new(big.Int).SetUint64(t.tokenID))
				cancel()
				if err != nil {
					results <- result{index: t.index, err: err}
					continue
				}
				info, ok := normalizeBoxInfo(t.tokenID, raw)
				if ok && wantOpened && info.OpenedAt == 0 {
					ok = false
				}
				if ok && !wantOpened && info.OpenedAt != 0 {
					ok = false
				}
				results <- result{index: t.index, info: info, ok: ok}
			}
		}()
	}

	go func() {
		for i := 0; i < total; i++ {
			tasks <- task{
				index:   i,
				tokenID: startTokenID + uint64(i),
			}
		}
		close(tasks)
		wg.Wait()
		close(results)
	}()

	type indexed struct {
		idx  int
		info BlindBoxNFTBoxInfo
	}
	collected := make([]indexed, 0, min(total, 64))
	var firstErr error
	for r := range results {
		if r.err != nil {
			if firstErr == nil {
				firstErr = r.err
			}
			continue
		}
		if !r.ok {
			continue
		}
		collected = append(collected, indexed{r.index, r.info})
	}
	if firstErr != nil {
		return nil, firstErr
	}

	sort.Slice(collected, func(i, j int) bool { return collected[i].idx < collected[j].idx })
	res := make([]BlindBoxNFTBoxInfo, len(collected))
	for i := range collected {
		res[i] = collected[i].info
	}
	return res, nil
}

func syncAddressBoxState(ctx context.Context, s *TransactionService) {
	if s == nil {
		return
	}
	syncCtx, cancel := context.WithTimeout(ctx, addressBoxRepairSyncTimeout)
	defer cancel()

	_, _ = s.GetBoxMintEvent(syncCtx, &pb.GetBoxMintEventRequest{})
	_, _ = s.GetBoxTransferEvent(syncCtx, &pb.GetBoxTransferEventRequest{})
	_, _ = s.GetBoxListEvent(syncCtx, &pb.GetBoxListEventRequest{})
	_, _ = s.GetBoxUnListEvent(syncCtx, &pb.GetBoxUnListEventRequest{})
	_, _ = s.GetBoxBuyEvent(syncCtx, &pb.GetBoxBuyEventRequest{})
	_, _ = s.GetBoxOpenEvent(syncCtx, &pb.GetBoxOpenEventRequest{})
	s.ac.UpdateBox(syncCtx, &pb.UpdateBoxRequest{})
}

func newAllInfoMintSummary() allInfoMintSummary {
	s := allInfoMintSummary{
		totalSum:        new(big.Int),
		rangeSum:        new(big.Int),
		notOpenSum:      new(big.Int),
		totalOpenReward: new(big.Int),
	}
	for i := range s.tierSums {
		s.tierSums[i] = new(big.Int)
	}
	return s
}

func addBigInt(dst **big.Int, src *big.Int) {
	if *dst == nil {
		*dst = new(big.Int)
	}
	if src == nil {
		return
	}
	(*dst).Add(*dst, src)
}

func formatWei18String(v *big.Int) string {
	if v == nil {
		return "0.000000000000000000"
	}
	sign := ""
	n := new(big.Int).Set(v)
	if n.Sign() < 0 {
		sign = "-"
		n.Abs(n)
	}
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	intPart := new(big.Int).Quo(n, divisor)
	fracPart := new(big.Int).Mod(n, divisor)
	frac := fracPart.String()
	if len(frac) < 18 {
		frac = strings.Repeat("0", 18-len(frac)) + frac
	}
	return sign + intPart.String() + "." + frac
}

func allInfoCacheKey(start, end uint64) string {
	return currentNftContract().Hex() + ":" + strconv.FormatUint(start, 10) + ":" + strconv.FormatUint(end, 10)
}

func (s *TransactionService) getCachedAllInfoSummary(start, end uint64) (allInfoMintSummary, bool) {
	key := allInfoCacheKey(start, end)
	s.allInfoCacheMu.RLock()
	entry, ok := s.allInfoCache[key]
	s.allInfoCacheMu.RUnlock()
	if !ok || time.Now().After(entry.expiresAt) {
		return allInfoMintSummary{}, false
	}
	return entry.summary, true
}

func (s *TransactionService) setCachedAllInfoSummary(start, end uint64, summary allInfoMintSummary) {
	key := allInfoCacheKey(start, end)
	s.allInfoCacheMu.Lock()
	s.allInfoCache[key] = allInfoChainCacheEntry{
		summary:   summary,
		expiresAt: time.Now().Add(allInfoChainCacheTTL),
	}
	s.allInfoCacheMu.Unlock()
}

func accumulateAllInfoSummary(summary *allInfoMintSummary, box BlindBoxNFTBoxInfo, start, end uint64) {
	if summary == nil || box.TokenId == nil || box.TokenId.Sign() <= 0 {
		return
	}
	usdt := box.UsdtPaid
	if usdt == nil {
		usdt = new(big.Int)
	}
	reward := box.Reward
	if reward == nil {
		reward = new(big.Int)
	}

	summary.totalCount++
	addBigInt(&summary.totalSum, usdt)

	if box.MintedAt >= start && box.MintedAt <= end {
		summary.rangeCount++
		addBigInt(&summary.rangeSum, usdt)
	}

	if box.OpenedAt == 0 {
		summary.notOpenCount++
		addBigInt(&summary.notOpenSum, usdt)
	}

	if box.RewardSetAt > 0 {
		addBigInt(&summary.totalOpenReward, reward)
	}

	tier := int(box.Tier)
	if tier >= 1 && tier <= 3 {
		summary.tierCounts[tier]++
		addBigInt(&summary.tierSums[tier], usdt)
	}
}

func (s *TransactionService) loadMintSummaryFromChain(ctx context.Context, start, end uint64) (allInfoMintSummary, error) {
	if cached, ok := s.getCachedAllInfoSummary(start, end); ok {
		return cached, nil
	}

	urls := bscReadURLs()
	if len(urls) == 0 {
		return allInfoMintSummary{}, fmt.Errorf("no RPC urls")
	}

	var lastErr error
	for _, url := range urls {
		urlErr := error(nil)
		dialCtx, dialCancel := context.WithTimeout(ctx, eventSyncDialTimeout)
		client, err := ethclient.DialContext(dialCtx, url)
		dialCancel()
		if err != nil {
			lastErr = err
			continue
		}

		nftAddr := common.HexToAddress(conf.GetContract("nft"))
		instance, err := NewNft(nftAddr, client)
		if err != nil {
			client.Close()
			lastErr = err
			continue
		}

		callCtx, cancel := context.WithTimeout(ctx, allInfoChainReadTimeout)
		total, err := instance.MintedTokenIdsLength(&bind.CallOpts{Context: callCtx})
		cancel()
		if err != nil {
			client.Close()
			lastErr = err
			continue
		}
		if total == nil || total.Sign() <= 0 {
			client.Close()
			summary := newAllInfoMintSummary()
			s.setCachedAllInfoSummary(start, end, summary)
			return summary, nil
		}

		totalCount := total.Uint64()
		summary := newAllInfoMintSummary()
		for offset := uint64(0); offset < totalCount; offset += allInfoMintBatchSize {
			count := allInfoMintBatchSize
			if remain := totalCount - offset; remain < count {
				count = remain
			}
			pageCtx, pageCancel := context.WithTimeout(ctx, allInfoMintPageTimeout)
			boxes, err := instance.GetMintedByPage(
				&bind.CallOpts{Context: pageCtx},
				new(big.Int).SetUint64(offset),
				new(big.Int).SetUint64(count),
			)
			pageCancel()
			if err != nil {
				urlErr = err
				break
			}
			for _, box := range boxes {
				accumulateAllInfoSummary(&summary, box, start, end)
			}
		}
		client.Close()
		if urlErr != nil {
			lastErr = urlErr
			continue
		}
		s.setCachedAllInfoSummary(start, end, summary)
		return summary, nil
	}

	if lastErr == nil {
		lastErr = fmt.Errorf("all RPCs failed")
	}
	return allInfoMintSummary{}, lastErr
}

func applyExcludeTokenIDsToMinted(minted []*biz.NftMinted, exclude []uint64) []*biz.NftMinted {
	if len(minted) == 0 || len(exclude) == 0 {
		return minted
	}
	excluded := make(map[uint64]struct{}, len(exclude))
	for _, id := range exclude {
		excluded[id] = struct{}{}
	}
	filtered := make([]*biz.NftMinted, 0, len(minted))
	for _, v := range minted {
		if _, skip := excluded[v.TokenID]; skip {
			continue
		}
		filtered = append(filtered, v)
	}
	return filtered
}

func paginateMintedRows(minted []*biz.NftMinted, page uint64) []*biz.NftMinted {
	if page == 0 {
		page = 1
	}
	start := int((page - 1) * addressBoxPageSize)
	if start >= len(minted) {
		return make([]*biz.NftMinted, 0)
	}
	end := start + addressBoxPageSize
	if end > len(minted) {
		end = len(minted)
	}
	return minted[start:end]
}

func normalizePrivateKeyHex(raw string) string {
	raw = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, raw)
	raw = strings.TrimPrefix(raw, "0x")
	raw = strings.TrimPrefix(raw, "0X")
	return raw
}

func parsePrivateKeyHex(raw string) (*ecdsa.PrivateKey, error) {
	normalized := normalizePrivateKeyHex(raw)
	if normalized == "" {
		return nil, fmt.Errorf("empty private key")
	}
	return crypto.HexToECDSA(normalized)
}

func readRPCURLs() []string {
	urls := conf.GetBscRpcUrls()
	if len(urls) == 0 {
		return bscReadURLs()
	}
	return urls
}

func getStakeRelationABI() (abi.ABI, error) {
	stakeRelationABIOnce.Do(func() {
		stakeRelationABIInst, stakeRelationABIErr = abi.JSON(strings.NewReader(stakeRelationABIJSON))
	})
	return stakeRelationABIInst, stakeRelationABIErr
}

func (s *TransactionService) getCachedVault(address string) (vault string, fresh bool, hasStale bool) {
	s.vaultCacheMu.RLock()
	entry, ok := s.vaultCache[address]
	s.vaultCacheMu.RUnlock()
	if !ok {
		return "", false, false
	}
	if time.Now().Before(entry.expiresAt) {
		return entry.vault, true, true
	}
	return entry.vault, false, true
}

func (s *TransactionService) setCachedVault(address, vault string, ttl time.Duration) {
	s.vaultCacheMu.Lock()
	s.vaultCache[address] = vaultCacheEntry{
		vault:     vault,
		expiresAt: time.Now().Add(ttl),
	}
	s.vaultCacheMu.Unlock()
}

func (s *TransactionService) getCachedRelation(address string) (biz.UserRelationSummary, bool) {
	s.relationCacheMu.RLock()
	entry, ok := s.relationCache[address]
	s.relationCacheMu.RUnlock()
	if !ok || time.Now().After(entry.expiresAt) {
		return biz.UserRelationSummary{}, false
	}
	return entry.summary, true
}

func (s *TransactionService) setCachedRelation(summary biz.UserRelationSummary, ttl time.Duration) {
	s.relationCacheMu.Lock()
	s.relationCache[normalizeAddress(summary.Address)] = relationCacheEntry{
		summary:   summary,
		expiresAt: time.Now().Add(ttl),
	}
	s.relationCacheMu.Unlock()
}

func (s *TransactionService) getCachedMarketPrice(tokenID uint64) (float64, bool) {
	s.priceCacheMu.RLock()
	entry, ok := s.priceCache[tokenID]
	s.priceCacheMu.RUnlock()
	if !ok || time.Now().After(entry.expiresAt) {
		return 0, false
	}
	return entry.price, true
}

func (s *TransactionService) setCachedMarketPrice(tokenID uint64, price float64, ttl time.Duration) {
	s.priceCacheMu.Lock()
	s.priceCache[tokenID] = priceCacheEntry{
		price:     price,
		expiresAt: time.Now().Add(ttl),
	}
	s.priceCacheMu.Unlock()
}

func dialFirstEthClient(ctx context.Context, urls []string) *ethclient.Client {
	for _, url := range urls {
		dialCtx, cancel := context.WithTimeout(ctx, addressBoxEthReadTimeout)
		client, err := ethclient.DialContext(dialCtx, url)
		cancel()
		if err == nil {
			return client
		}
	}
	return nil
}

func (s *TransactionService) loadSecondaryBuyPrices(ctx context.Context, tokenIDs []uint64) map[uint64]float64 {
	prices := make(map[uint64]float64, len(tokenIDs))
	missing := make([]uint64, 0, len(tokenIDs))
	for _, tokenID := range tokenIDs {
		if price, ok := s.getCachedMarketPrice(tokenID); ok {
			prices[tokenID] = price
			continue
		}
		missing = append(missing, tokenID)
	}
	if len(missing) == 0 {
		return prices
	}

	client := dialFirstEthClient(ctx, bscReadURLs())
	if client == nil {
		return prices
	}
	defer client.Close()

	instance, err := NewNft(common.HexToAddress(conf.GetContract("nft")), client)
	if err != nil {
		return prices
	}

	var (
		wg  sync.WaitGroup
		mu  sync.Mutex
		sem = make(chan struct{}, 4)
	)

	for _, tokenID := range missing {
		tokenID := tokenID
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			callCtx, cancel := context.WithTimeout(ctx, addressBoxEthReadTimeout)
			defer cancel()
			tmpRes, err := instance.SecondaryBuyPrice(&bind.CallOpts{Context: callCtx}, new(big.Int).SetUint64(tokenID))
			if err != nil || tmpRes == nil {
				return
			}
			price := BigIntToFloat64(tmpRes, 18)
			mu.Lock()
			prices[tokenID] = price
			mu.Unlock()
			s.setCachedMarketPrice(tokenID, price, marketPriceCacheTTL)
		}()
	}
	wg.Wait()
	return prices
}

func fetchVaultOfFromURL(ctx context.Context, url string, user common.Address) (string, bool) {
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return "", false
	}
	defer client.Close()

	instance, err := NewFactory(common.HexToAddress(conf.GetContract("vault_factory")), client)
	if err != nil {
		return "", false
	}
	vault, err := instance.VaultOf(&bind.CallOpts{Context: ctx}, user)
	if err != nil {
		return "", false
	}
	if strings.EqualFold(vault.Hex(), zeroAddressHex) {
		return "", true
	}
	return vault.Hex(), true
}

func resolveFirstVaultAddress(ctx context.Context, address string) (string, bool) {
	if !common.IsHexAddress(address) {
		return "", false
	}
	urls := bscReadURLs()
	if len(urls) == 0 {
		return "", false
	}

	raceCtx, cancel := context.WithTimeout(ctx, addressBoxVaultResolveWall)
	defer cancel()

	type result struct {
		vault string
	}
	results := make(chan result, len(urls))
	var wg sync.WaitGroup
	user := common.HexToAddress(address)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			callCtx, callCancel := context.WithTimeout(raceCtx, addressBoxEthReadTimeout)
			defer callCancel()
			vault, ok := fetchVaultOfFromURL(callCtx, url, user)
			if !ok {
				return
			}
			select {
			case results <- result{vault: vault}:
				cancel()
			case <-raceCtx.Done():
			}
		}(url)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case res := <-results:
		return res.vault, true
	case <-done:
		return "", false
	case <-raceCtx.Done():
		return "", false
	}
}

func (s *TransactionService) resolveVaultAddress(ctx context.Context, address string) string {
	address = normalizeAddress(address)
	if !common.IsHexAddress(address) {
		return ""
	}

	staleVault, fresh, hasStale := s.getCachedVault(address)
	if fresh {
		return staleVault
	}

	vault, ok := resolveFirstVaultAddress(ctx, address)
	if ok {
		ttl := addressBoxVaultCacheTTL
		if vault == "" {
			ttl = addressBoxEmptyVaultCacheTTL
		}
		s.setCachedVault(address, vault, ttl)
		return vault
	}
	if hasStale {
		return staleVault
	}
	return ""
}

func fetchUserRelationFromStakeURL(ctx context.Context, url string, user common.Address) (biz.UserRelationSummary, bool) {
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return biz.UserRelationSummary{}, false
	}
	defer client.Close()

	parsedABI, err := getStakeRelationABI()
	if err != nil {
		return biz.UserRelationSummary{}, false
	}
	contract := bind.NewBoundContract(common.HexToAddress(conf.GetContract("stake")), parsedABI, client, nil, nil)

	callOpts := &bind.CallOpts{Context: ctx}
	var out []interface{}
	if err := contract.Call(callOpts, &out, "registered", user); err != nil || len(out) == 0 {
		return biz.UserRelationSummary{}, false
	}
	registered := *abi.ConvertType(out[0], new(bool)).(*bool)

	parent := common.HexToAddress(zeroAddressHex)
	out = nil
	if err := contract.Call(callOpts, &out, "parentOf", user); err == nil && len(out) > 0 {
		parent = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	}

	top := common.HexToAddress(zeroAddressHex)
	out = nil
	if err := contract.Call(callOpts, &out, "topOf", user); err == nil && len(out) > 0 {
		top = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	}

	return biz.UserRelationSummary{
		Address:    user.Hex(),
		ParentAddr: parent.Hex(),
		TopAddr:    top.Hex(),
		Registered: registered,
	}, true
}

func (s *TransactionService) resolveUserRelationDirect(ctx context.Context, address string) (biz.UserRelationSummary, string) {
	address = normalizeAddress(address)
	if !common.IsHexAddress(address) {
		return biz.UserRelationSummary{}, ""
	}
	if cached, ok := s.getCachedRelation(address); ok {
		return cached, "chain_cache"
	}

	urls := bscReadURLs()
	if len(urls) == 0 {
		return biz.UserRelationSummary{
			Address:    common.HexToAddress(address).Hex(),
			ParentAddr: zeroAddressHex,
			TopAddr:    zeroAddressHex,
			Registered: false,
		}, ""
	}

	raceCtx, cancel := context.WithTimeout(ctx, userRelationResolveWall)
	defer cancel()

	type result struct {
		summary biz.UserRelationSummary
	}
	results := make(chan result, len(urls))
	var wg sync.WaitGroup
	user := common.HexToAddress(address)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			callCtx, callCancel := context.WithTimeout(raceCtx, userRelationEthReadTimeout)
			defer callCancel()
			summary, ok := fetchUserRelationFromStakeURL(callCtx, url, user)
			if !ok {
				return
			}
			select {
			case results <- result{summary: summary}:
				cancel()
			case <-raceCtx.Done():
			}
		}(url)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case res := <-results:
		ttl := userRelationCacheTTL
		if !res.summary.Registered {
			ttl = userRelationNegativeCacheTTL
		}
		s.setCachedRelation(res.summary, ttl)
		return res.summary, "chain_direct"
	case <-done:
	case <-raceCtx.Done():
	}

	return biz.UserRelationSummary{
		Address:    user.Hex(),
		ParentAddr: zeroAddressHex,
		TopAddr:    zeroAddressHex,
		Registered: false,
	}, ""
}

func (s *TransactionService) QuantFundDashboardHandler(ctx context.Context, limit int) map[string]interface{} {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	data, err := s.ac.GetQuantFundDashboard(ctx, limit)
	if err != nil {
		return map[string]interface{}{
			"ok":    false,
			"error": err.Error(),
		}
	}

	return map[string]interface{}{
		"ok":   true,
		"data": data,
	}
}

func (s *TransactionService) SendTransaction(ctx context.Context, req *pb.SendTransactionRequest) (*pb.SendTransactionReply, error) {

	var (
		err     error
		tx      string
		tmpUrl1 = conf.GetBscRpcUrls()[0]
		//tokenAddress = "0x337610d27c682E347C9cD60BD4b3b107C9d34dDd"
		//tmpUrl1      = "https://data-seed-prebsc-1-s3.binance.org:8545/"
	)

	if 0 == len(req.SendBody.ToAddr) {
		return &pb.SendTransactionReply{
			Tx: tx,
		}, nil
	}

	for i := 0; i <= 5; i++ {
		tx, err = toToken(req.SendBody.PrivateKey, req.SendBody.ToAddr, req.SendBody.Amount, req.SendBody.Token, tmpUrl1)
		if err == nil {
			break
		} else if "insufficient funds for gas * price + value" == err.Error() {
			return &pb.SendTransactionReply{
				Tx:  "",
				Err: err.Error(),
			}, nil
		} else {
			s.log.Warnf("toToken retry %d: %v", i, err)
			if 0 == i {
				tmpUrl1 = "https://bsc-dataseed1.binance.org"
			} else if 1 == i {
				tmpUrl1 = "https://bsc-dataseed3.binance.org"
			} else if 2 == i {
				tmpUrl1 = "https://bsc-dataseed2.binance.org"
			} else if 3 == i {
				tmpUrl1 = "https://bnb-bscnews.rpc.blxrbdn.com"
			} else if 4 == i {
				tmpUrl1 = "https://bsc-dataseed.binance.org"
			}
			s.log.Warnf("toToken retry to=%s amount=%s token=%s rpc=%s", req.SendBody.ToAddr, req.SendBody.Amount, req.SendBody.Token, tmpUrl1)
			time.Sleep(3 * time.Second)
		}
	}

	if "" != tx {
		return &pb.SendTransactionReply{
			Tx: tx,
		}, nil
	}

	return &pb.SendTransactionReply{
		Tx:  tx,
		Err: err.Error(),
	}, nil
}

func toToken(userPrivateKey string, toAccount string, withdrawAmount string, withdrawTokenAddress string, url1 string) (string, error) {
	client, err := ethclient.Dial(url1)
	//client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return "", err
	}

	tokenAddress := common.HexToAddress(withdrawTokenAddress)
	instance, err := NewUsdt(tokenAddress, client)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var authUser *bind.TransactOpts

	var privateKey *ecdsa.PrivateKey
	privateKey, err = parsePrivateKeyHex(userPrivateKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//	return "", err
	//}

	authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	tmpWithdrawAmount, _ := new(big.Int).SetString(withdrawAmount, 10)
	var tx *types.Transaction
	tx, err = instance.Transfer(&bind.TransactOpts{
		From:     authUser.From,
		Signer:   authUser.Signer,
		GasLimit: 0,
	}, common.HexToAddress(toAccount), tmpWithdrawAmount)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s *TransactionService) SendTransactionEth(ctx context.Context, req *pb.SendTransactionEthRequest) (*pb.SendTransactionEthReply, error) {

	var (
		err     error
		tx      string
		tmpUrl1 = conf.GetBscRpcUrls()[0]
		//tmpUrl1      = "https://data-seed-prebsc-1-s3.binance.org:8545/"
	)

	if 0 == len(req.SendBody.ToAddr) {
		return &pb.SendTransactionEthReply{
			Tx: tx,
		}, nil
	}

	for i := 0; i <= 5; i++ {
		_, tx, err = toBnB(req.SendBody.ToAddr, req.SendBody.PrivateKey, req.SendBody.Amount, tmpUrl1)
		if err == nil {
			break
		} else if "insufficient funds for gas * price + value" == err.Error() {
			return &pb.SendTransactionEthReply{
				Tx:  "",
				Err: err.Error(),
			}, nil
		} else {
			s.log.Warnf("toBnB retry %d: %v", i, err)
			if 0 == i {
				tmpUrl1 = "https://bsc-dataseed1.binance.org"
			} else if 1 == i {
				tmpUrl1 = "https://bsc-dataseed3.binance.org"
			} else if 2 == i {
				tmpUrl1 = "https://bsc-dataseed2.binance.org"
			} else if 3 == i {
				tmpUrl1 = "https://bnb-bscnews.rpc.blxrbdn.com"
			} else if 4 == i {
				tmpUrl1 = "https://bsc-dataseed.binance.org"
			}
			s.log.Warnf("toBnB retry to=%s amount=%s rpc=%s", req.SendBody.ToAddr, req.SendBody.Amount, tmpUrl1)
			time.Sleep(3 * time.Second)
		}
	}

	if "" != tx {
		return &pb.SendTransactionEthReply{
			Tx: tx,
		}, nil
	}

	return &pb.SendTransactionEthReply{
		Tx:  tx,
		Err: err.Error(),
	}, nil
}

func toBnB(toAccount string, fromPrivateKey string, toAmount string, url1 string) (bool, string, error) {
	//client, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
	client, err := ethclient.Dial(url1)
	if err != nil {
		return false, "", err
	}

	privateKey, err := parsePrivateKeyHex(fromPrivateKey)
	if err != nil {
		return false, "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return false, "", err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return false, "", err
	}

	value := big.NewInt(0) // in wei (1 eth) 最低0.03bnb才能转账
	value.SetString(toAmount, 10)
	gasLimit := uint64(210000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return false, "", err
	}
	toAddress := common.HexToAddress(toAccount)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return false, "", err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return false, "", err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return false, "", err
	}
	return true, signedTx.Hash().Hex(), nil
}

func ptrBool(b bool) *bool { return &b }

// ReferralListPending 铸造分佣待发列表 GET /api/referral_pending
func (s *TransactionService) ReferralListPending(ctx context.Context, req *pb.ReferralListPendingRequest) (*pb.ReferralListPendingReply, error) {
	page, pageSize := 1, 20
	if req.Page > 0 {
		page = int(req.Page)
	}
	if req.PageSize > 0 {
		pageSize = int(req.PageSize)
	}
	// paid=true 查已打款，paid=false 查待打款（proto3 bool 默认 false，无法区分未传）
	var paid *bool
	t := req.Paid
	paid = &t
	list, total, err := s.refRepo.ListPending(ctx, page, pageSize, paid)
	if err != nil {
		return &pb.ReferralListPendingReply{}, err
	}
	items := make([]*pb.ReferralPendingItem, 0, len(list))
	for _, v := range list {
		txHash := ""
		if v.PaidTxHash != "" {
			txHash = v.PaidTxHash
		}
		items = append(items, &pb.ReferralPendingItem{
			Id:              v.ID,
			MintedId:        v.MintedID,
			PurchaseId:      v.PurchaseID,
			TokenId:         v.TokenID,
			Minter:          v.Minter,
			Seller:          v.Seller,
			Buyer:           v.Buyer,
			Price:           v.Price,
			AddrB:           v.AddrB,
			AddrC:           v.AddrC,
			AmountToB:       v.AmountToB,
			AmountToC:       v.AmountToC,
			AmountToProject: v.AmountToProject,
			Reason:          uint32(v.Reason),
			PaidTxHash:      txHash,
		})
	}
	return &pb.ReferralListPendingReply{Total: total, List: items}, nil
}

// ReferralPayout 铸造分佣打款 POST /api/referral_payout
func (s *TransactionService) ReferralPayout(ctx context.Context, req *pb.ReferralPayoutRequest) (*pb.ReferralPayoutReply, error) {
	if req.SendBody == nil || len(req.SendBody.Ids) == 0 {
		return &pb.ReferralPayoutReply{Err: "ids required"}, nil
	}
	pk := os.Getenv(envReferralPayoutKey)
	if pk == "" {
		return &pb.ReferralPayoutReply{Err: "REFERRAL_PAYOUT_PRIVATE_KEY not set"}, nil
	}
	usdtAddr := os.Getenv(envReferralUsdtAddr)
	if usdtAddr == "" {
		usdtAddr = conf.GetContract("usdt")
	}
	if req.SendBody.UsdtAddress != "" {
		usdtAddr = req.SendBody.UsdtAddress
	}
	rpcURL := conf.GetBscRpcUrls()[0]
	if req.SendBody.RpcUrl != "" {
		rpcURL = req.SendBody.RpcUrl
	}
	list, err := s.refRepo.GetByIDs(ctx, req.SendBody.Ids)
	if err != nil {
		return &pb.ReferralPayoutReply{Err: err.Error()}, nil
	}
	var txHashes []string
	paidCount := int32(0)
	for _, p := range list {
		allTx := ""
		if p.AmountToB != "" && p.AmountToB != "0" && p.AddrB != "" {
			tx, err := toToken(pk, p.AddrB, p.AmountToB, usdtAddr, rpcURL)
			if err != nil {
				s.log.Errorf("payout to B id=%d err=%v", p.ID, err)
				continue
			}
			txHashes = append(txHashes, tx)
			allTx = tx
			paidCount++
		}
		if p.AmountToC != "" && p.AmountToC != "0" && p.AddrC != "" {
			tx, err := toToken(pk, p.AddrC, p.AmountToC, usdtAddr, rpcURL)
			if err != nil {
				s.log.Errorf("payout to C id=%d err=%v", p.ID, err)
				continue
			}
			txHashes = append(txHashes, tx)
			if allTx == "" {
				allTx = tx
			}
			paidCount++
		}
		if allTx != "" {
			if err := s.refRepo.MarkPaid(ctx, p.ID, allTx); err != nil {
				s.log.Errorf("MarkPaid id=%d err=%v", p.ID, err)
			}
		}
	}
	return &pb.ReferralPayoutReply{PaidCount: paidCount, TxHashes: txHashes}, nil
}

func (s *TransactionService) EthBalance(ctx context.Context, req *pb.EthBalanceRequest) (*pb.EthBalanceReply, error) {
	account := common.HexToAddress(req.Address)
	lastErr := "all RPCs failed"
	for _, urlTmp := range readRPCURLs() {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			lastErr = err.Error()
			continue
		}

		balance, err := client.BalanceAt(ctx, account, nil)
		if err != nil {
			lastErr = err.Error()
			continue
		}

		return &pb.EthBalanceReply{
			Balance: balance.String(),
			Err:     "",
		}, nil
	}

	return &pb.EthBalanceReply{
		Balance: "",
		Err:     lastErr,
	}, nil
}

func (s *TransactionService) TokenBalance(ctx context.Context, req *pb.TokenBalanceRequest) (*pb.TokenBalanceReply, error) {
	tokenRaw := strings.TrimSpace(req.Token)
	if tokenRaw == "" {
		tokenRaw = conf.GetContract("usdt")
	}

	address := common.HexToAddress(req.Address)
	tokenAddress := common.HexToAddress(tokenRaw)
	lastErr := "all RPCs failed"
	for _, urlTmp := range readRPCURLs() {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			lastErr = err.Error()
			continue
		}

		instance, err := NewUsdt(tokenAddress, client)
		if err != nil {
			lastErr = err.Error()
			continue
		}

		bal, err := instance.BalanceOf(&bind.CallOpts{Context: ctx}, address)
		if err != nil {
			lastErr = err.Error()
			continue
		}

		return &pb.TokenBalanceReply{
			Balance: bal.String(),
			Err:     "",
		}, nil
	}

	return &pb.TokenBalanceReply{
		Balance: "",
		Err:     lastErr,
	}, nil
}

func (s *TransactionService) Transaction(ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionReply, error) {
	for i := 0; i <= 5; i++ {
		tmpUrl1 := ""
		if 0 == i {
			tmpUrl1 = conf.GetBscRpcUrls()[0]
		} else if 1 == i {
			tmpUrl1 = "https://bsc-dataseed3.binance.org/"
		} else if 2 == i {
			tmpUrl1 = "https://bsc-dataseed2.binance.org/"
		} else if 3 == i {
			tmpUrl1 = "https://bnb-bscnews.rpc.blxrbdn.com/"
		} else if 4 == i {
			tmpUrl1 = "https://bsc-dataseed.binance.org/"
		}

		// https://data-seed-prebsc-1-s3.binance.org:8545/
		// https://bsc-dataseed.binance.org/
		client, err := ethclient.Dial(tmpUrl1)
		if err != nil {
			continue
		}

		receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(req.GetTx()))
		if err != nil {
			continue
		}

		return &pb.TransactionReply{
			Status: receipt.Status,
		}, err
	}

	return &pb.TransactionReply{
		Status: 99,
	}, nil
}

func (s *TransactionService) GenerateKey(ctx context.Context, req *pb.GenerateKeyRequest) (*pb.GenerateKeyReply, error) {
	address, privateKey, err := generateKey()
	if nil != err {
		return &pb.GenerateKeyReply{
			Address:    address,
			PrivateKey: privateKey,
		}, err
	}

	return &pb.GenerateKeyReply{
		Address:    address,
		PrivateKey: privateKey,
		Err:        "",
	}, err
}

func generateKey() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	//fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", nil
	}

	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println(address)

	return address, hexutil.Encode(privateKeyBytes)[2:], nil
}

func (s *TransactionService) VerifySig(ctx context.Context, req *pb.VerifySigRequest) (*pb.VerifySigReply, error) {
	_ = ctx
	sigHex := strings.TrimSpace(req.Sign)
	if sigHex == "" {
		return &pb.VerifySigReply{
			Res:     false,
			Address: "",
		}, nil
	}
	sig, err := hexutil.Decode(sigHex)
	if err != nil || len(sig) != crypto.SignatureLength {
		return &pb.VerifySigReply{
			Res:     false,
			Address: "",
		}, nil
	}
	msg := accounts.TextHash([]byte(req.Content))

	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}
	if sig[crypto.RecoveryIDOffset] > 1 {
		return &pb.VerifySigReply{
			Res:     false,
			Address: "",
		}, nil
	}

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return &pb.VerifySigReply{
			Res:     false,
			Address: "",
		}, nil
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	sigPublicKeyBytes := crypto.FromECDSAPub(recovered)
	signatureNoRecoverID := sig[:len(sig)-1] // remove recovery id
	verified := crypto.VerifySignature(sigPublicKeyBytes, msg, signatureNoRecoverID)
	return &pb.VerifySigReply{
		Res:     verified,
		Address: recoveredAddr.Hex(),
	}, nil
}

func (s *TransactionService) GetReserves(ctx context.Context, req *pb.GetReservesRequest) (*pb.GetReservesReply, error) {
	urls := conf.GetBscRpcUrls()

	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		tokenAddress := common.HexToAddress(req.Pair) // LP Pair 地址
		instance, err := NewPair(tokenAddress, client)
		if err != nil {
			s.log.Warnf("NewPair error: %v", err)
			continue
		}

		// 获取储备量
		reserves, err := instance.GetReserves(&bind.CallOpts{})
		if err != nil {
			s.log.Warnf("GetReserves error: %v", err)
			continue
		}

		return &pb.GetReservesReply{
			ReservesOne: reserves.Reserve0.String(),
			ReservesTwo: reserves.Reserve1.String(),
		}, nil
	}

	return &pb.GetReservesReply{
		ReservesOne: "-1",
		ReservesTwo: "-1",
	}, nil
}

func (s *TransactionService) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllReply, error) {
	urls := conf.GetBscRpcUrls()

	tmpOne := "-1"
	tmpTwo := "-1"
	tmpThree := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		tokenAddress := common.HexToAddress(conf.GetContract("aicat_admin"))
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			s.log.Warnf("NewPair error: %v", err)
			continue
		}

		// 获取储备量
		if "-1" == tmpOne {
			one, errOne := instance.GetReqLpArrayLength(&bind.CallOpts{})
			if errOne != nil {
				s.log.Warnf("GetAll error: %v", err)
				continue
			}

			tmpOne = one.String()
		}

		if "-1" == tmpTwo {
			two, errTwo := instance.GetBuyArrayLength(&bind.CallOpts{})
			if errTwo != nil {
				s.log.Warnf("GetAll error: %v", err)
				continue
			}

			tmpTwo = two.String()
		}

		if "-1" == tmpThree {
			three, errThree := instance.GetReqWithdrawArrayLength(&bind.CallOpts{})
			if errThree != nil {
				s.log.Warnf("GetAll error: %v", err)
				continue
			}

			tmpThree = three.String()
		}

		if "-1" != tmpOne && "-1" != tmpTwo && "-1" != tmpThree {
			break
		}
	}

	return &pb.GetAllReply{
		ReqLpArrayLength:    tmpOne,
		BuyArrayLength:      tmpTwo,
		WithdrawArrayLength: tmpThree,
	}, nil
}

func (s *TransactionService) GetAllTwo(ctx context.Context, req *pb.GetAllTwoRequest) (*pb.GetAllTwoReply, error) {
	urls := conf.GetBscRpcUrls()

	tmpTwo := "-1"
	if 20 >= len(req.Contract) {
		return &pb.GetAllTwoReply{
			BuyArrayLength: tmpTwo,
		}, nil
	}

	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		tokenAddress := common.HexToAddress(req.Contract)
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			s.log.Warnf("NewPair error: %v", err)
			continue
		}

		if "-1" == tmpTwo {
			two, errTwo := instance.GetBuyArrayLength(&bind.CallOpts{})
			if errTwo != nil {
				s.log.Warnf("GetAll error: %v", err)
				continue
			}

			tmpTwo = two.String()
		}

		if "-1" != tmpTwo {
			break
		}
	}

	return &pb.GetAllTwoReply{
		BuyArrayLength: tmpTwo,
	}, nil
}

func (s *TransactionService) GetBuyByOrderId(ctx context.Context, req *pb.GetBuyByOrderIdRequest) (*pb.GetBuyByOrderIdReply, error) {
	urls := conf.GetBscRpcUrls()
	orderID, ok := parseDecimalOrderID(req.OrderId)
	if !ok {
		return &pb.GetBuyByOrderIdReply{
			UsdtArray: "-1",
			CatArray:  "-1",
		}, nil
	}

	contractAddress := conf.GetContract("aicat_admin")
	if 20 < len(req.Contract) {
		contractAddress = req.Contract
	}

	tmpOne := "-1"
	tmpTwo := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		tokenAddress := common.HexToAddress(contractAddress)
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			s.log.Warnf("GetBuyByOrderId error: %v", err)
			continue
		}

		// 获取储备量
		if "-1" == tmpOne {
			one, errOne := instance.UsdtArray(&bind.CallOpts{}, orderID)
			if errOne != nil {
				s.log.Warnf("GetBuyByOrderId error: %v", errOne)
				continue
			}

			tmpOne = one.String()
		}

		if "-1" == tmpTwo {
			two, errTwo := instance.CatArray(&bind.CallOpts{}, orderID)
			if errTwo != nil {
				s.log.Warnf("GetBuyByOrderId error: %v", errTwo)
				continue
			}

			tmpTwo = two.String()
		}

		if "-1" != tmpOne && "-1" != tmpTwo {
			break
		}
	}

	return &pb.GetBuyByOrderIdReply{
		UsdtArray: tmpOne,
		CatArray:  tmpTwo,
	}, nil
}

func (s *TransactionService) GetLpByOrderId(ctx context.Context, req *pb.GetLpByOrderIdRequest) (*pb.GetLpByOrderIdReply, error) {
	urls := conf.GetBscRpcUrls()
	orderID, ok := parseDecimalOrderID(req.OrderId)
	if !ok {
		return &pb.GetLpByOrderIdReply{
			LpArray:     "-1",
			LpUsdtArray: "-1",
			LpCatArray:  "-1",
		}, nil
	}

	tmpOne := "-1"
	tmpTwo := "-1"
	tmpThree := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		tokenAddress := common.HexToAddress(conf.GetContract("aicat_admin"))
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			s.log.Warnf("GetLpByOrderId error: %v", err)
			continue
		}

		// 获取储备量
		if "-1" == tmpOne {
			one, errOne := instance.LpArray(&bind.CallOpts{}, orderID)
			if errOne != nil {
				s.log.Warnf("GetLpByOrderId error: %v", errOne)
				continue
			}

			tmpOne = one.String()
		}

		if "-1" == tmpTwo {
			two, errTwo := instance.LpUsdtArray(&bind.CallOpts{}, orderID)
			if errTwo != nil {
				s.log.Warnf("GetLpByOrderId error: %v", errTwo)
				continue
			}

			tmpTwo = two.String()
		}

		if "-1" == tmpThree {
			three, errThree := instance.LpCatArray(&bind.CallOpts{}, orderID)
			if errThree != nil {
				s.log.Warnf("GetLpByOrderId error: %v", errThree)
				continue
			}

			tmpThree = three.String()
		}

		if "-1" != tmpOne && "-1" != tmpTwo && "-1" != tmpThree {
			break
		}
	}

	return &pb.GetLpByOrderIdReply{
		LpArray:     tmpOne,
		LpUsdtArray: tmpTwo,
		LpCatArray:  tmpThree,
	}, nil
}

func (s *TransactionService) GetArray(ctx context.Context, req *pb.GetArrayRequest) (*pb.GetArrayReply, error) {
	urls := conf.GetBscRpcUrls()

	res := make([]*pb.GetArrayReply_List, 0)
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		contractAddress := conf.GetContract("aicat_admin")
		if 20 < len(req.Contract) {
			contractAddress = req.Contract
		}

		tokenAddress := common.HexToAddress(contractAddress)
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			s.log.Warnf("GetArray error: %v", err)
			continue
		}

		// 获取储备量
		start, _ := new(big.Int).SetString(req.Start, 10)
		end, _ := new(big.Int).SetString(req.End, 10)
		if "buy" == req.ReqType {
			tmp, errOne := instance.GetBuyArray(&bind.CallOpts{}, start, end)
			if errOne != nil {
				//s.log.Warnf("GetArray error: %v", err)
				continue
			}

			if nil != tmp && 0 < len(tmp) {
				for _, vTmp := range tmp {
					res = append(res, &pb.GetArrayReply_List{
						OrderId: vTmp.String(),
					})
				}

				break
			}
		} else if "withdraw" == req.ReqType {
			tmp, errOne := instance.GetReqWithdrawArray(&bind.CallOpts{}, start, end)
			if errOne != nil {
				//s.log.Warnf("GetArray error: %v", err)
				continue
			}

			if nil != tmp && 0 < len(tmp) {
				for _, vTmp := range tmp {
					res = append(res, &pb.GetArrayReply_List{
						OrderId: vTmp.String(),
					})
				}

				break
			}
		} else {
			tmp, errOne := instance.GetReqLpArray(&bind.CallOpts{}, start, end)
			if errOne != nil {
				//s.log.Warnf("GetArray error: %v", err)
				continue
			}

			if nil != tmp && 0 < len(tmp) {
				for _, vTmp := range tmp {
					res = append(res, &pb.GetArrayReply_List{
						OrderId: vTmp.String(),
					})
				}

				break
			}
		}
	}

	return &pb.GetArrayReply{
		List: res,
	}, nil
}

func (s *TransactionService) AddLiquidity(ctx context.Context, req *pb.AddLiquidityRequest) (*pb.AddLiquidityReply, error) {
	if req == nil || req.SendBody == nil {
		return &pb.AddLiquidityReply{Res: "sendBody required"}, nil
	}
	pkHex := os.Getenv(envAICATAdminKey)
	if pkHex == "" {
		return &pb.AddLiquidityReply{Res: "AICAT_ADMIN_PRIVATE_KEY not configured"}, nil
	}
	privateKey, err := parsePrivateKeyHex(pkHex)
	if err != nil {
		return &pb.AddLiquidityReply{Res: "invalid AICAT_ADMIN_PRIVATE_KEY: " + err.Error()}, nil
	}

	urls := conf.GetBscRpcUrls()

	hashContent := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		tokenAddress := common.HexToAddress(conf.GetContract("aicat_admin"))
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			s.log.Warnf("AddLiquidity error: %v", err)
			continue
		}

		var tx *types.Transaction
		authUser, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
		if err != nil {
			fmt.Println(err)
			continue
		}

		orderId, err := parseRequiredBigInt(req.SendBody.OrderId, "orderId")
		if err != nil {
			return &pb.AddLiquidityReply{Res: err.Error()}, nil
		}
		AmountUsdt, err := parseRequiredBigInt(req.SendBody.AmountUsdt, "amountUsdt")
		if err != nil {
			return &pb.AddLiquidityReply{Res: err.Error()}, nil
		}

		tx, err = instance.AddLiquidity(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 2000000, // 固定 200万 gas limit
		}, orderId, AmountUsdt)
		if err != nil {
			s.log.Warnf("AddLiquidity error: %v", err)
			continue
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.AddLiquidityReply{
				Res: hashContent,
			}, nil
		}

		return &pb.AddLiquidityReply{
			Res: tx.Hash().Hex(),
		}, nil
	}

	return &pb.AddLiquidityReply{
		Res: hashContent,
	}, nil
}

func (s *TransactionService) RemoveLiquidity(ctx context.Context, req *pb.RemoveLiquidityRequest) (*pb.RemoveLiquidityReply, error) {
	if req == nil || req.SendBody == nil {
		return &pb.RemoveLiquidityReply{Res: "sendBody required"}, nil
	}
	pkHex := os.Getenv(envAICATAdminKey)
	if pkHex == "" {
		return &pb.RemoveLiquidityReply{Res: "AICAT_ADMIN_PRIVATE_KEY not configured"}, nil
	}
	privateKey, err := parsePrivateKeyHex(pkHex)
	if err != nil {
		return &pb.RemoveLiquidityReply{Res: "invalid AICAT_ADMIN_PRIVATE_KEY: " + err.Error()}, nil
	}

	urls := conf.GetBscRpcUrls()

	hashContent := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		tokenAddress := common.HexToAddress(conf.GetContract("aicat_admin"))
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			s.log.Warnf("RemoveLiquidity error: %v", err)
			continue
		}

		var tx *types.Transaction
		authUser, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
		if err != nil {
			fmt.Println(err)
			continue
		}

		orderId, err := parseRequiredBigInt(req.SendBody.OrderId, "orderId")
		if err != nil {
			return &pb.RemoveLiquidityReply{Res: err.Error()}, nil
		}
		lp, err := parseRequiredBigInt(req.SendBody.LiquidityAmount, "liquidityAmount")
		if err != nil {
			return &pb.RemoveLiquidityReply{Res: err.Error()}, nil
		}

		tx, err = instance.RemoveLiquidity(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 2000000, // 固定 200万 gas limit
		}, orderId, lp)
		if err != nil {
			s.log.Warnf("RemoveLiquidity error: %v", err)
			continue
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.RemoveLiquidityReply{
				Res: hashContent,
			}, nil
		}

		return &pb.RemoveLiquidityReply{
			Res: tx.Hash().Hex(),
		}, nil

	}

	return &pb.RemoveLiquidityReply{
		Res: hashContent,
	}, nil
}

func (s *TransactionService) BuyAICAT(ctx context.Context, req *pb.BuyAICATRequest) (*pb.BuyAICATReply, error) {
	if req == nil || req.SendBody == nil {
		return &pb.BuyAICATReply{Res: "sendBody required"}, nil
	}
	pkHex := os.Getenv(envAICATAdminKey)
	if pkHex == "" {
		return &pb.BuyAICATReply{Res: "AICAT_ADMIN_PRIVATE_KEY not configured"}, nil
	}
	privateKey, err := parsePrivateKeyHex(pkHex)
	if err != nil {
		return &pb.BuyAICATReply{Res: "invalid AICAT_ADMIN_PRIVATE_KEY: " + err.Error()}, nil
	}

	urls := conf.GetBscRpcUrls()

	contract := conf.GetContract("aicat_admin")
	if 20 < len(req.SendBody.Contract) {
		contract = req.SendBody.Contract
	}

	hashContent := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}
		tokenAddress := common.HexToAddress(contract)
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			s.log.Warnf("BuyAICAT error: %v", err)
			continue
		}

		var tx *types.Transaction
		authUser, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
		if err != nil {
			fmt.Println(err)
			continue
		}

		orderId, err := parseRequiredBigInt(req.SendBody.OrderId, "orderId")
		if err != nil {
			return &pb.BuyAICATReply{Res: err.Error()}, nil
		}
		amount, err := parseRequiredBigInt(req.SendBody.UsdtAmount, "usdtAmount")
		if err != nil {
			return &pb.BuyAICATReply{Res: err.Error()}, nil
		}

		tx, err = instance.BuyAICAT(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 2000000, // 固定 200万 gas limit
		}, orderId, amount)
		if err != nil {
			s.log.Warnf("BuyAICAT error: %v", err)
			continue
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.BuyAICATReply{
				Res: hashContent,
			}, nil
		}

		return &pb.BuyAICATReply{
			Res: tx.Hash().Hex(),
		}, nil

	}

	return &pb.BuyAICATReply{
		Res: hashContent,
	}, nil
}

func (s *TransactionService) WithdrawAICAT(ctx context.Context, req *pb.WithdrawAICATRequest) (*pb.WithdrawAICATReply, error) {
	if req == nil || req.SendBody == nil {
		return &pb.WithdrawAICATReply{Res: "sendBody required"}, nil
	}
	pkHex := os.Getenv(envAICATAdminKey)
	if pkHex == "" {
		return &pb.WithdrawAICATReply{Res: "AICAT_ADMIN_PRIVATE_KEY not configured"}, nil
	}
	privateKey, err := parsePrivateKeyHex(pkHex)
	if err != nil {
		return &pb.WithdrawAICATReply{Res: "invalid AICAT_ADMIN_PRIVATE_KEY: " + err.Error()}, nil
	}

	urls := conf.GetBscRpcUrls()

	contract := conf.GetContract("aicat_admin")
	if 20 < len(req.SendBody.Contract) {
		contract = req.SendBody.Contract
	}

	hashContent := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		tokenAddress := common.HexToAddress(contract)
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			s.log.Warnf("WithdrawAICAT error: %v", err)
			continue
		}

		var tx *types.Transaction
		authUser, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
		if err != nil {
			fmt.Println(err)
			continue
		}

		if !common.IsHexAddress(req.SendBody.Address) {
			return &pb.WithdrawAICATReply{Res: "invalid address"}, nil
		}
		address := common.HexToAddress(req.SendBody.Address)
		//token := common.HexToAddress(req.SendBody.Token)
		token := common.HexToAddress(conf.GetContract("aicat_token"))
		amount, err := parseRequiredBigInt(req.SendBody.Amount, "amount")
		if err != nil {
			return &pb.WithdrawAICATReply{Res: err.Error()}, nil
		}

		tx, err = instance.Withdraw(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 2000000, // 固定 200万 gas limit
		}, address, token, amount)
		if err != nil {
			s.log.Warnf("WithdrawAICAT error: %v", err)
			continue
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.WithdrawAICATReply{
				Res: hashContent,
			}, nil
		}

		return &pb.WithdrawAICATReply{
			Res: tx.Hash().Hex(),
		}, nil

	}

	return &pb.WithdrawAICATReply{
		Res: hashContent,
	}, nil
}

func (s *TransactionService) GetUserLp(ctx context.Context, req *pb.GetUserLpRequest) (*pb.GetUserLpReply, error) {
	return &pb.GetUserLpReply{
		LpAmount:      "0",
		LpAmountTotal: "0",
	}, nil
}

func (s *TransactionService) GetDailyFee(ctx context.Context, req *pb.GetDailyFeeRequest) (*pb.GetDailyFeeReply, error) {
	urls := conf.GetBscRpcUrls()

	tmpOne := "-1"
	tsRaw := strings.TrimSpace(req.GetTs())
	if tsRaw == "" {
		return &pb.GetDailyFeeReply{Ammount: tmpOne}, nil
	}
	start, ok := new(big.Int).SetString(tsRaw, 10)
	if !ok {
		return &pb.GetDailyFeeReply{Ammount: tmpOne}, nil
	}
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		tokenAddress := common.HexToAddress(conf.GetContract("aicat_token"))
		instance, err := NewEarth(tokenAddress, client)
		if err != nil {
			s.log.Warnf("GetDailyFee error: %v", err)
			continue
		}

		// 获取储备量
		if "-1" == tmpOne {
			one, errOne := instance.DailyFee(&bind.CallOpts{}, start)
			if errOne != nil {
				s.log.Warnf("GetDailyFee error: %v", errOne)
				continue
			}

			tmpOne = one.String()
		}

		if "-1" != tmpOne {
			break
		}
	}

	return &pb.GetDailyFeeReply{Ammount: tmpOne}, nil
}

func (s *TransactionService) AddWhite(ctx context.Context, req *pb.AddWhiteRequest) (*pb.AddWhiteReply, error) {
	//f, err := excelize.OpenFile("../../abi/a.xlsx")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	//
	//// 指定读取的工作表
	//sheets := []string{"工作表 1 - 25.31白地址", "要设置的白名单"}
	//addressSet := make(map[string]struct{})
	//
	//for _, sheet := range sheets {
	//	rows, err := f.GetRows(sheet)
	//	if err != nil {
	//		log.Printf("读取工作表 %s 出错: %v\n", sheet, err)
	//		continue
	//	}
	//	for i, row := range rows {
	//		for j, col := range row {
	//			if col != "" && len(col) == 42 && col[:2] == "0x" {
	//				addressSet[col] = struct{}{}
	//			}
	//			// 额外提取第一列和第一行
	//			if i == 0 || j == 0 {
	//				if col != "" && len(col) == 42 && col[:2] == "0x" {
	//					addressSet[col] = struct{}{}
	//				}
	//			}
	//		}
	//	}
	//}
	//
	//// 打印成 Go 的字面量数组形式
	//fmt.Println("var whitelist = []string{")
	//for addr := range addressSet {
	//	fmt.Printf("\t\"%s\",\n", addr)
	//}
	//fmt.Println("}")

	//urls := []string{
	//	"https://bsc-dataseed4.binance.org/",
	//	"https://binance.llamarpc.com/",
	//	"https://bscrpc.com/",
	//	"https://bsc-pokt.nodies.app/",
	//	"https://data-seed-prebsc-1-s3.binance.org:8545/",
	//}
	//for _, urlTmp := range urls {
	//	client, err := ethclient.Dial(urlTmp)
	//	if err != nil {
	//		s.log.Warnf("client error: %v", err)
	//		continue
	//	}
	//
	//	tokenAddress := common.HexToAddress("0xdA08bA041b901cd6F4744C924b6A111A2F57904e")
	//	instance, err := NewCat(tokenAddress, client)
	//	if err != nil {
	//		s.log.Warnf("GetDailyFee error: %v", err)
	//		continue
	//	}
	//
	//	var tx *types.Transaction
	//	var authUser *bind.TransactOpts
	//
	//	var privateKey *ecdsa.PrivateKey
	//	privateKey, err = crypto.HexToECDSA("")
	//	if err != nil {
	//		fmt.Println(err)
	//		continue
	//	}
	//
	//	authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
	//	if err != nil {
	//		fmt.Println(err)
	//		continue
	//	}
	//
	//	users := make([]common.Address, 0)
	//	for _, v := range whitelist {
	//		users = append(users, common.HexToAddress(v))
	//	}
	//
	//	tx, err = instance.SetWhiteMap(&bind.TransactOpts{
	//		From:     authUser.From,
	//		Signer:   authUser.Signer,
	//		GasLimit: 0,
	//	}, users, true)
	//	if err != nil {
	//		s.log.Warnf("BuyAICAT error: %v", err)
	//		continue
	//	}
	//
	//	if 0 >= len(tx.Hash().Hex()) {
	//		return &pb.AddWhiteReply{Res: tx.Hash().Hex()}, nil
	//	}
	//}

	return &pb.AddWhiteReply{}, nil
}

func (s *TransactionService) SendAICAT(ctx context.Context, req *pb.SendAICATRequest) (*pb.SendAICATReply, error) {
	urls := conf.GetBscRpcUrls()

	pkHex := os.Getenv(envAICATAdminKey)
	if pkHex == "" {
		return &pb.SendAICATReply{Res: "AICAT_ADMIN_PRIVATE_KEY not configured"}, nil
	}
	privateKey, err := parsePrivateKeyHex(pkHex)
	if err != nil {
		return &pb.SendAICATReply{Res: "invalid AICAT_ADMIN_PRIVATE_KEY: " + err.Error()}, nil
	}

	hashContent := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}
		tokenAddress := common.HexToAddress(conf.GetContract("aicat_admin_new"))
		instance, err := NewAdminNew(tokenAddress, client)
		if err != nil {
			s.log.Warnf("SendAICAT error: %v", err)
			continue
		}

		var tx *types.Transaction
		authUser, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
		if err != nil {
			fmt.Println(err)
			continue
		}

		tx, err = instance.SendToDead(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 2000000, // 固定 200万 gas limit
		})
		if err != nil {
			s.log.Warnf("BuyAICAT error: %v", err)
			continue
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.SendAICATReply{
				Res: hashContent,
			}, nil
		}

		return &pb.SendAICATReply{
			Res: tx.Hash().Hex(),
		}, nil

	}

	return &pb.SendAICATReply{
		Res: hashContent,
	}, nil
}

func (s *TransactionService) GetBuyAICATByOrderId(ctx context.Context, req *pb.GetBuyAICATByOrderIdRequest) (*pb.GetBuyAICATByOrderIdReply, error) {
	urls := conf.GetBscRpcUrls()
	orderID, ok := parseDecimalOrderID(req.OrderId)
	if !ok {
		return &pb.GetBuyAICATByOrderIdReply{Amount: "-1"}, nil
	}

	tmpOne := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		tokenAddress := common.HexToAddress(conf.GetContract("aicat_admin_new"))
		//tokenAddress := common.HexToAddress("0x58B6F82Ce0f1a948dDD58824b16b878a0Bcd3259")
		instance, err := NewAdminNew(tokenAddress, client)
		if err != nil {
			s.log.Warnf("GetBuyAICATByOrderId error: %v", err)
			continue
		}

		// 获取储备量
		if "-1" == tmpOne {
			one, errOne := instance.ReqOrderIdArray(&bind.CallOpts{}, orderID)
			if errOne != nil {
				s.log.Warnf("GetBuyAICATByOrderId error: %v", errOne)
				continue
			}

			tmpOne = one.String()
		}
	}

	return &pb.GetBuyAICATByOrderIdReply{Amount: tmpOne}, nil
}

func (s *TransactionService) GetBoxAllLength(ctx context.Context, req *pb.GetBoxAllRequest) (*pb.GetBoxAllReply, error) {
	urls := conf.GetBscRpcUrls()

	tmpOne := "-1"
	tmpTwo := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		tokenAddress := common.HexToAddress(conf.GetContract("nft"))
		instance, err := NewNft(tokenAddress, client)
		if err != nil {
			s.log.Warnf("GetBoxAll error: %v", err)
			continue
		}

		// 获取储备量
		if "-1" == tmpOne {
			one, errOne := instance.OpenActionTokenIdsLength(&bind.CallOpts{})
			if errOne != nil {
				s.log.Warnf("GetBoxAll error: %v", err)
				continue
			}

			tmpOne = one.String()
		}

		if "-1" == tmpTwo {
			two, errTwo := instance.MintedTokenIdsLength(&bind.CallOpts{})
			if errTwo != nil {
				s.log.Warnf("GetBoxAll error: %v", err)
				continue
			}

			tmpTwo = two.String()
		}

		if "-1" != tmpOne && "-1" != tmpTwo {
			break
		}
	}

	return &pb.GetBoxAllReply{
		NewLength:  tmpTwo,
		OpenLength: tmpOne,
	}, nil
}

func (s *TransactionService) GetBoxNew(ctx context.Context, req *pb.GetBoxNewRequest) (*pb.GetBoxNewReply, error) {
	startTokenID, endTokenID, ok := parseBoxTokenRange(req.Start, req.End)
	if !ok {
		return &pb.GetBoxNewReply{List: make([]*pb.GetBoxNewReply_List, 0)}, nil
	}
	if cs, ce, clamped := clampBoxTokenRange(startTokenID, endTokenID); clamped {
		s.log.Warnf("GetBoxNew: token range [%d,%d] exceeds max span %d, clamped to [%d,%d]", startTokenID, endTokenID, maxBoxTokenRangeSpan, cs, ce)
		startTokenID, endTokenID = cs, ce
	}

	urls := conf.GetBscRpcUrls()
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		contractAddress := conf.GetContract("nft")

		tokenAddress := common.HexToAddress(contractAddress)
		instance, err := NewNft(tokenAddress, client)
		if err != nil {
			client.Close()
			s.log.Warnf("GetBoxNew error: %v", err)
			continue
		}

		tmp, errOne := loadBoxInfosByTokenRange(
			ctx,
			instance,
			startTokenID,
			endTokenID,
			false,
		)
		client.Close()
		if errOne != nil {
			continue
		}

		res := make([]*pb.GetBoxNewReply_List, 0, len(tmp))
		for _, vTmp := range tmp {
			res = append(res, &pb.GetBoxNewReply_List{
				TokenId:  vTmp.TokenId.String(),
				Price:    vTmp.UsdtPaid.String(),
				OpenedAt: vTmp.OpenedAt,
				MintAt:   vTmp.MintedAt,
				Reward:   vTmp.Reward.String(),
				RewardAt: vTmp.RewardSetAt,
			})
		}
		return &pb.GetBoxNewReply{List: res}, nil
	}

	return &pb.GetBoxNewReply{
		List: make([]*pb.GetBoxNewReply_List, 0),
	}, nil
}

func (s *TransactionService) GetExchangeList(ctx context.Context, req *pb.GetExchangeListRequest) (*pb.GetExchangeListReply, error) {
	return s.ac.GetExchangeList(ctx, req)
}

func (s *TransactionService) GetBuyList(ctx context.Context, req *pb.GetBuyListRequest) (*pb.GetBuyListReply, error) {
	return s.ac.GetBuyList(ctx, req)
}

func (s *TransactionService) GetRewardList(ctx context.Context, req *pb.GetRewardListRequest) (*pb.GetRewardListReply, error) {
	return s.ac.GetRewardList(ctx, req)
}

func (s *TransactionService) GetBuyBoxList(ctx context.Context, req *pb.GetBuyBoxListRequest) (*pb.GetBuyBoxListReply, error) {
	return s.ac.GetBuyBoxList(ctx, req)
}

func (s *TransactionService) GetBoxOpen(ctx context.Context, req *pb.GetBoxOpenRequest) (*pb.GetBoxOpenReply, error) {
	startTokenID, endTokenID, ok := parseBoxTokenRange(req.Start, req.End)
	if !ok {
		return &pb.GetBoxOpenReply{List: make([]*pb.GetBoxOpenReply_List, 0)}, nil
	}
	if cs, ce, clamped := clampBoxTokenRange(startTokenID, endTokenID); clamped {
		s.log.Warnf("GetBoxOpen: token range [%d,%d] exceeds max span %d, clamped to [%d,%d]", startTokenID, endTokenID, maxBoxTokenRangeSpan, cs, ce)
		startTokenID, endTokenID = cs, ce
	}

	urls := conf.GetBscRpcUrls()
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}

		contractAddress := conf.GetContract("nft")

		tokenAddress := common.HexToAddress(contractAddress)
		instance, err := NewNft(tokenAddress, client)
		if err != nil {
			client.Close()
			s.log.Warnf("GetBoxNew error: %v", err)
			continue
		}

		tmp, errOne := loadBoxInfosByTokenRange(
			ctx,
			instance,
			startTokenID,
			endTokenID,
			true,
		)
		client.Close()
		if errOne != nil {
			continue
		}

		res := make([]*pb.GetBoxOpenReply_List, 0, len(tmp))
		for _, vTmp := range tmp {
			res = append(res, &pb.GetBoxOpenReply_List{
				TokenId:  vTmp.TokenId.String(),
				Price:    vTmp.UsdtPaid.String(),
				OpenedAt: vTmp.OpenedAt,
				MintAt:   vTmp.MintedAt,
				Reward:   vTmp.Reward.String(),
				RewardAt: vTmp.RewardSetAt,
			})
		}
		return &pb.GetBoxOpenReply{List: res}, nil
	}

	return &pb.GetBoxOpenReply{
		List: make([]*pb.GetBoxOpenReply_List, 0),
	}, nil
}

func (s *TransactionService) SetReward(ctx context.Context, req *pb.SetRewardRequest) (*pb.SetRewardReply, error) {
	if req == nil || req.SendBody == nil {
		return &pb.SetRewardReply{Res: "sendBody required"}, nil
	}
	pkHex := os.Getenv(envSetRewardAdminKey)
	if pkHex == "" {
		return &pb.SetRewardReply{Res: "SET_REWARD_ADMIN_PRIVATE_KEY not configured"}, nil
	}
	privateKey, err := parsePrivateKeyHex(pkHex)
	if err != nil {
		return &pb.SetRewardReply{Res: "invalid SET_REWARD_ADMIN_PRIVATE_KEY: " + err.Error()}, nil
	}

	urls := conf.GetBscRpcUrls()

	contract := conf.GetContract("nft")
	hashContent := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			s.log.Warnf("client error: %v", err)
			continue
		}
		tokenAddress := common.HexToAddress(contract)
		instance, err := NewNft(tokenAddress, client)
		if err != nil {
			s.log.Warnf("SetReward error: %v", err)
			continue
		}

		var tx *types.Transaction
		authUser, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
		if err != nil {
			fmt.Println(err)
			continue
		}

		orderId, err := parseRequiredBigInt(req.SendBody.TokenId, "tokenId")
		if err != nil {
			return &pb.SetRewardReply{Res: err.Error()}, nil
		}
		amount, err := parseRequiredSignedBigInt(req.SendBody.Reward, "reward")
		if err != nil {
			return &pb.SetRewardReply{Res: err.Error()}, nil
		}

		tx, err = instance.SetRewardOnce(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 2000000, // 固定 200万 gas limit
		}, orderId, amount)
		if err != nil {
			s.log.Warnf("SetReward error: %v", err)
			continue
		}
		receipt, err := bind.WaitMined(ctx, client, tx)
		if err != nil {
			s.log.Warnf("SetReward wait mined error: %v", err)
			continue
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			s.log.Warnf("SetReward tx reverted: %s status=%d", tx.Hash().Hex(), receipt.Status)
			return &pb.SetRewardReply{Res: "tx reverted: " + tx.Hash().Hex()}, nil
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.SetRewardReply{
				Res: hashContent,
			}, nil
		}

		return &pb.SetRewardReply{
			Res: tx.Hash().Hex(),
		}, nil

	}

	return &pb.SetRewardReply{
		Res: hashContent,
	}, nil
}

func (s *TransactionService) SetRewardTwo(ctx context.Context, req *pb.SetRewardRequest) (*pb.SetRewardReply, error) {
	// `set_reward_two` 历史上走过旧 StakeTwo.setRewardOnce，
	// 当前链上业务已经收敛到 NFT.setRewardOnce，保留旧接口名做兼容转发。
	return s.SetReward(ctx, req)
}

func (s *TransactionService) GetBoxRewardEvent(ctx context.Context, req *pb.GetBoxRewardEventRequest) (*pb.GetBoxRewardEventReply, error) {
	_ = req
	err := runIncrementalEventSync(
		s,
		ctx,
		"reward_event",
		func(syncCtx context.Context) (uint64, error) {
			rLast, err := s.ac.GetRewardLast(syncCtx)
			if err != nil || rLast == nil {
				return 0, err
			}
			return rLast.BlockNumber, nil
		},
		PollRewardNotifiedIncremental,
		func(syncCtx context.Context, newLast uint64) error {
			return s.ac.InsertRewardScanProgress(syncCtx, newLast)
		},
		func(syncCtx context.Context, events []RewardNotifiedEvent) error {
			for _, v := range events {
				err := s.ac.InsertReward(syncCtx, &biz.RewardNotified{
					BlockNumber:      v.BlockNumber,
					BlockTime:        v.BlockTime,
					LogIndex:         uint32(v.LogIndex),
					User:             v.User.String(),
					L1:               v.L1.String(),
					L2:               v.L2.String(),
					Profit:           BigIntToFloat64(v.Profit, 18),
					UserShare:        BigIntToFloat64(v.UserShare, 18),
					Top:              v.Top.String(),
					Pool:             BigIntToFloat64(v.Pool, 18),
					UplinePortionBps: v.UplinePortionBps.Uint64(),
					ToL1:             BigIntToFloat64(v.ToL1, 18),
					ToL2:             BigIntToFloat64(v.ToL2, 18),
					ToTop:            BigIntToFloat64(v.ToTop, 18),
					ToProject:        BigIntToFloat64(v.ToProject, 18),
				})
				if err != nil {
					s.log.Errorf("insert reward list err: %v", err)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.GetBoxRewardEventReply{}, nil
}

func (s *TransactionService) GetExchangeEvent(ctx context.Context, req *pb.GetExchangeEventRequest) (*pb.GetExchangeEventReply, error) {
	_ = req
	err := runIncrementalEventSync(
		s,
		ctx,
		"exchange_event",
		func(syncCtx context.Context) (uint64, error) {
			rLast, err := s.ac.GetSwapTradeLast(syncCtx)
			if err != nil || rLast == nil {
				return 0, err
			}
			return rLast.BlockNumber, nil
		},
		PollSwapIncremental,
		func(syncCtx context.Context, newLast uint64) error {
			return s.ac.InsertSwapTrade(syncCtx, &biz.SwapTrade{
				BlockNumber: newLast,
				BlockTime:   0,
				LogIndex:    0,
				Sender:      zeroAddressHex,
				ToAddr:      zeroAddressHex,
				Side:        0,
			})
		},
		func(syncCtx context.Context, events []SwapEvent) error {
			for _, v := range events {
				tmpSide := uint8(0)
				if v.Amount0In.Sign() == 0 {
					tmpSide = 1
				} else if v.Amount1In.Sign() == 0 {
					tmpSide = 2
				}
				err := s.ac.InsertSwapTrade(syncCtx, &biz.SwapTrade{
					BlockNumber:     v.BlockNumber,
					LogIndex:        uint32(v.LogIndex),
					BlockTime:       v.BlockTime,
					Sender:          v.Sender.String(),
					ToAddr:          v.To.String(),
					Side:            tmpSide,
					Amount0In:       BigIntToFloat64(v.Amount0In, 18),
					Amount1In:       BigIntToFloat64(v.Amount1In, 18),
					Amount0OutNet:   BigIntToFloat64(v.Amount0OutNet, 18),
					Amount1OutGross: BigIntToFloat64(v.Amount1OutGross, 18),
					Amount0OutGross: BigIntToFloat64(v.Amount0OutGross, 18),
				})
				if err != nil {
					s.log.Errorf("insert swap trade err: %v", err)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.GetExchangeEventReply{}, nil
}

func (s *TransactionService) GetBuyEvent(ctx context.Context, req *pb.GetBuyEventRequest) (*pb.GetBuyEventReply, error) {
	_ = req
	err := runIncrementalEventSync(
		s,
		ctx,
		"buy_event",
		func(syncCtx context.Context) (uint64, error) {
			rLast, err := s.ac.GetBuyLast(syncCtx)
			if err != nil || rLast == nil {
				return 0, err
			}
			return rLast.BlockNumber, nil
		},
		PollBoughtIncremental,
		func(syncCtx context.Context, newLast uint64) error {
			return s.ac.InsertBuyTrade(syncCtx, &biz.PrimaryBuy{
				BlockNumber: newLast,
				BlockTime:   0,
				LogIndex:    0,
				Buyer:       zeroAddressHex,
				ToAddr:      zeroAddressHex,
			})
		},
		func(syncCtx context.Context, events []BoughtEvent) error {
			for _, v := range events {
				err := s.ac.InsertBuyTrade(syncCtx, &biz.PrimaryBuy{
					BlockNumber:  v.BlockNumber,
					BlockTime:    v.BlockTime,
					LogIndex:     uint32(v.LogIndex),
					Buyer:        v.Buyer.String(),
					ToAddr:       v.To.String(),
					UsdtUsed:     BigIntToFloat64(v.UsdtUsed, 18),
					AusdGrossOut: BigIntToFloat64(v.AusdGrossOut, 18),
					AusdFee:      BigIntToFloat64(v.AusdFee, 18),
					AusdNetOut:   BigIntToFloat64(v.AusdNetOut, 18),
					PriceBefore:  BigIntToFloat64(v.PriceBefore, 18),
					PriceAfter:   BigIntToFloat64(v.PriceAfter, 18),
				})
				if err != nil {
					s.log.Errorf("insert buy err: %v", err)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.GetBuyEventReply{}, nil
}

func (s *TransactionService) GetSellEvent(ctx context.Context, req *pb.GetSellEventRequest) (*pb.GetSellEventReply, error) {
	_ = req
	err := runIncrementalEventSync(
		s,
		ctx,
		"sell_event",
		func(syncCtx context.Context) (uint64, error) {
			rLast, err := s.ac.GetSellLast(syncCtx)
			if err != nil || rLast == nil {
				return 0, err
			}
			return rLast.BlockNumber, nil
		},
		PollSoldIncremental,
		func(syncCtx context.Context, newLast uint64) error {
			return s.ac.InsertSellTrade(syncCtx, &biz.PrimarySell{
				BlockNumber: newLast,
				BlockTime:   0,
				LogIndex:    0,
				Seller:      zeroAddressHex,
				ToAddr:      zeroAddressHex,
			})
		},
		func(syncCtx context.Context, events []SoldEvent) error {
			for _, v := range events {
				err := s.ac.InsertSellTrade(syncCtx, &biz.PrimarySell{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    uint32(v.LogIndex),
					Seller:      v.Seller.String(),
					ToAddr:      v.To.String(),
					AusdGrossIn: BigIntToFloat64(v.AusdGrossIn, 18),
					AusdBurn:    BigIntToFloat64(v.AusdBurn, 18),
					AusdFee:     BigIntToFloat64(v.AusdFee, 18),
					UsdtOut:     BigIntToFloat64(v.UsdtOut, 18),
					PriceBefore: BigIntToFloat64(v.PriceBefore, 18),
					PriceAfter:  BigIntToFloat64(v.PriceAfter, 18),
				})
				if err != nil {
					s.log.Errorf("insert sell err: %v", err)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.GetSellEventReply{}, nil
}

func (s *TransactionService) GetBoxBuyEvent(ctx context.Context, req *pb.GetBoxBuyEventRequest) (*pb.GetBoxBuyEventReply, error) {
	_ = req
	err := runIncrementalEventSync(
		s,
		ctx,
		"box_buy_event",
		func(syncCtx context.Context) (uint64, error) {
			rLast, err := s.ac.GetNftMarketPurchaseLast(syncCtx)
			if err != nil || rLast == nil {
				return 0, err
			}
			return rLast.BlockNumber, nil
		},
		PollPurchasedIncremental,
		func(syncCtx context.Context, newLast uint64) error {
			return s.ac.InsertNftMarketPurchase(syncCtx, &biz.NftMarketPurchase{
				BlockNumber: newLast,
				BlockTime:   0,
				LogIndex:    0,
				Buyer:       zeroAddressHex,
				Seller:      zeroAddressHex,
				TokenID:     0,
				PriceUSDT:   0,
				FeePaidInB:  0,
				FeeUSDT:     0,
				FeeB:        0,
				CheckStatus: 1,
			})
		},
		func(syncCtx context.Context, events []PurchasedEvent) error {
			for _, v := range events {
				tmpPaid := uint8(1)
				if v.FeePaidInB {
					tmpPaid = 2
				}
				err := s.ac.InsertNftMarketPurchase(syncCtx, &biz.NftMarketPurchase{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    v.LogIndex,
					Buyer:       v.Buyer.String(),
					Seller:      v.Seller.String(),
					TokenID:     v.TokenID.Uint64(),
					PriceUSDT:   BigIntToFloat64(v.PriceUSDT, 18),
					FeePaidInB:  tmpPaid,
					FeeUSDT:     BigIntToFloat64(v.FeeUSDT, 18),
					FeeB:        BigIntToFloat64(v.FeeB, 18),
					CheckStatus: 0,
				})
				if err != nil {
					s.log.Errorf("insert market purchase err: %v", err)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.GetBoxBuyEventReply{}, nil
}

// SyncMintedFromContractHandler 从合约 view 同步铸造数据到 nft_minted，返回同步结果
func (s *TransactionService) SyncMintedFromContractHandler(ctx context.Context) map[string]interface{} {
	out := map[string]interface{}{"synced": 0, "error": ""}
	urls := bscReadURLs()
	if len(urls) == 0 {
		out["error"] = "no RPC urls"
		return out
	}
	var lastErr error
	for _, url := range urls {
		client, err := ethclient.DialContext(ctx, url)
		if err != nil {
			lastErr = fmt.Errorf("dial %s: %w", url, err)
			continue
		}
		n, err := SyncMintedFromContract(ctx, s.ac, client)
		client.Close()
		if err != nil {
			lastErr = err
			s.log.Warnf("sync_mint_from_contract %s: %v", url, err)
			continue
		}
		out["synced"] = n
		out["error"] = ""
		return out
	}
	out["error"] = "all RPCs failed: " + lastErr.Error()
	return out
}

// DebugMintSync 调试接口：返回铸造事件拉取参数与结果，排查 nft_minted 为空
func (s *TransactionService) DebugMintSync(ctx context.Context) map[string]interface{} {
	nftAddr := currentNftContract()
	res := map[string]interface{}{
		"nft_contract": nftAddr.Hex(),
		"deploy_block": conf.GetDeployBlockNftMint(DeployBlockNFTMint),
		"db_last":      nil,
		"head":         nil,
		"from":         nil,
		"to":           nil,
		"events_count": 0,
		"error":        "",
	}
	rLast, err := s.ac.GetNftMintedLast(ctx)
	if err == nil && rLast != nil {
		res["db_last"] = rLast.BlockNumber
	}
	urls := bscReadURLs()
	if len(urls) == 0 {
		res["error"] = "no RPC urls"
		return res
	}
	client, err := ethclient.DialContext(ctx, urls[0])
	if err != nil {
		res["error"] = "dial: " + err.Error()
		return res
	}
	defer client.Close()
	head, err := client.BlockNumber(ctx)
	if err != nil {
		res["error"] = "BlockNumber: " + err.Error()
		return res
	}
	res["head"] = head
	safeTo := head - Confirmations
	deployBlock := conf.GetDeployBlockNftMint(DeployBlockNFTMint)
	last := uint64(0)
	if rLast != nil {
		last = rLast.BlockNumber
	}
	from := last + 1
	if from < deployBlock {
		from = deployBlock
	}
	if safeTo < from {
		res["from"] = from
		res["to"] = safeTo
		res["error"] = "safeTo < from, no range to fetch"
		return res
	}
	to := from + MaxBlocksPerCronTick - 1
	if to > safeTo {
		to = safeTo
	}
	res["from"] = from
	res["to"] = to
	evs, _, err := PollMintedIncremental(ctx, client, last)
	if err != nil {
		res["error"] = err.Error()
		return res
	}
	res["events_count"] = len(evs)
	if len(evs) > 0 {
		first := evs[0]
		res["first_event"] = map[string]interface{}{
			"block": first.BlockNumber,
			"to":    first.To.Hex(),
			"token": first.TokenID.String(),
		}
	}
	return res
}

func (s *TransactionService) GetBoxMintEvent(ctx context.Context, req *pb.GetBoxMintEventRequest) (*pb.GetBoxMintEventReply, error) {
	_ = req
	err := runIncrementalEventSync(
		s,
		ctx,
		"box_mint_event",
		func(syncCtx context.Context) (uint64, error) {
			rLast, err := s.ac.GetNftMintedLast(syncCtx)
			if err != nil || rLast == nil {
				return 0, err
			}
			return rLast.BlockNumber, nil
		},
		func(syncCtx context.Context, client *ethclient.Client, last uint64) ([]MintedEvent, uint64, error) {
			return PollMintedIncremental(syncCtx, client, last)
		},
		func(syncCtx context.Context, newLast uint64) error {
			return s.ac.InsertNftMinted(syncCtx, &biz.NftMinted{
				BlockNumber: newLast,
				BlockTime:   0,
				LogIndex:    0,
				ToAddr:      zeroAddressHex,
				MintAddr:    zeroAddressHex,
				TokenID:     0,
				Tier:        0,
				UsdtPaid:    0,
				Status:      0,
				ListedAt:    0,
				OpenStatus:  0,
				OpenedAt:    0,
			})
		},
		func(syncCtx context.Context, events []MintedEvent) error {
			for _, v := range events {
				err := s.ac.InsertNftMinted(syncCtx, &biz.NftMinted{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    v.LogIndex,
					ToAddr:      v.To.String(),
					MintAddr:    v.To.String(),
					TokenID:     v.TokenID.Uint64(),
					Tier:        uint64(v.Tier),
					UsdtPaid:    BigIntToFloat64(v.UsdtPaid, 18),
					Status:      0,
					ListedAt:    0,
					OpenStatus:  0,
					OpenedAt:    0,
				})
				if err != nil {
					s.log.Errorf("insert minted box err: %v", err)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.GetBoxMintEventReply{}, nil
}

func (s *TransactionService) GetBoxMintEventFix(ctx context.Context, req *pb.GetBoxMintEventRequest) (*pb.GetBoxMintEventReply, error) {
	_ = req
	syncCtx, cancel := eventSyncContext(ctx)
	defer cancel()

	rLast, err := s.ac.GetNftMintedLast(syncCtx)
	if err != nil {
		return nil, err
	}
	last := uint64(0)
	if rLast != nil {
		last = rLast.BlockNumber
	}

	client := dialFirstEthClient(syncCtx, bscReadURLs())
	if client == nil {
		return &pb.GetBoxMintEventReply{}, nil
	}
	defer client.Close()

	events, _, err := PollMintedIncremental(syncCtx, client, last, 4096)
	if err != nil {
		s.log.Warnf("box_mint_event_fix poll failed: %v", err)
		return &pb.GetBoxMintEventReply{}, nil
	}
	for _, v := range events {
		err = s.ac.UpdateNftMinted(syncCtx, v.TokenID.Uint64(), v.To.String())
		if err != nil {
			s.log.Errorf("update minted box err: %v", err)
		}
	}
	return &pb.GetBoxMintEventReply{}, nil
}

func (s *TransactionService) GetBoxListEvent(ctx context.Context, req *pb.GetBoxListEventRequest) (*pb.GetBoxListEventReply, error) {
	_ = req
	err := runIncrementalEventSync(
		s,
		ctx,
		"box_list_event",
		func(syncCtx context.Context) (uint64, error) {
			rLast, err := s.ac.GetNftMarketListedLast(syncCtx)
			if err != nil || rLast == nil {
				return 0, err
			}
			return rLast.BlockNumber, nil
		},
		PollListedIncremental,
		func(syncCtx context.Context, newLast uint64) error {
			return s.ac.InsertNftMarketListed(syncCtx, &biz.NftMarketListed{
				BlockNumber: newLast,
				BlockTime:   0,
				LogIndex:    0,
				Seller:      zeroAddressHex,
				TokenID:     0,
				Timestamp:   0,
				CheckStatus: 1,
			})
		},
		func(syncCtx context.Context, events []ListedEvent) error {
			for _, v := range events {
				err := s.ac.InsertNftMarketListed(syncCtx, &biz.NftMarketListed{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    v.LogIndex,
					Seller:      v.Seller.String(),
					TokenID:     v.TokenID.Uint64(),
					Timestamp:   v.Timestamp.Uint64(),
				})
				if err != nil {
					s.log.Errorf("insert list box err: %v", err)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.GetBoxListEventReply{}, nil
}

func (s *TransactionService) GetBoxUnListEvent(ctx context.Context, req *pb.GetBoxUnListEventRequest) (*pb.GetBoxUnListEventReply, error) {
	_ = req
	err := runIncrementalEventSync(
		s,
		ctx,
		"box_unlist_event",
		func(syncCtx context.Context) (uint64, error) {
			rLast, err := s.ac.GetNftMarketUnlistedLast(syncCtx)
			if err != nil || rLast == nil {
				return 0, err
			}
			return rLast.BlockNumber, nil
		},
		PollUnlistedIncremental,
		func(syncCtx context.Context, newLast uint64) error {
			return s.ac.InsertNftMarketUnlisted(syncCtx, &biz.NftMarketUnlisted{
				BlockNumber: newLast,
				BlockTime:   0,
				LogIndex:    0,
				Operator:    zeroAddressHex,
				TokenID:     0,
				CheckStatus: 1,
			})
		},
		func(syncCtx context.Context, events []UnlistedEvent) error {
			for _, v := range events {
				err := s.ac.InsertNftMarketUnlisted(syncCtx, &biz.NftMarketUnlisted{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    v.LogIndex,
					Operator:    v.Operator.String(),
					TokenID:     v.TokenID.Uint64(),
					CheckStatus: 0,
				})
				if err != nil {
					s.log.Errorf("insert un list box err: %v", err)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.GetBoxUnListEventReply{}, nil
}

func (s *TransactionService) GetBoxOpenEvent(ctx context.Context, req *pb.GetBoxOpenEventRequest) (*pb.GetBoxOpenEventReply, error) {
	_ = req
	err := runIncrementalEventSync(
		s,
		ctx,
		"box_open_event",
		func(syncCtx context.Context) (uint64, error) {
			rLast, err := s.ac.GetNftOpenedLast(syncCtx)
			if err != nil || rLast == nil {
				return 0, err
			}
			return rLast.BlockNumber, nil
		},
		PollOpenedIncremental,
		func(syncCtx context.Context, newLast uint64) error {
			return s.ac.InsertNftOpened(syncCtx, &biz.NftOpened{
				BlockNumber: newLast,
				BlockTime:   0,
				LogIndex:    0,
				UserAddr:    zeroAddressHex,
				TokenID:     0,
				OpenedAt:    0,
				Reward:      0,
				CheckStatus: 1,
			})
		},
		func(syncCtx context.Context, events []OpenedEvent) error {
			for _, v := range events {
				err := s.ac.InsertNftOpened(syncCtx, &biz.NftOpened{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    v.LogIndex,
					UserAddr:    v.Owner.String(),
					TokenID:     v.TokenID.Uint64(),
					OpenedAt:    v.Timestamp.Uint64(),
					Reward:      BigIntToFloat64(v.UsdtFeePaid, 18),
					CheckStatus: 0,
				})
				if err != nil {
					s.log.Errorf("insert open box err: %v", err)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.GetBoxOpenEventReply{}, nil
}

func (s *TransactionService) GetBoxTransferEvent(ctx context.Context, req *pb.GetBoxTransferEventRequest) (*pb.GetBoxTransferEventReply, error) {
	_ = req
	err := runIncrementalEventSync(
		s,
		ctx,
		"box_transfer_event",
		func(syncCtx context.Context) (uint64, error) {
			rLast, err := s.ac.GetNftTransferLast(syncCtx)
			if err != nil || rLast == nil {
				return 0, err
			}
			return rLast.BlockNumber, nil
		},
		PollTransferIncremental,
		func(syncCtx context.Context, newLast uint64) error {
			return s.ac.InsertNftTransfer(syncCtx, &biz.NftTransfer{
				BlockNumber: newLast,
				BlockTime:   0,
				LogIndex:    0,
				FromAddr:    zeroAddressHex,
				ToAddr:      zeroAddressHex,
				TokenID:     0,
				CheckStatus: 1,
			})
		},
		func(syncCtx context.Context, events []TransferEvent) error {
			for _, v := range events {
				err := s.ac.InsertNftTransfer(syncCtx, &biz.NftTransfer{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    v.LogIndex,
					FromAddr:    v.From.String(),
					ToAddr:      v.To.String(),
					TokenID:     v.TokenID.Uint64(),
					CheckStatus: 0,
				})
				if err != nil {
					s.log.Errorf("insert tran box err: %v", err)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.GetBoxTransferEventReply{}, nil
}

func (s *TransactionService) GetUserREvent(ctx context.Context, req *pb.GetUserREventRequest) (*pb.GetUserREventReply, error) {
	_ = req
	err := runIncrementalEventSync(
		s,
		ctx,
		"user_r_event",
		func(syncCtx context.Context) (uint64, error) {
			rLast, err := s.ac.GetUserRegisteredLast(syncCtx)
			if err != nil || rLast == nil {
				return 0, err
			}
			return rLast.BlockNumber, nil
		},
		PollRegisteredIncremental,
		func(syncCtx context.Context, newLast uint64) error {
			return s.ac.InsertUserRegistered(syncCtx, &biz.UserRegistered{
				BlockNumber: newLast,
				BlockTime:   0,
				LogIndex:    0,
				UserAddr:    zeroAddressHex,
				ParentAddr:  zeroAddressHex,
				TopAddr:     zeroAddressHex,
			})
		},
		func(syncCtx context.Context, events []RegisteredEvent) error {
			for _, v := range events {
				err := s.ac.InsertUserRegistered(syncCtx, &biz.UserRegistered{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    v.LogIndex,
					UserAddr:    v.User.String(),
					ParentAddr:  v.Parent.String(),
					TopAddr:     v.Top.String(),
				})
				if err != nil {
					s.log.Errorf("insert user r err: %v", err)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserREventReply{}, nil
}

func (s *TransactionService) UpdateBox(ctx context.Context, req *pb.UpdateBoxRequest) (*pb.UpdateBoxReply, error) {
	_ = ctx
	syncCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.ac.UpdateBox(syncCtx, req)
	return &pb.UpdateBoxReply{}, nil
}

func (s *TransactionService) GetAddressBox(ctx context.Context, req *pb.GetAddressBoxRequest) (*pb.GetAddressBoxReply, error) {
	urls := bscReadURLs()
	vault := s.resolveVaultAddress(ctx, req.Address)

	addresses := make([]string, 0, 2)
	addresses = append(addresses, req.Address)
	if vault != "" {
		addresses = append(addresses, vault)
	}

	repoReq := addressBoxRepoRequest(req)
	minted, err := s.ac.GetAddressBoxAll(ctx, addresses, repoReq)
	if err != nil {
		return nil, err
	}
	if len(minted) == 0 && req.Page <= 1 && strings.TrimSpace(req.Address) != "" {
		syncAddressBoxState(ctx, s)
		refreshed, refreshErr := s.ac.GetAddressBoxAll(ctx, addresses, repoReq)
		if refreshErr == nil {
			minted = refreshed
		}
	}

	// 前端出售成功后传入，在链/库同步前先从本页结果中排除（不必依赖本地缓存时长）
	minted = applyExcludeTokenIDsToMinted(minted, req.GetExcludeTokenIds())

	// 空数据场景直接返回，避免无意义的链上 owner 共识校验和价格读取拖慢响应
	if len(minted) == 0 {
		return &pb.GetAddressBoxReply{
			Count: 0,
			List:  make([]*pb.GetAddressBoxReply_List, 0),
		}, nil
	}

	// 多 RPC：单节点可能滞后 1 分钟内仍返回旧 owner；多节点 owner 不一致则视为已转手，从仓库隐藏
	const maxAddrBoxRPC = 5
	clients := dialEthClients(urls, maxAddrBoxRPC)
	defer func() {
		for _, c := range clients {
			c.Close()
		}
	}()

	nftAddr := common.HexToAddress(conf.GetContract("nft"))
	userAddr := common.HexToAddress(req.Address)
	var vaultAddr common.Address
	if vault != "" {
		vaultAddr = common.HexToAddress(vault)
	}

	filtered := make([]*biz.NftMinted, 0, len(minted))
	for _, v := range minted {
		if !ownerOfConsensusKeepForAddressBox(ctx, nftAddr, v.TokenID, clients, userAddr, vaultAddr, nftAddr, req.Num) {
			continue
		}
		filtered = append(filtered, v)
	}
	mergeNftMintedOpenFromChain(ctx, nftAddr, clients, filtered)
	filtered = filterMintedByOpenTab(filtered, req.OpenStatus)
	if req.OpenStatus == 1 && req.NumFour > 0 && len(filtered) > 1 {
		sort.SliceStable(filtered, func(i, j int) bool {
			ai, aj := filtered[i].OpenedAt, filtered[j].OpenedAt
			if req.NumFour > 1 {
				if ai != aj {
					return ai < aj
				}
				return filtered[i].TokenID < filtered[j].TokenID
			}
			if ai != aj {
				return ai > aj
			}
			return filtered[i].TokenID > filtered[j].TokenID
		})
	}
	totalCount := len(filtered)
	minted = paginateMintedRows(filtered, req.Page)

	if len(minted) == 0 {
		return &pb.GetAddressBoxReply{
			Count: uint64(totalCount),
			List:  make([]*pb.GetAddressBoxReply_List, 0),
		}, nil
	}

	tokenIDs := make([]uint64, 0, len(minted))
	for _, v := range minted {
		tokenIDs = append(tokenIDs, v.TokenID)
	}
	// 与 GetMarketList 一致：走 loadSecondaryBuyPrices（多 RPC 首连、短缓存、并发），避免仅用 clients[0] 且吞错导致恒为 0。
	secondaryFull := s.loadSecondaryBuyPrices(ctx, tokenIDs)

	res := make([]*pb.GetAddressBoxReply_List, 0, len(minted))
	for _, v := range minted {
		// priceUsdt = 铸造本金；priceUsdtNow = 链上 secondaryBuyPrice（本金 + 月化线性增值后的现价，如 20u → 约 20.02u）
		priceNow := v.UsdtPaid
		if full, ok := secondaryFull[v.TokenID]; ok && full > 0 {
			priceNow = full
		}

		res = append(res, &pb.GetAddressBoxReply_List{
			TokenId:       v.TokenID,
			PriceUsdt:     v.UsdtPaid,
			PriceUsdtNow:  priceNow,
			BlockTime:     v.BlockTime,
			ListTime:      v.ListedAt,
			Tier:          v.Tier,
			OpenStatus:    uint64(v.OpenStatus),
			OpenTime:      v.OpenedAt,
			Minter:        v.MintAddr,
			BoxRewardUsdt: v.BoxRewardUsdt,
			RewardSetAt:   v.RewardSetAt,
			ChainUsdtPaid: v.ChainUsdtPaid,
		})
	}

	return &pb.GetAddressBoxReply{
		Count: uint64(totalCount),
		List:  res,
	}, nil
}

func (s *TransactionService) AddressBoxV2Handler(ctx context.Context, address string, req *pb.GetAddressBoxRequest) map[string]interface{} {
	reply, err := s.GetAddressBox(ctx, req)
	if err != nil {
		return map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		}
	}

	if len(reply.List) == 0 {
		return map[string]interface{}{
			"code": 0,
			"data": map[string]interface{}{
				"count": reply.Count,
				"list":  make([]map[string]interface{}, 0),
			},
		}
	}

	vault := s.resolveVaultAddress(ctx, address)

	addresses := []string{address}
	if vault != "" {
		addresses = append(addresses, vault)
	}
	minted, repoErr := s.ac.GetAddressBoxAll(ctx, addresses, addressBoxRepoRequest(req))
	if repoErr != nil {
		return map[string]interface{}{
			"code":    500,
			"message": repoErr.Error(),
		}
	}

	rowsByToken := make(map[uint64]*biz.NftMinted, len(minted))
	for _, item := range minted {
		rowsByToken[item.TokenID] = item
	}

	list := make([]map[string]interface{}, 0, len(reply.List))
	for _, item := range reply.List {
		row := rowsByToken[item.TokenId]
		sourceType := "market_or_transfer"
		if row != nil && row.SourceType != "" {
			sourceType = row.SourceType
		}
		list = append(list, map[string]interface{}{
			"tokenId":       item.TokenId,
			"priceUsdt":     item.PriceUsdt,
			"priceUsdtNow":  item.PriceUsdtNow,
			"blockTime":     item.BlockTime,
			"listTime":      item.ListTime,
			"tier":          item.Tier,
			"openStatus":    item.OpenStatus,
			"openTime":      item.OpenTime,
			"minter":        item.Minter,
			"boxRewardUsdt": item.BoxRewardUsdt,
			"rewardSetAt":   item.RewardSetAt,
			"chainUsdtPaid": item.ChainUsdtPaid,
			"sourceType":    sourceType,
			"sourceIsSelf":  sourceType == "self_minted",
		})
	}

	return map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"count": reply.Count,
			"list":  list,
		},
	}
}

func dialEthClients(urls []string, max int) []*ethclient.Client {
	out := make([]*ethclient.Client, 0, max)
	for _, u := range urls {
		if len(out) >= max {
			break
		}
		c, err := ethclient.Dial(u)
		if err != nil {
			continue
		}
		out = append(out, c)
	}
	return out
}

// ownerOfConsensusKeepForAddressBox 在多个 RPC 上读 ownerOf：多节点 owner 不一致则剔除；一致则按 keepAddressBoxByOwner 判断。
// 若全部 RPC 的 ownerOf 均失败（常见：token 已不存在或节点短暂异常），不展示该行，避免库中幽灵 tokenId 长期出现在 get_address_box。
func ownerOfConsensusKeepForAddressBox(ctx context.Context, nftAddr common.Address, tokenID uint64, clients []*ethclient.Client, userAddr, vaultAddr, nftContract common.Address, reqNum uint64) bool {
	if len(clients) == 0 {
		return false
	}
	owners := make([]common.Address, 0, len(clients))
	for _, c := range clients {
		inst, err := NewNft(nftAddr, c)
		if inst == nil || err != nil {
			continue
		}
		callCtx, cancel := context.WithTimeout(ctx, addressBoxEthReadTimeout)
		owner, err := inst.OwnerOf(&bind.CallOpts{Context: callCtx}, new(big.Int).SetUint64(tokenID))
		cancel()
		if err != nil {
			continue
		}
		owners = append(owners, owner)
	}
	if len(owners) == 0 {
		return false
	}
	first := owners[0]
	for _, o := range owners[1:] {
		if o != first {
			return false
		}
	}
	return keepAddressBoxByOwner(first, userAddr, vaultAddr, nftContract, reqNum)
}

// keepAddressBoxByOwner 根据链上持有人与 tab 参数（num）决定是否展示。
// 已卖出：owner 不在 用户/金库/NFT 合约内 → 不展示。
// 已挂单（合约托管）：num=0 仓库未上架 → 不展示；num=1 或 全部 → 展示。
func keepAddressBoxByOwner(owner, userAddr, vaultAddr, nftAddr common.Address, reqNum uint64) bool {
	if owner == userAddr {
		return true
	}
	if vaultAddr != (common.Address{}) && owner == vaultAddr {
		return true
	}
	if owner == nftAddr {
		return reqNum != 0
	}
	return false
}

// addressBoxRepoRequest 未上架/已开盒 Tab 若仅按 DB open_status 查询，事件未入库时会漏数据；先按「全部」拉持有人记录再在内存里用链上 boxInfo 分 Tab。
func addressBoxRepoRequest(req *pb.GetAddressBoxRequest) *pb.GetAddressBoxRequest {
	if req == nil {
		return nil
	}
	if req.OpenStatus >= 2 {
		return req
	}
	cp := *req
	cp.OpenStatus = 2
	return &cp
}

// weiToUsdtFloat 将链上 uint256/int256（18 位小数）转为 float64 展示值。
func weiToUsdtFloat(w *big.Int) float64 {
	if w == nil {
		return 0
	}
	x := new(big.Float).SetInt(w)
	x.Quo(x, big.NewFloat(1e18))
	f, _ := x.Float64()
	return f
}

// mergeNftMintedOpenFromChain 读 NFT boxInfo：对齐 openedAt / 开盒 Tab，并写入本金与盈亏字段（与 PrincipalVault 是否已写入无关）。
func mergeNftMintedOpenFromChain(ctx context.Context, nftAddr common.Address, clients []*ethclient.Client, rows []*biz.NftMinted) {
	if len(rows) == 0 || len(clients) == 0 {
		return
	}
	var wg sync.WaitGroup
	sem := make(chan struct{}, 5)
	for _, r := range rows {
		r := r
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			var ok bool
			var openedAt uint64
			var usdtPaid, reward *big.Int
			var rewardSetAt uint64
			for _, c := range clients {
				inst, err := NewNft(nftAddr, c)
				if inst == nil || err != nil {
					continue
				}
				callCtx, cancel := context.WithTimeout(ctx, addressBoxEthReadTimeout)
				info, err := inst.BoxInfo(&bind.CallOpts{Context: callCtx}, new(big.Int).SetUint64(r.TokenID))
				cancel()
				if err != nil {
					continue
				}
				ok = true
				openedAt = info.OpenedAt
				usdtPaid = info.UsdtPaid
				reward = info.Reward
				rewardSetAt = info.RewardSetAt
				break
			}
			if !ok {
				return
			}
			r.ChainUsdtPaid = weiToUsdtFloat(usdtPaid)
			r.BoxRewardUsdt = weiToUsdtFloat(reward)
			r.RewardSetAt = rewardSetAt
			if openedAt > 0 {
				r.OpenStatus = 1
				if openedAt > r.OpenedAt {
					r.OpenedAt = openedAt
				}
			}
		}()
	}
	wg.Wait()
}

func filterMintedByOpenTab(rows []*biz.NftMinted, openStatus uint64) []*biz.NftMinted {
	if openStatus >= 2 {
		return rows
	}
	out := make([]*biz.NftMinted, 0, len(rows))
	for _, v := range rows {
		if openStatus == 1 {
			if v.OpenStatus == 1 {
				out = append(out, v)
			}
			continue
		}
		if v.OpenStatus != 1 {
			out = append(out, v)
		}
	}
	return out
}

func (s *TransactionService) GetMarketList(ctx context.Context, req *pb.GetMarketListRequest) (*pb.GetMarketListReply, error) {
	minted, count, err := s.ac.GetNftMintedPage(ctx, req)
	if err != nil {
		return nil, err
	}

	// 按 token_id 去重（防御：历史版本对每个 RPC 重复追加；或库里同 token 多行）
	seen := make(map[uint64]struct{}, len(minted))
	uniq := make([]*biz.NftMinted, 0, len(minted))
	for _, v := range minted {
		if v == nil {
			continue
		}
		if _, ok := seen[v.TokenID]; ok {
			continue
		}
		seen[v.TokenID] = struct{}{}
		uniq = append(uniq, v)
	}
	minted = uniq

	res := make([]*pb.GetMarketListReply_List, 0, len(minted))
	tokenIDs := make([]uint64, 0, len(minted))
	for _, v := range minted {
		tokenIDs = append(tokenIDs, v.TokenID)
	}
	prices := s.loadSecondaryBuyPrices(ctx, tokenIDs)

	for _, v := range minted {
		res = append(res, &pb.GetMarketListReply_List{
			TokenId:      v.TokenID,
			PriceUsdt:    v.UsdtPaid,
			PriceUsdtNow: prices[v.TokenID],
			BlockTime:    v.BlockTime,
			ListTime:     v.ListedAt,
			Tier:         v.Tier,
			Minter:       v.MintAddr,
		})
	}

	return &pb.GetMarketListReply{
		Count: uint64(count),
		List:  res,
	}, nil
}

func (s *TransactionService) GetSellBoxList(ctx context.Context, req *pb.GetSellBoxListRequest) (*pb.GetSellBoxListReply, error) {
	addresses := []string{req.Address}
	vault := s.resolveVaultAddress(ctx, req.Address)
	if vault != "" {
		addresses = append(addresses, vault)
	}

	minted, count, err := s.ac.GetSellBoxList(ctx, addresses, req)
	if err != nil {
		return nil, err
	}

	res := make([]*pb.GetSellBoxListReply_List, 0, len(minted))
	for _, v := range minted {
		res = append(res, &pb.GetSellBoxListReply_List{
			TokenId:   v.TokenID,
			PriceUsdt: v.PriceUSDT,
			FeeUsdt:   v.FeeUSDT,
			FeeB:      v.FeeB,
			BlockTime: v.BlockTime,
		})
	}

	return &pb.GetSellBoxListReply{
		Count: uint64(count),
		List:  res,
	}, nil
}

func (s *TransactionService) GetAllInfo(ctx context.Context, req *pb.GetAllInfoRequest) (*pb.GetAllInfoReply, error) {
	reply, err := s.ac.GetAllInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	summary, summaryErr := s.loadMintSummaryFromChain(ctx, req.Start, req.End)
	if summaryErr != nil {
		s.log.Warnf("GetAllInfo chain summary fallback to DB: %v", summaryErr)
		return reply, nil
	}

	reply.TodayMintedCount = summary.rangeCount
	reply.TodayMintedSum = formatWei18String(summary.rangeSum)
	reply.MintedCount = summary.totalCount
	reply.MintedSum = formatWei18String(summary.totalSum)
	reply.MintedCountNoOpen = summary.notOpenCount
	reply.MintedSumNoOpen = formatWei18String(summary.notOpenSum)
	reply.MintedCountOne = summary.tierCounts[1]
	reply.MintedSumOne = formatWei18String(summary.tierSums[1])
	reply.MintedCountTwo = summary.tierCounts[2]
	reply.MintedSumTwo = formatWei18String(summary.tierSums[2])
	reply.MintedCountThree = summary.tierCounts[3]
	reply.MintedSumThree = formatWei18String(summary.tierSums[3])
	reply.OpenReward = formatWei18String(summary.totalOpenReward)
	return reply, nil
}

// GetOpenedBoxProfitStats 按日开盒利润统计：已开盲盒的利润明细与汇总（可筛选某天/时间段）
func (s *TransactionService) GetOpenedBoxProfitStats(ctx context.Context, req *pb.GetOpenedBoxProfitStatsRequest) (*pb.GetOpenedBoxProfitStatsReply, error) {
	reply, err := s.ac.GetOpenedBoxProfitStats(ctx, req)
	if err != nil {
		return nil, err
	}
	for _, item := range reply.List {
		if reward, parseErr := strconv.ParseFloat(item.Reward, 64); parseErr == nil {
			item.Reward = formatReward4(reward)
		}
	}
	return reply, nil
}

// GetFeeIncomeStats 手续费收入统计：二级市场交易手续费（USDT/HRX）按时间段汇总
func (s *TransactionService) GetFeeIncomeStats(ctx context.Context, req *pb.GetFeeIncomeStatsRequest) (*pb.GetFeeIncomeStatsReply, error) {
	return s.ac.GetFeeIncomeStats(ctx, req)
}

func (s *TransactionService) UserProfitSummaryHandler(ctx context.Context, address string, start, end uint64) map[string]interface{} {
	summary := s.ac.GetUserProfitSummary(ctx, address, start, end)
	return map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"address":          summary.Address,
			"dailyOpenProfit":  summary.DailyOpenProfit,
			"totalOpenProfit":  summary.TotalOpenProfit,
			"totalMintProfit":  summary.TotalMintProfit,
			"settledOpenCount": summary.SettledOpenCount,
			"todayOpenCount":   summary.TodayOpenCount,
		},
	}
}

func (s *TransactionService) UserRelationSummaryHandler(ctx context.Context, address string) map[string]interface{} {
	address = strings.TrimSpace(address)
	if !common.IsHexAddress(address) {
		return map[string]interface{}{
			"code":    400,
			"message": "invalid address",
		}
	}
	relation := s.ac.GetUserRelationSummary(ctx, address)
	syncSource := "db"
	if !relation.Registered {
		direct, source := s.resolveUserRelationDirect(ctx, address)
		if source != "" {
			relation = &direct
			syncSource = source
		} else {
			syncSource = "not_found"
		}
	}
	return map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"address":    relation.Address,
			"parentAddr": relation.ParentAddr,
			"topAddr":    relation.TopAddr,
			"registered": relation.Registered,
			"syncSource": syncSource,
		},
	}
}

func (s *TransactionService) RuntimeConfigListHandler(ctx context.Context, prefix string) map[string]interface{} {
	items, err := s.ac.ListRuntimeConfigs(ctx, strings.TrimSpace(prefix))
	if err != nil {
		return map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		}
	}
	list := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		list = append(list, map[string]interface{}{
			"key":         item.Key,
			"value":       item.Value,
			"description": item.Description,
		})
	}
	return map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"list": list,
		},
	}
}

func (s *TransactionService) RuntimeConfigUpsertHandler(ctx context.Context, key, value, description string) map[string]interface{} {
	key = strings.TrimSpace(key)
	if key == "" {
		return map[string]interface{}{
			"code":    400,
			"message": "key is required",
		}
	}
	if err := s.ac.UpsertRuntimeConfig(ctx, key, value, description); err != nil {
		return map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"key":         key,
			"value":       value,
			"description": description,
		},
	}
}

func (s *TransactionService) syncUserRelationFromChain(ctx context.Context, address string) (bool, error) {
	urls := conf.GetBscRpcUrlsForPolling()
	if len(urls) == 0 {
		urls = conf.GetBscRpcUrls()
	}
	user := common.HexToAddress(address)
	for _, urlTmp := range urls {
		client, err := ethclient.DialContext(ctx, urlTmp)
		if err != nil {
			continue
		}
		ev, fetchErr := FetchLatestRegisteredByUser(ctx, client, RegisterContract, user, DeployBlockRegistered)
		client.Close()
		if fetchErr != nil {
			continue
		}
		if ev == nil {
			continue
		}
		if err := s.ac.InsertUserRegistered(ctx, &biz.UserRegistered{
			BlockNumber: ev.BlockNumber,
			BlockTime:   ev.BlockTime,
			LogIndex:    ev.LogIndex,
			UserAddr:    ev.User.Hex(),
			ParentAddr:  ev.Parent.Hex(),
			TopAddr:     ev.Top.Hex(),
		}); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

const (
	eventSyncHTTPTimeout = 12 * time.Second
	eventSyncDialTimeout = 1800 * time.Millisecond
)

func eventSyncContext(ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithTimeout(ctx, eventSyncHTTPTimeout)
}

func runIncrementalEventSync[T any](
	s *TransactionService,
	ctx context.Context,
	label string,
	getLast func(context.Context) (uint64, error),
	poll func(context.Context, *ethclient.Client, uint64) ([]T, uint64, error),
	onEmpty func(context.Context, uint64) error,
	onEvents func(context.Context, []T) error,
) error {
	syncCtx, cancel := eventSyncContext(ctx)
	defer cancel()

	last, err := getLast(syncCtx)
	if err != nil {
		return err
	}

	urls := bscReadURLs()
	if len(urls) == 0 {
		return nil
	}

	var lastErr error
	for _, url := range urls {
		dialCtx, dialCancel := context.WithTimeout(syncCtx, eventSyncDialTimeout)
		client, err := ethclient.DialContext(dialCtx, url)
		dialCancel()
		if err != nil {
			lastErr = err
			continue
		}

		events, newLast, err := poll(syncCtx, client, last)
		client.Close()
		if err != nil {
			lastErr = err
			s.log.Warnf("%s poll via %s failed: %v", label, url, err)
			if syncCtx.Err() != nil {
				break
			}
			continue
		}
		if newLast <= last {
			return nil
		}
		if len(events) == 0 {
			if onEmpty != nil {
				return onEmpty(syncCtx, newLast)
			}
			return nil
		}
		return onEvents(syncCtx, events)
	}

	if lastErr != nil {
		s.log.Warnf("%s sync skipped: %v", label, lastErr)
	}
	return nil
}

func ParseUint64Query(values url.Values, key string) uint64 {
	raw := strings.TrimSpace(values.Get(key))
	if raw == "" {
		return 0
	}
	v, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	// DeployBlock TODO: 你把这里替换成实际部署块高（从 BscScan/OKLink 的 Contract Creation 里看）
	DeployBlock uint64 = 73406834

	// Confirmations 防回滚确认数：最新块往回退 6 个块
	Confirmations uint64 = 6

	// QueryStep 每次 FilterLogs 的分段大小（BSC 公共 RPC 经常会限制范围，建议 2k~5k）
	QueryStep uint64 = 2000

	// MaxBlocksPerCronTick 计划任务约每分钟跑一次时，单次 Poll 最多推进的块数；过大易超 1 分钟，过小追历史慢。
	MaxBlocksPerCronTick uint64 = 6000

	// envScanBaseBlock 全链游标统一起点：当各表已处理块高 last 低于此值时，先视为已处理到该块（下一格从 last+1 扫，即至少从 该值+1）。
	// 清库/重扫时设为当前链附近块高。未设置或 "0" 则关闭（仅按原 DB 与 DeployBlock* 逻辑）。生产勿随意设置，会跳过 [last, base) 区间日志。
	envScanBaseBlock = "REQUESTETH_SCAN_BASE_BLOCK"
)

// scanBaseBlockLowerBound 从环境读扫块下界；0 表示关闭。
func scanBaseBlockLowerBound() uint64 {
	s := strings.TrimSpace(os.Getenv(envScanBaseBlock))
	if s == "" {
		return 0
	}
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

// bumpLastToScanBaseIfNeeded 把 DB 游标抬到全链重扫起点（与 DeployBlock 取大在后续 from 里自然实现）。
func bumpLastToScanBaseIfNeeded(last uint64) uint64 {
	b := scanBaseBlockLowerBound()
	if b == 0 || last >= b {
		return last
	}
	return b
}

// maxAdvanceBlocksOnEmptyFilter 当 FilterLogs 在本窗口返回 0 条时，哨兵（onEmpty）最多把游标推进到的块跨度。
// 避免 [from,to] 一次可达 MaxBlocksPerCronTick 时若 RPC 误返回空则整段被标为已处理、永久漏事件。
// 有事件时 capNewLastWhenNoEvents 仍返回完整 to，行为与改前一致。
const maxAdvanceBlocksOnEmptyFilter = uint64(2000)

func capNewLastWhenNoEvents(nEvents int, from, to uint64) uint64 {
	if nEvents > 0 {
		return to
	}
	if to < from {
		return to
	}
	span := to - from + 1
	if span <= maxAdvanceBlocksOnEmptyFilter {
		return to
	}
	return from + maxAdvanceBlocksOnEmptyFilter - 1
}

// syncCron*：单次 HTTP 扫链目标在 1 分钟内结束（与 curl --max-time 58 对齐），避免每分钟叠跑占满 RPC。
var (
	syncCronHTTPWall   = 48 * time.Second
	syncCronCtxTimeout = 55 * time.Second
	syncCronOuterMax   = 8
)

// rpcThrottle 全局 RPC 限速：所有 FilterLogs 调用共享，防止并发突发超过节点配额。
var (
	rpcMu       sync.Mutex
	rpcLastCall time.Time
	rpcMinGap   = 80 * time.Millisecond // ~12 calls/sec，10 个同步进程共享
)

func rpcThrottle() {
	rpcMu.Lock()
	defer rpcMu.Unlock()
	if gap := time.Since(rpcLastCall); gap < rpcMinGap {
		time.Sleep(rpcMinGap - gap)
	}
	rpcLastCall = time.Now()
}

var (
	DividendContract = common.HexToAddress(conf.GetContract("dividend"))
)

// 仅包含 RewardNotified 事件（够用就行）
const umbrellaStakeDividendABI = `[
  {
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "user",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "profit",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "userShare",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "l1",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "l2",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "top",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "pool",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "uplinePortionBps",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "toL1",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "toL2",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "toTop",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "toProject",
				"type": "uint256"
			}
		],
		"name": "RewardNotified",
		"type": "event"
	}
]`

type RewardNotifiedEvent struct {
	// 用于唯一定位/落库去重
	BlockNumber uint64
	BlockTime   uint64
	TxHash      common.Hash
	LogIndex    uint

	// indexed
	User common.Address
	L1   common.Address
	L2   common.Address

	// data（非 indexed）
	Profit           *big.Int
	UserShare        *big.Int
	Top              common.Address
	Pool             *big.Int
	UplinePortionBps *big.Int
	ToL1             *big.Int
	ToL2             *big.Int
	ToTop            *big.Int
	ToProject        *big.Int
}

func parseRewardNotified(a abi.ABI, lg types.Log) (RewardNotifiedEvent, error) {
	ev, ok := a.Events["RewardNotified"]
	if !ok {
		return RewardNotifiedEvent{}, fmt.Errorf("event RewardNotified not found in ABI")
	}
	if len(lg.Topics) != 4 {
		return RewardNotifiedEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return RewardNotifiedEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := RewardNotifiedEvent{
		BlockNumber: lg.BlockNumber,
		TxHash:      lg.TxHash,
		LogIndex:    lg.Index,

		User: common.BytesToAddress(lg.Topics[1].Bytes()),
		L1:   common.BytesToAddress(lg.Topics[2].Bytes()),
		L2:   common.BytesToAddress(lg.Topics[3].Bytes()),
	}

	// data 部分按事件里“非 indexed inputs”的顺序解
	values, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return RewardNotifiedEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	// 顺序：
	// profit, userShare, top, pool, uplinePortionBps, toL1, toL2, toTop, toProject
	if len(values) != 9 {
		return RewardNotifiedEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(values))
	}

	out.Profit = values[0].(*big.Int)
	out.UserShare = values[1].(*big.Int)
	out.Top = values[2].(common.Address)
	out.Pool = values[3].(*big.Int)
	out.UplinePortionBps = values[4].(*big.Int)
	out.ToL1 = values[5].(*big.Int)
	out.ToL2 = values[6].(*big.Int)
	out.ToTop = values[7].(*big.Int)
	out.ToProject = values[8].(*big.Int)

	return out, nil
}

func fillRewardNotifiedTimes(ctx context.Context, client *ethclient.Client, evs []RewardNotifiedEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}
		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

func fetchRewardNotifiedLogWindow(ctx context.Context, client *ethclient.Client, contract common.Address, eventID common.Hash, parsedABI abi.ABI, start, end uint64, res *[]RewardNotifiedEvent) error {
	if end < start {
		return nil
	}
	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(end),
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{eventID}},
	}
	logs, err := filterLogsRetryGeneric(ctx, client, q)
	if err != nil {
		if isFilterLogsLimitErr(err) && start < end {
			mid := start + (end-start)/2
			if err := fetchRewardNotifiedLogWindow(ctx, client, contract, eventID, parsedABI, start, mid, res); err != nil {
				return err
			}
			return fetchRewardNotifiedLogWindow(ctx, client, contract, eventID, parsedABI, mid+1, end, res)
		}
		return fmt.Errorf("FilterLogs RewardNotified [%d,%d]: %w", start, end, err)
	}
	for _, lg := range logs {
		ev, err := parseRewardNotified(parsedABI, lg)
		if err != nil {
			return fmt.Errorf("parse RewardNotified tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
		}
		*res = append(*res, ev)
	}
	return nil
}

// FetchRewardNotifiedByRange 按 [fromBlock, toBlock]（包含边界）拉取 RewardNotified 事件，内部自动分段 QueryStep
func FetchRewardNotifiedByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock uint64, toBlock uint64) ([]RewardNotifiedEvent, error) {

	if toBlock < fromBlock {
		return []RewardNotifiedEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(umbrellaStakeDividendABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	eventID := parsedABI.Events["RewardNotified"].ID

	res := make([]RewardNotifiedEvent, 0, 64)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}
		if err := fetchRewardNotifiedLogWindow(ctx, client, contract, eventID, parsedABI, start, end, &res); err != nil {
			return nil, err
		}
	}

	// ✅ 统一补时间戳
	if err := fillRewardNotifiedTimes(ctx, client, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PollRewardNotifiedIncremental 你描述的“增量拉取 + 最新块往回 6”核心逻辑
// lastProcessedFromDB：DB 里记录的“已处理到的最高块”（建议含义：你上次成功跑完时的 safeTo）
func PollRewardNotifiedIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []RewardNotifiedEvent, newLastProcessed uint64, err error) {

	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		// 链太新/异常情况
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	lastProcessedFromDB = bumpLastToScanBaseIfNeeded(lastProcessedFromDB)

	from := lastProcessedFromDB + 1
	if from < DeployBlock {
		from = DeployBlock
	}

	if safeTo < from {
		// 没有可确认的新块
		return []RewardNotifiedEvent{}, lastProcessedFromDB, nil
	}

	to := from + MaxBlocksPerCronTick - 1
	if to > safeTo {
		to = safeTo
	}

	evs, err := FetchRewardNotifiedByRange(ctx, client, DividendContract, from, to)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}

	// 只要你把 evs 全部落库成功（建议 txHash+logIndex 唯一键），再把 DB 的 lastProcessed 更新为本次扫到的 to
	return evs, capNewLastWhenNoEvents(len(evs), from, to), nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func BigIntToFloat64(x *big.Int, decimals int) float64 {
	if x == nil {
		return 0
	}
	r := new(big.Rat).SetInt(x)
	if decimals > 0 {
		r.Quo(r, new(big.Rat).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)))
	}
	f, _ := r.Float64()
	return f
}

const (
	// DeployBlockTwo TODO: 改成 SwapHrxUsdt 实际部署块高
	DeployBlockTwo uint64 = 72859368
)

const swapHrxUsdtABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "sender", "type": "address"},
      {"indexed": false, "internalType": "uint256", "name": "amount0In", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "amount1In", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "amount0OutGross", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "amount0OutNet", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "amount1OutGross", "type": "uint256"},
      {"indexed": true,  "internalType": "address", "name": "to", "type": "address"}
    ],
    "name": "Swap",
    "type": "event"
  }
]`

type SwapEvent struct {
	// 唯一定位（建议落库用 txHash+logIndex 去重）
	BlockNumber uint64
	TxHash      common.Hash
	LogIndex    uint

	// ✅ 秒级时间戳（区块级，同区块所有交易相同）
	BlockTime uint64

	// indexed
	Sender common.Address
	To     common.Address

	// data
	Amount0In       *big.Int
	Amount1In       *big.Int
	Amount0OutGross *big.Int
	Amount0OutNet   *big.Int
	Amount1OutGross *big.Int
}

func parseSwap(a abi.ABI, lg types.Log) (SwapEvent, error) {
	ev, ok := a.Events["Swap"]
	if !ok {
		return SwapEvent{}, fmt.Errorf("event Swap not found in ABI")
	}
	if len(lg.Topics) != 3 {
		return SwapEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return SwapEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := SwapEvent{
		BlockNumber: lg.BlockNumber,
		TxHash:      lg.TxHash,
		LogIndex:    lg.Index,

		Sender: common.BytesToAddress(lg.Topics[1].Bytes()),
		To:     common.BytesToAddress(lg.Topics[2].Bytes()),
	}

	vals, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return SwapEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	if len(vals) != 5 {
		return SwapEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(vals))
	}

	out.Amount0In = vals[0].(*big.Int)
	out.Amount1In = vals[1].(*big.Int)
	out.Amount0OutGross = vals[2].(*big.Int)
	out.Amount0OutNet = vals[3].(*big.Int)
	out.Amount1OutGross = vals[4].(*big.Int)
	return out, nil
}

func fetchSwapLogWindow(ctx context.Context, client *ethclient.Client, contract common.Address, swapID common.Hash, parsedABI abi.ABI, start, end uint64, res *[]SwapEvent) error {
	if end < start {
		return nil
	}
	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(end),
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{swapID}},
	}
	logs, err := filterLogsRetryGeneric(ctx, client, q)
	if err != nil {
		if isFilterLogsLimitErr(err) && start < end {
			mid := start + (end-start)/2
			if err := fetchSwapLogWindow(ctx, client, contract, swapID, parsedABI, start, mid, res); err != nil {
				return err
			}
			return fetchSwapLogWindow(ctx, client, contract, swapID, parsedABI, mid+1, end, res)
		}
		return fmt.Errorf("FilterLogs Swap [%d,%d]: %w", start, end, err)
	}
	for _, lg := range logs {
		ev, err := parseSwap(parsedABI, lg)
		if err != nil {
			return fmt.Errorf("parse Swap tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
		}
		*res = append(*res, ev)
	}
	return nil
}

// ✅ 批量填充 blockTime：缓存 blockNumber -> header.Time（秒）
func fillBlockTimes(ctx context.Context, client *ethclient.Client, evs []SwapEvent) error {
	cache := make(map[uint64]uint64, 256)

	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}
		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

// FetchSwapByRange 按 [fromBlock, toBlock] 拉取 Swap，并填充 BlockTime(秒)
func FetchSwapByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]SwapEvent, error) {
	if toBlock < fromBlock {
		return []SwapEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(swapHrxUsdtABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	swapID := parsedABI.Events["Swap"].ID

	res := make([]SwapEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}
		if err := fetchSwapLogWindow(ctx, client, contract, swapID, parsedABI, start, end, &res); err != nil {
			return nil, err
		}
	}

	// ✅ 统一补时间戳
	if err := fillBlockTimes(ctx, client, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PollSwapIncremental：增量拉取 + latest-6
func PollSwapIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []SwapEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	lastProcessedFromDB = bumpLastToScanBaseIfNeeded(lastProcessedFromDB)

	from := lastProcessedFromDB + 1
	if from < DeployBlockTwo {
		from = DeployBlockTwo
	}
	if safeTo < from {
		return []SwapEvent{}, lastProcessedFromDB, nil
	}

	to := from + MaxBlocksPerCronTick - 1
	if to > safeTo {
		to = safeTo
	}

	evs, err := FetchSwapByRange(ctx, client, currentSwapContract(), from, to)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, capNewLastWhenNoEvents(len(evs), from, to), nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	DeployBlockThree uint64 = 72528236
	DeployBlockFour  uint64 = 72869154
)

var (
	PrimaryMarketContract = common.HexToAddress(conf.GetContract("primary_market"))
)

// 仅包含 Bought / Sold 事件（够用就行）
const primaryMarketABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "buyer", "type": "address"},
      {"indexed": true,  "internalType": "address", "name": "to", "type": "address"},
      {"indexed": false, "internalType": "uint256", "name": "usdtUsed", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "ausdGrossOut", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "ausdFee", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "ausdNetOut", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "priceBefore", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "priceAfter", "type": "uint256"}
    ],
    "name": "Bought",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "seller", "type": "address"},
      {"indexed": true,  "internalType": "address", "name": "to", "type": "address"},
      {"indexed": false, "internalType": "uint256", "name": "ausdGrossIn", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "ausdFee", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "ausdBurn", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "usdtOut", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "priceBefore", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "priceAfter", "type": "uint256"}
    ],
    "name": "Sold",
    "type": "event"
  }
]`

// -------------------- Bought --------------------

type BoughtEvent struct {
	BlockNumber uint64
	TxHash      common.Hash
	LogIndex    uint

	// ✅ 秒级时间戳（区块级）
	BlockTime uint64

	// indexed
	Buyer common.Address
	To    common.Address

	// data
	UsdtUsed     *big.Int
	AusdGrossOut *big.Int
	AusdFee      *big.Int
	AusdNetOut   *big.Int
	PriceBefore  *big.Int
	PriceAfter   *big.Int
}

func parseBought(a abi.ABI, lg types.Log) (BoughtEvent, error) {
	ev, ok := a.Events["Bought"]
	if !ok {
		return BoughtEvent{}, fmt.Errorf("event Bought not found in ABI")
	}
	// topics: [topic0=eventID, topic1=buyer, topic2=to]
	if len(lg.Topics) != 3 {
		return BoughtEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return BoughtEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := BoughtEvent{
		BlockNumber: lg.BlockNumber,
		TxHash:      lg.TxHash,
		LogIndex:    lg.Index,

		Buyer: common.BytesToAddress(lg.Topics[1].Bytes()),
		To:    common.BytesToAddress(lg.Topics[2].Bytes()),
	}

	// 非 indexed inputs 顺序：
	// usdtUsed, ausdGrossOut, ausdFee, ausdNetOut, priceBefore, priceAfter
	values, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return BoughtEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	if len(values) != 6 {
		return BoughtEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(values))
	}

	out.UsdtUsed = values[0].(*big.Int)
	out.AusdGrossOut = values[1].(*big.Int)
	out.AusdFee = values[2].(*big.Int)
	out.AusdNetOut = values[3].(*big.Int)
	out.PriceBefore = values[4].(*big.Int)
	out.PriceAfter = values[5].(*big.Int)

	return out, nil
}

func fetchBoughtLogWindow(ctx context.Context, client *ethclient.Client, contract common.Address, boughtID common.Hash, parsedABI abi.ABI, start, end uint64, res *[]BoughtEvent) error {
	if end < start {
		return nil
	}
	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(end),
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{boughtID}},
	}
	logs, err := filterLogsRetryGeneric(ctx, client, q)
	if err != nil {
		if isFilterLogsLimitErr(err) && start < end {
			mid := start + (end-start)/2
			if err := fetchBoughtLogWindow(ctx, client, contract, boughtID, parsedABI, start, mid, res); err != nil {
				return err
			}
			return fetchBoughtLogWindow(ctx, client, contract, boughtID, parsedABI, mid+1, end, res)
		}
		return fmt.Errorf("FilterLogs Bought [%d,%d]: %w", start, end, err)
	}
	for _, lg := range logs {
		ev, err := parseBought(parsedABI, lg)
		if err != nil {
			return fmt.Errorf("parse Bought tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
		}
		*res = append(*res, ev)
	}
	return nil
}

// FetchBoughtByRange 按 [fromBlock, toBlock] 拉取 Bought，自动分段 QueryStep，并填 BlockTime
func FetchBoughtByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]BoughtEvent, error) {
	if toBlock < fromBlock {
		return []BoughtEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(primaryMarketABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	boughtID := parsedABI.Events["Bought"].ID

	res := make([]BoughtEvent, 0, 128)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}
		if err := fetchBoughtLogWindow(ctx, client, contract, boughtID, parsedABI, start, end, &res); err != nil {
			return nil, err
		}
	}

	if err := fillBoughtTimes(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

func fillBoughtTimes(ctx context.Context, client *ethclient.Client, evs []BoughtEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}
		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

// PollBoughtIncremental 增量拉取 + latest-6
func PollBoughtIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []BoughtEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	lastProcessedFromDB = bumpLastToScanBaseIfNeeded(lastProcessedFromDB)

	from := lastProcessedFromDB + 1
	if from < DeployBlockThree {
		from = DeployBlockThree
	}
	if safeTo < from {
		return []BoughtEvent{}, lastProcessedFromDB, nil
	}

	to := from + MaxBlocksPerCronTick - 1
	if to > safeTo {
		to = safeTo
	}

	evs, err := FetchBoughtByRange(ctx, client, PrimaryMarketContract, from, to)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, capNewLastWhenNoEvents(len(evs), from, to), nil
}

// -------------------- Sold --------------------

type SoldEvent struct {
	BlockNumber uint64
	TxHash      common.Hash
	LogIndex    uint

	// ✅ 秒级时间戳（区块级）
	BlockTime uint64

	// indexed
	Seller common.Address
	To     common.Address

	// data
	AusdGrossIn *big.Int
	AusdFee     *big.Int
	AusdBurn    *big.Int
	UsdtOut     *big.Int
	PriceBefore *big.Int
	PriceAfter  *big.Int
}

func parseSold(a abi.ABI, lg types.Log) (SoldEvent, error) {
	ev, ok := a.Events["Sold"]
	if !ok {
		return SoldEvent{}, fmt.Errorf("event Sold not found in ABI")
	}
	// topics: [topic0=eventID, topic1=seller, topic2=to]
	if len(lg.Topics) != 3 {
		return SoldEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return SoldEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := SoldEvent{
		BlockNumber: lg.BlockNumber,
		TxHash:      lg.TxHash,
		LogIndex:    lg.Index,

		Seller: common.BytesToAddress(lg.Topics[1].Bytes()),
		To:     common.BytesToAddress(lg.Topics[2].Bytes()),
	}

	// 非 indexed inputs 顺序：
	// ausdGrossIn, ausdFee, ausdBurn, usdtOut, priceBefore, priceAfter
	values, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return SoldEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	if len(values) != 6 {
		return SoldEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(values))
	}

	out.AusdGrossIn = values[0].(*big.Int)
	out.AusdFee = values[1].(*big.Int)
	out.AusdBurn = values[2].(*big.Int)
	out.UsdtOut = values[3].(*big.Int)
	out.PriceBefore = values[4].(*big.Int)
	out.PriceAfter = values[5].(*big.Int)

	return out, nil
}

func fetchSoldLogWindow(ctx context.Context, client *ethclient.Client, contract common.Address, soldID common.Hash, parsedABI abi.ABI, start, end uint64, res *[]SoldEvent) error {
	if end < start {
		return nil
	}
	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(end),
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{soldID}},
	}
	logs, err := filterLogsRetryGeneric(ctx, client, q)
	if err != nil {
		if isFilterLogsLimitErr(err) && start < end {
			mid := start + (end-start)/2
			if err := fetchSoldLogWindow(ctx, client, contract, soldID, parsedABI, start, mid, res); err != nil {
				return err
			}
			return fetchSoldLogWindow(ctx, client, contract, soldID, parsedABI, mid+1, end, res)
		}
		return fmt.Errorf("FilterLogs Sold [%d,%d]: %w", start, end, err)
	}
	for _, lg := range logs {
		ev, err := parseSold(parsedABI, lg)
		if err != nil {
			return fmt.Errorf("parse Sold tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
		}
		*res = append(*res, ev)
	}
	return nil
}

// FetchSoldByRange 按 [fromBlock, toBlock] 拉取 Sold，自动分段 QueryStep，并填 BlockTime
func FetchSoldByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]SoldEvent, error) {
	if toBlock < fromBlock {
		return []SoldEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(primaryMarketABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	soldID := parsedABI.Events["Sold"].ID

	res := make([]SoldEvent, 0, 128)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}
		if err := fetchSoldLogWindow(ctx, client, contract, soldID, parsedABI, start, end, &res); err != nil {
			return nil, err
		}
	}

	if err := fillSoldTimes(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

func fillSoldTimes(ctx context.Context, client *ethclient.Client, evs []SoldEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}
		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

// PollSoldIncremental 增量拉取 + latest-6
func PollSoldIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []SoldEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	lastProcessedFromDB = bumpLastToScanBaseIfNeeded(lastProcessedFromDB)

	from := lastProcessedFromDB + 1
	if from < DeployBlockFour {
		from = DeployBlockFour
	}
	if safeTo < from {
		return []SoldEvent{}, lastProcessedFromDB, nil
	}

	to := from + MaxBlocksPerCronTick - 1
	if to > safeTo {
		to = safeTo
	}

	evs, err := FetchSoldByRange(ctx, client, PrimaryMarketContract, from, to)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, capNewLastWhenNoEvents(len(evs), from, to), nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	// DeployBlockNFT TODO: 改成 BlindBoxNFT 实际部署块高
	DeployBlockNFT uint64 = 72530801
)

const blindBoxPurchasedABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "buyer", "type": "address"},
      {"indexed": true,  "internalType": "address", "name": "seller", "type": "address"},
      {"indexed": true,  "internalType": "uint256", "name": "tokenId", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "priceUSDT", "type": "uint256"},
      {"indexed": false, "internalType": "bool",    "name": "feePaidInB", "type": "bool"},
      {"indexed": false, "internalType": "uint256", "name": "feeUSDT", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "feeB", "type": "uint256"}
    ],
    "name": "Purchased",
    "type": "event"
  }
]`

type PurchasedEvent struct {
	// 唯一定位（建议落库用 blockNumber+logIndex 去重）
	BlockNumber uint64
	LogIndex    uint

	// ✅ 秒级时间戳（区块级，同区块所有交易相同）
	BlockTime uint64

	// indexed
	Buyer   common.Address
	Seller  common.Address
	TokenID *big.Int

	// data（原样 bigint）
	PriceUSDT  *big.Int
	FeePaidInB bool
	FeeUSDT    *big.Int
	FeeB       *big.Int
}

func parsePurchased(a abi.ABI, lg types.Log) (PurchasedEvent, error) {
	ev, ok := a.Events["Purchased"]
	if !ok {
		return PurchasedEvent{}, fmt.Errorf("event Purchased not found in ABI")
	}
	// topics: [eventId, buyer, seller, tokenId]
	if len(lg.Topics) != 4 {
		return PurchasedEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return PurchasedEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := PurchasedEvent{
		BlockNumber: lg.BlockNumber,
		LogIndex:    lg.Index,

		Buyer:   common.BytesToAddress(lg.Topics[1].Bytes()),
		Seller:  common.BytesToAddress(lg.Topics[2].Bytes()),
		TokenID: new(big.Int).SetBytes(lg.Topics[3].Bytes()),
	}

	vals, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return PurchasedEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	// data: priceUSDT, feePaidInB, feeUSDT, feeB
	if len(vals) != 4 {
		return PurchasedEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(vals))
	}

	price, ok := vals[0].(*big.Int)
	if !ok {
		return PurchasedEvent{}, fmt.Errorf("bad type priceUSDT: %T", vals[0])
	}
	paidInB, ok := vals[1].(bool)
	if !ok {
		return PurchasedEvent{}, fmt.Errorf("bad type feePaidInB: %T", vals[1])
	}
	feeU, ok := vals[2].(*big.Int)
	if !ok {
		return PurchasedEvent{}, fmt.Errorf("bad type feeUSDT: %T", vals[2])
	}
	feeB, ok := vals[3].(*big.Int)
	if !ok {
		return PurchasedEvent{}, fmt.Errorf("bad type feeB: %T", vals[3])
	}

	out.PriceUSDT = price
	out.FeePaidInB = paidInB
	out.FeeUSDT = feeU
	out.FeeB = feeB

	return out, nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func filterLogsRetryPurchased(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		rpcThrottle()
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func headerByNumberRetryPurchased(ctx context.Context, client *ethclient.Client, bn uint64) (*types.Header, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err == nil {
			return h, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// fetchPurchasedLogWindow 拉取 [start,end] Purchased 日志；遇 limit exceeded 二分（与 Opened 一致）
func fetchPurchasedLogWindow(ctx context.Context, client *ethclient.Client, contract common.Address, purchasedID common.Hash, parsedABI abi.ABI, start, end uint64, res *[]PurchasedEvent) error {
	if end < start {
		return nil
	}
	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(end),
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{purchasedID}},
	}
	logs, err := filterLogsRetryPurchased(ctx, client, q)
	if err != nil {
		if isFilterLogsLimitErr(err) && start < end {
			mid := start + (end-start)/2
			if err := fetchPurchasedLogWindow(ctx, client, contract, purchasedID, parsedABI, start, mid, res); err != nil {
				return err
			}
			return fetchPurchasedLogWindow(ctx, client, contract, purchasedID, parsedABI, mid+1, end, res)
		}
		return fmt.Errorf("FilterLogs [%d,%d]: %w", start, end, err)
	}
	for _, lg := range logs {
		ev, err := parsePurchased(parsedABI, lg)
		if err != nil {
			return fmt.Errorf("parse Purchased tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
		}
		*res = append(*res, ev)
	}
	return nil
}

// ✅ 批量填充 blockTime：缓存 blockNumber -> header.Time（秒）
func fillBlockTimesPurchased(ctx context.Context, client *ethclient.Client, evs []PurchasedEvent) error {
	cache := make(map[uint64]uint64, 256)

	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}

		h, err := headerByNumberRetryPurchased(ctx, client, bn)
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}

		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

// FetchPurchasedByRange 按 [fromBlock, toBlock] 拉取 Purchased，并填充 BlockTime(秒)
func FetchPurchasedByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]PurchasedEvent, error) {
	if toBlock < fromBlock {
		return []PurchasedEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(blindBoxPurchasedABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	purchasedID := parsedABI.Events["Purchased"].ID

	res := make([]PurchasedEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}
		if err := fetchPurchasedLogWindow(ctx, client, contract, purchasedID, parsedABI, start, end, &res); err != nil {
			return nil, err
		}
	}

	// ✅ 统一补时间戳（HeaderByNumber 也最多重试 5 次）
	if err := fillBlockTimesPurchased(ctx, client, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PollPurchasedIncremental：增量拉取 + latest-Confirmations
func PollPurchasedIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []PurchasedEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	lastProcessedFromDB = bumpLastToScanBaseIfNeeded(lastProcessedFromDB)

	from := lastProcessedFromDB + 1
	if from < DeployBlockNFT {
		from = DeployBlockNFT
	}
	if safeTo < from {
		return []PurchasedEvent{}, lastProcessedFromDB, nil
	}

	to := from + MaxBlocksPerCronTick - 1
	if to > safeTo {
		to = safeTo
	}

	evs, err := FetchPurchasedByRange(ctx, client, currentNftContract(), from, to)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, capNewLastWhenNoEvents(len(evs), from, to), nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	DeployBlockNFTListed uint64 = 73402672 // TODO: 改成你想开始监听的块高
)

const blindBoxListedABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "seller", "type": "address"},
      {"indexed": true,  "internalType": "uint256", "name": "tokenId", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "timestamp", "type": "uint256"}
    ],
    "name": "Listed",
    "type": "event"
  }
]`

type ListedEvent struct {
	BlockNumber uint64
	LogIndex    uint
	BlockTime   uint64 // 秒级（区块时间）

	// indexed
	Seller  common.Address
	TokenID *big.Int

	// data
	Timestamp *big.Int
}

func parseListed(a abi.ABI, lg types.Log) (ListedEvent, error) {
	ev, ok := a.Events["Listed"]
	if !ok {
		return ListedEvent{}, fmt.Errorf("event Listed not found in ABI")
	}
	// topics: [eventId, seller, tokenId]
	if len(lg.Topics) != 3 {
		return ListedEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return ListedEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := ListedEvent{
		BlockNumber: lg.BlockNumber,
		LogIndex:    lg.Index,
		Seller:      common.BytesToAddress(lg.Topics[1].Bytes()),
		TokenID:     new(big.Int).SetBytes(lg.Topics[2].Bytes()),
	}

	vals, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return ListedEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	// data: timestamp
	if len(vals) != 1 {
		return ListedEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(vals))
	}

	ts, ok := vals[0].(*big.Int)
	if !ok {
		return ListedEvent{}, fmt.Errorf("bad type timestamp: %T", vals[0])
	}
	out.Timestamp = ts

	return out, nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func filterLogsRetryListed(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		rpcThrottle()
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func headerByNumberRetryListed(ctx context.Context, client *ethclient.Client, bn uint64) (*types.Header, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err == nil {
			return h, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

func fillBlockTimesListed(ctx context.Context, client *ethclient.Client, evs []ListedEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}

		h, err := headerByNumberRetryListed(ctx, client, bn)
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}

		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

func fetchListedLogWindow(ctx context.Context, client *ethclient.Client, contract common.Address, listedID common.Hash, parsedABI abi.ABI, start, end uint64, res *[]ListedEvent) error {
	if end < start {
		return nil
	}
	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(end),
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{listedID}},
	}
	logs, err := filterLogsRetryListed(ctx, client, q)
	if err != nil {
		if isFilterLogsLimitErr(err) && start < end {
			mid := start + (end-start)/2
			if err := fetchListedLogWindow(ctx, client, contract, listedID, parsedABI, start, mid, res); err != nil {
				return err
			}
			return fetchListedLogWindow(ctx, client, contract, listedID, parsedABI, mid+1, end, res)
		}
		return fmt.Errorf("FilterLogs [%d,%d]: %w", start, end, err)
	}
	for _, lg := range logs {
		ev, err := parseListed(parsedABI, lg)
		if err != nil {
			return fmt.Errorf("parse Listed tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
		}
		*res = append(*res, ev)
	}
	return nil
}

func FetchListedByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]ListedEvent, error) {
	if toBlock < fromBlock {
		return []ListedEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(blindBoxListedABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	listedID := parsedABI.Events["Listed"].ID

	res := make([]ListedEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}
		if err := fetchListedLogWindow(ctx, client, contract, listedID, parsedABI, start, end, &res); err != nil {
			return nil, err
		}
	}

	// ✅ HeaderByNumber 也最多重试 5 次
	if err := fillBlockTimesListed(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

func PollListedIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []ListedEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	lastProcessedFromDB = bumpLastToScanBaseIfNeeded(lastProcessedFromDB)

	deployListed := conf.GetDeployBlockNftListed(DeployBlockNFTListed)
	from := lastProcessedFromDB + 1
	if from < deployListed {
		from = deployListed
	}
	if safeTo < from {
		return []ListedEvent{}, lastProcessedFromDB, nil
	}

	to := from + MaxBlocksPerCronTick - 1
	if to > safeTo {
		to = safeTo
	}

	evs, err := FetchListedByRange(ctx, client, currentNftContract(), from, to)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, capNewLastWhenNoEvents(len(evs), from, to), nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	DeployBlockNFTUnlisted uint64 = 73530659 // TODO: 改成你想开始监听的块高
)

const blindBoxUnlistedABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "operator", "type": "address"},
      {"indexed": true,  "internalType": "uint256", "name": "tokenId", "type": "uint256"}
    ],
    "name": "Unlisted",
    "type": "event"
  }
]`

type UnlistedEvent struct {
	BlockNumber uint64
	LogIndex    uint
	BlockTime   uint64 // 秒级

	// indexed
	Operator common.Address
	TokenID  *big.Int
}

func parseUnlisted(a abi.ABI, lg types.Log) (UnlistedEvent, error) {
	ev, ok := a.Events["Unlisted"]
	if !ok {
		return UnlistedEvent{}, fmt.Errorf("event Unlisted not found in ABI")
	}
	// topics: [eventId, operator, tokenId]
	if len(lg.Topics) != 3 {
		return UnlistedEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return UnlistedEvent{}, fmt.Errorf("topic0 mismatch")
	}

	return UnlistedEvent{
		BlockNumber: lg.BlockNumber,
		LogIndex:    lg.Index,
		Operator:    common.BytesToAddress(lg.Topics[1].Bytes()),
		TokenID:     new(big.Int).SetBytes(lg.Topics[2].Bytes()),
	}, nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func filterLogsRetryUnlisted(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		rpcThrottle()
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func headerByNumberRetryUnlisted(ctx context.Context, client *ethclient.Client, bn uint64) (*types.Header, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err == nil {
			return h, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

func fillBlockTimesUnlisted(ctx context.Context, client *ethclient.Client, evs []UnlistedEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}

		h, err := headerByNumberRetryUnlisted(ctx, client, bn)
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}

		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

func fetchUnlistedLogWindow(ctx context.Context, client *ethclient.Client, contract common.Address, unlistedID common.Hash, parsedABI abi.ABI, start, end uint64, res *[]UnlistedEvent) error {
	if end < start {
		return nil
	}
	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(end),
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{unlistedID}},
	}
	logs, err := filterLogsRetryUnlisted(ctx, client, q)
	if err != nil {
		if isFilterLogsLimitErr(err) && start < end {
			mid := start + (end-start)/2
			if err := fetchUnlistedLogWindow(ctx, client, contract, unlistedID, parsedABI, start, mid, res); err != nil {
				return err
			}
			return fetchUnlistedLogWindow(ctx, client, contract, unlistedID, parsedABI, mid+1, end, res)
		}
		return fmt.Errorf("FilterLogs [%d,%d]: %w", start, end, err)
	}
	for _, lg := range logs {
		ev, err := parseUnlisted(parsedABI, lg)
		if err != nil {
			return fmt.Errorf("parse Unlisted tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
		}
		*res = append(*res, ev)
	}
	return nil
}

func FetchUnlistedByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]UnlistedEvent, error) {
	if toBlock < fromBlock {
		return []UnlistedEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(blindBoxUnlistedABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	unlistedID := parsedABI.Events["Unlisted"].ID

	res := make([]UnlistedEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}
		if err := fetchUnlistedLogWindow(ctx, client, contract, unlistedID, parsedABI, start, end, &res); err != nil {
			return nil, err
		}
	}

	// ✅ HeaderByNumber 也最多重试 5 次
	if err := fillBlockTimesUnlisted(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

func PollUnlistedIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []UnlistedEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	lastProcessedFromDB = bumpLastToScanBaseIfNeeded(lastProcessedFromDB)

	from := lastProcessedFromDB + 1
	if from < DeployBlockNFTUnlisted {
		from = DeployBlockNFTUnlisted
	}
	if safeTo < from {
		return []UnlistedEvent{}, lastProcessedFromDB, nil
	}

	to := from + MaxBlocksPerCronTick - 1
	if to > safeTo {
		to = safeTo
	}

	evs, err := FetchUnlistedByRange(ctx, client, currentNftContract(), from, to)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, capNewLastWhenNoEvents(len(evs), from, to), nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/* =========================
   Minted (铸造) 事件
   ========================= */

const (
	DeployBlockNFTMint uint64 = 72539344
)

const blindBoxMintedABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "to", "type": "address"},
      {"indexed": true,  "internalType": "uint256", "name": "tokenId", "type": "uint256"},
      {"indexed": true,  "internalType": "uint8",   "name": "tier", "type": "uint8"},
      {"indexed": false, "internalType": "uint256", "name": "usdtPaid", "type": "uint256"}
    ],
    "name": "Minted",
    "type": "event"
  }
]`

type MintedEvent struct {
	BlockNumber uint64
	LogIndex    uint
	BlockTime   uint64 // 秒级

	// indexed
	To      common.Address
	TokenID *big.Int

	// data
	Tier     uint8
	UsdtPaid *big.Int
}

func parseMinted(a abi.ABI, lg types.Log) (MintedEvent, error) {
	ev, ok := a.Events["Minted"]
	if !ok {
		return MintedEvent{}, fmt.Errorf("event Minted not found in ABI")
	}
	// topics: [eventId, to, tokenId, tier]
	if len(lg.Topics) != 4 {
		return MintedEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return MintedEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := MintedEvent{
		BlockNumber: lg.BlockNumber,
		LogIndex:    lg.Index,

		To:      common.BytesToAddress(lg.Topics[1].Bytes()),
		TokenID: new(big.Int).SetBytes(lg.Topics[2].Bytes()),
	}

	// tier 是 indexed，topic 里是 32 bytes uint
	tierBI := new(big.Int).SetBytes(lg.Topics[3].Bytes())
	if tierBI.Sign() < 0 || tierBI.BitLen() > 8 {
		return MintedEvent{}, fmt.Errorf("tier out of uint8: %s", tierBI.String())
	}
	out.Tier = uint8(tierBI.Uint64())

	// 非 indexed data：usdtPaid
	vals, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return MintedEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	if len(vals) != 1 {
		return MintedEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(vals))
	}

	usdtPaid, ok := vals[0].(*big.Int)
	if !ok {
		return MintedEvent{}, fmt.Errorf("bad type usdtPaid: %T", vals[0])
	}
	out.UsdtPaid = usdtPaid

	return out, nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func filterLogsRetryMinted(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		rpcThrottle()
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// fetchMintedLogWindow 拉取 [start,end] Minted 日志；遇 limit exceeded 二分（与 Purchased/Opened 一致）
func fetchMintedLogWindow(ctx context.Context, client *ethclient.Client, contract common.Address, mintedID common.Hash, parsedABI abi.ABI, start, end uint64, res *[]MintedEvent) error {
	if end < start {
		return nil
	}
	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(end),
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{mintedID}},
	}
	logs, err := filterLogsRetryMinted(ctx, client, q)
	if err != nil {
		if isFilterLogsLimitErr(err) && start < end {
			mid := start + (end-start)/2
			if err := fetchMintedLogWindow(ctx, client, contract, mintedID, parsedABI, start, mid, res); err != nil {
				return err
			}
			return fetchMintedLogWindow(ctx, client, contract, mintedID, parsedABI, mid+1, end, res)
		}
		return fmt.Errorf("FilterLogs [%d,%d]: %w", start, end, err)
	}
	for _, lg := range logs {
		ev, err := parseMinted(parsedABI, lg)
		if err != nil {
			return fmt.Errorf("parse Minted tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
		}
		*res = append(*res, ev)
	}
	return nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func headerByNumberRetryMinted(ctx context.Context, client *ethclient.Client, bn uint64) (*types.Header, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err == nil {
			return h, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

func fillBlockTimesMinted(ctx context.Context, client *ethclient.Client, evs []MintedEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}

		h, err := headerByNumberRetryMinted(ctx, client, bn)
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}

		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

// QueryStepMinted 铸造事件单次 FilterLogs 区块数（Binance 节点约 128 较稳）
const QueryStepMinted uint64 = 128

// QueryStepOpened 单次 FilterLogs 块窗；专用/高配节点可调大以减少 RPC 次数，遇 limit exceeded 仍会二分缩小
const QueryStepOpened uint64 = 1024

// FetchMintedByRange 按 [fromBlock, toBlock] 拉取 Minted，并填充 BlockTime(秒)
func FetchMintedByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]MintedEvent, error) {
	if toBlock < fromBlock {
		return []MintedEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(blindBoxMintedABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	mintedID := parsedABI.Events["Minted"].ID

	res := make([]MintedEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStepMinted {
		end := start + QueryStepMinted - 1
		if end > toBlock {
			end = toBlock
		}
		if err := fetchMintedLogWindow(ctx, client, contract, mintedID, parsedABI, start, end, &res); err != nil {
			return nil, err
		}
	}

	// ✅ HeaderByNumber 也最多重试 5 次
	if err := fillBlockTimesMinted(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

// PollMintedIncremental：增量拉取 + latest-Confirmations。maxSpan 可选：>0 时覆盖单次最大块窗（如 GetBoxMintEventFix 需更宽窗）。
func PollMintedIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64, maxSpan ...uint64) (events []MintedEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	lastProcessedFromDB = bumpLastToScanBaseIfNeeded(lastProcessedFromDB)

	deployBlock := conf.GetDeployBlockNftMint(DeployBlockNFTMint)
	from := lastProcessedFromDB + 1
	if from < deployBlock {
		from = deployBlock
	}
	if safeTo < from {
		return []MintedEvent{}, lastProcessedFromDB, nil
	}

	span := MaxBlocksPerCronTick
	if len(maxSpan) > 0 && maxSpan[0] > 0 {
		span = maxSpan[0]
	}
	to := from + span - 1
	if to > safeTo {
		to = safeTo
	}

	evs, err := FetchMintedByRange(ctx, client, currentNftContract(), from, to)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, capNewLastWhenNoEvents(len(evs), from, to), nil
}

// SyncMintedFromContract 从合约 view 函数同步铸造数据到 nft_minted，绕过 eth_getLogs 限制
func SyncMintedFromContract(ctx context.Context, ac *biz.AppUsecase, client *ethclient.Client) (synced int, err error) {
	nftAddr := currentNftContract()
	instance, err := NewNft(nftAddr, client)
	if err != nil {
		return 0, fmt.Errorf("NewNft: %w", err)
	}
	total, err := instance.MintedTokenIdsLength(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, fmt.Errorf("MintedTokenIdsLength: %w", err)
	}
	if total.Sign() <= 0 {
		return 0, nil
	}
	const batchSize = 50
	synced = 0
	for start := big.NewInt(0); start.Cmp(total) < 0; start.Add(start, big.NewInt(batchSize)) {
		count := batchSize
		if remaining := new(big.Int).Sub(total, start); remaining.Cmp(big.NewInt(batchSize)) < 0 {
			count = int(remaining.Int64())
		}
		boxes, err := instance.GetMintedByPage(&bind.CallOpts{Context: ctx}, start, big.NewInt(int64(count)))
		if err != nil {
			return synced, fmt.Errorf("GetMintedByPage start=%s: %w", start.String(), err)
		}
		tokenIds := make([]uint64, 0, len(boxes))
		for _, b := range boxes {
			if b.TokenId != nil {
				tokenIds = append(tokenIds, b.TokenId.Uint64())
			}
		}
		if len(tokenIds) == 0 {
			continue
		}
		existing, err := ac.GetNftMintedByTokenIds(ctx, tokenIds)
		if err != nil {
			return synced, fmt.Errorf("GetNftMintedByTokenIds: %w", err)
		}
		for _, b := range boxes {
			if b.TokenId == nil {
				continue
			}
			tid := b.TokenId.Uint64()
			if existing[tid] != nil {
				continue
			}
			minter, err := callMinterOf(ctx, client, tid)
			if err != nil {
				return synced, fmt.Errorf("minterOf(%d): %w", tid, err)
			}
			if err := ac.InsertNftMinted(ctx, &biz.NftMinted{
				BlockNumber: 0,
				BlockTime:   b.MintedAt,
				LogIndex:    0,
				ToAddr:      minter,
				MintAddr:    minter,
				TokenID:     tid,
				Tier:        uint64(b.Tier),
				UsdtPaid:    toFloat(b.UsdtPaid),
			}); err != nil {
				return synced, fmt.Errorf("InsertNftMinted token=%d: %w", tid, err)
			}
			synced++
		}
	}
	return synced, nil
}

func toFloat(b *big.Int) float64 {
	if b == nil {
		return 0
	}
	f, _ := new(big.Float).SetInt(b).Quo(new(big.Float).SetInt(b), big.NewFloat(1e18)).Float64()
	return f
}

const minterOfABI = `[{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"minterOf","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`

func callMinterOf(ctx context.Context, client *ethclient.Client, tokenId uint64) (string, error) {
	parsed, err := abi.JSON(strings.NewReader(minterOfABI))
	if err != nil {
		return "", err
	}
	input, err := parsed.Pack("minterOf", new(big.Int).SetUint64(tokenId))
	if err != nil {
		return "", err
	}
	nftAddr := currentNftContract()
	msg := ethereum.CallMsg{To: &nftAddr, Data: input}
	out, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return "", err
	}
	if len(out) == 0 {
		return "", fmt.Errorf("empty result")
	}
	unpacked, err := parsed.Unpack("minterOf", out)
	if err != nil {
		return "", err
	}
	if len(unpacked) < 1 {
		return "", fmt.Errorf("no output")
	}
	converted := abi.ConvertType(unpacked[0], new(common.Address))
	if converted == nil {
		return "", fmt.Errorf("convert minterOf result failed")
	}
	addr, ok := converted.(*common.Address)
	if !ok || addr == nil {
		return "", fmt.Errorf("invalid minterOf address type")
	}
	return addr.Hex(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	DeployBlockNFTOpened uint64 = 73406511 // TODO: 按你的实际部署/开启监听块高改
)

const blindBoxOpenedABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "owner", "type": "address"},
      {"indexed": true,  "internalType": "uint256", "name": "tokenId", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "timestamp", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "usdtFeePaid", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "feeBPaid", "type": "uint256"}
    ],
    "name": "Opened",
    "type": "event"
  }
]`

type OpenedEvent struct {
	BlockNumber uint64
	LogIndex    uint
	BlockTime   uint64 // 秒级（区块时间）

	// indexed
	Owner   common.Address
	TokenID *big.Int

	// data
	Timestamp   *big.Int
	UsdtFeePaid *big.Int
	FeeBPaid    *big.Int
}

func parseOpened(a abi.ABI, lg types.Log) (OpenedEvent, error) {
	ev, ok := a.Events["Opened"]
	if !ok {
		return OpenedEvent{}, fmt.Errorf("event Opened not found in ABI")
	}
	// topics: [eventId, owner, tokenId]
	if len(lg.Topics) != 3 {
		return OpenedEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return OpenedEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := OpenedEvent{
		BlockNumber: lg.BlockNumber,
		LogIndex:    lg.Index,

		Owner:   common.BytesToAddress(lg.Topics[1].Bytes()),
		TokenID: new(big.Int).SetBytes(lg.Topics[2].Bytes()),
	}

	vals, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return OpenedEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	// data: timestamp, usdtFeePaid, feeBPaid (新) 或 timestamp, bnbPaid (旧兼容)
	if len(vals) == 3 {
		ts, ok := vals[0].(*big.Int)
		if !ok {
			return OpenedEvent{}, fmt.Errorf("bad type timestamp: %T", vals[0])
		}
		usdtFee, ok := vals[1].(*big.Int)
		if !ok {
			return OpenedEvent{}, fmt.Errorf("bad type usdtFeePaid: %T", vals[1])
		}
		feeB, ok := vals[2].(*big.Int)
		if !ok {
			return OpenedEvent{}, fmt.Errorf("bad type feeBPaid: %T", vals[2])
		}
		out.Timestamp = ts
		out.UsdtFeePaid = usdtFee
		out.FeeBPaid = feeB
	} else if len(vals) == 2 {
		ts, ok := vals[0].(*big.Int)
		if !ok {
			return OpenedEvent{}, fmt.Errorf("bad type timestamp: %T", vals[0])
		}
		bnbPaid, ok := vals[1].(*big.Int)
		if !ok {
			return OpenedEvent{}, fmt.Errorf("bad type bnbPaid: %T", vals[1])
		}
		out.Timestamp = ts
		out.UsdtFeePaid = bnbPaid
		out.FeeBPaid = big.NewInt(0)
	} else {
		return OpenedEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(vals))
	}

	return out, nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func filterLogsRetryOpened(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		rpcThrottle()
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

func isFilterLogsLimitErr(err error) bool {
	if err == nil {
		return false
	}
	s := strings.ToLower(err.Error())
	return (strings.Contains(s, "limit") && strings.Contains(s, "exceed")) ||
		strings.Contains(s, "returned more than") ||
		strings.Contains(s, "too many")
}

// filterLogsRetryGeneric 供 Bought/Sold/Swap/Reward 等共用，避免单次 FilterLogs 失败后无重试
func filterLogsRetryGeneric(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	const maxTry = 5
	const delay = 400 * time.Millisecond
	var lastErr error
	for i := 0; i < maxTry; i++ {
		rpcThrottle()
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// fetchOpenedLogWindow 拉取 [start,end] 内 Opened 日志；遇 limit exceeded 时二分区间（与 Mint 的公共节点策略一致）
func fetchOpenedLogWindow(ctx context.Context, client *ethclient.Client, contract common.Address, openedID common.Hash, parsedABI abi.ABI, start, end uint64, res *[]OpenedEvent) error {
	if end < start {
		return nil
	}
	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(end),
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{openedID}},
	}
	logs, err := filterLogsRetryOpened(ctx, client, q)
	if err != nil {
		if isFilterLogsLimitErr(err) && start < end {
			mid := start + (end-start)/2
			if err := fetchOpenedLogWindow(ctx, client, contract, openedID, parsedABI, start, mid, res); err != nil {
				return err
			}
			return fetchOpenedLogWindow(ctx, client, contract, openedID, parsedABI, mid+1, end, res)
		}
		return fmt.Errorf("FilterLogs [%d,%d]: %w", start, end, err)
	}
	for _, lg := range logs {
		ev, err := parseOpened(parsedABI, lg)
		if err != nil {
			return fmt.Errorf("parse log tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
		}
		*res = append(*res, ev)
	}
	return nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func headerByNumberRetryOpened(ctx context.Context, client *ethclient.Client, bn uint64) (*types.Header, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err == nil {
			return h, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

func fillBlockTimesOpened(ctx context.Context, client *ethclient.Client, evs []OpenedEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}

		h, err := headerByNumberRetryOpened(ctx, client, bn)
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}

		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

func FetchOpenedByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]OpenedEvent, error) {
	if toBlock < fromBlock {
		return []OpenedEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(blindBoxOpenedABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	openedID := parsedABI.Events["Opened"].ID

	res := make([]OpenedEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStepOpened {
		end := start + QueryStepOpened - 1
		if end > toBlock {
			end = toBlock
		}
		if err := fetchOpenedLogWindow(ctx, client, contract, openedID, parsedABI, start, end, &res); err != nil {
			return nil, err
		}
	}

	// ✅ HeaderByNumber 也最多重试 5 次
	if err := fillBlockTimesOpened(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

func PollOpenedIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []OpenedEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	lastProcessedFromDB = bumpLastToScanBaseIfNeeded(lastProcessedFromDB)

	from := lastProcessedFromDB + 1
	if from < DeployBlockNFTOpened {
		from = DeployBlockNFTOpened
	}
	if safeTo < from {
		return []OpenedEvent{}, lastProcessedFromDB, nil
	}

	to := from + MaxBlocksPerCronTick - 1
	if to > safeTo {
		to = safeTo
	}

	evs, err := FetchOpenedByRange(ctx, client, currentNftContract(), from, to)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, capNewLastWhenNoEvents(len(evs), from, to), nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	DeployBlockNFTTransfer uint64 = 72530801 // TODO: 按你的实际部署/开启监听块高改
)

const erc721TransferABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "from", "type": "address"},
      {"indexed": true,  "internalType": "address", "name": "to",   "type": "address"},
      {"indexed": true,  "internalType": "uint256", "name": "tokenId", "type": "uint256"}
    ],
    "name": "Transfer",
    "type": "event"
  }
]`

type TransferEvent struct {
	BlockNumber uint64
	LogIndex    uint
	BlockTime   uint64 // 秒级（区块时间）

	// indexed
	From    common.Address
	To      common.Address
	TokenID *big.Int
}

func parseTransfer(a abi.ABI, lg types.Log) (TransferEvent, error) {
	ev, ok := a.Events["Transfer"]
	if !ok {
		return TransferEvent{}, fmt.Errorf("event Transfer not found in ABI")
	}
	// topics: [eventId, from, to, tokenId]
	if len(lg.Topics) != 4 {
		return TransferEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return TransferEvent{}, fmt.Errorf("topic0 mismatch")
	}

	return TransferEvent{
		BlockNumber: lg.BlockNumber,
		LogIndex:    lg.Index,

		From:    common.BytesToAddress(lg.Topics[1].Bytes()),
		To:      common.BytesToAddress(lg.Topics[2].Bytes()),
		TokenID: new(big.Int).SetBytes(lg.Topics[3].Bytes()),
	}, nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func filterLogsRetry(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		rpcThrottle()
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func headerByNumberRetry(ctx context.Context, client *ethclient.Client, bn uint64) (*types.Header, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err == nil {
			return h, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

func fillBlockTimesTransfer(ctx context.Context, client *ethclient.Client, evs []TransferEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}

		h, err := headerByNumberRetry(ctx, client, bn)
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}

		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

func fetchTransferLogWindow(ctx context.Context, client *ethclient.Client, contract common.Address, transferID common.Hash, parsedABI abi.ABI, start, end uint64, res *[]TransferEvent) error {
	if end < start {
		return nil
	}
	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(end),
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{transferID}},
	}
	logs, err := filterLogsRetry(ctx, client, q)
	if err != nil {
		if isFilterLogsLimitErr(err) && start < end {
			mid := start + (end-start)/2
			if err := fetchTransferLogWindow(ctx, client, contract, transferID, parsedABI, start, mid, res); err != nil {
				return err
			}
			return fetchTransferLogWindow(ctx, client, contract, transferID, parsedABI, mid+1, end, res)
		}
		return fmt.Errorf("FilterLogs Transfer [%d,%d]: %w", start, end, err)
	}
	for _, lg := range logs {
		ev, err := parseTransfer(parsedABI, lg)
		if err != nil {
			return fmt.Errorf("parse Transfer tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
		}
		*res = append(*res, ev)
	}
	return nil
}

func FetchTransferByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]TransferEvent, error) {
	if toBlock < fromBlock {
		return []TransferEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(erc721TransferABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	transferID := parsedABI.Events["Transfer"].ID

	res := make([]TransferEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}
		if err := fetchTransferLogWindow(ctx, client, contract, transferID, parsedABI, start, end, &res); err != nil {
			return nil, err
		}
	}

	// ✅ HeaderByNumber 也最多重试 5 次
	if err := fillBlockTimesTransfer(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

func PollTransferIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []TransferEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	lastProcessedFromDB = bumpLastToScanBaseIfNeeded(lastProcessedFromDB)

	from := lastProcessedFromDB + 1
	deployTransfer := conf.GetDeployBlockNftTransfer(DeployBlockNFTTransfer)
	if from < deployTransfer {
		from = deployTransfer
	}
	if safeTo < from {
		return []TransferEvent{}, lastProcessedFromDB, nil
	}

	to := from + MaxBlocksPerCronTick - 1
	if to > safeTo {
		to = safeTo
	}

	evs, err := FetchTransferByRange(ctx, client, currentNftContract(), from, to)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, capNewLastWhenNoEvents(len(evs), from, to), nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	DeployBlockRegistered uint64 = 72530801 // TODO: 改成触发 Registered 的合约实际部署块高
)

var (
	RegisterContract = common.HexToAddress(conf.GetContract("register"))
)

const registeredABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true, "internalType": "address", "name": "user",   "type": "address"},
      {"indexed": true, "internalType": "address", "name": "parent", "type": "address"},
      {"indexed": true, "internalType": "address", "name": "top",    "type": "address"}
    ],
    "name": "Registered",
    "type": "event"
  }
]`

type RegisteredEvent struct {
	BlockNumber uint64
	LogIndex    uint
	BlockTime   uint64 // 秒级（区块时间）

	// indexed
	User   common.Address
	Parent common.Address
	Top    common.Address
}

func parseRegistered(a abi.ABI, lg types.Log) (RegisteredEvent, error) {
	ev, ok := a.Events["Registered"]
	if !ok {
		return RegisteredEvent{}, fmt.Errorf("event Registered not found in ABI")
	}
	// topics: [eventId, user, parent, top]
	if len(lg.Topics) != 4 {
		return RegisteredEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return RegisteredEvent{}, fmt.Errorf("topic0 mismatch")
	}

	return RegisteredEvent{
		BlockNumber: lg.BlockNumber,
		LogIndex:    lg.Index,
		User:        common.BytesToAddress(lg.Topics[1].Bytes()),
		Parent:      common.BytesToAddress(lg.Topics[2].Bytes()),
		Top:         common.BytesToAddress(lg.Topics[3].Bytes()),
	}, nil
}

// ✅ 简单重试：最多 5 次（局部变量）
func filterLogsRetryRegistered(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		rpcThrottle()
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// ✅ 简单重试：最多 5 次（局部变量）
func headerByNumberRetryRegistered(ctx context.Context, client *ethclient.Client, bn uint64) (*types.Header, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err == nil {
			return h, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

func fillBlockTimesRegistered(ctx context.Context, client *ethclient.Client, evs []RegisteredEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}

		h, err := headerByNumberRetryRegistered(ctx, client, bn)
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}

		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

func fetchRegisteredLogWindow(ctx context.Context, client *ethclient.Client, contract common.Address, evID common.Hash, parsedABI abi.ABI, start, end uint64, res *[]RegisteredEvent) error {
	if end < start {
		return nil
	}
	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(start),
		ToBlock:   new(big.Int).SetUint64(end),
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{evID}},
	}
	logs, err := filterLogsRetryRegistered(ctx, client, q)
	if err != nil {
		if isFilterLogsLimitErr(err) && start < end {
			mid := start + (end-start)/2
			if err := fetchRegisteredLogWindow(ctx, client, contract, evID, parsedABI, start, mid, res); err != nil {
				return err
			}
			return fetchRegisteredLogWindow(ctx, client, contract, evID, parsedABI, mid+1, end, res)
		}
		return fmt.Errorf("FilterLogs Registered [%d,%d]: %w", start, end, err)
	}
	for _, lg := range logs {
		ev, err := parseRegistered(parsedABI, lg)
		if err != nil {
			return fmt.Errorf("parse Registered tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
		}
		*res = append(*res, ev)
	}
	return nil
}

func FetchRegisteredByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]RegisteredEvent, error) {
	if toBlock < fromBlock {
		return []RegisteredEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(registeredABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	evID := parsedABI.Events["Registered"].ID

	res := make([]RegisteredEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}
		if err := fetchRegisteredLogWindow(ctx, client, contract, evID, parsedABI, start, end, &res); err != nil {
			return nil, err
		}
	}

	if err := fillBlockTimesRegistered(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

func FetchLatestRegisteredByUser(ctx context.Context, client *ethclient.Client, contract common.Address, user common.Address, fromBlock uint64) (*RegisteredEvent, error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, nil
	}
	safeTo := head - Confirmations
	parsedABI, err := abi.JSON(strings.NewReader(registeredABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	evID := parsedABI.Events["Registered"].ID
	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(safeTo),
		Addresses: []common.Address{contract},
		Topics: [][]common.Hash{
			{evID},
			{common.BytesToHash(user.Bytes())},
		},
	}
	logs, err := filterLogsRetryRegistered(ctx, client, q)
	if err != nil {
		return nil, err
	}
	if len(logs) == 0 {
		return nil, nil
	}
	last := logs[len(logs)-1]
	ev, err := parseRegistered(parsedABI, last)
	if err != nil {
		return nil, err
	}
	h, err := headerByNumberRetryRegistered(ctx, client, ev.BlockNumber)
	if err != nil {
		return nil, err
	}
	ev.BlockTime = h.Time
	return &ev, nil
}

func PollRegisteredIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []RegisteredEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	lastProcessedFromDB = bumpLastToScanBaseIfNeeded(lastProcessedFromDB)

	from := lastProcessedFromDB + 1
	if from < DeployBlockRegistered {
		from = DeployBlockRegistered
	}
	if safeTo < from {
		return []RegisteredEvent{}, lastProcessedFromDB, nil
	}

	to := from + MaxBlocksPerCronTick - 1
	if to > safeTo {
		to = safeTo
	}

	evs, err := FetchRegisteredByRange(ctx, client, RegisterContract, from, to)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, capNewLastWhenNoEvents(len(evs), from, to), nil
}
