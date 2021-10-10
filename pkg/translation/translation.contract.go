package translation

type TemplateData = map[string]interface{}
type Translator = func(param TranslatorDto) (translation string)
