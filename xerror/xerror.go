package xerror

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type Error interface {
	Error() string
}

// XError 是一个可携带上下文信息的错误类型。
type XError struct {
	Msg   string
	Cause error
	File  string
	Line  int
	Func  string
}

// NewXError 创建新错误（带文件、行号、函数信息）
func NewXError(msg string) *XError {
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	return &XError{
		Msg:  msg,
		File: filepath.Base(file),
		Line: line,
		Func: shortFunc(fn.Name()),
	}
}

// XError 实现 error 接口
func (e *XError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s:%d [%s] %s | %v", e.File, e.Line, e.Func, e.Msg, e.Cause)
	}

	return fmt.Sprintf("%s:%d [%s] %s", e.File, e.Line, e.Func, e.Msg)
}

// Unwrap 支持 errors.Unwrap / errors.Is / errors.As
func (e *XError) Unwrap() error {
	return e.Cause
}

// Wrap 在原有错误上包裹一层上下文信息
func Wrap(err error, msg string) *XError {
	if err == nil {
		return nil
	}

	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	return &XError{
		Msg:   msg,
		Cause: err,
		File:  filepath.Base(file),
		Line:  line,
		Func:  shortFunc(fn.Name()),
	}
}

// shortFunc 去除冗余包路径
func shortFunc(name string) string {
	if i := strings.LastIndex(name, "/"); i != -1 {
		return name[i+1:]
	}

	return name
}

// FormatStack 递归输出错误堆栈（可用于日志打印）
func FormatStack(err error) string {
	if err == nil {
		return ""
	}

	var XError *XError
	var b strings.Builder
	for e := err; e != nil; e = errors.Unwrap(e) {
		switch {
		case errors.As(e, &XError):
			b.WriteString(e.Error())
			b.WriteString("\n")
		}
	}

	return b.String()
}

// FirstXError 获取最外层错误
func FirstXError(err error) *XError {
	var xerr *XError
	if errors.As(err, &xerr) {
		return xerr
	}
	return nil
}

// Range 循环处理
func Range(err error, handle func(er error)) {
	for e := err; e != nil; e = errors.Unwrap(e) {
		handle(e)
	}
}
