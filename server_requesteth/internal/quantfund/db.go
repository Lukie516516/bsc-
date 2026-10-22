package quantfund

import (
	"fmt"
	"sync/atomic"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var taskSeq uint64

func OpenDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&QuantFundLedger{}, &QuantTransferTask{}, &QuantProfitResult{})
}

func NewTaskNo() string {
	seq := atomic.AddUint64(&taskSeq, 1)
	now := time.Now()
	return "QTF" + now.Format("20060102150405") + fmt.Sprintf("%09d%06d", now.Nanosecond(), seq%1000000)
}
