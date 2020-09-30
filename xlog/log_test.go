package xlog

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	// default log
	Info(White, "白色 info", Reset, WhiteBright, "高亮 info", Reset, "恢复默认颜色", WhiteDelLine, "删除线", Reset, WhiteUnderLine, "下划线", Reset, WhiteBevel, "斜体 info", Reset, WhiteBg, "背景", Reset)
	Debug(Blue, "蓝色 debug", Reset, BlueBright, "高亮 debug", Reset, "恢复默认颜色", BlueDelLine, "删除线", Reset, BlueUnderLine, "下划线", Reset, BlueBevel, "斜体 debug", Reset, BlueBg, "背景", Reset)
	Warn(Yellow, "黄色 warning", Reset, YellowBright, "高亮 warning", Reset, "恢复默认颜色", YellowDelLine, "删除线", Reset, YellowUnderLine, "下划线", Reset, YellowBevel, "斜体 warning", Reset, YellowBg, "背景", Reset)
	Error(Red, "红色 error", Reset, RedBright, "高亮 error", Reset, "恢复默认颜色", RedDelLine, "删除线", Reset, RedUnderLine, "下划线", Reset, RedBevel, "斜体 error", Reset, RedBg, "背景", Reset)

	fmt.Println()

	// zap log
	Zap().Info("zap info")
	Zap().Debug("zap debug")
	Zap().Warn("zap warn")
	Zap().Error("zap error")
}
