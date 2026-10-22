package data

import (
	"context"
	stdlog "log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"requestEth/internal/quantfund"

	"requestEth/internal/biz"
	"requestEth/internal/conf"
)

// EnvSkipAutoMigrate 为 "1" 时跳过 GORM AutoMigrate（生产库表已存在且数据量大时，迁移可能长时间锁表导致 API 无法启动）。
// 首次部署或表结构变更时请勿设置，或单独执行迁移后再启动 API。
const EnvSkipAutoMigrate = "REQUESTETH_SKIP_AUTOMIGRATE"

// EnvDatabaseSource 若设置则覆盖配置文件中的 data.database.source（仅主服务；便于本机 XAMPP 等不写死密码进 yaml）。
const EnvDatabaseSource = "REQUESTETH_DATABASE_SOURCE"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewTransaction, NewUserRepo, NewReferralPendingRepo)

type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// 用来承载事务的上下文
type contextTxKey struct{}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:  db,
		rdb: rdb,
	}, cleanup, nil
}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// ExecTx gorm Transaction
func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// DB 根据此方法来判断当前的 db 是不是使用 事务的 DB
func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

// NewDB .
func NewDB(c *conf.Data) *gorm.DB {
	dsn := ""
	if dbconf := c.GetDatabase(); dbconf != nil {
		dsn = dbconf.GetSource()
	}
	if v := os.Getenv(EnvDatabaseSource); strings.TrimSpace(v) != "" {
		dsn = strings.TrimSpace(v)
	}
	stdlog.Print("[requestEth] NewDB: enter (next: sql.log -> gorm.Open -> AutoMigrate)")
	logPath := filepath.Join("log", "sql.log")
	if _, err := os.Stat(logPath); err != nil {
		logPath = filepath.Join("..", "..", "log", "sql.log")
	}
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("failed opening sql.log: %v", err)
		panic("failed opening sql.log")
	}

	// 终端打印输入 sql 执行记录
	newLogger := logger.New(
		stdlog.New(f, "\r\n", stdlog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢查询 SQL 阈值
			Colorful:      true,        // 禁用彩色打印
			//IgnoreRecordNotFoundError: false,
			LogLevel: logger.Error, // Log lever
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{
			//SingularTable: true, // 表名是否加 s
		},
	})

	if err != nil {
		log.Errorf("failed opening MySQL: %v", err)
		panic("failed to connect database")
	}

	// 自动迁移：创建 nft_minted、nft_opened 等表（链上事件同步、referral、blind_settle 依赖）
	// 表多或磁盘慢时可能耗时数分钟；未结束前 HTTP 不会监听，docker logs 若只有一行多半是卡在这里
	if os.Getenv(EnvSkipAutoMigrate) == "1" {
		stdlog.Print("[requestEth] database AutoMigrate skipped (" + EnvSkipAutoMigrate + "=1)")
	} else {
		stdlog.Print("[requestEth] database AutoMigrate start...")
		if err := db.AutoMigrate(
			&SwapTrade{}, &PrimaryBuy{}, &PrimarySell{}, &RewardNotified{}, &RewardDetail{},
			&NftMarketPurchase{}, &NftMinted{}, &NftMarketListed{}, &NftMarketUnlisted{},
			&NftOpened{}, &NftTransfer{}, &UserRegistered{}, &ReferralPending{}, &RuntimeConfig{},
			&quantfund.QuantFundLedger{}, &quantfund.QuantTransferTask{}, &quantfund.QuantProfitResult{},
		); err != nil {
			log.Errorf("auto migrate failed: %v", err)
			panic("failed to migrate database")
		}
		stdlog.Print("[requestEth] database AutoMigrate done")
	}

	return db
}

func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		DialTimeout:  time.Second * 2,
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	return rdb
}

// Paginate 分页
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
