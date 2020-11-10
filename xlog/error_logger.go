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

type ErrorLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (e *ErrorLogger) logOut(format *string, v ...interface{}) {
	e.once.Do(func() {
		e.init()
	})
	if format != nil {
		e.logger.Output(3, fmt.Sprintf(*format, v...))
		//i.logger.Writer().Write(stack())
		return
	}
	e.logger.Output(3, fmt.Sprintln(v...))
	//i.logger.Writer().Write(stack())
}

func (e *ErrorLogger) init() {
	version := strings.Split(runtime.Version(), ".")
	fmt.Println("version:", version)
	fmt.Println("version:", version[1])
	if gotil.String2Int(version[1]) >= 14 {
		e.logger = log.New(os.Stdout, Red+" [ERROR] "+Reset, log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
		return
	}
	e.logger = log.New(os.Stdout, Red+" [ERROR] >> "+Reset, log.Lshortfile|log.Ldate|log.Lmicroseconds)
}

//func stack() (bs []byte) {
//	var buf [2 << 10]byte
//	runtime.Stack(buf[:], false)
//	return buf[:]
//}
