package translation

type TemplateData = map[string]interface{}
type Translator = func(param TranslatorParam) string

type TranslationService interface {
	Translate(param TranslatorParam) string
}

type TranslatorParam struct {
	MessageId string
	Template  TemplateData
}
