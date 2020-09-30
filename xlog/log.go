package xlog

var (
	debugLog xLogger = &DebugLogger{}
	infoLog  xLogger = &InfoLogger{}
	warnLog  xLogger = &WarnLogger{}
	errLog   xLogger = &ErrorLogger{}
)

var (
	Reset = Color([]byte{27, 91, 48, 109})
	// 标准
	White   = Color([]byte{27, 91, 51, 48, 109}) // 白色
	Red     = Color([]byte{27, 91, 51, 49, 109}) // 红色
	Green   = Color([]byte{27, 91, 51, 50, 109}) // 绿色
	Yellow  = Color([]byte{27, 91, 51, 51, 109}) // 黄色
	Blue    = Color([]byte{27, 91, 51, 52, 109}) // 蓝色
	Magenta = Color([]byte{27, 91, 51, 53, 109}) // 紫色
	Cyan    = Color([]byte{27, 91, 51, 54, 109}) // 青色
	// 高亮
	WhiteBright   = Color([]byte{27, 91, 49, 59, 51, 48, 109})
	RedBright     = Color([]byte{27, 91, 49, 59, 51, 49, 109})
	GreenBright   = Color([]byte{27, 91, 49, 59, 51, 50, 109})
	YellowBright  = Color([]byte{27, 91, 49, 59, 51, 51, 109})
	BlueBright    = Color([]byte{27, 91, 49, 59, 51, 52, 109})
	MagentaBright = Color([]byte{27, 91, 49, 59, 51, 53, 109})
	CyanBright    = Color([]byte{27, 91, 49, 59, 51, 54, 109})
	// 英文斜体
	WhiteBevel   = Color([]byte{27, 91, 51, 59, 51, 48, 109})
	RedBevel     = Color([]byte{27, 91, 51, 59, 51, 49, 109})
	GreenBevel   = Color([]byte{27, 91, 51, 59, 51, 50, 109})
	YellowBevel  = Color([]byte{27, 91, 51, 59, 51, 51, 109})
	BlueBevel    = Color([]byte{27, 91, 51, 59, 51, 52, 109})
	MagentaBevel = Color([]byte{27, 91, 51, 59, 51, 53, 109})
	CyanBevel    = Color([]byte{27, 91, 51, 59, 51, 54, 109})
	// 下划线
	WhiteUnderLine   = Color([]byte{27, 91, 52, 59, 51, 48, 109})
	RedUnderLine     = Color([]byte{27, 91, 52, 59, 51, 49, 109})
	GreenUnderLine   = Color([]byte{27, 91, 52, 59, 51, 50, 109})
	YellowUnderLine  = Color([]byte{27, 91, 52, 59, 51, 51, 109})
	BlueUnderLine    = Color([]byte{27, 91, 52, 59, 51, 52, 109})
	MagentaUnderLine = Color([]byte{27, 91, 52, 59, 51, 53, 109})
	CyanUnderLine    = Color([]byte{27, 91, 52, 59, 51, 54, 109})
	// 背景色
	WhiteBg   = Color([]byte{27, 91, 55, 59, 51, 48, 109})
	RedBg     = Color([]byte{27, 91, 55, 59, 51, 49, 109})
	GreenBg   = Color([]byte{27, 91, 55, 59, 51, 50, 109})
	YellowBg  = Color([]byte{27, 91, 55, 59, 51, 51, 109})
	BlueBg    = Color([]byte{27, 91, 55, 59, 51, 52, 109})
	MagentaBg = Color([]byte{27, 91, 55, 59, 51, 53, 109})
	CyanBg    = Color([]byte{27, 91, 55, 59, 51, 54, 109})
	// 删除线
	WhiteDelLine   = Color([]byte{27, 91, 57, 59, 51, 48, 109})
	RedDelLine     = Color([]byte{27, 91, 57, 59, 51, 49, 109})
	GreenDelLine   = Color([]byte{27, 91, 57, 59, 51, 50, 109})
	YellowDelLine  = Color([]byte{27, 91, 57, 59, 51, 51, 109})
	BlueDelLine    = Color([]byte{27, 91, 57, 59, 51, 52, 109})
	MagentaDelLine = Color([]byte{27, 91, 57, 59, 51, 53, 109})
	CyanDelLine    = Color([]byte{27, 91, 57, 59, 51, 54, 109})
)

type Color string

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
