package translation

type TemplateData = map[string]interface{}
type Translator = func(param TranslatorDto) string

type TranslationService interface {
	Translate(param TranslatorDto) string
}
