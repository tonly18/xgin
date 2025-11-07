package xglobal

import "github.com/gin-gonic/gin"

const (
	UserId   = "user_id"
	ClientIp = "client_ip"
	TraceId  = "trace_id"
)

var (
	GinMode string = gin.DebugMode // gin的运行模式: debug、release、test'
	LogFile string                 // log路径
)
