package config

import (
	"github.com/spf13/viper"
)

type ViperConfig struct {
	FileName string
}

func (s *ViperConfig) LoadConfiguration() error {
	viper.SetConfigFile(s.FileName)
	err := viper.ReadInConfig()
	return err
}

func (s *ViperConfig) GetString(key string) string {
	value := viper.GetString(key)
	return value
}

func (s *ViperConfig) GetInt(key string) int {
	value := viper.GetInt(key)
	return value
}

func (s *ViperConfig) Get(key string) interface{} {
	value := viper.Get(key)
	return value
}

func (s *ViperConfig) Set(key string, value interface{}) {
	viper.Set(key, value)
}

func (s *ViperConfig) SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}
