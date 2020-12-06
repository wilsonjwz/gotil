package orm

import (
	"database/sql"

	"github.com/iGoogle-ink/gotil/xtime"
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
	DSN            string         // data source name.
	MaxOpenConn    int            // pool
	MaxIdleConn    int            // pool
	MaxConnTimeout xtime.Duration // connect max life time. exp:10s、2m、1m10s
	ShowSQL        bool
}
