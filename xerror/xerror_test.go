package xerror

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestTXError(t *testing.T) {
	err := HandleUserRequest(404)

	res := errors.Is(err, io.EOF)
	fmt.Println("---res:", res)
	fmt.Println("---res:", errors.Is(err, io.ErrClosedPipe))

	fmt.Println("捕获到错误：")
	fmt.Println(FormatStack(err))

	errs := strings.Split(FormatStack(err), "\n")
	for _, v := range errs {
		fmt.Println(v)
	}

	Range(err, func(er error) {
		fmt.Println("------er:", er)
	})
}

func HandleUserRequest(id int) *XError {
	err := GetUserInfo(id)
	if err != nil {
		return Wrap(err, "333")
	}
	return nil
}

func GetUserInfo(id int) *XError {
	err := GetUserByID(id)
	if err != nil {
		return Wrap(err, "222")
	}
	return nil
}

func GetUserByID(id int) *XError {
	//return NewXError("111")
	return Wrap(io.EOF, "111")
}
