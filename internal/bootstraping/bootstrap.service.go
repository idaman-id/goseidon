package bootstraping

import (
	"idaman.id/storage/internal/repository"
	"idaman.id/storage/pkg/config"
)

/*
	@todo register i18n
*/
func Setup() error {
	configService, err := config.NewConfig(config.CONFIG_VIPER)
	isFailedCreateConfig := err != nil
	if isFailedCreateConfig {
		return err
	}

	config.Service = configService
	err = config.Init()
	isFailedInitConfig := err != nil
	if isFailedInitConfig {
		return err
	}

	err = repository.Init(repository.DATABASE_MONGO)
	isFailedInitDatabase := err != nil
	if isFailedInitDatabase {
		return err
	}

	// i18nBundle := i18n.NewBundle(language.English)
	// i18nBundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	// i18nBundle.MustLoadMessageFile("pkg/translation/status.en.json")
	// i18nBundle.MustLoadMessageFile("pkg/translation/status.id.json")

	// localizer := createLocalizer(i18nBundle)
	// dependency := &handler.Dependency{
	// 	getLocalizer: localizer,
	// 	getLocale:    localeParser,
	// }

	return nil
}
