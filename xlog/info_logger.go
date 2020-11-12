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

type InfoLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (i *InfoLogger) logOut(col *ColorType, format *string, v ...interface{}) {
	i.once.Do(func() {
		i.init()
	})
	if col != nil {
		if format != nil {
			i.logger.Output(3, string(*col)+fmt.Sprintf(*format, v...))
			return
		}
		i.logger.Output(3, string(*col)+fmt.Sprintln(v...))
		return
	}
	if format != nil {
		i.logger.Output(3, fmt.Sprintf(*format, v...))
		return
	}
	i.logger.Output(3, fmt.Sprintln(v...))
}

func (i *InfoLogger) init() {
	if gotil.String2Int(strings.Split(runtime.Version(), ".")[1]) >= 14 {
		i.logger = log.New(os.Stdout, "[INFO] >> ", 64|log.Lshortfile|log.Ldate|log.Lmicroseconds)
		return
	}
	i.logger = log.New(os.Stdout, "[INFO] ", log.Lshortfile|log.Ldate|log.Lmicroseconds)
}
