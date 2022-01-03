package config

import (
	config_viper "idaman.id/storage/internal/config-viper"
	app_error "idaman.id/storage/internal/error"
)

func NewConfig(provider string) (ConfigService, error) {
	isProviderSupported := provider == CONFIG_VIPER
	if !isProviderSupported {
		err := app_error.NewNotSupportedError("Config")
		return nil, err
	}

	config := &config_viper.ViperConfig{
		FileName: ".env",
	}

	return config, nil
}

func InitConfig(s ConfigService) error {

	err := s.LoadConfiguration()
	isFailedLoadConfig := err != nil
	if isFailedLoadConfig {
		return err
	}

	setDefaultData(s)

	return nil
}

func setDefaultData(config ConfigService) {
	config.SetDefault("APP_HOST", "localhost")
	config.SetDefault("APP_PORT", 3000)
	config.SetDefault("APP_DEFAULT_LOCALE", "en")
	config.SetDefault("MIN_UPLOADED_FILE", 1)
	config.SetDefault("MAX_UPLOADED_FILE", 5)
	config.SetDefault("MIN_FILE_SIZE", 1)
	config.SetDefault("MAX_FILE_SIZE", 134217728)
}
