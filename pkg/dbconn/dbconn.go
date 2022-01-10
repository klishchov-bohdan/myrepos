package dbconn

import (
	"gorm.io/gorm"
	"time"
)

type DBConn interface {
	SetConnPool(maxConns, idleConns int, lifeTimeConns time.Duration) error
	GetDB() *gorm.DB
}
