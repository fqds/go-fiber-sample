package config

import (
	"github.com/spf13/viper"
)

var Config = viper.New()

func SetConfig() {
	Config.SetConfigFile("config.yaml")
	Config.ReadInConfig()
}