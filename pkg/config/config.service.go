package config

import "idaman.id/storage/pkg/app"

var (
	Service ConfigService
)

func Init() error {

	isServiceAvailable := Service != nil
	if !isServiceAvailable {
		return &app.NotFoundError{
			Message: app.STATUS_NOT_FOUND,
			Context: "Config",
		}
	}

	err := Service.LoadConfiguration()
	isFailedLoadConfig := err != nil
	if isFailedLoadConfig {
		return err
	}

	setDefaultData(Service)

	return nil
}

func CreateConfig(provider string) (ConfigService, error) {
	isProviderSupported := provider == CONFIG_VIPER
	if !isProviderSupported {
		err := &app.NotSupportedError{
			Message: app.STATUS_NOT_SUPPORTED,
			Context: "Config",
		}
		return nil, err
	}

	config := &ViperConfig{
		FileName: ".env",
	}

	return config, nil
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
