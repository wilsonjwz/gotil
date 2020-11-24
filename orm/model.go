package orm

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	ErrNoRow                = sql.ErrNoRows
	ErrRecordNotFound       = gorm.ErrRecordNotFound
	ErrCantStartTransaction = gorm.ErrCantStartTransaction
	ErrInvalidSQL           = gorm.ErrInvalidSQL
	ErrInvalidTransaction   = gorm.ErrInvalidTransaction
	ErrUnaddressable        = gorm.ErrUnaddressable
)

// MySQLConfig mysql config.
type MySQLConfig struct {
	DSN            string        // data source name.
	MaxOpenConn    int           // pool
	MaxIdleConn    int           // pool
	MaxConnTimeout time.Duration // connect max life time. second
	ShowSQL        bool
}
