package iface

import (
	"github.com/tonly18/xgin/request"
	"github.com/tonly18/xgin/response"
	"github.com/tonly18/xgin/xerror"
)

type IHandler interface {
	PreHandler(*request.Request)
	Handler(*request.Request) (*response.Response, xerror.Error)
	PostHandler(*request.Request)
}
