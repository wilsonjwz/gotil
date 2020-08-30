package limit

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/iGoogle-ink/gotil/web"
)

func TestInitServer(t *testing.T) {
	// 解开注释测试

	//c := &web.Config{
	//	Port:  ":2233",
	//	Limit: nil,
	//}
	//g := web.InitServer(c)
	//initRoute(g.Gin)
	//g.Start()
	//
	//ch := make(chan os.Signal)
	//signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	//for {
	//	si := <-ch
	//	switch si {
	//	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
	//		xlog.Warnf("get a signal %s, stop the process", si.String())
	//		// todo something
	//		return
	//	case syscall.SIGHUP:
	//	default:
	//		return
	//	}
	//}
}

func initRoute(g *gin.Engine) {
	g.GET("/a", func(c *gin.Context) {
		web.JSON(c, "a", nil)
	})
	g.GET("/b", func(c *gin.Context) {
		web.JSON(c, "b", nil)
	})
	g.GET("/c", func(c *gin.Context) {
		web.JSON(c, "c", nil)
	})
	g.GET("/d", func(c *gin.Context) {
		web.JSON(c, "d", nil)
	})
}
