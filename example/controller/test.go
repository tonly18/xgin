package controller

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
	"github.com/tonly18/xgin/example/service"
	"github.com/tonly18/xgin/handler"
	"github.com/tonly18/xgin/request"
	"github.com/tonly18/xgin/response"
	"github.com/tonly18/xgin/xerror"
)

type TestHandler struct {
	handler.BaseHandle
}

func (c *TestHandler) Handler(req *request.Request) (*response.Response, xerror.Error) {
	fmt.Println("TestHandler start...", time.Now().Format(time.DateTime))

	x := req.GetCtx().Query("x")
	ix := cast.ToInt(x)

	testService := service.NewTestService(req)
	data, xerr := testService.GetData(ix)
	if xerr != nil {
		return nil, xerror.Wrap(xerr, "test-handler-error")
	}

	return &response.Response{
		Data: data,
	}, nil
}

type TestNotHandler struct {
	handler.BaseHandle
}

func (c *TestNotHandler) Handler(req *request.Request) (*response.Response, xerror.Error) {
	fmt.Println("TestHandler start...", time.Now().Format(time.DateTime))

	testService := service.NewTestService(req)
	data, xerr := testService.GetData(0)
	if xerr != nil {
		return nil, xerror.Wrap(xerr, "test-handler-error")
	}

	return &response.Response{
		Data: data,
	}, nil
}

type ABCHandler struct {
	handler.BaseHandle
}

func (c *ABCHandler) Handler(req *request.Request) (*response.Response, xerror.Error) {
	fmt.Println("TestHandler start...", time.Now().Format(time.DateTime))

	return &response.Response{
		Data: "ABCHandler",
	}, nil
}
