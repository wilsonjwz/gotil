package xlog

var (
	debugLog xLogger = &DebugLogger{}
	infoLog  xLogger = &InfoLogger{}
	warnLog  xLogger = &WarnLogger{}
	errLog   xLogger = &ErrorLogger{}
)

var (
	Reset = string([]byte{27, 91, 48, 109})
	// 标准
	White   = string([]byte{27, 91, 51, 48, 109}) // 白色
	Red     = string([]byte{27, 91, 51, 49, 109}) // 红色
	Green   = string([]byte{27, 91, 51, 50, 109}) // 绿色
	Yellow  = string([]byte{27, 91, 51, 51, 109}) // 黄色
	Blue    = string([]byte{27, 91, 51, 52, 109}) // 蓝色
	Magenta = string([]byte{27, 91, 51, 53, 109}) // 紫色
	Cyan    = string([]byte{27, 91, 51, 54, 109}) // 青色
	// 高亮
	WhiteBright   = string([]byte{27, 91, 49, 59, 51, 48, 109})
	RedBright     = string([]byte{27, 91, 49, 59, 51, 49, 109})
	GreenBright   = string([]byte{27, 91, 49, 59, 51, 50, 109})
	YellowBright  = string([]byte{27, 91, 49, 59, 51, 51, 109})
	BlueBright    = string([]byte{27, 91, 49, 59, 51, 52, 109})
	MagentaBright = string([]byte{27, 91, 49, 59, 51, 53, 109})
	CyanBright    = string([]byte{27, 91, 49, 59, 51, 54, 109})
	// 英文斜体
	WhiteBevel   = string([]byte{27, 91, 51, 59, 51, 48, 109})
	RedBevel     = string([]byte{27, 91, 51, 59, 51, 49, 109})
	GreenBevel   = string([]byte{27, 91, 51, 59, 51, 50, 109})
	YellowBevel  = string([]byte{27, 91, 51, 59, 51, 51, 109})
	BlueBevel    = string([]byte{27, 91, 51, 59, 51, 52, 109})
	MagentaBevel = string([]byte{27, 91, 51, 59, 51, 53, 109})
	CyanBevel    = string([]byte{27, 91, 51, 59, 51, 54, 109})
	// 下划线
	WhiteUnderLine   = string([]byte{27, 91, 52, 59, 51, 48, 109})
	RedUnderLine     = string([]byte{27, 91, 52, 59, 51, 49, 109})
	GreenUnderLine   = string([]byte{27, 91, 52, 59, 51, 50, 109})
	YellowUnderLine  = string([]byte{27, 91, 52, 59, 51, 51, 109})
	BlueUnderLine    = string([]byte{27, 91, 52, 59, 51, 52, 109})
	MagentaUnderLine = string([]byte{27, 91, 52, 59, 51, 53, 109})
	CyanUnderLine    = string([]byte{27, 91, 52, 59, 51, 54, 109})
	// 背景色
	WhiteBg   = string([]byte{27, 91, 55, 59, 51, 48, 109})
	RedBg     = string([]byte{27, 91, 55, 59, 51, 49, 109})
	GreenBg   = string([]byte{27, 91, 55, 59, 51, 50, 109})
	YellowBg  = string([]byte{27, 91, 55, 59, 51, 51, 109})
	BlueBg    = string([]byte{27, 91, 55, 59, 51, 52, 109})
	MagentaBg = string([]byte{27, 91, 55, 59, 51, 53, 109})
	CyanBg    = string([]byte{27, 91, 55, 59, 51, 54, 109})
	// 删除线
	WhiteDelLine   = string([]byte{27, 91, 57, 59, 51, 48, 109})
	RedDelLine     = string([]byte{27, 91, 57, 59, 51, 49, 109})
	GreenDelLine   = string([]byte{27, 91, 57, 59, 51, 50, 109})
	YellowDelLine  = string([]byte{27, 91, 57, 59, 51, 51, 109})
	BlueDelLine    = string([]byte{27, 91, 57, 59, 51, 52, 109})
	MagentaDelLine = string([]byte{27, 91, 57, 59, 51, 53, 109})
	CyanDelLine    = string([]byte{27, 91, 57, 59, 51, 54, 109})
)

type xLogger interface {
	logOut(format *string, args ...interface{})
}

func Info(args ...interface{}) {
	infoLog.logOut(nil, args...)
}

func Infof(format string, args ...interface{}) {
	infoLog.logOut(&format, args...)
}

func Debug(args ...interface{}) {
	debugLog.logOut(nil, args...)
}

func Debugf(format string, args ...interface{}) {
	debugLog.logOut(&format, args...)
}

func Warn(args ...interface{}) {
	warnLog.logOut(nil, args...)
}

func Warnf(format string, args ...interface{}) {
	warnLog.logOut(&format, args...)
}

func Error(args ...interface{}) {
	errLog.logOut(nil, args...)
}

func Errorf(format string, args ...interface{}) {
	errLog.logOut(&format, args...)
}
