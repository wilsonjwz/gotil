package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iGoogle-ink/gotil/ecode"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORS gin middleware cors
func (g *GinEngine) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") // 请求头部
		if origin == "" {
			origin = c.Request.Host
		}
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			// 允许跨域返回的Header
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Content-Type, Token, Session")
			// 允许的方法
			c.Header("Access-Control-Allow-Methods", "POST, PUT ,GET, OPTIONS, DELETE, HEAD, TRACE, UPDATE")
			// 运行客户端解析的Header
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			// 缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			// 允许客户端传递校验信息，cookie
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// Recovery gin middleware recovery
func (g *GinEngine) Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				r, ok := err.(error)
				if !ok {
					r = fmt.Errorf("%v", err)
				}
				stack := make([]byte, 4<<10)
				length := runtime.Stack(stack, false)
				stackStr := string(stack[:length])

				param := &RecoverInfo{
					Time:  time.Now().Format("2006-01-02 15:04:05"),
					Url:   c.Request.URL.Path,
					Err:   r.Error(),
					Query: c.Request.URL.Query(),
					Stack: stackStr,
				}
				data, _ := json.Marshal(param)
				log.Println("[GinPanic] >> ", string(data))

				c.Error(r) // nolint: errcheck
				c.Abort()
				JSON(c, nil, ecode.ServerErr)
			}
		}()
		c.Next()
	}
}

// Recover echo middleware recover
func (e *EchoEngine) Recover() echo.MiddlewareFunc {
	return e.recoverWithConfig(middleware.DefaultRecoverConfig)
}

func (e *EchoEngine) recoverWithConfig(config middleware.RecoverConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultRecoverConfig.Skipper
	}
	if config.StackSize == 0 {
		config.StackSize = middleware.DefaultRecoverConfig.StackSize

	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}
			defer func() error {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					stack := make([]byte, config.StackSize)
					stackStr := ""
					length := runtime.Stack(stack, false)
					if !config.DisablePrintStack {
						stackStr = string(stack[:length])
					}

					//c.Response().Header().Set("Content-Type", "application/json;charset=UTF-8")
					//c.Response().Write([]byte(returnMsg))

					param := &RecoverInfo{
						Time:  time.Now().Format("2006-01-02 15:04:05"),
						Url:   c.Request().URL.Path,
						Err:   err.Error(),
						Query: c.QueryParams(),
						Stack: stackStr,
					}
					data, _ := json.Marshal(param)
					log.Println("[EchoPanic] >> ", string(data))

					return c.JSON(http.StatusOK, CommonRsp{
						Code:    11503,
						Message: "服务异常,请稍后再试。",
						Data:    nil,
					})

				}
				return nil
			}()
			return next(c)
		}
	}
}

// Logger echo middleware logger
func (e *EchoEngine) Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper:          middleware.DefaultSkipper,
		Format:           "[ECHO] ${time_custom} | ${status} | ${remote_ip} | ${method}      ${uri}\n",
		CustomTimeFormat: "2006/01/02 - 15:04:05",
	})
}
