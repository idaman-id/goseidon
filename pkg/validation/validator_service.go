package validation

/*
@todo
instantiate validationService according to a given provider
*/
func NewValidator(provider string) ValidationService {

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
