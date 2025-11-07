package config

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestConfigFromFile(t *testing.T) {
	configFile := "/Users/tonly/Documents/item/services/payment-service/conf/config_dev.toml"

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("FileName: %s, Error: %v", configFile, err))
	}

	var config *ConfigStruck = &ConfigStruck{}
	if err := viper.Unmarshal(config); err != nil {
		panic(fmt.Errorf("FileName: %s, Error: %v", configFile, err))
	}

	fmt.Printf("xglobal.Http: %+v\n", config.Http)
	fmt.Printf("xglobal.MySql: %+v\n", config.MySql)
	fmt.Printf("xglobal.Redis: %+v\n", config.Redis)
	fmt.Printf("xglobal.Kafka: %+v\n", config.Kafka)
	fmt.Printf("xglobal.Log: %+v\n", config.Log)
	fmt.Printf("xglobal.Test: %+v\n", config.Test)
}
