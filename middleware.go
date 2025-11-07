package xgin

import (
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tonly18/xgin/logger"
	"github.com/tonly18/xgin/xglobal"
)

func defaultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//start time
		startTime := time.Now()

		//IP
		c.Set(xglobal.ClientIp, c.ClientIP())

		//trace id
		traceId := c.Request.Header.Get("trace_id")
		if traceId == "" {
			traceId = uuid.New().String()
		}
		c.Set(xglobal.TraceId, traceId)

		//打印请求参数
		printParams(c)

		c.Next()

		// 日志格式
		logger.Infof(c, `[URI:%s | Method:%s | Status Code:%d | Execution Time(ms):%d]`, c.Request.RequestURI, c.Request.Method, c.Writer.Status(), time.Since(startTime).Milliseconds())
	}
}

// 打印请求参数
func printParams(c *gin.Context) {
	var bodyData string
	query := c.Request.URL.Query()
	bodyBytes, _ := c.GetRawData()
	if len(bodyBytes) > 0 {
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		bodyData = strings.ReplaceAll(string(bodyBytes), `"`, "")
		bodyData = strings.ReplaceAll(bodyData, "\n", "")
		bodyData = strings.ReplaceAll(bodyData, " ", "")
	}

	logger.Infof(c, `[URI:%s | Method:%s | Query:%v | Body:%v]`, c.Request.RequestURI, c.Request.Method, query.Encode(), bodyData)
}
