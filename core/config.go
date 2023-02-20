package core

import (
	"log"

	"github.com/spf13/viper"
)

var ConfigViper *viper.Viper

func init() {
	ConfigViper = viper.New()
	ConfigViper.SetConfigName("config")
	ConfigViper.SetConfigType("yaml")
	ConfigViper.AddConfigPath("./config")

	err := ConfigViper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %w", err)
	}
}
