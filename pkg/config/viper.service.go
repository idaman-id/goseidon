package config

import (
	"github.com/spf13/viper"
)

type ViperConfig struct {
	FileName string
}

func (config *ViperConfig) LoadConfiguration() error {
	viper.SetConfigFile(config.FileName)
	err := viper.ReadInConfig()
	return err
}

func (config *ViperConfig) GetString(key string) string {
	value := viper.GetString(key)
	return value
}

func (config *ViperConfig) GetInt(key string) int {
	value := viper.GetInt(key)
	return value
}

func (config *ViperConfig) Get(key string) interface{} {
	value := viper.Get(key)
	return value
}

func (config *ViperConfig) Set(key string, value interface{}) {
	viper.Set(key, value)
}

func (config *ViperConfig) SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}
