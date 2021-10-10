package config

import (
	"idaman.id/storage/pkg/app"
)

var config Config

func InitConfig(provider string) error {
	configService, err := CreateConfig(provider)
	if err != nil {
		return err
	}

	config = configService

	err = loadConfiguration()
	if err != nil {
		return err
	}

	return nil
}

func CreateConfig(provider string) (Config, error) {
	if provider != CONFIG_VIPER {
		return nil, &app.NotSupportedError{
			Message: app.STATUS_NOT_SUPPORTED,
			Context: "Config",
		}
	}

	config := &ViperConfig{
		fileName: ".env",
	}
	return config, nil
}

func loadConfiguration() error {
	err := config.loadConfiguration()
	if err != nil {
		return err
	}

	config.SetDefault("MIN_UPLOADED_FILE", 1)
	config.SetDefault("MAX_UPLOADED_FILE", 5)
	return nil
}

func GetString(key string) string {
	return config.GetString(key)
}

func GetInt(key string) int {
	return config.GetInt(key)
}

func Get(key string) interface{} {
	return config.Get(key)
}

func Set(key string, value interface{}) {
	config.Set(key, value)
}

func SetDefault(key string, value interface{}) {
	config.SetDefault(key, value)
}
