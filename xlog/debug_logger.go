package xlog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/iGoogle-ink/gotil"
)

type DebugLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (i *DebugLogger) logOut(format *string, v ...interface{}) {
	i.once.Do(func() {
		i.init()
	})
	if format != nil {
		i.logger.Output(3, fmt.Sprintf(*format, v...))
		return
	}
	i.logger.Output(3, fmt.Sprintln(v...))
}

func (i *DebugLogger) init() {
	version := runtime.Version()
	vv := strings.Split(version, ".")
	fmt.Println("version:", vv)
	fmt.Println("version:", vv[1])
	if gotil.String2Int(vv[1]) >= 14 {
		i.logger = log.New(os.Stdout, Cyan+" [DEBUG] "+Reset, log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
		return
	}
	i.logger = log.New(os.Stdout, Cyan+" [DEBUG] >> "+Reset, log.Lshortfile|log.Ldate|log.Lmicroseconds)
}
