package viper

import (
	"github.com/spf13/viper"
)

// ConfigInterface interface for Viper
type ConfigInterface interface {
	LoadConfig() error
	GetStringValue(key string) string
	GetIntValue(key string) int
	GetPort() int
}

type Config struct{}

// LoadConfig load configuration file
func (c *Config) LoadConfig() error {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.AutomaticEnv()

	return nil
}

func (c *Config) GetStringValue(key string) string {
	value := viper.GetString(key)
	return value
}

func (c *Config) GetIntValue(key string) int {
	value := viper.GetInt(key)
	return value
}

var ViperConfig = &Config{}
