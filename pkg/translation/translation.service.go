package translation

import "github.com/nicksnyder/go-i18n/v2/i18n"

func translateById(localizer *i18n.Localizer, id string, template TemplateData) string {
	translation, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: template,
	})
	if err != nil {
		return id
	}
	return translation
}

func CreateSimpleTranslator(localizer *i18n.Localizer) Translator {
	return func(param TranslatorDto) string {
		return translateById(localizer, param.Id, param.Template)
	}
}
