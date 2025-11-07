package middleware

import (
	"errors"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tonly18/xgin/logger"
	"github.com/tonly18/xgin/xglobal"
)

// RecoverPanic 中间件: recover处理panic
func RecoverPanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(c, fmt.Sprintf(`[middleware panic] Error(1): %v`, err))
				logger.Error(c, fmt.Sprintf(`[middleware panic] User Id: %v, Client IP: %v`, c.GetString("user_id"), c.GetString("client_ip")))
				for i := 1; i < 20; i++ {
					if pc, file, line, ok := runtime.Caller(i); ok {
						fname := runtime.FuncForPC(pc).Name() //获取函数名
						logger.Error(c, fmt.Sprintf(`[middleware panic] goroutine:%v, file:%s, function name:%s, line:%d`, pc, file, fname, line))
					}
				}
				logger.Error(c, fmt.Sprintf(`[middleware panic] Error(2): %v`, err))

				//error: broken pipe error
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seError := strings.ToLower(se.Error())
						if strings.Contains(seError, "broken pipe") {
							brokenPipe = true
						} else if strings.Contains(seError, "connection reset by peer") {
							brokenPipe = true
						}
						logger.Error(c, fmt.Sprintf(`[middleware panic] broken pipe is error se error: %v`, se))
					}
				}
				if brokenPipe {
					logger.Error(c, fmt.Sprintf(`[middleware panic] broken pipe is error: %v`, err))
					c.Abort()
				}
			}
		}()

		c.Set(xglobal.UserId, "222333")

		// before request
		c.Next()
		// after request
	}
}
