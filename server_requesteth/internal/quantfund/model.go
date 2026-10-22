package quantfund

import "time"

const (
	LedgerStatusPendingTransfer uint8 = iota
	LedgerStatusTransferLocked
	LedgerStatusTransferred
	LedgerStatusTransferFailed
	LedgerStatusResultReady
	LedgerStatusSettled
	LedgerStatusSettlementFailed
)

const (
	ResultStatusPendingSettle uint8 = iota
	ResultStatusSettleLocked
	ResultStatusSettled
	ResultStatusSettleFailed
)

type QuantFundLedger struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	MintedID uint64 `gorm:"column:minted_id;not null;index:idx_minted_id"`
	TokenID  uint64 `gorm:"column:token_id;not null;uniqueIndex:uk_token_id"`

	UserAddr string `gorm:"column:user_addr;type:varchar(42);not null;default:''"`
	MintAddr string `gorm:"column:mint_addr;type:varchar(42);not null;default:''"`

	UsdtPaid      string `gorm:"column:usdt_paid;type:decimal(65,18);not null;default:0"`
	QuantBps      int    `gorm:"column:quant_bps;not null;default:9500"`
	QuantAmount   string `gorm:"column:quant_amount;type:decimal(65,18);not null;default:0"`
	ReserveAmount string `gorm:"column:reserve_amount;type:decimal(65,18);not null;default:0"`

	BlockNumber uint64 `gorm:"column:block_number;not null;index:idx_status_block,priority:2"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`

	Status           uint8  `gorm:"column:status;not null;default:0;index:idx_status_block,priority:1"`
	TransferTaskNo   string `gorm:"column:transfer_task_no;type:varchar(64);not null;default:'';index:idx_transfer_task_no"`
	TransferTxHash   string `gorm:"column:transfer_tx_hash;type:varchar(80);not null;default:''"`
	ResultTaskNo     string `gorm:"column:result_task_no;type:varchar(64);not null;default:'';index:idx_result_task_no"`
	SettlementTxHash string `gorm:"column:settlement_tx_hash;type:varchar(80);not null;default:''"`
	Remark           string `gorm:"column:remark;type:varchar(255);not null;default:''"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (QuantFundLedger) TableName() string { return "quant_fund_ledger" }

type QuantTransferTask struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	TaskNo      string `gorm:"column:task_no;type:varchar(64);not null;uniqueIndex:uk_task_no"`
	LedgerCount int    `gorm:"column:ledger_count;not null;default:0"`
	QuantBps    int    `gorm:"column:quant_bps;not null;default:9500"`
	TotalAmount string `gorm:"column:total_amount;type:decimal(65,18);not null;default:0"`

	FromAddr string `gorm:"column:from_addr;type:varchar(42);not null;default:''"`
	ToAddr   string `gorm:"column:to_addr;type:varchar(42);not null;default:''"`
	TxHash   string `gorm:"column:tx_hash;type:varchar(80);not null;default:''"`

	Status uint8  `gorm:"column:status;not null;default:0;index:idx_task_status"`
	Remark string `gorm:"column:remark;type:varchar(255);not null;default:''"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (QuantTransferTask) TableName() string { return "quant_transfer_task" }

type QuantProfitResult struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	TaskNo   string `gorm:"column:task_no;type:varchar(64);not null;uniqueIndex:uk_result_task_no"`
	LedgerID uint64 `gorm:"column:ledger_id;not null;uniqueIndex:uk_result_ledger_id"`
	TokenID  uint64 `gorm:"column:token_id;not null;index:idx_result_token_id"`

	UserAddr string `gorm:"column:user_addr;type:varchar(42);not null;default:''"`
	MintAddr string `gorm:"column:mint_addr;type:varchar(42);not null;default:''"`

	QuantAmount      string `gorm:"column:quant_amount;type:decimal(65,18);not null;default:0"`
	ProfitBps        int    `gorm:"column:profit_bps;not null;default:0"`
	ProfitAmount     string `gorm:"column:profit_amount;type:decimal(65,18);not null;default:0"`
	ReturnAmount     string `gorm:"column:return_amount;type:decimal(65,18);not null;default:0"`
	SettlementReward string `gorm:"column:settlement_reward;type:decimal(65,18);not null;default:0"`
	ResultMode       string `gorm:"column:result_mode;type:varchar(32);not null;default:'simulate'"`

	Status           uint8  `gorm:"column:status;not null;default:0;index:idx_result_status"`
	SettlementTxHash string `gorm:"column:settlement_tx_hash;type:varchar(80);not null;default:''"`
	Remark           string `gorm:"column:remark;type:varchar(255);not null;default:''"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (QuantProfitResult) TableName() string { return "quant_profit_result" }

type MintedRow struct {
	ID          uint64 `gorm:"column:id"`
	TokenID     uint64 `gorm:"column:token_id"`
	ToAddr      string `gorm:"column:to_addr"`
	MintAddr    string `gorm:"column:mint_addr"`
	UsdtPaid    string `gorm:"column:usdt_paid"`
	BlockNumber uint64 `gorm:"column:block_number"`
	BlockTime   uint64 `gorm:"column:block_time"`
}
