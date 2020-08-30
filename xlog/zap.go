package xlog

import (
	"sync"
	"time"

	"github.com/iGoogle-ink/gotil"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
	c      zap.Config
	once   sync.Once
}

var z = &Logger{}

func Zap() *Logger {
	z.once.Do(func() {
		z.initZap()
	})
	return z
}

func (l *Logger) Info(args ...interface{}) {
	if l.Sugar != nil {
		l.Sugar.Info(args...)
		return
	}
	infoLog.logOut(nil, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.Sugar != nil {
		l.Sugar.Infof(format, args...)
		return
	}
	infoLog.logOut(&format, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	if l.Sugar != nil {
		l.Sugar.Debug(args...)
		return
	}
	debugLog.logOut(nil, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.Sugar != nil {
		l.Sugar.Debugf(format, args...)
		return
	}
	debugLog.logOut(&format, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	if l.Sugar != nil {
		l.Sugar.Warn(args...)
		return
	}
	warnLog.logOut(nil, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.Sugar != nil {
		l.Sugar.Warnf(format, args...)
		return
	}
	warnLog.logOut(&format, args...)
}

func (l *Logger) Error(args ...interface{}) {
	if l.Sugar != nil {
		l.Sugar.Error(args...)
		return
	}
	errLog.logOut(nil, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.Sugar != nil {
		l.Sugar.Errorf(format, args...)
		return
	}
	errLog.logOut(&format, args...)
}

func (l *Logger) initZap() {
	var err error
	l.c = zap.NewProductionConfig()
	l.c.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	l.c.EncoderConfig.EncodeTime = timeEncoder
	l.Logger, err = l.c.Build()
	if err != nil {
		Errorf("l.initZap(),err:%+v", err)
		return
	}
	l.Sugar = l.Logger.Sugar()
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	encodeTimeLayout(t, gotil.TimeLayout_1, enc)
}

func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}
