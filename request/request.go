package request

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tonly18/xgin/xglobal"
)

// Request 请求
type Request struct {
	ctx *gin.Context
}

func NewRequest(c *gin.Context) *Request {
	return &Request{
		ctx: c,
	}
}

func (r *Request) GetCtx() *gin.Context {
	return r.ctx
}

func (r *Request) GetTraceID() string {
	return r.ctx.GetString(xglobal.TraceId)
}

func (r *Request) ClientIP() string {
	return r.ctx.ClientIP()
}

func (r *Request) Deadline() (deadline time.Time, ok bool) {
	return r.ctx.Deadline()
}

func (r *Request) Done() <-chan struct{} {
	return r.ctx.Done()
}

func (r *Request) Err() error {
	return r.ctx.Err()
}

func (r *Request) Value(key any) any {
	return r.ctx.Value(key)
}
