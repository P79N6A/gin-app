package config

import (
	"github.com/spf13/viper"
	"fmt"
)

var logConfig, appConfig *viper.Viper

func init() {
	logConfig = viper.New()
	logConfig.SetConfigType("yaml")
	logConfig.AddConfigPath("/Users/zhangbingbing/go/src/gin-app/config")
	logConfig.AddConfigPath(".")
	logConfig.SetConfigName("log")
	err := logConfig.ReadInConfig()
	if err != nil {
		fmt.Println("load log config error", err)
	}

	appConfig = viper.New()
	appConfig.SetConfigType("yaml")
	appConfig.AddConfigPath("/Users/zhangbingbing/go/src/gin-app/config")
	appConfig.AddConfigPath(".")
	appConfig.SetConfigName("app")
	err = appConfig.ReadInConfig()
	if err != nil {
		fmt.Println("load app config error", err)
	}
}

func GetLogConfig() map[string]interface{} {
	return logConfig.AllSettings()
}

func GetAppConfig() map[string]interface{} {
	return appConfig.AllSettings()
}
