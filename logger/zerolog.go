package logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/tonly18/xgin/xglobal"

	"github.com/rs/zerolog"
)

// logger
var logger zerolog.Logger

func Init(file ...string) {
	output := os.Stdout
	if len(file) > 0 {
		fs, err := os.OpenFile(file[0], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Sprintf("open log file is error: %v", err))
		}
		runtime.SetFinalizer(fs, func(f *os.File) {
			f.Close()
		})
		output = fs
	}

	// 初始化 console logger
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		_, file, line, _ = runtime.Caller(6)
		return fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	logger = zerolog.New(output).With().Caller().Timestamp().Logger().Hook(&ZeroLogHook{})
}

func Debug(ctx context.Context, args ...any) {
	logger.Debug().Fields(map[string]any{
		"ip":       ctx.Value(xglobal.ClientIp),
		"trace_id": ctx.Value(xglobal.TraceId),
	}).Msg(fmt.Sprint(args...))
}
func Debugf(ctx context.Context, format string, args ...any) {
	logger.Debug().Fields(map[string]any{
		"ip":       ctx.Value(xglobal.ClientIp),
		"trace_id": ctx.Value(xglobal.TraceId),
	}).Msgf(format, args...)
}

func Info(ctx context.Context, args ...any) {
	logger.Info().Fields(map[string]any{
		"ip":       ctx.Value(xglobal.ClientIp),
		"trace_id": ctx.Value(xglobal.TraceId),
	}).Msg(fmt.Sprint(args...))
}
func Infof(ctx context.Context, format string, args ...any) {
	logger.Info().Fields(map[string]any{
		"ip":       ctx.Value(xglobal.ClientIp),
		"trace_id": ctx.Value(xglobal.TraceId),
	}).Msgf(format, args...)
}

func Warning(ctx context.Context, args ...any) {
	logger.Warn().Fields(map[string]any{
		"ip":       ctx.Value(xglobal.ClientIp),
		"trace_id": ctx.Value(xglobal.TraceId),
	}).Msg(fmt.Sprint(args...))
}
func Warningf(ctx context.Context, format string, args ...any) {
	logger.Warn().Fields(map[string]any{
		"ip":       ctx.Value(xglobal.ClientIp),
		"trace_id": ctx.Value(xglobal.TraceId),
	}).Msgf(format, args...)
}

func Error(ctx context.Context, args ...any) {
	logger.Error().Fields(map[string]any{
		"ip":       ctx.Value(xglobal.ClientIp),
		"trace_id": ctx.Value(xglobal.TraceId),
	}).Msg(fmt.Sprint(args...))
}
func Errorf(ctx context.Context, format string, args ...any) {
	logger.Error().Fields(map[string]any{
		"ip":       ctx.Value(xglobal.ClientIp),
		"trace_id": ctx.Value(xglobal.TraceId),
	}).Msgf(format, args...)
}
