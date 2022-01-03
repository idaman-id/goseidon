package config_viper

import (
	"github.com/spf13/viper"
)

type viperConfig struct {
	FileName string
}

func (s *viperConfig) GetString(key string) string {
	value := viper.GetString(key)
	return value
}

func (s *viperConfig) GetInt(key string) int {
	value := viper.GetInt(key)
	return value
}

func (s *viperConfig) Get(key string) interface{} {
	value := viper.Get(key)
	return value
}

func (s *viperConfig) Set(key string, value interface{}) {
	viper.Set(key, value)
}

func (s *viperConfig) SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}

func (s *viperConfig) loadConfig() error {
	viper.SetConfigFile(s.FileName)
	err := viper.ReadInConfig()
	return err
}

func NewViperConfig(fn string) (*viperConfig, error) {
	s := &viperConfig{
		FileName: fn,
	}

	err := s.loadConfig()
	if err != nil {
		return nil, err
	}

	s.SetDefault("APP_HOST", "localhost")
	s.SetDefault("APP_PORT", 3000)
	s.SetDefault("APP_DEFAULT_LOCALE", "en")
	s.SetDefault("MIN_UPLOADED_FILE", 1)
	s.SetDefault("MAX_UPLOADED_FILE", 5)
	s.SetDefault("MIN_FILE_SIZE", 1)
	s.SetDefault("MAX_FILE_SIZE", 134217728)

	return s, nil
}
