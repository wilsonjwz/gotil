package web

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/iGoogle-ink/gotil/limit"
	"github.com/labstack/echo/v4"
)

func TestInitEcho(t *testing.T) {
	c := &Config{
		Port: ":2234",
		Limit: &limit.Config{
			Rate:       0, // 0 速率不限流
			BucketSize: 100,
		},
	}
	e := InitEcho(c)
	e.Release()
	e.Echo.Use(e.Logger())
	e.Echo.Use(e.Recover())

	initRouteE(e.Echo)

	e.Start()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			// todo something

			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func initRouteE(e *echo.Echo) {
	e.GET("/echo/ping", func(c echo.Context) error {
		//err := ecode.New(2323, "asdsda")
		JSON(c, "echo data", nil)
		return nil
	})

	e.GET("/echo/file", func(c echo.Context) error {
		//err := ecode.New(2323, "asdsda")
		File(c, "echo_test.go", "echo.go")
		return nil
	})
}
