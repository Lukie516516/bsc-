package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// ReferralPending 铸造分佣待发记录（B推广者0.3% + D团队长0.2%，不满足时0.5%归项目方）
// 来源：铸造盲盒(nft_minted)，二级市场购买不分
type ReferralPending struct {
	ID uint64 `gorm:"primaryKey;autoIncrement;column:id"`

	MintedID   uint64 `gorm:"column:minted_id;index:idx_minted;comment:关联 nft_minted.id（铸造分佣来源），应用层防重"`
	PurchaseID uint64 `gorm:"column:purchase_id;default:0;index:idx_purchase;comment:历史兼容 nft_market_purchase.id，新记录为0"`
	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	TokenID     uint64 `gorm:"column:token_id;not null;index:idx_token_id"`

	Minter string `gorm:"column:minter;type:varchar(42);not null;comment:铸造者 MintAddr"`
	Seller string `gorm:"column:seller;type:varchar(42);default:0x0;comment:铸造场景为0x0"`
	Buyer  string `gorm:"column:buyer;type:varchar(42);default:0x0;comment:铸造场景为0x0"`
	Price  string `gorm:"column:price;type:varchar(78);not null;comment:铸造金额 usdt_paid wei"`

	// 推荐人：B=推广者 parentOf(minter)，D=团队长 topOf(minter) 存于 addr_c
	AddrB string `gorm:"column:addr_b;type:varchar(42);not null;comment:推广者B 0.3%"`
	AddrC string `gorm:"column:addr_c;type:varchar(42);not null;comment:团队长D 0.2% (isTopEligible:无上级+质押300HRX)"`

	// 待发金额（wei 字符串）
	AmountToB string `gorm:"column:amount_to_b;type:varchar(78);not null;default:0;comment:给B的USDT 0.3%"`
	AmountToC string `gorm:"column:amount_to_c;type:varchar(78);not null;default:0;comment:给D的USDT 0.2%"`
	AmountToProject string `gorm:"column:amount_to_project;type:varchar(78);not null;default:0;comment:沉淀项目方"`

	// 原因：0=正常 1=仅D不合格 2=仅B不合格 3=无上级 4=B和D都不合格
	Reason   uint8 `gorm:"column:reason;not null;default:0"`
	PaidAt   *time.Time `gorm:"column:paid_at;comment:打款时间"`
	PaidTxHash string `gorm:"column:paid_tx_hash;type:varchar(66);comment:打款交易哈希"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

// TableName 表名
func (ReferralPending) TableName() string {
	return "referral_pending"
}

// ReferralPendingRepo 铸造分佣数据仓库
type ReferralPendingRepo struct {
	data *Data
	log  *log.Helper
}

// NewReferralPendingRepo .
func NewReferralPendingRepo(data *Data, logger log.Logger) *ReferralPendingRepo {
	return &ReferralPendingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Insert 插入待发记录
func (r *ReferralPendingRepo) Insert(ctx context.Context, m *ReferralPending) error {
	if err := r.data.DB(ctx).Create(m).Error; err != nil {
		return errors.New(500, "CREATE_REFERRAL_PENDING_ERROR", err.Error())
	}
	return nil
}

// ExistsByMintedID 是否已有该铸造的分佣记录
func (r *ReferralPendingRepo) ExistsByMintedID(ctx context.Context, mintedID uint64) (bool, error) {
	if mintedID == 0 {
		return false, nil
	}
	var c int64
	if err := r.data.DB(ctx).Model(&ReferralPending{}).Where("minted_id = ?", mintedID).Count(&c).Error; err != nil {
		return false, err
	}
	return c > 0, nil
}

// ExistsByPurchaseID 是否已有该 purchase 的分佣记录（历史兼容）
func (r *ReferralPendingRepo) ExistsByPurchaseID(ctx context.Context, purchaseID uint64) (bool, error) {
	var c int64
	if err := r.data.DB(ctx).Model(&ReferralPending{}).Where("purchase_id = ?", purchaseID).Count(&c).Error; err != nil {
		return false, err
	}
	return c > 0, nil
}

// ListPending 分页查询待打款记录
func (r *ReferralPendingRepo) ListPending(ctx context.Context, page, pageSize int, paid *bool) ([]*ReferralPending, int64, error) {
	var list []*ReferralPending
	var total int64

	q := r.data.DB(ctx).Model(&ReferralPending{})
	if paid != nil {
		if *paid {
			q = q.Where("paid_at IS NOT NULL")
		} else {
			q = q.Where("paid_at IS NULL")
		}
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * pageSize
	if pageSize <= 0 {
		pageSize = 20
	}
	if err := q.Order("id ASC").Offset(offset).Limit(pageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// MarkPaid 标记已打款
func (r *ReferralPendingRepo) MarkPaid(ctx context.Context, id uint64, txHash string) error {
	now := time.Now()
	res := r.data.DB(ctx).Model(&ReferralPending{}).Where("id = ?", id).
		Updates(map[string]interface{}{"paid_at": now, "paid_tx_hash": txHash, "updated_at": now})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// GetByIDs 按 referral_pending.id 批量查询
func (r *ReferralPendingRepo) GetByIDs(ctx context.Context, ids []uint64) ([]*ReferralPending, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var list []*ReferralPending
	if err := r.data.DB(ctx).Where("id IN ? AND paid_at IS NULL", ids).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// GetByPurchaseIDs 按 purchase_id 批量查询
func (r *ReferralPendingRepo) GetByPurchaseIDs(ctx context.Context, ids []uint64) (map[uint64]*ReferralPending, error) {
	var list []*ReferralPending
	if err := r.data.DB(ctx).Where("purchase_id IN ?", ids).Find(&list).Error; err != nil {
		return nil, err
	}
	m := make(map[uint64]*ReferralPending)
	for _, v := range list {
		m[v.PurchaseID] = v
	}
	return m, nil
}
