package translation

import "github.com/nicksnyder/go-i18n/v2/i18n"

func translateById(id string, localizer *i18n.Localizer) (translation string) {
	translation, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: id,
	})
	if err != nil {
		return id
	}
	return
}

func CreateSimpleTranslator(localizer *i18n.Localizer) Translator {
	return func(id string) (translation string) {
		return translateById(id, localizer)
	}
}
