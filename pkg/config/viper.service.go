package config

import (
	"github.com/spf13/viper"
)

type ViperConfig struct {
	fileName string
}

func (config *ViperConfig) LoadConfiguration() error {
	viper.SetConfigFile(config.fileName)
	err := viper.ReadInConfig()
	return err
}

func (config *ViperConfig) GetString(key string) string {
	var value string
	keyValue := viper.Get(key)
	value, _ = keyValue.(string)
	return value
}

func (config *ViperConfig) GetInt(key string) int {
	var value int
	value = viper.GetInt(key)
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
