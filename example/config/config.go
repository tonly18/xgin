package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigStruck struct {
	Http *HttpConfig
}

type HttpConfig struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	Env  string `json:"env"`
	Mode string `json:"mode"`
}

// Config data
var Config = ConfigStruck{}

// Init 初始化配置信息
func Init() {
	// file xglobal
	if err := parseConfigFromToml(&Config); err != nil {
		panic(fmt.Sprintf("xglobal init from file error:%v", err))
	}

	fmt.Printf("config.Http: %+v\n", Config.Http)
}

// 获取配置文件并解析到指定的变量
func parseConfigFromToml(config any) error {
	viper.SetConfigFile("./conf/config_local.toml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	//parse
	return viper.Unmarshal(config)
}
