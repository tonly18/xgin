package logger

import (
	"github.com/rs/zerolog"
)

// ZeroLogHook hook
type ZeroLogHook struct{}

func (h *ZeroLogHook) Run(logger *zerolog.Event, level zerolog.Level, msg string) {}
