package data

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"requestEth/internal/biz"
	"time"
)

type SwapTrade struct {
	ID uint64 `gorm:"primarykey;type:bigint unsigned;comment:主键"`

	BlockNumber uint64 `gorm:"type:bigint unsigned;not null;index:idx_block_number;index:idx_block_log,priority:1;comment:区块高度"`
	LogIndex    uint32 `gorm:"type:int unsigned;not null;default:0;index:idx_block_log,priority:2;comment:logIndex(排序用)"`
	BlockTime   uint64 `gorm:"type:bigint unsigned;not null;index:idx_block_number;index:idx_block_log,priority:1;comment:区块高度"`

	Sender string `gorm:"type:varchar(42);not null;index:idx_sender_block,priority:1;comment:事件sender(0x...)"`
	ToAddr string `gorm:"column:to_addr;type:varchar(42);not null;index:idx_to_block,priority:1;comment:事件to(0x...)"`

	Side uint8 `gorm:"type:tinyint unsigned;not null;default:0;index:idx_side_block,priority:1;comment:方向1=BUY(DL)2=SELL(DL)0=UNKNOWN"`

	Amount0In       float64 `gorm:"column:amount0_in;type:decimal(65,18);not null;default:0;comment:DL in"`
	Amount1In       float64 `gorm:"column:amount1_in;type:decimal(65,18);not null;default:0;comment:OTHER in"`
	Amount0OutGross float64 `gorm:"column:amount0_out_gross;type:decimal(65,18);not null;default:0;comment:DL out"`
	Amount0OutNet   float64 `gorm:"column:amount0_out_net;type:decimal(65,18);not null;default:0;comment:DL out net"`
	Amount1OutGross float64 `gorm:"column:amount1_out_gross;type:decimal(65,18);not null;default:0;comment:OTHER out"`

	CreatedAt time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;comment:更新时间"`
}

func (SwapTrade) TableName() string { return "swap_trade" }

type PrimaryBuy struct {
	ID uint64 `gorm:"primarykey;type:bigint unsigned;comment:主键"`

	BlockNumber uint64 `gorm:"column:block_number;type:bigint unsigned;not null;index:idx_block_log,priority:1;comment:区块高度"`
	BlockTime   uint64 `gorm:"column:block_time;type:bigint unsigned;not null;index:idx_block_time;comment:区块时间(秒)"`
	LogIndex    uint32 `gorm:"column:log_index;type:int unsigned;not null;default:0;index:idx_block_log,priority:2;comment:logIndex(排序用)"`

	Buyer  string `gorm:"column:buyer;type:varchar(42);not null;index:idx_buyer;comment:buyer"`
	ToAddr string `gorm:"column:to_addr;type:varchar(42);not null;index:idx_to;comment:to"`

	UsdtUsed     float64 `gorm:"column:usdt_used;type:decimal(65,18);not null;default:0;comment:usdtUsed"`
	AusdGrossOut float64 `gorm:"column:ausd_gross_out;type:decimal(65,18);not null;default:0;comment:ausdGrossOut"`
	AusdFee      float64 `gorm:"column:ausd_fee;type:decimal(65,18);not null;default:0;comment:ausdFee"`
	AusdNetOut   float64 `gorm:"column:ausd_net_out;type:decimal(65,18);not null;default:0;comment:ausdNetOut"`
	PriceBefore  float64 `gorm:"column:price_before;type:decimal(65,18);not null;default:0;comment:priceBefore"`
	PriceAfter   float64 `gorm:"column:price_after;type:decimal(65,18);not null;default:0;comment:priceAfter"`

	CreatedAt time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;comment:更新时间"`
}

func (PrimaryBuy) TableName() string { return "primary_buy" }

type PrimarySell struct {
	ID uint64 `gorm:"primarykey;type:bigint unsigned;comment:主键"`

	BlockNumber uint64 `gorm:"column:block_number;type:bigint unsigned;not null;index:idx_block_log,priority:1;comment:区块高度"`
	BlockTime   uint64 `gorm:"column:block_time;type:bigint unsigned;not null;index:idx_block_time;comment:区块时间(秒)"`
	LogIndex    uint32 `gorm:"column:log_index;type:int unsigned;not null;default:0;index:idx_block_log,priority:2;comment:logIndex(排序用)"`

	Seller string `gorm:"column:seller;type:varchar(42);not null;index:idx_seller;comment:seller"`
	ToAddr string `gorm:"column:to_addr;type:varchar(42);not null;index:idx_to;comment:to"`

	AusdGrossIn float64 `gorm:"column:ausd_gross_in;type:decimal(65,18);not null;default:0;comment:ausdGrossIn"`
	AusdFee     float64 `gorm:"column:ausd_fee;type:decimal(65,18);not null;default:0;comment:ausdFee"`
	AusdBurn    float64 `gorm:"column:ausd_burn;type:decimal(65,18);not null;default:0;comment:ausdBurn"`
	UsdtOut     float64 `gorm:"column:usdt_out;type:decimal(65,18);not null;default:0;comment:usdtOut"`
	PriceBefore float64 `gorm:"column:price_before;type:decimal(65,18);not null;default:0;comment:priceBefore"`
	PriceAfter  float64 `gorm:"column:price_after;type:decimal(65,18);not null;default:0;comment:priceAfter"`

	CreatedAt time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;comment:更新时间"`
}

func (PrimarySell) TableName() string { return "primary_sell" }

type RewardNotified struct {
	ID uint64 `gorm:"primarykey;type:bigint unsigned;comment:主键"`

	BlockNumber uint64 `gorm:"column:block_number;type:bigint unsigned;not null;index:idx_block_log,priority:1;comment:区块高度"`
	BlockTime   uint64 `gorm:"column:block_time;type:bigint unsigned;not null;index:idx_block_time;comment:区块时间(秒)"`
	LogIndex    uint32 `gorm:"column:log_index;type:int unsigned;not null;default:0;index:idx_block_log,priority:2;comment:logIndex(排序用)"`

	User string `gorm:"column:user;type:varchar(42);not null;index:idx_user_time,priority:1;comment:indexed user"`
	L1   string `gorm:"column:l1;type:varchar(42);not null;index:idx_l1_time,priority:1;comment:indexed l1"`
	L2   string `gorm:"column:l2;type:varchar(42);not null;index:idx_l2_time,priority:1;comment:indexed l2"`

	Profit    float64 `gorm:"column:profit;type:decimal(65,18);not null;default:0;comment:profit"`
	UserShare float64 `gorm:"column:user_share;type:decimal(65,18);not null;default:0;comment:userShare"`
	Top       string  `gorm:"column:top;type:varchar(42);not null;index:idx_top_time,priority:1;comment:top"`

	Pool             float64 `gorm:"column:pool;type:decimal(65,18);not null;default:0;comment:pool"`
	UplinePortionBps uint64  `gorm:"column:upline_portion_bps;type:bigint unsigned;not null;default:0;comment:uplinePortionBps"`

	ToL1      float64 `gorm:"column:to_l1;type:decimal(65,18);not null;default:0;comment:toL1"`
	ToL2      float64 `gorm:"column:to_l2;type:decimal(65,18);not null;default:0;comment:toL2"`
	ToTop     float64 `gorm:"column:to_top;type:decimal(65,18);not null;default:0;comment:toTop"`
	ToProject float64 `gorm:"column:to_project;type:decimal(65,18);not null;default:0;comment:toProject"`

	CreatedAt time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;comment:更新时间"`
}

func (RewardNotified) TableName() string { return "reward_notified" }

type RewardDetail struct {
	ID         uint64    `gorm:"primarykey;type:bigint unsigned;comment:主键"`
	User       string    `gorm:"column:user;type:varchar(42);not null;index:idx_user_time,priority:1;comment:indexed user"`
	Amount     float64   `gorm:"column:amount;type:decimal(65,18);not null;default:0;comment:amount"`
	Reason     uint64    `gorm:"column:reason;type:bigint unsigned;not null;default:0;comment:reason"`
	NotifiedId uint64    `gorm:"column:notified_id;type:int unsigned;not null;default:0;comment:notified_id"`
	CreatedAt  time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null;comment:更新时间"`
	BlockTime  uint64    `gorm:"column:block_time;type:bigint unsigned;not null;index:idx_block_time;comment:区块时间(秒)"`
}

func (RewardDetail) TableName() string { return "reward_detail" }

// data/nft_market_purchase.go
type NftMarketPurchase struct {
	ID uint64 `gorm:"primaryKey;autoIncrement;column:id"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null"`

	Buyer   string `gorm:"column:buyer;type:varchar(42);not null"`
	Seller  string `gorm:"column:seller;type:varchar(42);not null"`
	TokenID uint64 `gorm:"column:token_id;not null"`

	PriceUSDT  float64 `gorm:"column:price_usdt;type:decimal(65,18);not null"`
	FeePaidInB uint8   `gorm:"column:fee_paid_in_b;not null"`
	FeeUSDT    float64 `gorm:"column:fee_usdt;type:decimal(65,18);not null"`
	FeeB       float64 `gorm:"column:fee_b;type:decimal(65,18);not null"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"` // 你表没有也无所谓
	CheckStatus uint64    `gorm:"not null"`
}

func (NftMarketPurchase) TableName() string { return "nft_market_purchase" }

type NftMinted struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`

	BlockNumber uint64 `gorm:"column:block_number;not null" json:"block_number"`
	BlockTime   uint64 `gorm:"column:block_time;not null" json:"block_time"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0" json:"log_index"`

	ToAddr  string `gorm:"column:to_addr;type:varchar(42);not null" json:"to_addr"`
	TokenID uint64 `gorm:"column:token_id;not null" json:"token_id"`

	Tier     uint64  `gorm:"column:tier;not null;default:0" json:"tier"`
	UsdtPaid float64 `gorm:"column:usdt_paid;not null;default:0" json:"usdt_paid"`

	Status   uint8  `gorm:"column:status;not null;default:0" json:"status"`
	ListedAt uint64 `gorm:"column:listed_at;not null;default:0" json:"listed_at"`

	OpenStatus uint8  `gorm:"column:open_status;not null;default:0" json:"open_status"`
	OpenedAt   uint64 `gorm:"column:opened_at;not null;default:0" json:"opened_at"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	CheckStatus uint64    `gorm:"not null"`
	CheckTime   uint64    `gorm:"not null"`

	MintAddr string `gorm:"column:mint_addr;type:varchar(42);not null" json:"mint_addr"`
}

func (NftMinted) TableName() string { return "nft_minted" }

type NftMarketListed struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	Seller    string `gorm:"column:seller;type:varchar(42);not null"`
	TokenID   uint64 `gorm:"column:token_id;not null"`
	Timestamp uint64 `gorm:"column:timestamp;not null;default:0"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CheckStatus uint64    `gorm:"not null"`
}

func (NftMarketListed) TableName() string { return "nft_market_listed" }

type NftMarketUnlisted struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	Operator string `gorm:"column:operator;type:varchar(42);not null"`
	TokenID  uint64 `gorm:"column:token_id;not null"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CheckStatus uint64    `gorm:"not null"`
}

func (NftMarketUnlisted) TableName() string { return "nft_market_unlisted" }

type NftOpened struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	UserAddr string `gorm:"column:user_addr;type:varchar(42);not null"`
	TokenID  uint64 `gorm:"column:token_id;not null" json:"token_id"`

	OpenedAt uint64  `gorm:"column:opened_at;not null;default:0"`
	Reward   float64 `gorm:"column:reward;type:decimal(65,18);not null;default:0"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CheckStatus uint64    `gorm:"not null"`
}

func (NftOpened) TableName() string { return "nft_opened" }

type NftTransfer struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	FromAddr string `gorm:"column:from_addr;type:varchar(42);not null"`
	ToAddr   string `gorm:"column:to_addr;type:varchar(42);not null"`
	TokenID  uint64 `gorm:"column:token_id;not null"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CheckStatus uint64    `gorm:"not null"`
}

func (NftTransfer) TableName() string { return "nft_transfer" }

type UserRegistered struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	UserAddr   string `gorm:"column:user_addr;type:varchar(42);not null"`
	ParentAddr string `gorm:"column:parent_addr;type:varchar(42);not null"`
	TopAddr    string `gorm:"column:top_addr;type:varchar(42);not null"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (UserRegistered) TableName() string { return "user_registered" }

type RuntimeConfig struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	ConfigKey   string    `gorm:"column:config_key;type:varchar(128);not null;uniqueIndex"`
	ConfigValue string    `gorm:"column:config_value;type:text;not null"`
	Description string    `gorm:"column:description;type:varchar(255);not null;default:''"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (RuntimeConfig) TableName() string { return "runtime_config" }

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *UserRepo) GetSwapTradeLast(ctx context.Context) (*biz.SwapTrade, error) {
	var v *SwapTrade

	if err := u.data.DB(ctx).Table("swap_trade").Order("block_number DESC, id DESC").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "SWAP_TRADE_ERROR", err.Error())
	}

	return &biz.SwapTrade{
		ID:              v.ID,
		BlockNumber:     v.BlockNumber,
		LogIndex:        v.LogIndex,
		BlockTime:       v.BlockTime,
		Sender:          v.Sender,
		ToAddr:          v.ToAddr,
		Side:            v.Side,
		Amount0In:       v.Amount0In,
		Amount1In:       v.Amount1In,
		Amount0OutNet:   v.Amount0OutNet,
		Amount1OutGross: v.Amount1OutGross,
		Amount0OutGross: v.Amount0OutGross,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetSwapTrade(ctx context.Context, start, end uint64) ([]*biz.SwapTrade, error) {
	var s []*SwapTrade

	res := make([]*biz.SwapTrade, 0)
	if err := u.data.DB(ctx).Table("swap_trade").Where("block_time >=?", start).Where("block_time <=?", end).Find(&s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "SWAP_TRADE_ERROR", err.Error())
	}

	for _, v := range s {
		res = append(res, &biz.SwapTrade{
			ID:              v.ID,
			BlockNumber:     v.BlockNumber,
			LogIndex:        v.LogIndex,
			BlockTime:       v.BlockTime,
			Sender:          v.Sender,
			ToAddr:          v.ToAddr,
			Side:            v.Side,
			Amount0In:       v.Amount0In,
			Amount1In:       v.Amount1In,
			Amount0OutNet:   v.Amount0OutNet,
			Amount1OutGross: v.Amount1OutGross,
			Amount0OutGross: v.Amount0OutGross,
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
		})
	}
	return res, nil
}

// InsertSwapTrade .
func (u *UserRepo) InsertSwapTrade(ctx context.Context, iData *biz.SwapTrade) error {
	var (
		err error
		s   SwapTrade
	)

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.Side = iData.Side
	s.Sender = iData.Sender
	s.Amount1In = iData.Amount1In
	s.Amount0OutGross = iData.Amount0OutGross
	s.Amount0OutNet = iData.Amount0OutNet
	s.Amount1OutGross = iData.Amount1OutGross
	s.Amount0In = iData.Amount0In
	s.ToAddr = iData.ToAddr
	s.LogIndex = iData.LogIndex

	err = u.data.DB(ctx).Table("swap_trade").Create(&s).Error
	if err != nil {
		return errors.New(500, "CREATE_SWAP_TRADE_ERROR", "信息创建失败")
	}

	return nil
}

func (u *UserRepo) GetPrimaryBuyLast(ctx context.Context) (*biz.PrimaryBuy, error) {
	var v PrimaryBuy

	if err := u.data.DB(ctx).Table("primary_buy").Order("block_number DESC, id DESC").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "PRIMARY_BUY_ERROR", err.Error())
	}

	return &biz.PrimaryBuy{
		ID:           v.ID,
		BlockNumber:  v.BlockNumber,
		BlockTime:    v.BlockTime,
		LogIndex:     v.LogIndex,
		Buyer:        v.Buyer,
		ToAddr:       v.ToAddr,
		UsdtUsed:     v.UsdtUsed,
		AusdGrossOut: v.AusdGrossOut,
		AusdFee:      v.AusdFee,
		AusdNetOut:   v.AusdNetOut,
		PriceBefore:  v.PriceBefore,
		PriceAfter:   v.PriceAfter,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetPrimaryBuy(ctx context.Context, start, end uint64) ([]*biz.PrimaryBuy, error) {
	var s []PrimaryBuy
	res := make([]*biz.PrimaryBuy, 0)

	if err := u.data.DB(ctx).
		Table("primary_buy").
		Where("block_time > ?", 0).
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Order("block_time asc, id asc").
		Find(&s).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "PRIMARY_BUY_ERROR", err.Error())
	}

	for _, v := range s {
		res = append(res, &biz.PrimaryBuy{
			ID:           v.ID,
			BlockNumber:  v.BlockNumber,
			BlockTime:    v.BlockTime,
			LogIndex:     v.LogIndex,
			Buyer:        v.Buyer,
			ToAddr:       v.ToAddr,
			UsdtUsed:     v.UsdtUsed,
			AusdGrossOut: v.AusdGrossOut,
			AusdFee:      v.AusdFee,
			AusdNetOut:   v.AusdNetOut,
			PriceBefore:  v.PriceBefore,
			PriceAfter:   v.PriceAfter,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}
	return res, nil
}

// InsertPrimaryBuy .
func (u *UserRepo) InsertPrimaryBuy(ctx context.Context, iData *biz.PrimaryBuy) error {
	var s PrimaryBuy

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex
	s.Buyer = iData.Buyer
	s.ToAddr = iData.ToAddr

	s.UsdtUsed = iData.UsdtUsed
	s.AusdGrossOut = iData.AusdGrossOut
	s.AusdFee = iData.AusdFee
	s.AusdNetOut = iData.AusdNetOut
	s.PriceBefore = iData.PriceBefore
	s.PriceAfter = iData.PriceAfter

	if err := u.data.DB(ctx).Table("primary_buy").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_PRIMARY_BUY_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) GetPrimarySellLast(ctx context.Context) (*biz.PrimarySell, error) {
	var v PrimarySell

	if err := u.data.DB(ctx).Table("primary_sell").Order("block_number DESC, id DESC").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "PRIMARY_SELL_ERROR", err.Error())
	}

	return &biz.PrimarySell{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,
		Seller:      v.Seller,
		ToAddr:      v.ToAddr,
		AusdGrossIn: v.AusdGrossIn,
		AusdFee:     v.AusdFee,
		AusdBurn:    v.AusdBurn,
		UsdtOut:     v.UsdtOut,
		PriceBefore: v.PriceBefore,
		PriceAfter:  v.PriceAfter,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetPrimarySell(ctx context.Context, start, end uint64) ([]*biz.PrimarySell, error) {
	var s []PrimarySell
	res := make([]*biz.PrimarySell, 0)

	if err := u.data.DB(ctx).
		Table("primary_sell").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Order("block_time asc, id asc").
		Find(&s).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "PRIMARY_SELL_ERROR", err.Error())
	}

	for _, v := range s {
		res = append(res, &biz.PrimarySell{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,
			Seller:      v.Seller,
			ToAddr:      v.ToAddr,
			AusdGrossIn: v.AusdGrossIn,
			AusdFee:     v.AusdFee,
			AusdBurn:    v.AusdBurn,
			UsdtOut:     v.UsdtOut,
			PriceBefore: v.PriceBefore,
			PriceAfter:  v.PriceAfter,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return res, nil
}

// InsertPrimarySell .
func (u *UserRepo) InsertPrimarySell(ctx context.Context, iData *biz.PrimarySell) error {
	var s PrimarySell

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex
	s.Seller = iData.Seller
	s.ToAddr = iData.ToAddr

	s.AusdGrossIn = iData.AusdGrossIn
	s.AusdFee = iData.AusdFee
	s.AusdBurn = iData.AusdBurn
	s.UsdtOut = iData.UsdtOut
	s.PriceBefore = iData.PriceBefore
	s.PriceAfter = iData.PriceAfter

	if err := u.data.DB(ctx).Table("primary_sell").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_PRIMARY_SELL_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) GetRewardNotifiedLast(ctx context.Context) (*biz.RewardNotified, error) {
	var v RewardNotified

	if err := u.data.DB(ctx).Table("reward_notified").Order("block_number DESC, id DESC").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "REWARD_NOTIFIED_ERROR", err.Error())
	}

	return &biz.RewardNotified{
		ID:               v.ID,
		BlockNumber:      v.BlockNumber,
		BlockTime:        v.BlockTime,
		LogIndex:         v.LogIndex,
		User:             v.User,
		L1:               v.L1,
		L2:               v.L2,
		Profit:           v.Profit,
		UserShare:        v.UserShare,
		Top:              v.Top,
		Pool:             v.Pool,
		UplinePortionBps: v.UplinePortionBps,
		ToL1:             v.ToL1,
		ToL2:             v.ToL2,
		ToTop:            v.ToTop,
		ToProject:        v.ToProject,
		CreatedAt:        v.CreatedAt,
		UpdatedAt:        v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetRewardNotified(ctx context.Context, start, end uint64) ([]*biz.RewardNotified, error) {
	var s []RewardNotified
	res := make([]*biz.RewardNotified, 0)

	if err := u.data.DB(ctx).
		Table("reward_notified").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Order("block_time asc, id asc").
		Find(&s).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "REWARD_NOTIFIED_ERROR", err.Error())
	}

	for _, v := range s {
		res = append(res, &biz.RewardNotified{
			ID:               v.ID,
			BlockNumber:      v.BlockNumber,
			BlockTime:        v.BlockTime,
			LogIndex:         v.LogIndex,
			User:             v.User,
			L1:               v.L1,
			L2:               v.L2,
			Profit:           v.Profit,
			UserShare:        v.UserShare,
			Top:              v.Top,
			Pool:             v.Pool,
			UplinePortionBps: v.UplinePortionBps,
			ToL1:             v.ToL1,
			ToL2:             v.ToL2,
			ToTop:            v.ToTop,
			ToProject:        v.ToProject,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		})
	}
	return res, nil
}

func (u *UserRepo) GetRewardNotifiedByIds(ctx context.Context, ids []uint64) (map[uint64]*biz.RewardNotified, error) {
	var s []RewardNotified
	res := make(map[uint64]*biz.RewardNotified, 0)

	if err := u.data.DB(ctx).
		Table("reward_notified").
		Where("id in(?)", ids).
		Find(&s).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "REWARD_NOTIFIED_ERROR", err.Error())
	}

	for _, v := range s {
		res[v.ID] = &biz.RewardNotified{
			ID:               v.ID,
			BlockNumber:      v.BlockNumber,
			BlockTime:        v.BlockTime,
			LogIndex:         v.LogIndex,
			User:             v.User,
			L1:               v.L1,
			L2:               v.L2,
			Profit:           v.Profit,
			UserShare:        v.UserShare,
			Top:              v.Top,
			Pool:             v.Pool,
			UplinePortionBps: v.UplinePortionBps,
			ToL1:             v.ToL1,
			ToL2:             v.ToL2,
			ToTop:            v.ToTop,
			ToProject:        v.ToProject,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		}
	}
	return res, nil
}

func (u *UserRepo) InsertRewardNotified(ctx context.Context, iData *biz.RewardNotified) error {
	var s RewardNotified

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.User = iData.User
	s.L1 = iData.L1
	s.L2 = iData.L2

	s.Profit = iData.Profit
	s.UserShare = iData.UserShare
	s.Top = iData.Top

	s.Pool = iData.Pool
	s.UplinePortionBps = iData.UplinePortionBps

	s.ToL1 = iData.ToL1
	s.ToL2 = iData.ToL2
	s.ToTop = iData.ToTop
	s.ToProject = iData.ToProject

	if err := u.data.DB(ctx).Table("reward_notified").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_NOTIFIED_ERROR", "信息创建失败")
	}

	var one RewardDetail
	one.User = s.User
	one.Reason = 1
	one.NotifiedId = s.ID
	one.Amount = s.UserShare
	one.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&one).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	var two RewardDetail
	two.User = s.L1
	two.Reason = 2
	two.NotifiedId = s.ID
	two.Amount = s.ToL1
	two.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&two).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	var three RewardDetail
	three.User = s.L2
	three.Reason = 3
	three.NotifiedId = s.ID
	three.Amount = s.ToL2
	three.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&three).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	var four RewardDetail
	four.User = s.Top
	four.Reason = 4
	four.NotifiedId = s.ID
	four.Amount = s.ToTop
	four.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&four).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	var five RewardDetail
	five.User = "0x144351b1a5Af4538f53556633D8f10495aA564A8"
	five.Reason = 5
	five.NotifiedId = s.ID
	five.Amount = s.ToProject
	five.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&five).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	return nil
}

// InsertRewardNotifiedScanProgress 仅写入 reward_notified 块游标，不写 reward_detail（与 InsertRewardNotified 区分，供空窗推进游标）
func (u *UserRepo) InsertRewardNotifiedScanProgress(ctx context.Context, blockNumber uint64) error {
	z := "0x0000000000000000000000000000000000000000"
	var s RewardNotified
	s.BlockNumber = blockNumber
	s.BlockTime = 0
	s.LogIndex = 0
	s.User = z
	s.L1 = z
	s.L2 = z
	s.Top = z
	if err := u.data.DB(ctx).Table("reward_notified").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_NOTIFIED_PROGRESS_ERROR", err.Error())
	}
	return nil
}

// GetUserRewardByUserIdPage .
func (u *UserRepo) GetUserRewardByUserIdPage(ctx context.Context, b *biz.Pagination, address string, reason uint64) ([]*biz.RewardDetail, error, int64) {
	var (
		count   int64
		rewards []*RewardDetail
	)

	res := make([]*biz.RewardDetail, 0)

	instance := u.data.db.Where("user", address).Table("reward_detail").Order("id desc")
	if 0 < reason {
		instance = instance.Where("reason=?", reason)
	}

	instance = instance.Count(&count)

	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found"), 0
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error()), 0
	}

	for _, reward := range rewards {
		res = append(res, &biz.RewardDetail{
			ID:         reward.ID,
			User:       reward.User,
			Amount:     reward.Amount,
			Reason:     reward.Reason,
			BlockTime:  reward.BlockTime,
			NotifiedId: reward.NotifiedId,
			CreatedAt:  reward.CreatedAt,
			UpdatedAt:  reward.UpdatedAt,
		})
	}

	return res, nil, count
}

// GetNftMarketPurchaseByAddressPage
// side: 0=buyer或seller  1=buyer  2=seller
func (u *UserRepo) GetNftMarketPurchaseByAddressPage(
	ctx context.Context,
	b *biz.Pagination,
	address string,
	addressTwo []string,
	side uint64,
) ([]*biz.NftMarketPurchase, error, int64) {

	var (
		count int64
		rows  []*NftMarketPurchase
	)

	res := make([]*biz.NftMarketPurchase, 0)

	instance := u.data.DB(ctx).Table("nft_market_purchase").Order("id desc")

	if address != "" {
		if side == 1 {
			instance = instance.Where("buyer = ?", address)
		} else if side == 2 {
			instance = instance.Where("seller = ?", address)
		} else if side == 3 {
			instance = instance.Where("seller in (?)", addressTwo)
		} else {
			instance = instance.Where("(buyer = ? OR seller = ?)", address, address)
		}
	}

	instance = instance.Count(&count)

	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&rows).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("NFT_PURCHASE_NOT_FOUND", "purchase not found"), 0
		}
		return nil, errors.New(500, "NFT_PURCHASE_PAGE_ERROR", err.Error()), 0
	}

	for _, r := range rows {
		res = append(res, &biz.NftMarketPurchase{
			ID:          r.ID,
			BlockNumber: r.BlockNumber,
			BlockTime:   r.BlockTime,
			LogIndex:    r.LogIndex,

			Buyer:   r.Buyer,
			Seller:  r.Seller,
			TokenID: r.TokenID,

			PriceUSDT:  r.PriceUSDT,
			FeePaidInB: r.FeePaidInB,
			FeeUSDT:    r.FeeUSDT,
			FeeB:       r.FeeB,

			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}

	return res, nil, count
}

func (u *UserRepo) InsertNftMarketPurchase(ctx context.Context, iData *biz.NftMarketPurchase) error {
	var s NftMarketPurchase

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.Buyer = iData.Buyer
	s.Seller = iData.Seller
	s.TokenID = iData.TokenID

	s.PriceUSDT = iData.PriceUSDT
	s.FeePaidInB = iData.FeePaidInB
	s.FeeUSDT = iData.FeeUSDT
	s.FeeB = iData.FeeB
	s.CheckStatus = iData.CheckStatus

	if err := u.data.DB(ctx).Table("nft_market_purchase").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_MARKET_PURCHASE_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) GetNftMarketPurchaseLast(ctx context.Context) (*biz.NftMarketPurchase, error) {
	// 二级市场 Purchased 扫块游标：按 block_number，勿仅用 id desc（否则最新 id 行块高较旧时游标回退）
	var v NftMarketPurchase
	if err := u.data.DB(ctx).Table("nft_market_purchase").Order("block_number DESC, id DESC").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_MARKET_PURCHASE_ERROR", err.Error())
	}
	if v.BlockNumber > 0 {
		return nftMarketPurchaseModelToBiz(&v), nil
	}
	var sent NftMarketPurchase
	err := u.data.DB(ctx).Table("nft_market_purchase").Where("token_id = ?", 0).Order("block_number DESC, id DESC").First(&sent).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nftMarketPurchaseModelToBiz(&v), nil
		}
		return nil, errors.New(500, "NFT_MARKET_PURCHASE_ERROR", err.Error())
	}
	return nftMarketPurchaseModelToBiz(&sent), nil
}

func nftMarketPurchaseModelToBiz(v *NftMarketPurchase) *biz.NftMarketPurchase {
	return &biz.NftMarketPurchase{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		Buyer:   v.Buyer,
		Seller:  v.Seller,
		TokenID: v.TokenID,

		PriceUSDT:  v.PriceUSDT,
		FeePaidInB: v.FeePaidInB,
		FeeUSDT:    v.FeeUSDT,
		FeeB:       v.FeeB,

		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		CheckStatus: v.CheckStatus,
	}
}

func (u *UserRepo) UpdateNftMinted(ctx context.Context, tokenId uint64, mintAddr string) error {
	res := u.data.DB(ctx).Table("nft_minted").Where("token_id=?", tokenId).
		Updates(map[string]interface{}{
			"mint_addr":  mintAddr,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_MT_ERROR", "修改失败")
	}
	return nil
}

func (u *UserRepo) InsertNftMinted(ctx context.Context, iData *biz.NftMinted) error {
	// token_id > 0：可能已有 sync_mint_from_contract 占位行、或同 token 重复扫到 Mint 事件。
	// 旧逻辑在「行已存在且 block_number>0」时直接 return，会吞掉事件里的 to_addr/mint_addr，导致 get_address_box 永远对不上链。
	if iData.TokenID > 0 {
		var existing NftMinted
		err := u.data.DB(ctx).Table("nft_minted").
			Where("token_id = ?", iData.TokenID).
			Order("id DESC").First(&existing).Error
		if err == nil {
			if iData.BlockNumber == 0 {
				return nil
			}
			needMerge := false
			if existing.BlockNumber == 0 {
				needMerge = true
			} else if iData.BlockNumber > existing.BlockNumber ||
				(iData.BlockNumber == existing.BlockNumber && iData.LogIndex > existing.LogIndex) {
				needMerge = true
			} else if iData.BlockNumber == existing.BlockNumber && iData.LogIndex == existing.LogIndex &&
				!strings.EqualFold(strings.TrimSpace(iData.ToAddr), strings.TrimSpace(existing.ToAddr)) {
				// 同块同 log，仅持有人等字段与链上事件不一致时补齐
				needMerge = true
			}
			if needMerge {
				u.data.DB(ctx).Table("nft_minted").Where("id = ?", existing.ID).
					Updates(map[string]interface{}{
						"block_number": iData.BlockNumber,
						"block_time":   iData.BlockTime,
						"log_index":    iData.LogIndex,
						"to_addr":      iData.ToAddr,
						"mint_addr":    iData.MintAddr,
						"tier":         iData.Tier,
						"usdt_paid":    iData.UsdtPaid,
						"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
					})
			}
			return nil
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(500, "QUERY_NFT_MINTED_ERROR", err.Error())
		}
	}

	var s NftMinted
	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex
	s.ToAddr = iData.ToAddr
	s.TokenID = iData.TokenID
	s.Tier = iData.Tier
	s.UsdtPaid = iData.UsdtPaid
	s.MintAddr = iData.MintAddr

	if err := u.data.DB(ctx).Table("nft_minted").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_MINTED_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) GetNftMintedLast(ctx context.Context) (*biz.NftMinted, error) {
	// 铸造扫块游标：按 block_number 取进度，勿仅用 id desc。
	// sync_mint_from_contract 会插入 block_number=0 的行且 id 往往最大，若用 id desc 会得到 last=0，
	// PollMintedIncremental 会从部署块重扫，且与旧哨兵（如 8880 万）脱节，新铸长期进不了库。
	var v NftMinted
	if err := u.data.DB(ctx).Table("nft_minted").Order("block_number DESC, id DESC").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_MINTED_ERROR", err.Error())
	}
	if v.BlockNumber > 0 {
		return nftMintedModelToBiz(&v), nil
	}
	var sentinel NftMinted
	err := u.data.DB(ctx).Table("nft_minted").Where("token_id = ?", 0).Order("block_number DESC, id DESC").First(&sentinel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nftMintedModelToBiz(&v), nil
		}
		return nil, errors.New(500, "NFT_MINTED_ERROR", err.Error())
	}
	return nftMintedModelToBiz(&sentinel), nil
}

func nftMintedModelToBiz(v *NftMinted) *biz.NftMinted {
	return &biz.NftMinted{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		ToAddr:   v.ToAddr,
		MintAddr: v.MintAddr,
		TokenID:  v.TokenID,

		Tier:     v.Tier,
		UsdtPaid: v.UsdtPaid,

		Status:   v.Status,
		ListedAt: v.ListedAt,

		OpenStatus: v.OpenStatus,
		OpenedAt:   v.OpenedAt,

		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		CheckStatus: v.CheckStatus,
		CheckTime:   v.CheckTime,
	}
}

func (u *UserRepo) InsertNftMarketListed(ctx context.Context, iData *biz.NftMarketListed) error {
	var s NftMarketListed

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.Seller = iData.Seller
	s.TokenID = iData.TokenID
	s.Timestamp = iData.Timestamp
	s.CheckStatus = iData.CheckStatus

	if err := u.data.DB(ctx).Table("nft_market_listed").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_MARKET_LISTED_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftMarketListedLast(ctx context.Context) (*biz.NftMarketListed, error) {
	var v NftMarketListed

	// 按最大块高取游标（与分块扫描进度一致）；勿仅用 id desc，否则可能与块序不一致
	if err := u.data.DB(ctx).Table("nft_market_listed").Order("block_number desc, id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_MARKET_LISTED_ERROR", err.Error())
	}

	return &biz.NftMarketListed{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		Seller:    v.Seller,
		TokenID:   v.TokenID,
		Timestamp: v.Timestamp,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) InsertNftMarketUnlisted(ctx context.Context, iData *biz.NftMarketUnlisted) error {
	var s NftMarketUnlisted

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.Operator = iData.Operator
	s.TokenID = iData.TokenID
	s.CheckStatus = iData.CheckStatus

	if err := u.data.DB(ctx).Table("nft_market_unlisted").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_MARKET_UNLISTED_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftMarketUnlistedLast(ctx context.Context) (*biz.NftMarketUnlisted, error) {
	var v NftMarketUnlisted

	if err := u.data.DB(ctx).Table("nft_market_unlisted").Order("block_number DESC, id DESC").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_MARKET_UNLISTED_ERROR", err.Error())
	}

	return &biz.NftMarketUnlisted{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		Operator: v.Operator,
		TokenID:  v.TokenID,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) InsertNftOpened(ctx context.Context, iData *biz.NftOpened) error {
	var s NftOpened

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.UserAddr = iData.UserAddr
	s.TokenID = iData.TokenID

	s.OpenedAt = iData.OpenedAt
	s.Reward = iData.Reward
	s.CheckStatus = iData.CheckStatus

	if err := u.data.DB(ctx).Table("nft_opened").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_OPENED_ERROR", err.Error())
	}
	return nil
}

func (u *UserRepo) GetNftOpenedLast(ctx context.Context) (*biz.NftOpened, error) {
	// Opened 扫块游标：按 block_number，勿仅用 id desc
	var v NftOpened
	if err := u.data.DB(ctx).Table("nft_opened").Order("block_number DESC, id DESC").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_OPENED_ERROR", err.Error())
	}
	if v.BlockNumber > 0 {
		return nftOpenedModelToBiz(&v), nil
	}
	var sent NftOpened
	err := u.data.DB(ctx).Table("nft_opened").Where("token_id = ?", 0).Order("block_number DESC, id DESC").First(&sent).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nftOpenedModelToBiz(&v), nil
		}
		return nil, errors.New(500, "NFT_OPENED_ERROR", err.Error())
	}
	return nftOpenedModelToBiz(&sent), nil
}

func nftOpenedModelToBiz(v *NftOpened) *biz.NftOpened {
	return &biz.NftOpened{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		UserAddr: v.UserAddr,
		TokenID:  v.TokenID,

		OpenedAt: v.OpenedAt,
		Reward:   v.Reward,

		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		CheckStatus: v.CheckStatus,
	}
}

func (u *UserRepo) InsertNftTransfer(ctx context.Context, iData *biz.NftTransfer) error {
	var s NftTransfer

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.FromAddr = iData.FromAddr
	s.ToAddr = iData.ToAddr
	s.TokenID = iData.TokenID
	s.CheckStatus = iData.CheckStatus

	if err := u.data.DB(ctx).Table("nft_transfer").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_TRANSFER_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftTransferLast(ctx context.Context) (*biz.NftTransfer, error) {
	var v NftTransfer

	if err := u.data.DB(ctx).Table("nft_transfer").Order("block_number DESC, id DESC").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	return &biz.NftTransfer{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		FromAddr: v.FromAddr,
		ToAddr:   v.ToAddr,
		TokenID:  v.TokenID,

		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		CheckStatus: v.CheckStatus,
	}, nil
}

func (u *UserRepo) GetNftTransferLastNoCheck(ctx context.Context) ([]*biz.NftTransfer, error) {
	var vL []*NftTransfer

	res := make([]*biz.NftTransfer, 0)
	if err := u.data.DB(ctx).Table("nft_transfer").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftTransfer{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			FromAddr: v.FromAddr,
			ToAddr:   v.ToAddr,
			TokenID:  v.TokenID,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftListLastNoCheck(ctx context.Context) ([]*biz.NftMarketListed, error) {
	var vL []*NftMarketListed

	res := make([]*biz.NftMarketListed, 0)
	if err := u.data.DB(ctx).Table("nft_market_listed").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftMarketListed{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			Seller:    v.Seller,
			TokenID:   v.TokenID,
			Timestamp: v.Timestamp,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftUnListLastNoCheck(ctx context.Context) ([]*biz.NftMarketUnlisted, error) {
	var vL []*NftMarketUnlisted

	res := make([]*biz.NftMarketUnlisted, 0)
	if err := u.data.DB(ctx).Table("nft_market_unlisted").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftMarketUnlisted{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			Operator: v.Operator,
			TokenID:  v.TokenID,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftBuyLastNoCheck(ctx context.Context) ([]*biz.NftMarketPurchase, error) {
	var vL []*NftMarketPurchase

	res := make([]*biz.NftMarketPurchase, 0)
	if err := u.data.DB(ctx).Table("nft_market_purchase").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftMarketPurchase{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			Buyer:   v.Buyer,
			Seller:  v.Seller,
			TokenID: v.TokenID,

			PriceUSDT:  v.PriceUSDT,
			FeePaidInB: v.FeePaidInB,
			FeeUSDT:    v.FeeUSDT,
			FeeB:       v.FeeB,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftOpenLastNoCheck(ctx context.Context) ([]*biz.NftOpened, error) {
	var vL []*NftOpened

	res := make([]*biz.NftOpened, 0)
	if err := u.data.DB(ctx).Table("nft_opened").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftOpened{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			UserAddr: v.UserAddr,
			TokenID:  v.TokenID,

			OpenedAt: v.OpenedAt,
			Reward:   v.Reward,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftMintedByTokenIds(ctx context.Context, tokenIds []uint64) (map[uint64]*biz.NftMinted, error) {
	res := make(map[uint64]*biz.NftMinted, 0)
	if len(tokenIds) == 0 {
		return res, nil
	}

	var vL []*NftMinted
	// 同一 token_id 多行时取 MAX(id)，与 GetNftMintedByAddressPage 一致。
	// 若用 Find 全表再覆盖 map，可能更新到旧行，列表仍用 MAX(id) 行 → 开盒后 open_status 仍显示未开。
	sub := u.data.DB(ctx).Table("nft_minted").Select("MAX(id)").Where("token_id IN (?)", tokenIds).Group("token_id")
	if err := u.data.DB(ctx).Table("nft_minted").Where("id IN (?)", sub).Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res[v.TokenID] = &biz.NftMinted{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			ToAddr:  v.ToAddr,
			TokenID: v.TokenID,

			Tier:     v.Tier,
			UsdtPaid: v.UsdtPaid,

			Status:   v.Status,
			ListedAt: v.ListedAt,

			OpenStatus: v.OpenStatus,
			OpenedAt:   v.OpenedAt,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			CheckTime: v.CheckTime,
		}
	}

	return res, nil
}

// UpdateNftMintedToAddress .
func (u *UserRepo) UpdateNftMintedToAddress(ctx context.Context, id, idT, check uint64, toAddr string) error {
	res := u.data.DB(ctx).Table("nft_transfer").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_T_ERROR", "修改失败")
	}

	if 1 == check {
		//.Where("check_time<=?", checkTime).
		resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).
			Updates(map[string]interface{}{
				"to_addr": toAddr,
				//"check_time": checkTime,
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
			return errors.New(500, "UPDATE_MINTED_ERROR", "修改失败")
		}
	}

	return nil
}

// UpdateNftMintedListStatus .
func (u *UserRepo) UpdateNftMintedListStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_listed").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_L_ERROR", "修改失败")
	}

	//.Where("check_time<=?", checkTime).
	resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).Where("check_time<?", checkTime).
		Updates(map[string]interface{}{
			"check_time": checkTime,
			"status":     1,
			"listed_at":  checkTime,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UPDATE_MINTED_L_ERROR", "修改失败")
	}

	return nil
}

// UpdateNftMintedUnListStatus .
func (u *UserRepo) UpdateNftMintedUnListStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_unlisted").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_L_ERROR", "修改失败")
	}

	//.Where("check_time<=?", checkTime).
	resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).Where("check_time<?", checkTime).
		Updates(map[string]interface{}{
			"check_time": checkTime,
			"status":     0,
			"listed_at":  0,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UPDATE_MINTED_U_L_ERROR", "修改失败")
	}

	return nil
}

func (u *UserRepo) UpdateUnlistedCheckStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_unlisted").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 2,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_CUL_ERROR", "修改失败")
	}

	return nil
}

func (u *UserRepo) UpdateListedCheckStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_listed").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 2,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_CL_ERROR", "修改失败")
	}

	return nil
}

func (u *UserRepo) UpdateBuyCheckStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_purchase").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 2,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_CB_ERROR", "修改失败")
	}

	return nil
}

// UpdateNftMintedBuyStatus 市场买入后须把持有人改为 buyer，否则 GetNftMintedByAddressPage 按 to_addr 筛不到买家仓库
func (u *UserRepo) UpdateNftMintedBuyStatus(ctx context.Context, id, idT, checkTime uint64, buyer string) error {
	res := u.data.DB(ctx).Table("nft_market_purchase").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_L_ERROR", "修改失败")
	}

	buyerAddr := strings.ToLower(strings.TrimSpace(buyer))

	//.Where("check_time<=?", checkTime).
	resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).Where("check_time<?", checkTime).
		Updates(map[string]interface{}{
			"check_time": checkTime,
			"status":     0,
			"listed_at":  0,
			"to_addr":    buyerAddr,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UPDATE_MINTED_B_ERROR", "修改失败")
	}

	return nil
}

// UpdateNftMintedOpenStatus .
func (u *UserRepo) UpdateNftMintedOpenStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_opened").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_L_ERROR", "修改失败")
	}

	var row struct {
		TokenID uint64 `gorm:"column:token_id"`
	}
	if err := u.data.DB(ctx).Table("nft_minted").Select("token_id").Where("id = ?", id).Take(&row).Error; err != nil {
		return errors.New(500, "UPDATE_MINTED_O_ERROR", "minted row not found")
	}
	qMinted := u.data.DB(ctx).Table("nft_minted").Where("token_id = ?", row.TokenID)
	resTwo := qMinted.Updates(map[string]interface{}{
		"open_status": 1,
		"opened_at":   checkTime,
		"updated_at":  time.Now().Format("2006-01-02 15:04:05"),
	})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UPDATE_MINTED_O_ERROR", "修改失败")
	}

	return nil
}

// GetNftMintedPage .
func (u *UserRepo) GetNftMintedPage(
	ctx context.Context,
	b *biz.Pagination,
	order uint64,
	orderTwo uint64,
	tier uint64,
) ([]*biz.NftMinted, error, int64) {

	var (
		count int64
		rows  []*NftMinted
	)

	res := make([]*biz.NftMinted, 0)

	// 同一 token_id 可能有多行（铸造+view 同步重复）；市场只展示每个 token 最新一条
	sub := u.data.DB(ctx).Table("nft_minted").Select("MAX(id)").Where("status = ? AND open_status = ?", 1, 0)
	if 0 < tier {
		sub = sub.Where("tier=?", tier)
	}
	sub = sub.Group("token_id")

	base := u.data.DB(ctx).Table("nft_minted").Where("id IN (?)", sub)

	if err := base.Count(&count).Error; err != nil {
		return nil, errors.New(500, "NFT_MINTED_ERROR", err.Error()), 0
	}

	// Count 与 Find 分开链，避免 GORM 链式状态异常
	q := u.data.DB(ctx).Table("nft_minted").Where("id IN (?)", sub)
	if 0 < order {
		q = q.Order("id asc")
	} else {
		q = q.Order("id desc")
	}

	if 0 < orderTwo {
		q = q.Order("tier asc")
	} else {
		q = q.Order("tier desc")
	}

	if err := q.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&rows).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("NFT_MINTED_NOT_FOUND", "minted not found"), 0
		}
		return nil, errors.New(500, "NFT_MINTED_ERROR", err.Error()), 0
	}

	for _, v := range rows {
		res = append(res, &biz.NftMinted{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			ToAddr:  v.ToAddr,
			TokenID: v.TokenID,

			Tier:     v.Tier,
			UsdtPaid: v.UsdtPaid,

			Status:   v.Status,
			ListedAt: v.ListedAt,

			OpenStatus: v.OpenStatus,
			OpenedAt:   v.OpenedAt,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			MintAddr:  v.MintAddr,
		})
	}

	return res, nil, count
}

func (u *UserRepo) GetQuantFundStatusSummary(ctx context.Context) ([]*biz.QuantFundStatusSummary, error) {
	type row struct {
		Status      uint8  `gorm:"column:status"`
		LedgerCount int64  `gorm:"column:ledger_count"`
		TotalAmount string `gorm:"column:total_amount"`
	}

	rows := make([]row, 0)
	if err := u.data.DB(ctx).Table("quant_fund_ledger").
		Select("status, COUNT(*) AS ledger_count, CAST(COALESCE(SUM(quant_amount), 0) AS CHAR) AS total_amount").
		Group("status").
		Order("status ASC").
		Scan(&rows).Error; err != nil {
		return nil, errors.New(500, "QUANT_FUND_SUMMARY_ERROR", err.Error())
	}

	res := make([]*biz.QuantFundStatusSummary, 0, len(rows))
	for _, v := range rows {
		res = append(res, &biz.QuantFundStatusSummary{
			Status:      v.Status,
			LedgerCount: v.LedgerCount,
			TotalAmount: defaultZero(v.TotalAmount),
		})
	}
	return res, nil
}

func (u *UserRepo) GetQuantResultStatusSummary(ctx context.Context) ([]*biz.QuantResultStatusSummary, error) {
	type row struct {
		Status      uint8  `gorm:"column:status"`
		ResultCount int64  `gorm:"column:result_count"`
		ProfitTotal string `gorm:"column:profit_total"`
		ReturnTotal string `gorm:"column:return_total"`
	}

	rows := make([]row, 0)
	if err := u.data.DB(ctx).Table("quant_profit_result").
		Select("status, COUNT(*) AS result_count, CAST(COALESCE(SUM(profit_amount), 0) AS CHAR) AS profit_total, CAST(COALESCE(SUM(return_amount), 0) AS CHAR) AS return_total").
		Group("status").
		Order("status ASC").
		Scan(&rows).Error; err != nil {
		return nil, errors.New(500, "QUANT_RESULT_SUMMARY_ERROR", err.Error())
	}

	res := make([]*biz.QuantResultStatusSummary, 0, len(rows))
	for _, v := range rows {
		res = append(res, &biz.QuantResultStatusSummary{
			Status:      v.Status,
			ResultCount: v.ResultCount,
			ProfitTotal: defaultZero(v.ProfitTotal),
			ReturnTotal: defaultZero(v.ReturnTotal),
		})
	}
	return res, nil
}

func (u *UserRepo) GetRecentQuantFundLedgers(ctx context.Context, limit int) ([]*biz.QuantFundLedgerView, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	rows := make([]*biz.QuantFundLedgerView, 0)
	if err := u.data.DB(ctx).Table("quant_fund_ledger").
		Select("id, minted_id, token_id, user_addr, mint_addr, CAST(usdt_paid AS CHAR) AS usdt_paid, quant_bps, CAST(quant_amount AS CHAR) AS quant_amount, CAST(reserve_amount AS CHAR) AS reserve_amount, block_number, block_time, status, transfer_task_no, transfer_tx_hash, result_task_no, settlement_tx_hash, remark, created_at, updated_at").
		Order("id DESC").
		Limit(limit).
		Scan(&rows).Error; err != nil {
		return nil, errors.New(500, "QUANT_FUND_LEDGER_ERROR", err.Error())
	}
	for _, v := range rows {
		v.UsdtPaid = defaultZero(v.UsdtPaid)
		v.QuantAmount = defaultZero(v.QuantAmount)
		v.ReserveAmount = defaultZero(v.ReserveAmount)
	}
	return rows, nil
}

func (u *UserRepo) GetRecentQuantProfitResults(ctx context.Context, limit int) ([]*biz.QuantProfitResultView, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	rows := make([]*biz.QuantProfitResultView, 0)
	if err := u.data.DB(ctx).Table("quant_profit_result").
		Select("id, task_no, ledger_id, token_id, user_addr, mint_addr, CAST(quant_amount AS CHAR) AS quant_amount, profit_bps, CAST(profit_amount AS CHAR) AS profit_amount, CAST(return_amount AS CHAR) AS return_amount, CAST(settlement_reward AS CHAR) AS settlement_reward, result_mode, status, settlement_tx_hash, remark, created_at, updated_at").
		Order("id DESC").
		Limit(limit).
		Scan(&rows).Error; err != nil {
		return nil, errors.New(500, "QUANT_RESULT_LIST_ERROR", err.Error())
	}
	for _, v := range rows {
		v.QuantAmount = defaultZero(v.QuantAmount)
		v.ProfitAmount = defaultZero(v.ProfitAmount)
		v.ReturnAmount = defaultZero(v.ReturnAmount)
		v.SettlementReward = defaultZero(v.SettlementReward)
	}
	return rows, nil
}

func (u *UserRepo) GetRecentQuantTransferTasks(ctx context.Context, limit int) ([]*biz.QuantTransferTaskView, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	rows := make([]*biz.QuantTransferTaskView, 0)
	if err := u.data.DB(ctx).Table("quant_transfer_task").
		Select("id, task_no, ledger_count, quant_bps, CAST(total_amount AS CHAR) AS total_amount, from_addr, to_addr, tx_hash, status, remark, created_at, updated_at").
		Order("id DESC").
		Limit(limit).
		Scan(&rows).Error; err != nil {
		return nil, errors.New(500, "QUANT_TRANSFER_TASK_ERROR", err.Error())
	}
	for _, v := range rows {
		v.TotalAmount = defaultZero(v.TotalAmount)
	}
	return rows, nil
}

func defaultZero(value string) string {
	if strings.TrimSpace(value) == "" {
		return "0"
	}
	return value
}

// nftMintedOpenStatusDBFilter 按 token 是否在表内任一行标记为已开盒过滤，避免同一 token_id 多行时 MAX(id) 行与 open 写在另一行导致已解锁列表为空。
func nftMintedOpenStatusDBFilter(db *gorm.DB, openStatus uint64) *gorm.DB {
	if openStatus >= 2 {
		return db
	}
	if openStatus == 1 {
		return db.Where("EXISTS (SELECT 1 FROM nft_minted m2 WHERE m2.token_id = nft_minted.token_id AND m2.open_status = ?)", 1)
	}
	return db.Where("NOT EXISTS (SELECT 1 FROM nft_minted m2 WHERE m2.token_id = nft_minted.token_id AND m2.open_status = ?)", 1)
}

func (u *UserRepo) enrichNftMintedRowsOpenedAt(ctx context.Context, rows []*NftMinted) {
	if len(rows) == 0 {
		return
	}
	tokens := make([]uint64, 0)
	for _, v := range rows {
		if v.OpenStatus != 1 {
			tokens = append(tokens, v.TokenID)
		}
	}
	if len(tokens) == 0 {
		return
	}
	type aggRow struct {
		TokenID  uint64 `gorm:"column:token_id"`
		OpenedAt uint64 `gorm:"column:opened_at"`
	}
	var aggs []aggRow
	if err := u.data.DB(ctx).Table("nft_minted").Select("token_id, MAX(opened_at) AS opened_at").Where("token_id IN ?", tokens).Where("open_status = ?", 1).Group("token_id").Find(&aggs).Error; err != nil {
		return
	}
	openAtByTok := make(map[uint64]uint64, len(aggs))
	for _, a := range aggs {
		openAtByTok[a.TokenID] = a.OpenedAt
	}
	for i := range rows {
		if rows[i].OpenStatus != 1 {
			if oa, ok := openAtByTok[rows[i].TokenID]; ok {
				rows[i].OpenStatus = 1
				rows[i].OpenedAt = oa
			}
		}
	}
}

// GetNftMintedByAddressPage .
func (u *UserRepo) GetNftMintedByAddressPage(
	ctx context.Context,
	b *biz.Pagination,
	start uint64,
	end uint64,
	address []string,
	status uint64,
	order uint64,
	tier uint64,
	openStatus uint64,
	openAtOrder uint64,
) ([]*biz.NftMinted, error, int64) {

	var (
		count int64
		rows  []*NftMinted
	)

	res := make([]*biz.NftMinted, 0)

	// 同一 token_id 多行时：子查询只取 MAX(id)（不在子查询里筛 open_status），否则「已开」写在旧行、MAX(id) 行仍是未开时，未解锁列表会一直包含该 token。
	sub := u.data.DB(ctx).Table("nft_minted").Select("MAX(id)").Where("to_addr in(?)", address)

	if 0 < start && 0 < end {
		sub = sub.Where("block_time >= ?", start).Where("block_time <= ?", end)
	}

	if 2 > status {
		if 1 == status {
			sub = sub.Where("status=?", 1)
		} else {
			// 仓库（仅未上架）
			sub = sub.Where("status=?", 0)
			sub = sub.Where("token_id > ?", 0)
			sub = sub.Where("token_id NOT IN (SELECT DISTINCT token_id FROM nft_minted WHERE status = 1)")
		}
	}

	if 0 < tier {
		sub = sub.Where("tier=?", tier)
	}

	sub = sub.Group("token_id")

	base := nftMintedOpenStatusDBFilter(u.data.DB(ctx).Table("nft_minted").Where("id IN (?)", sub), openStatus)

	if err := base.Count(&count).Error; err != nil {
		return nil, errors.New(500, "NFT_MINTED_ERROR", err.Error()), 0
	}
	if openStatus == 1 {
		if count > 10 {
			count = 10
		}
		if b.PageNum > 1 {
			return res, nil, count
		}
	}

	q := nftMintedOpenStatusDBFilter(u.data.DB(ctx).Table("nft_minted").Where("id IN (?)", sub), openStatus)
	if 0 < order {
		q = q.Order("id asc")
	} else {
		q = q.Order("id desc")
	}

	if 0 < openAtOrder {
		if 1 < openAtOrder {
			q = q.Order("opened_at asc")
		} else {
			q = q.Order("opened_at desc")
		}
	}
	if openStatus == 1 {
		q = q.Limit(10)
	}

	if err := q.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&rows).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("NFT_MINTED_NOT_FOUND", "minted not found"), 0
		}
		return nil, errors.New(500, "NFT_MINTED_ERROR", err.Error()), 0
	}

	if 2 > openStatus && openStatus == 1 {
		u.enrichNftMintedRowsOpenedAt(ctx, rows)
	}

	for _, v := range rows {
		sourceType := "market_or_transfer"
		for _, addr := range address {
			if strings.EqualFold(strings.TrimSpace(addr), strings.TrimSpace(v.MintAddr)) {
				sourceType = "self_minted"
				break
			}
		}
		res = append(res, &biz.NftMinted{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			ToAddr:  v.ToAddr,
			TokenID: v.TokenID,

			Tier:     v.Tier,
			UsdtPaid: v.UsdtPaid,

			Status:   v.Status,
			ListedAt: v.ListedAt,

			OpenStatus: v.OpenStatus,
			OpenedAt:   v.OpenedAt,

			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
			MintAddr:   v.MintAddr,
			SourceType: sourceType,
		})
	}

	return res, nil, count
}

func (u *UserRepo) GetNftMintedByAddressAll(
	ctx context.Context,
	start uint64,
	end uint64,
	address []string,
	status uint64,
	order uint64,
	tier uint64,
	openStatus uint64,
	openAtOrder uint64,
) ([]*biz.NftMinted, error) {
	rows := make([]*NftMinted, 0)
	res := make([]*biz.NftMinted, 0)

	sub := u.data.DB(ctx).Table("nft_minted").Select("MAX(id)").Where("to_addr in(?)", address)

	if 0 < start && 0 < end {
		sub = sub.Where("block_time >= ?", start).Where("block_time <= ?", end)
	}

	if 2 > status {
		if 1 == status {
			sub = sub.Where("status=?", 1)
		} else {
			sub = sub.Where("status=?", 0)
			sub = sub.Where("token_id > ?", 0)
			sub = sub.Where("token_id NOT IN (SELECT DISTINCT token_id FROM nft_minted WHERE status = 1)")
		}
	}

	if 0 < tier {
		sub = sub.Where("tier=?", tier)
	}

	sub = sub.Group("token_id")

	q := nftMintedOpenStatusDBFilter(u.data.DB(ctx).Table("nft_minted").Where("id IN (?)", sub), openStatus)
	if 0 < order {
		q = q.Order("id asc")
	} else {
		q = q.Order("id desc")
	}

	if 0 < openAtOrder {
		if 1 < openAtOrder {
			q = q.Order("opened_at asc")
		} else {
			q = q.Order("opened_at desc")
		}
	}
	if openStatus == 1 {
		q = q.Limit(10)
	}

	if err := q.Find(&rows).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("NFT_MINTED_NOT_FOUND", "minted not found")
		}
		return nil, errors.New(500, "NFT_MINTED_ERROR", err.Error())
	}

	if 2 > openStatus && openStatus == 1 {
		u.enrichNftMintedRowsOpenedAt(ctx, rows)
	}

	for _, v := range rows {
		sourceType := "market_or_transfer"
		for _, addr := range address {
			if strings.EqualFold(strings.TrimSpace(addr), strings.TrimSpace(v.MintAddr)) {
				sourceType = "self_minted"
				break
			}
		}
		res = append(res, &biz.NftMinted{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,
			ToAddr:      v.ToAddr,
			TokenID:     v.TokenID,
			Tier:        v.Tier,
			UsdtPaid:    v.UsdtPaid,
			Status:      v.Status,
			ListedAt:    v.ListedAt,
			OpenStatus:  v.OpenStatus,
			OpenedAt:    v.OpenedAt,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			MintAddr:    v.MintAddr,
			SourceType:  sourceType,
		})
	}

	return res, nil
}

func (u *UserRepo) InsertUserRegistered(ctx context.Context, iData *biz.UserRegistered) error {
	var existing UserRegistered
	err := u.data.DB(ctx).Table("user_registered").
		Where("LOWER(user_addr) = LOWER(?)", iData.UserAddr).
		Order("block_number desc, log_index desc, id desc").
		First(&existing).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(500, "CREATE_USER_REGISTERED_ERROR", err.Error())
	}

	incoming := UserRegistered{
		BlockNumber: iData.BlockNumber,
		BlockTime:   iData.BlockTime,
		LogIndex:    iData.LogIndex,
		UserAddr:    iData.UserAddr,
		ParentAddr:  iData.ParentAddr,
		TopAddr:     iData.TopAddr,
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := u.data.DB(ctx).Table("user_registered").Create(&incoming).Error; err != nil {
			return errors.New(500, "CREATE_USER_REGISTERED_ERROR", "信息创建失败")
		}
		return nil
	}

	shouldUpdate := incoming.BlockNumber > existing.BlockNumber ||
		(incoming.BlockNumber == existing.BlockNumber && incoming.LogIndex >= existing.LogIndex)
	if !shouldUpdate {
		return nil
	}

	if err := u.data.DB(ctx).Table("user_registered").Where("id = ?", existing.ID).Updates(map[string]interface{}{
		"block_number": incoming.BlockNumber,
		"block_time":   incoming.BlockTime,
		"log_index":    incoming.LogIndex,
		"user_addr":    incoming.UserAddr,
		"parent_addr":  incoming.ParentAddr,
		"top_addr":     incoming.TopAddr,
	}).Error; err != nil {
		return errors.New(500, "CREATE_USER_REGISTERED_ERROR", "信息更新失败")
	}
	return nil
}

func (u *UserRepo) GetUserRegisteredLast(ctx context.Context) (*biz.UserRegistered, error) {
	var v UserRegistered

	if err := u.data.DB(ctx).Table("user_registered").Order("block_number DESC, id DESC").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "USER_REGISTERED_ERROR", err.Error())
	}

	return &biz.UserRegistered{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		UserAddr:   v.UserAddr,
		ParentAddr: v.ParentAddr,
		TopAddr:    v.TopAddr,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetUserRegisteredByAddress(ctx context.Context, address string) (*biz.UserRegistered, error) {
	var v UserRegistered
	if err := u.data.DB(ctx).Table("user_registered").
		Where("LOWER(user_addr) = LOWER(?)", address).
		Order("block_number desc, log_index desc, id desc").
		First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "USER_REGISTERED_ERROR", err.Error())
	}
	return &biz.UserRegistered{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,
		UserAddr:    v.UserAddr,
		ParentAddr:  v.ParentAddr,
		TopAddr:     v.TopAddr,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}, nil
}

// GetUserRCount .
func (u *UserRepo) GetUserRCount(ctx context.Context) int64 {
	var (
		count int64
	)

	u.data.db.Table("user_registered").Where("block_time > ?", 0).Count(&count)
	return count
}

// GetUserRCountBySe .
func (u *UserRepo) GetUserRCountBySe(ctx context.Context, start, end uint64) int64 {
	var (
		count int64
	)

	u.data.db.Table("user_registered").Where("block_time > ?", 0).Where("block_time >= ?", start).Where("block_time <= ?", end).Count(&count)
	return count
}

// GetMintNftCountBySe .
func (u *UserRepo) GetMintNftCountBySe(ctx context.Context, start, end uint64) int64 {
	var (
		count int64
	)

	u.data.db.Table("nft_minted").Where("token_id > ?", 0).Where("block_time > ?", 0).Where("block_time >= ?", start).Where("block_time <= ?", end).Count(&count)
	return count
}

// GetMintNftUsdtPaidSumBySe 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetMintNftUsdtPaidSumBySe(ctx context.Context, start, end uint64) string {
	var sum string
	u.data.db.Table("nft_minted").
		Select("COALESCE(SUM(usdt_paid), 0)").
		Where("token_id > ?", 0).
		Where("block_time > ?", 0).
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Scan(&sum)
	return sum
}

// GetMintNftCount .
func (u *UserRepo) GetMintNftCount(paidType uint64) int64 {
	var (
		count int64
	)

	instance := u.data.db.Table("nft_minted").Where("token_id > ?", 0).Where("block_time > ?", 0)
	if 0 < paidType {
		instance = instance.Where("tier", paidType)
	}
	instance.Count(&count)
	return count
}

// GetMintNftNotOpenCount .
func (u *UserRepo) GetMintNftNotOpenCount(paidType uint64) int64 {
	var (
		count int64
	)

	instance := u.data.db.Table("nft_minted").Where("token_id > ?", 0).Where("block_time > ?", 0).Where("open_status=?", 0)
	if 0 < paidType {
		instance = instance.Where("tier", paidType)
	}
	instance.Count(&count)
	return count
}

// GetMintNftUsdtPaidSum 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetMintNftUsdtPaidSum(paidType uint64) string {
	var sum string
	instance := u.data.db.Table("nft_minted").Where("token_id > ?", 0).Where("block_time > ?", 0)
	if 0 < paidType {
		instance = instance.Where("tier", paidType)
	}

	instance.Select("COALESCE(SUM(usdt_paid), 0)").Scan(&sum)
	return sum
}

// GetMintNftNotOpenUsdtPaidSum 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetMintNftNotOpenUsdtPaidSum(paidType uint64) string {
	var sum string
	instance := u.data.db.Table("nft_minted").Where("token_id > ?", 0).Where("block_time > ?", 0).Where("open_status=?", 0)
	if 0 < paidType {
		instance = instance.Where("tier", paidType)
	}

	instance.Select("COALESCE(SUM(usdt_paid), 0)").Scan(&sum)
	return sum
}

// GetNftBuyCountBySe .
func (u *UserRepo) GetNftBuyCountBySe(ctx context.Context, start, end uint64) int64 {
	var (
		count int64
	)

	u.data.db.Table("nft_market_purchase").Where("block_time > ?", 0).Where("block_time >= ?", start).Where("block_time <= ?", end).Count(&count)
	return count
}

// GetNftBuySumBySe 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetNftBuySumBySe(ctx context.Context, start, end uint64) string {
	var sum string
	u.data.db.Table("nft_market_purchase").
		Select("COALESCE(SUM(price_usdt), 0)").
		Where("block_time > ?", 0).
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Scan(&sum)
	return sum
}

// GetNftBuyCount .
func (u *UserRepo) GetNftBuyCount() int64 {
	var (
		count int64
	)

	u.data.db.Table("nft_market_purchase").Where("block_time > ?", 0).Count(&count)
	return count
}

// GetNftBuySum 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetNftBuySum() string {
	var sum string
	u.data.db.Table("nft_market_purchase").
		Select("COALESCE(SUM(price_usdt), 0)").
		Where("block_time > ?", 0).
		Scan(&sum)
	return sum
}

// GetNftOpenCountBySe .
func (u *UserRepo) GetNftOpenCountBySe(ctx context.Context, start, end uint64) int64 {
	var (
		count int64
	)

	u.data.db.Table("nft_opened").Where("token_id > ?", 0).Where("block_time > ?", 0).Where("block_time >= ?", start).Where("block_time <= ?", end).Count(&count)
	return count
}

// GetNftOpenSumB 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetNftOpenSum() string {
	var sum string
	u.data.db.Table("nft_opened").
		Select("COALESCE(SUM(reward), 0)").
		Where("token_id > ?", 0).
		Where("block_time > ?", 0).
		Scan(&sum)
	return sum
}

// GetNftOpenSumBySe 统计时间段内已开盲盒的 reward 总和（check_status=1 已结算）
func (u *UserRepo) GetNftOpenSumBySe(ctx context.Context, start, end uint64) string {
	var sum string
	u.data.DB(ctx).Table("nft_opened").
		Select("COALESCE(SUM(reward), 0)").
		Where("token_id > ?", 0).
		Where("check_status = ?", 1).
		Where("block_time > ?", 0).
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Scan(&sum)
	return sum
}

func (u *UserRepo) GetUserOpenProfitSumBySe(ctx context.Context, address string, start, end uint64) string {
	var sum string
	q := u.data.DB(ctx).Table("nft_opened").
		Select("COALESCE(SUM(reward), 0)").
		Where("user_addr = ?", address).
		Where("check_status = ?", 1)
	if start > 0 && end > 0 {
		q = q.Where("block_time >= ?", start).Where("block_time <= ?", end)
	}
	q.Scan(&sum)
	return defaultZero(sum)
}

func (u *UserRepo) GetUserOpenProfitSum(ctx context.Context, address string) string {
	return u.GetUserOpenProfitSumBySe(ctx, address, 0, 0)
}

func (u *UserRepo) GetUserMintProfitSum(ctx context.Context, address string) string {
	var sum string
	u.data.DB(ctx).Table("nft_opened as o").
		Joins("JOIN nft_minted m ON o.token_id = m.token_id").
		Select("COALESCE(SUM(o.reward), 0)").
		Where("m.mint_addr = ?", address).
		Where("o.check_status = ?", 1).
		Scan(&sum)
	return defaultZero(sum)
}

func (u *UserRepo) GetUserOpenedCountBySe(ctx context.Context, address string, start, end uint64) int64 {
	var count int64
	q := u.data.DB(ctx).Table("nft_opened").
		Where("user_addr = ?", address).
		Where("check_status = ?", 1)
	if start > 0 && end > 0 {
		q = q.Where("block_time >= ?", start).Where("block_time <= ?", end)
	}
	q.Count(&count)
	return count
}

func (u *UserRepo) GetUserOpenedCount(ctx context.Context, address string) int64 {
	return u.GetUserOpenedCountBySe(ctx, address, 0, 0)
}

// GetOpenedBoxListBySe 已结算盲盒列表（含铸造本金），按 block_time 倒序，支持分页
func (u *UserRepo) GetOpenedBoxListBySe(ctx context.Context, b *biz.Pagination, start, end uint64) ([]*biz.OpenedBoxWithMint, error, int64) {
	var count int64
	type Row struct {
		TokenID   uint64
		UserAddr  string
		Reward    float64
		BlockTime uint64
		UsdtPaid  float64
	}
	base := u.data.DB(ctx).Table("nft_opened as o").
		Select("o.token_id, o.user_addr, o.reward, o.block_time, COALESCE(m.usdt_paid, 0) as usdt_paid").
		Joins("LEFT JOIN nft_minted m ON o.token_id = m.token_id").
		Where("o.token_id > ?", 0).
		Where("o.check_status = ?", 1).
		Where("o.block_time > ?", 0).
		Where("o.block_time >= ?", start).
		Where("o.block_time <= ?", end)
	base.Count(&count)
	var rows []Row
	q := base.Order("o.block_time DESC")
	if err := q.Scopes(Paginate(b.PageNum, b.PageSize)).Scan(&rows).Error; err != nil {
		return nil, err, 0
	}
	res := make([]*biz.OpenedBoxWithMint, 0, len(rows))
	for _, r := range rows {
		res = append(res, &biz.OpenedBoxWithMint{
			TokenID:   r.TokenID,
			UserAddr:  r.UserAddr,
			Reward:    r.Reward,
			BlockTime: r.BlockTime,
			UsdtPaid:  r.UsdtPaid,
		})
	}
	return res, nil, count
}

// GetOpenedBoxProfitLossCountBySe 时间段内盈利/亏损盲盒数量（check_status=1）
func (u *UserRepo) GetOpenedBoxProfitLossCountBySe(ctx context.Context, start, end uint64) (profitCount, lossCount int64) {
	u.data.DB(ctx).Table("nft_opened").
		Where("token_id > ?", 0).
		Where("check_status = ?", 1).
		Where("block_time > ?", 0).
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Where("reward > 0").
		Count(&profitCount)
	u.data.DB(ctx).Table("nft_opened").
		Where("token_id > ?", 0).
		Where("check_status = ?", 1).
		Where("block_time > ?", 0).
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Where("reward < 0").
		Count(&lossCount)
	return profitCount, lossCount
}

// GetNftMarketPurchaseFeeSumBySe 二级市场交易手续费汇总（时间段）
func (u *UserRepo) GetNftMarketPurchaseFeeSumBySe(ctx context.Context, start, end uint64) (feeUsdtSum, feeBSum string) {
	u.data.DB(ctx).Table("nft_market_purchase").
		Select("COALESCE(SUM(fee_usdt), 0)").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Scan(&feeUsdtSum)
	u.data.DB(ctx).Table("nft_market_purchase").
		Select("COALESCE(SUM(fee_b), 0)").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Scan(&feeBSum)
	return feeUsdtSum, feeBSum
}

func (u *UserRepo) ListRuntimeConfigs(ctx context.Context, prefix string) ([]*biz.RuntimeConfigEntry, error) {
	var rows []RuntimeConfig
	q := u.data.DB(ctx).Table("runtime_config").Order("config_key asc")
	if strings.TrimSpace(prefix) != "" {
		q = q.Where("config_key LIKE ?", prefix+"%")
	}
	if err := q.Find(&rows).Error; err != nil {
		return nil, errors.New(500, "RUNTIME_CONFIG_ERROR", err.Error())
	}
	res := make([]*biz.RuntimeConfigEntry, 0, len(rows))
	for _, row := range rows {
		res = append(res, &biz.RuntimeConfigEntry{
			Key:         row.ConfigKey,
			Value:       row.ConfigValue,
			Description: row.Description,
		})
	}
	return res, nil
}

func (u *UserRepo) UpsertRuntimeConfig(ctx context.Context, key, value, description string) error {
	var row RuntimeConfig
	err := u.data.DB(ctx).Table("runtime_config").Where("config_key = ?", key).First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(500, "RUNTIME_CONFIG_ERROR", err.Error())
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		row = RuntimeConfig{
			ConfigKey:   key,
			ConfigValue: value,
			Description: description,
		}
		if err := u.data.DB(ctx).Table("runtime_config").Create(&row).Error; err != nil {
			return errors.New(500, "RUNTIME_CONFIG_ERROR", err.Error())
		}
		return nil
	}
	updates := map[string]interface{}{
		"config_value": value,
	}
	if description != "" {
		updates["description"] = description
	}
	if err := u.data.DB(ctx).Table("runtime_config").Where("id = ?", row.ID).Updates(updates).Error; err != nil {
		return errors.New(500, "RUNTIME_CONFIG_ERROR", err.Error())
	}
	return nil
}
