package config

import (
	config_viper "idaman.id/storage/internal/config-viper"
)

func NewConfigService() (ConfigService, error) {
	return config_viper.NewViperConfig(".env")
}
