package handler

import (
	"github.com/tonly18/xgin/request"
	"github.com/tonly18/xgin/response"
	"github.com/tonly18/xgin/xerror"
)

type BaseHandle struct{}

func (c *BaseHandle) PreHandler(*request.Request) {}
func (c *BaseHandle) Handler(*request.Request) (*response.Response, *xerror.XError) {
	return &response.Response{}, nil
}
func (c *BaseHandle) PostHandler(*request.Request) {}
