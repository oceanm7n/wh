package config

import (
	"github.com/spf13/viper"
)

// Initializes the config
func Init() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
