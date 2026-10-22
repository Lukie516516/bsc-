package biz

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	pb "requestEth/api/requestEth/v1"

	"requestEth/internal/conf"
)

type SwapTrade struct {
	ID              uint64
	BlockNumber     uint64
	LogIndex        uint32
	BlockTime       uint64
	Sender          string
	ToAddr          string
	Side            uint8
	Amount0In       float64
	Amount1In       float64
	Amount0OutNet   float64
	Amount1OutGross float64
	Amount0OutGross float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type PrimaryBuy struct {
	ID           uint64
	BlockNumber  uint64
	BlockTime    uint64
	LogIndex     uint32
	Buyer        string
	ToAddr       string
	UsdtUsed     float64
	AusdGrossOut float64
	AusdFee      float64
	AusdNetOut   float64
	PriceBefore  float64
	PriceAfter   float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PrimarySell struct {
	ID          uint64
	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint32
	Seller      string
	ToAddr      string
	AusdGrossIn float64
	AusdFee     float64
	AusdBurn    float64
	UsdtOut     float64
	PriceBefore float64
	PriceAfter  float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RewardNotified struct {
	ID               uint64
	BlockNumber      uint64
	BlockTime        uint64
	LogIndex         uint32
	User             string
	L1               string
	L2               string
	Profit           float64
	UserShare        float64
	Top              string
	Pool             float64
	UplinePortionBps uint64
	ToL1             float64
	ToL2             float64
	ToTop            float64
	ToProject        float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type NftMarketPurchase struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	Buyer   string
	Seller  string
	TokenID uint64

	PriceUSDT  float64
	FeePaidInB uint8
	FeeUSDT    float64
	FeeB       float64

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CheckStatus uint64
}

type RewardDetail struct {
	ID         uint64
	User       string
	Amount     float64
	Reason     uint64
	NotifiedId uint64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	BlockTime  uint64
}

type NftMinted struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	ToAddr   string
	MintAddr string
	TokenID  uint64

	Tier     uint64
	UsdtPaid float64

	Status   uint8
	ListedAt uint64

	OpenStatus uint8
	OpenedAt   uint64

	// 链上 boxInfo（GetAddressBox 合并 RPC 写入；与 DB 无关）
	ChainUsdtPaid float64 // usdtPaid / 1e18
	BoxRewardUsdt float64 // reward / 1e18（int256）
	RewardSetAt   uint64  // 0=链上未结算

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CheckStatus uint64
	CheckTime   uint64
	SourceType  string
}

type UserProfitSummary struct {
	Address          string
	DailyOpenProfit  string
	TotalOpenProfit  string
	TotalMintProfit  string
	SettledOpenCount uint64
	TodayOpenCount   uint64
}

type UserRelationSummary struct {
	Address    string
	ParentAddr string
	TopAddr    string
	Registered bool
}

type RuntimeConfigEntry struct {
	Key         string
	Value       string
	Description string
}

type QuantFundStatusSummary struct {
	Status      uint8
	LedgerCount int64
	TotalAmount string
}

type QuantFundLedgerView struct {
	ID               uint64
	MintedID         uint64
	TokenID          uint64
	UserAddr         string
	MintAddr         string
	UsdtPaid         string
	QuantBps         int
	QuantAmount      string
	ReserveAmount    string
	BlockNumber      uint64
	BlockTime        uint64
	Status           uint8
	TransferTaskNo   string
	TransferTxHash   string
	ResultTaskNo     string
	SettlementTxHash string
	Remark           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type QuantTransferTaskView struct {
	ID          uint64
	TaskNo      string
	LedgerCount int
	QuantBps    int
	TotalAmount string
	FromAddr    string
	ToAddr      string
	TxHash      string
	Status      uint8
	Remark      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type QuantResultStatusSummary struct {
	Status      uint8
	ResultCount int64
	ProfitTotal string
	ReturnTotal string
}

type QuantProfitResultView struct {
	ID               uint64
	TaskNo           string
	LedgerID         uint64
	TokenID          uint64
	UserAddr         string
	MintAddr         string
	QuantAmount      string
	ProfitBps        int
	ProfitAmount     string
	ReturnAmount     string
	SettlementReward string
	ResultMode       string
	Status           uint8
	SettlementTxHash string
	Remark           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type QuantFundDashboard struct {
	PendingCount        int64
	PendingAmount       string
	LockedCount         int64
	LockedAmount        string
	TransferredCount    int64
	TransferredAmount   string
	FailedCount         int64
	FailedAmount        string
	ResultReadyCount    int64
	ResultReadyAmount   string
	SettledCount        int64
	SettledAmount       string
	SettleFailedCount   int64
	SettleFailedAmount  string
	ResultPendingCount  int64
	ResultPendingProfit string
	ResultSettledCount  int64
	ResultSettledProfit string
	ResultFailedCount   int64
	ResultFailedProfit  string
	RecentLedgers       []*QuantFundLedgerView
	RecentTasks         []*QuantTransferTaskView
	RecentResults       []*QuantProfitResultView
	StatusSummary       []*QuantFundStatusSummary
	ResultSummary       []*QuantResultStatusSummary
}

type NftMarketListed struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	Seller    string
	TokenID   uint64
	Timestamp uint64

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CheckStatus uint64
}

type NftMarketUnlisted struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	Operator string
	TokenID  uint64

	CreatedAt time.Time
	UpdatedAt time.Time

	CheckStatus uint64
}

type NftOpened struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	UserAddr string
	TokenID  uint64

	OpenedAt uint64
	Reward   float64

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CheckStatus uint64
}

// OpenedBoxWithMint 已开盲盒+铸造信息（用于按日利润统计）
type OpenedBoxWithMint struct {
	TokenID   uint64
	UserAddr  string
	Reward    float64
	BlockTime uint64
	UsdtPaid  float64
}

type NftTransfer struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	FromAddr string
	ToAddr   string
	TokenID  uint64

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CheckStatus uint64
}

type UserRegistered struct {
	ID uint64

	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint

	UserAddr   string
	ParentAddr string
	TopAddr    string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepo interface {
	GetSwapTradeLast(ctx context.Context) (*SwapTrade, error)
	GetSwapTrade(ctx context.Context, start, end uint64) ([]*SwapTrade, error)
	InsertSwapTrade(ctx context.Context, iData *SwapTrade) error
	GetPrimaryBuyLast(ctx context.Context) (*PrimaryBuy, error)
	GetPrimaryBuy(ctx context.Context, start, end uint64) ([]*PrimaryBuy, error)
	InsertPrimaryBuy(ctx context.Context, iData *PrimaryBuy) error
	GetPrimarySellLast(ctx context.Context) (*PrimarySell, error)
	GetPrimarySell(ctx context.Context, start, end uint64) ([]*PrimarySell, error)
	InsertPrimarySell(ctx context.Context, iData *PrimarySell) error
	GetRewardNotifiedLast(ctx context.Context) (*RewardNotified, error)
	GetRewardNotified(ctx context.Context, start, end uint64) ([]*RewardNotified, error)
	GetRewardNotifiedByIds(ctx context.Context, ids []uint64) (map[uint64]*RewardNotified, error)
	InsertRewardNotified(ctx context.Context, iData *RewardNotified) error
	// InsertRewardNotifiedScanProgress 仅 reward_notified 游标，不写 reward_detail（空窗扫块）
	InsertRewardNotifiedScanProgress(ctx context.Context, blockNumber uint64) error
	GetUserRewardByUserIdPage(ctx context.Context, b *Pagination, address string, reason uint64) ([]*RewardDetail, error, int64)
	GetNftMarketPurchaseByAddressPage(ctx context.Context, b *Pagination, address string, addressTwo []string, side uint64) ([]*NftMarketPurchase, error, int64)
	GetNftMarketPurchaseLast(ctx context.Context) (*NftMarketPurchase, error)
	InsertNftMarketPurchase(ctx context.Context, iData *NftMarketPurchase) error
	InsertNftMinted(ctx context.Context, iData *NftMinted) error
	UpdateNftMinted(ctx context.Context, tokenId uint64, mintAddr string) error
	GetNftMintedLast(ctx context.Context) (*NftMinted, error)
	InsertNftMarketListed(ctx context.Context, iData *NftMarketListed) error
	GetNftMarketListedLast(ctx context.Context) (*NftMarketListed, error)
	InsertNftMarketUnlisted(ctx context.Context, iData *NftMarketUnlisted) error
	GetNftMarketUnlistedLast(ctx context.Context) (*NftMarketUnlisted, error)
	InsertNftOpened(ctx context.Context, iData *NftOpened) error
	GetNftOpenedLast(ctx context.Context) (*NftOpened, error)
	InsertNftTransfer(ctx context.Context, iData *NftTransfer) error
	GetNftTransferLast(ctx context.Context) (*NftTransfer, error)
	GetNftTransferLastNoCheck(ctx context.Context) ([]*NftTransfer, error)
	GetNftMintedByTokenIds(ctx context.Context, tokenIds []uint64) (map[uint64]*NftMinted, error)
	UpdateNftMintedToAddress(ctx context.Context, id, idT, check uint64, toAddr string) error
	GetNftListLastNoCheck(ctx context.Context) ([]*NftMarketListed, error)
	GetNftUnListLastNoCheck(ctx context.Context) ([]*NftMarketUnlisted, error)
	GetNftBuyLastNoCheck(ctx context.Context) ([]*NftMarketPurchase, error)
	GetNftOpenLastNoCheck(ctx context.Context) ([]*NftOpened, error)
	UpdateNftMintedListStatus(ctx context.Context, id, idT, checkTime uint64) error
	UpdateNftMintedUnListStatus(ctx context.Context, id, idT, checkTime uint64) error
	UpdateNftMintedBuyStatus(ctx context.Context, id, idT, checkTime uint64, buyer string) error
	UpdateNftMintedOpenStatus(ctx context.Context, id, idT, checkTime uint64) error
	UpdateUnlistedCheckStatus(ctx context.Context, id, idT, checkTime uint64) error
	UpdateListedCheckStatus(ctx context.Context, id, idT, checkTime uint64) error
	UpdateBuyCheckStatus(ctx context.Context, id, idT, checkTime uint64) error
	GetNftMintedByAddressPage(ctx context.Context, b *Pagination, start uint64, end uint64, address []string, status uint64, order uint64, tier uint64, openStatus uint64, openAtOrder uint64) ([]*NftMinted, error, int64)
	GetNftMintedByAddressAll(ctx context.Context, start uint64, end uint64, address []string, status uint64, order uint64, tier uint64, openStatus uint64, openAtOrder uint64) ([]*NftMinted, error)
	GetNftMintedPage(ctx context.Context, b *Pagination, order uint64, orderTwo uint64, tier uint64) ([]*NftMinted, error, int64)
	GetUserRegisteredLast(ctx context.Context) (*UserRegistered, error)
	InsertUserRegistered(ctx context.Context, iData *UserRegistered) error
	GetUserRegisteredByAddress(ctx context.Context, address string) (*UserRegistered, error)
	GetUserRCount(ctx context.Context) int64
	GetUserRCountBySe(ctx context.Context, start, end uint64) int64
	GetMintNftCountBySe(ctx context.Context, start, end uint64) int64
	GetMintNftUsdtPaidSumBySe(ctx context.Context, start, end uint64) string
	GetMintNftCount(paidType uint64) int64
	GetMintNftNotOpenCount(paidType uint64) int64
	GetMintNftUsdtPaidSum(paidType uint64) string
	GetMintNftNotOpenUsdtPaidSum(paidType uint64) string
	GetNftBuyCountBySe(ctx context.Context, start, end uint64) int64
	GetNftBuySumBySe(ctx context.Context, start, end uint64) string
	GetNftBuySum() string
	GetNftBuyCount() int64
	GetNftOpenCountBySe(ctx context.Context, start, end uint64) int64
	GetNftOpenSum() string
	GetNftOpenSumBySe(ctx context.Context, start, end uint64) string
	GetOpenedBoxListBySe(ctx context.Context, b *Pagination, start, end uint64) ([]*OpenedBoxWithMint, error, int64)
	GetOpenedBoxProfitLossCountBySe(ctx context.Context, start, end uint64) (profitCount, lossCount int64)
	GetNftMarketPurchaseFeeSumBySe(ctx context.Context, start, end uint64) (feeUsdtSum, feeBSum string)
	GetUserOpenProfitSumBySe(ctx context.Context, address string, start, end uint64) string
	GetUserOpenProfitSum(ctx context.Context, address string) string
	GetUserMintProfitSum(ctx context.Context, address string) string
	GetUserOpenedCountBySe(ctx context.Context, address string, start, end uint64) int64
	GetUserOpenedCount(ctx context.Context, address string) int64
	ListRuntimeConfigs(ctx context.Context, prefix string) ([]*RuntimeConfigEntry, error)
	UpsertRuntimeConfig(ctx context.Context, key, value, description string) error
	GetQuantFundStatusSummary(ctx context.Context) ([]*QuantFundStatusSummary, error)
	GetRecentQuantFundLedgers(ctx context.Context, limit int) ([]*QuantFundLedgerView, error)
	GetRecentQuantTransferTasks(ctx context.Context, limit int) ([]*QuantTransferTaskView, error)
	GetQuantResultStatusSummary(ctx context.Context) ([]*QuantResultStatusSummary, error)
	GetRecentQuantProfitResults(ctx context.Context, limit int) ([]*QuantProfitResultView, error)
}

// AppUsecase is an app usecase.
type AppUsecase struct {
	userRepo UserRepo
	tx       Transaction
	log      *log.Helper
}

// NewAppUsecase new a app usecase.
func NewAppUsecase(userRepo UserRepo, tx Transaction, logger log.Logger) *AppUsecase {
	return &AppUsecase{userRepo: userRepo, tx: tx, log: log.NewHelper(logger)}
}

func (ac *AppUsecase) GetSwapTradeLast(ctx context.Context) (*SwapTrade, error) {
	var (
		rLast *SwapTrade
		err   error
	)
	rLast, err = ac.userRepo.GetSwapTradeLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertSwapTrade(ctx context.Context, trade *SwapTrade) error {
	var (
		err error
	)

	err = ac.userRepo.InsertSwapTrade(ctx, trade)

	return err
}

func (ac *AppUsecase) GetBuyLast(ctx context.Context) (*PrimaryBuy, error) {
	var (
		rLast *PrimaryBuy
		err   error
	)
	rLast, err = ac.userRepo.GetPrimaryBuyLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertBuyTrade(ctx context.Context, trade *PrimaryBuy) error {
	var (
		err error
	)

	err = ac.userRepo.InsertPrimaryBuy(ctx, trade)

	return err
}

func (ac *AppUsecase) GetSellLast(ctx context.Context) (*PrimarySell, error) {
	var (
		rLast *PrimarySell
		err   error
	)
	rLast, err = ac.userRepo.GetPrimarySellLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertSellTrade(ctx context.Context, trade *PrimarySell) error {
	var (
		err error
	)

	err = ac.userRepo.InsertPrimarySell(ctx, trade)

	return err
}

func (ac *AppUsecase) GetRewardLast(ctx context.Context) (*RewardNotified, error) {
	var (
		rLast *RewardNotified
		err   error
	)
	rLast, err = ac.userRepo.GetRewardNotifiedLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertReward(ctx context.Context, trade *RewardNotified) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertRewardNotified(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "分红写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) InsertRewardScanProgress(ctx context.Context, blockNumber uint64) error {
	return ac.userRepo.InsertRewardNotifiedScanProgress(ctx, blockNumber)
}

func (ac *AppUsecase) GetNftMarketPurchaseLast(ctx context.Context) (*NftMarketPurchase, error) {
	var (
		rLast *NftMarketPurchase
		err   error
	)
	rLast, err = ac.userRepo.GetNftMarketPurchaseLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertNftMarketPurchase(ctx context.Context, trade *NftMarketPurchase) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertNftMarketPurchase(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetNftMarketListedLast(ctx context.Context) (*NftMarketListed, error) {
	var (
		rLast *NftMarketListed
		err   error
	)
	rLast, err = ac.userRepo.GetNftMarketListedLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertNftMarketListed(ctx context.Context, trade *NftMarketListed) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertNftMarketListed(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetNftMarketUnlistedLast(ctx context.Context) (*NftMarketUnlisted, error) {
	var (
		rLast *NftMarketUnlisted
		err   error
	)
	rLast, err = ac.userRepo.GetNftMarketUnlistedLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertNftMarketUnlisted(ctx context.Context, trade *NftMarketUnlisted) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertNftMarketUnlisted(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetNftTransferLast(ctx context.Context) (*NftTransfer, error) {
	var (
		rLast *NftTransfer
		err   error
	)
	rLast, err = ac.userRepo.GetNftTransferLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertNftTransfer(ctx context.Context, trade *NftTransfer) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertNftTransfer(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetNftOpenedLast(ctx context.Context) (*NftOpened, error) {
	var (
		rLast *NftOpened
		err   error
	)
	rLast, err = ac.userRepo.GetNftOpenedLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertNftOpened(ctx context.Context, trade *NftOpened) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertNftOpened(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetNftMintedLast(ctx context.Context) (*NftMinted, error) {
	var (
		rLast *NftMinted
		err   error
	)
	rLast, err = ac.userRepo.GetNftMintedLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) GetNftMintedByTokenIds(ctx context.Context, tokenIds []uint64) (map[uint64]*NftMinted, error) {
	return ac.userRepo.GetNftMintedByTokenIds(ctx, tokenIds)
}

func (ac *AppUsecase) UpdateNftMinted(ctx context.Context, tokenId uint64, mintAddr string) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.UpdateNftMinted(ctx, tokenId, mintAddr)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) InsertNftMinted(ctx context.Context, trade *NftMinted) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertNftMinted(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "买盲盒写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetQuantFundDashboard(ctx context.Context, limit int) (*QuantFundDashboard, error) {
	summary, err := ac.userRepo.GetQuantFundStatusSummary(ctx)
	if err != nil {
		return nil, err
	}
	resultSummary, err := ac.userRepo.GetQuantResultStatusSummary(ctx)
	if err != nil {
		return nil, err
	}
	ledgers, err := ac.userRepo.GetRecentQuantFundLedgers(ctx, limit)
	if err != nil {
		return nil, err
	}
	tasks, err := ac.userRepo.GetRecentQuantTransferTasks(ctx, limit)
	if err != nil {
		return nil, err
	}
	results, err := ac.userRepo.GetRecentQuantProfitResults(ctx, limit)
	if err != nil {
		return nil, err
	}

	res := &QuantFundDashboard{
		PendingAmount:       "0",
		LockedAmount:        "0",
		TransferredAmount:   "0",
		FailedAmount:        "0",
		ResultReadyAmount:   "0",
		SettledAmount:       "0",
		SettleFailedAmount:  "0",
		ResultPendingProfit: "0",
		ResultSettledProfit: "0",
		ResultFailedProfit:  "0",
		RecentLedgers:       ledgers,
		RecentTasks:         tasks,
		RecentResults:       results,
		StatusSummary:       summary,
		ResultSummary:       resultSummary,
	}
	for _, item := range summary {
		switch item.Status {
		case 0:
			res.PendingCount = item.LedgerCount
			res.PendingAmount = item.TotalAmount
		case 1:
			res.LockedCount = item.LedgerCount
			res.LockedAmount = item.TotalAmount
		case 2:
			res.TransferredCount = item.LedgerCount
			res.TransferredAmount = item.TotalAmount
		case 3:
			res.FailedCount = item.LedgerCount
			res.FailedAmount = item.TotalAmount
		case 4:
			res.ResultReadyCount = item.LedgerCount
			res.ResultReadyAmount = item.TotalAmount
		case 5:
			res.SettledCount = item.LedgerCount
			res.SettledAmount = item.TotalAmount
		case 6:
			res.SettleFailedCount = item.LedgerCount
			res.SettleFailedAmount = item.TotalAmount
		}
	}
	for _, item := range resultSummary {
		switch item.Status {
		case 0:
			res.ResultPendingCount = item.ResultCount
			res.ResultPendingProfit = item.ProfitTotal
		case 2:
			res.ResultSettledCount = item.ResultCount
			res.ResultSettledProfit = item.ProfitTotal
		case 3:
			res.ResultFailedCount = item.ResultCount
			res.ResultFailedProfit = item.ProfitTotal
		}
	}

	return res, nil
}

func (ac *AppUsecase) GetUserRegisteredLast(ctx context.Context) (*UserRegistered, error) {
	var (
		rLast *UserRegistered
		err   error
	)
	rLast, err = ac.userRepo.GetUserRegisteredLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertUserRegistered(ctx context.Context, trade *UserRegistered) error {
	var (
		err error
	)

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.InsertUserRegistered(ctx, trade)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "注册写入mysql错误")
		return err
	}

	return err
}

func (ac *AppUsecase) GetUserRegisteredByAddress(ctx context.Context, address string) (*UserRegistered, error) {
	return ac.userRepo.GetUserRegisteredByAddress(ctx, address)
}

func (ac *AppUsecase) GetUserRelationSummary(ctx context.Context, address string) *UserRelationSummary {
	if address == "" {
		return &UserRelationSummary{}
	}
	row, err := ac.userRepo.GetUserRegisteredByAddress(ctx, address)
	if err != nil || row == nil {
		return &UserRelationSummary{
			Address:    address,
			ParentAddr: "0x0000000000000000000000000000000000000000",
			TopAddr:    "0x0000000000000000000000000000000000000000",
			Registered: false,
		}
	}
	return &UserRelationSummary{
		Address:    row.UserAddr,
		ParentAddr: row.ParentAddr,
		TopAddr:    row.TopAddr,
		Registered: true,
	}
}

func (ac *AppUsecase) ListRuntimeConfigs(ctx context.Context, prefix string) ([]*RuntimeConfigEntry, error) {
	return ac.userRepo.ListRuntimeConfigs(ctx, prefix)
}

func (ac *AppUsecase) UpsertRuntimeConfig(ctx context.Context, key, value, description string) error {
	return ac.userRepo.UpsertRuntimeConfig(ctx, key, value, description)
}

func (ac *AppUsecase) GetExchangeList(ctx context.Context, req *pb.GetExchangeListRequest) (*pb.GetExchangeListReply, error) {
	res := make([]*pb.GetExchangeListReply_List, 0)

	var (
		swapTrades []*SwapTrade
		err        error
	)

	swapTrades, err = ac.userRepo.GetSwapTrade(ctx, req.Start, req.End)
	if err != nil {
		return nil, err
	}

	for _, v := range swapTrades {

		var tmpPrice float64
		if 1 == v.Side && 0 < v.Amount0OutGross {
			tmpPrice = v.Amount1In / v.Amount0OutGross
		} else if 2 == v.Side && 0 < v.Amount0In {
			tmpPrice = v.Amount1OutGross / v.Amount0In
		} else {
			continue
		}

		res = append(res, &pb.GetExchangeListReply_List{
			BlockTime: v.BlockTime,
			Price:     tmpPrice,
		})
	}

	return &pb.GetExchangeListReply{List: res}, nil
}

func (ac *AppUsecase) GetBuyList(ctx context.Context, req *pb.GetBuyListRequest) (*pb.GetBuyListReply, error) {
	res := make([]*pb.GetBuyListReply_List, 0)

	var (
		buyList []*PrimaryBuy
		err     error
	)

	buyList, err = ac.userRepo.GetPrimaryBuy(ctx, req.Start, req.End)
	if err != nil {
		return nil, err
	}

	for _, v := range buyList {
		res = append(res, &pb.GetBuyListReply_List{
			BlockTime:    v.BlockTime,
			Price:        v.PriceAfter,
			UsdtUse:      v.UsdtUsed,
			AusdGrossOut: v.AusdGrossOut,
		})
	}

	return &pb.GetBuyListReply{List: res}, nil
}

type Pagination struct {
	PageNum  int
	PageSize int
}

func (ac *AppUsecase) GetRewardList(ctx context.Context, req *pb.GetRewardListRequest) (*pb.GetRewardListReply, error) {
	res := make([]*pb.GetRewardListReply_List, 0)

	var (
		userRewards []*RewardDetail
		count       int64
		err         error
	)

	if 0 > req.Reason || 5 < req.Reason || 0 >= len(req.Address) {
		return &pb.GetRewardListReply{
			Count: 0,
			List:  res,
		}, nil
	}

	userRewards, err, count = ac.userRepo.GetUserRewardByUserIdPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	}, req.Address, req.Reason)
	if nil != err {
		return &pb.GetRewardListReply{
			Count: uint64(count),
			List:  res,
		}, err
	}

	tmpIds := make([]uint64, 0)
	for _, vUserReward := range userRewards {
		tmpIds = append(tmpIds, vUserReward.NotifiedId)
	}

	var (
		rn map[uint64]*RewardNotified
	)

	rn, err = ac.userRepo.GetRewardNotifiedByIds(ctx, tmpIds)
	if nil != err {
		return &pb.GetRewardListReply{
			Count: uint64(count),
			List:  res,
		}, err
	}

	for _, vUserReward := range userRewards {
		if _, ok := rn[vUserReward.NotifiedId]; !ok {
			continue
		}

		res = append(res, &pb.GetRewardListReply_List{
			User:      rn[vUserReward.NotifiedId].User,
			Reason:    vUserReward.Reason,
			Amount:    vUserReward.Amount,
			BlockTime: vUserReward.BlockTime,
		})
	}

	return &pb.GetRewardListReply{
		Count: uint64(count),
		List:  res,
	}, nil
}

func (ac *AppUsecase) GetBuyBoxList(ctx context.Context, req *pb.GetBuyBoxListRequest) (*pb.GetBuyBoxListReply, error) {
	res := make([]*pb.GetBuyBoxListReply_List, 0)

	var (
		records []*NftMarketPurchase
		count   int64
		err     error
	)

	records, err, count = ac.userRepo.GetNftMarketPurchaseByAddressPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	}, req.Address, nil, 1)
	if nil != err {
		return &pb.GetBuyBoxListReply{
			Count: uint64(count),
			List:  res,
		}, err
	}

	for _, v := range records {
		res = append(res, &pb.GetBuyBoxListReply_List{
			User:      v.Buyer,
			TokenId:   v.TokenID,
			PriceUsdt: v.PriceUSDT,
			FeeUsdt:   v.FeeUSDT,
			FeeB:      v.FeeB,
			BlockTime: v.BlockTime,
		})
	}

	return &pb.GetBuyBoxListReply{
		Count: uint64(count),
		List:  res,
	}, nil
}

func (ac *AppUsecase) UpdateBox(ctx context.Context, req *pb.UpdateBoxRequest) {

	var (
		res       []*NftTransfer
		resTwo    []*NftMarketListed
		resThree  []*NftMarketUnlisted
		resFour   []*NftMarketPurchase
		resFive   []*NftOpened
		resMinted map[uint64]*NftMinted
		err       error
	)

	tokenIdsMap := make(map[uint64]uint64, 0)

	// 交易
	res, err = ac.userRepo.GetNftTransferLastNoCheck(ctx)
	if nil != err {
		return
	}

	for _, v := range res {
		tokenIdsMap[v.TokenID] = v.TokenID
	}

	// 上架
	resTwo, err = ac.userRepo.GetNftListLastNoCheck(ctx)
	if nil != err {
		return
	}

	for _, v := range resTwo {
		tokenIdsMap[v.TokenID] = v.TokenID
	}

	// 下
	resThree, err = ac.userRepo.GetNftUnListLastNoCheck(ctx)
	if nil != err {
		return
	}

	for _, v := range resThree {
		tokenIdsMap[v.TokenID] = v.TokenID
	}

	// 买
	resFour, err = ac.userRepo.GetNftBuyLastNoCheck(ctx)
	if nil != err {
		return
	}

	for _, v := range resFour {
		tokenIdsMap[v.TokenID] = v.TokenID
	}

	// 开
	resFive, err = ac.userRepo.GetNftOpenLastNoCheck(ctx)
	if nil != err {
		return
	}

	for _, v := range resFive {
		tokenIdsMap[v.TokenID] = v.TokenID
	}

	// 整合
	tokenIds := make([]uint64, 0)
	for _, v := range tokenIdsMap {
		tokenIds = append(tokenIds, v)
	}
	if 0 >= len(tokenIds) {
		return
	}

	resMinted, err = ac.userRepo.GetNftMintedByTokenIds(ctx, tokenIds)
	if nil != err || 0 >= len(resMinted) {
		return
	}

	for _, v := range res {
		if _, ok := resMinted[v.TokenID]; !ok {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].BlockTime {
			continue
		}

		//if v.BlockTime <= resMinted[v.TokenID].CheckTime {
		//	fmt.Println("不是最新状态", v, resMinted[v.TokenID])
		//	continue
		//}

		// 上架时 NFT 转入合约托管（Transfer 的 to 为 NFT 合约），不把 to_addr 改成合约地址，由 Listed 事件分支写 status/listed_at
		tmpCheck := uint64(1)
		nftContract := strings.ToLower(strings.TrimSpace(conf.GetContract("nft")))
		if nftContract != "" && strings.ToLower(v.ToAddr) == nftContract {
			tmpCheck = 0
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.UpdateNftMintedToAddress(ctx, resMinted[v.TokenID].ID, v.ID, tmpCheck, v.ToAddr)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "写入mysql错误")
			return
		}

	}

	for _, v := range resTwo {
		if v.TokenID == 0 {
			continue
		}
		if _, ok := resMinted[v.TokenID]; !ok {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].BlockTime {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].CheckTime {
			ac.userRepo.UpdateListedCheckStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			fmt.Println("不是最新状态2", v, resMinted[v.TokenID])
			continue
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.UpdateNftMintedListStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "写入mysql错误")
			ac.userRepo.UpdateListedCheckStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			return
		}
	}

	for _, v := range resThree {
		if _, ok := resMinted[v.TokenID]; !ok {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].BlockTime {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].CheckTime {
			ac.userRepo.UpdateUnlistedCheckStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			fmt.Println("不是最新状态3", v, resMinted[v.TokenID])
			continue
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.UpdateNftMintedUnListStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			ac.userRepo.UpdateUnlistedCheckStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			fmt.Println(err, "写入mysql错误")
			return
		}
	}

	for _, v := range resFour {
		if _, ok := resMinted[v.TokenID]; !ok {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].BlockTime {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].CheckTime {
			ac.userRepo.UpdateBuyCheckStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			fmt.Println("不是最新状态4", v, resMinted[v.TokenID])
			continue
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.UpdateNftMintedBuyStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime, v.Buyer)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			ac.userRepo.UpdateBuyCheckStatus(ctx, resMinted[v.TokenID].ID, v.ID, v.BlockTime)
			fmt.Println(err, "写入mysql错误")
			return
		}
	}

	for _, v := range resFive {
		if _, ok := resMinted[v.TokenID]; !ok {
			continue
		}

		if v.BlockTime < resMinted[v.TokenID].BlockTime {
			continue
		}

		// opened_at 与链上 Opened 事件 timestamp 一致；仅用区块时间会与 nft_opened.opened_at 不一致
		openedAt := v.OpenedAt
		if openedAt == 0 {
			openedAt = v.BlockTime
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.UpdateNftMintedOpenStatus(ctx, resMinted[v.TokenID].ID, v.ID, openedAt)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "写入mysql错误")
			return
		}
	}
}

func (ac *AppUsecase) GetAddressBox(ctx context.Context, address []string, req *pb.GetAddressBoxRequest) ([]*NftMinted, int64, error) {
	var (
		minted []*NftMinted
		count  int64
		err    error
	)

	minted, err, count = ac.userRepo.GetNftMintedByAddressPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, req.Start, req.End, address, req.Num, req.NumThree, req.NumTwo, req.OpenStatus, req.NumFour)

	return minted, count, err
}

func (ac *AppUsecase) GetAddressBoxAll(ctx context.Context, address []string, req *pb.GetAddressBoxRequest) ([]*NftMinted, error) {
	return ac.userRepo.GetNftMintedByAddressAll(ctx, req.Start, req.End, address, req.Num, req.NumThree, req.NumTwo, req.OpenStatus, req.NumFour)
}

func (ac *AppUsecase) GetUserProfitSummary(ctx context.Context, address string, start, end uint64) *UserProfitSummary {
	if address == "" {
		return &UserProfitSummary{}
	}
	return &UserProfitSummary{
		Address:          address,
		DailyOpenProfit:  ac.userRepo.GetUserOpenProfitSumBySe(ctx, address, start, end),
		TotalOpenProfit:  ac.userRepo.GetUserOpenProfitSum(ctx, address),
		TotalMintProfit:  ac.userRepo.GetUserMintProfitSum(ctx, address),
		SettledOpenCount: uint64(ac.userRepo.GetUserOpenedCount(ctx, address)),
		TodayOpenCount:   uint64(ac.userRepo.GetUserOpenedCountBySe(ctx, address, start, end)),
	}
}

func (ac *AppUsecase) GetNftMintedPage(ctx context.Context, req *pb.GetMarketListRequest) ([]*NftMinted, int64, error) {
	var (
		minted []*NftMinted
		count  int64
		err    error
	)

	minted, err, count = ac.userRepo.GetNftMintedPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, req.Num, req.NumThree, req.NumTwo)

	return minted, count, err
}

func (ac *AppUsecase) GetSellBoxList(ctx context.Context, address []string, req *pb.GetSellBoxListRequest) ([]*NftMarketPurchase, int64, error) {
	var (
		records []*NftMarketPurchase
		count   int64
		err     error
	)

	records, err, count = ac.userRepo.GetNftMarketPurchaseByAddressPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, req.Address, address, 3)

	return records, count, err
}

func (ac *AppUsecase) GetAllInfo(ctx context.Context, req *pb.GetAllInfoRequest) (*pb.GetAllInfoReply, error) {

	return &pb.GetAllInfoReply{
		UserCount:         uint64(ac.userRepo.GetUserRCount(ctx)),
		TodayUserCount:    uint64(ac.userRepo.GetUserRCountBySe(ctx, req.Start, req.End)),
		TodayMintedCount:  uint64(ac.userRepo.GetMintNftCountBySe(ctx, req.Start, req.End)),
		TodayMintedSum:    ac.userRepo.GetMintNftUsdtPaidSumBySe(ctx, req.Start, req.End),
		MintedCount:       uint64(ac.userRepo.GetMintNftCount(0)),
		MintedSum:         ac.userRepo.GetMintNftUsdtPaidSum(0),
		MintedCountOne:    uint64(ac.userRepo.GetMintNftCount(1)),
		MintedSumOne:      ac.userRepo.GetMintNftUsdtPaidSum(1),
		MintedCountTwo:    uint64(ac.userRepo.GetMintNftCount(2)),
		MintedSumTwo:      ac.userRepo.GetMintNftUsdtPaidSum(2),
		MintedCountThree:  uint64(ac.userRepo.GetMintNftCount(3)),
		MintedSumThree:    ac.userRepo.GetMintNftUsdtPaidSum(3),
		BuyCount:          uint64(ac.userRepo.GetNftBuyCount()),
		BuySum:            ac.userRepo.GetNftBuySum(),
		TodayBuyCount:     uint64(ac.userRepo.GetNftBuyCountBySe(ctx, req.Start, req.End)),
		TodayBuySum:       ac.userRepo.GetNftBuySumBySe(ctx, req.Start, req.End),
		OpenReward:        ac.userRepo.GetNftOpenSum(),
		MintedCountNoOpen: uint64(ac.userRepo.GetMintNftNotOpenCount(0)),
		MintedSumNoOpen:   ac.userRepo.GetMintNftNotOpenUsdtPaidSum(0),
	}, nil
}

// GetOpenedBoxProfitStats 按日开盒利润统计：已开盲盒明细与汇总（可筛选时间段）
func (ac *AppUsecase) GetOpenedBoxProfitStats(ctx context.Context, req *pb.GetOpenedBoxProfitStatsRequest) (*pb.GetOpenedBoxProfitStatsReply, error) {
	page, pageSize := int(req.Page), int(req.PageSize)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100
	}
	b := &Pagination{PageNum: page, PageSize: pageSize}

	list, err, total := ac.userRepo.GetOpenedBoxListBySe(ctx, b, req.Start, req.End)
	if err != nil {
		return nil, err
	}
	profitCount, lossCount := ac.userRepo.GetOpenedBoxProfitLossCountBySe(ctx, req.Start, req.End)
	totalProfit := ac.userRepo.GetNftOpenSumBySe(ctx, req.Start, req.End)

	items := make([]*pb.GetOpenedBoxProfitStatsReply_Item, 0, len(list))
	for _, v := range list {
		items = append(items, &pb.GetOpenedBoxProfitStatsReply_Item{
			TokenId:   v.TokenID,
			UserAddr:  v.UserAddr,
			Reward:    fmt.Sprintf("%.18g", v.Reward),
			UsdtPaid:  fmt.Sprintf("%.18g", v.UsdtPaid),
			BlockTime: v.BlockTime,
		})
	}
	return &pb.GetOpenedBoxProfitStatsReply{
		Total:       total,
		TotalProfit: totalProfit,
		ProfitCount: uint64(profitCount),
		LossCount:   uint64(lossCount),
		List:        items,
	}, nil
}

// GetFeeIncomeStats 手续费收入统计：二级市场交易手续费（USDT/HRX）按时间段汇总
func (ac *AppUsecase) GetFeeIncomeStats(ctx context.Context, req *pb.GetFeeIncomeStatsRequest) (*pb.GetFeeIncomeStatsReply, error) {
	feeUsdtSum, feeBSum := ac.userRepo.GetNftMarketPurchaseFeeSumBySe(ctx, req.Start, req.End)
	count := ac.userRepo.GetNftBuyCountBySe(ctx, req.Start, req.End)
	return &pb.GetFeeIncomeStatsReply{
		FeeUsdtSum: feeUsdtSum,
		FeeBSum:    feeBSum,
		Count:      uint64(count),
	}, nil
}
