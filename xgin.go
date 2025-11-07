package xgin

import (
	"github.com/tonly18/xgin/logger"
	"github.com/tonly18/xgin/xglobal"
)

func SetConfig(conf *Config) {
	if conf.GinMode != "" {
		xglobal.GinMode = conf.GinMode
	}
	if conf.ConfigFile != "" {
		xglobal.LogFile = conf.ConfigFile
	}

	//init zero log
	logger.Init(xglobal.LogFile)
}
