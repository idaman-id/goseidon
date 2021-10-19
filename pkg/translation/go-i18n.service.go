package translation

import "github.com/nicksnyder/go-i18n/v2/i18n"

type GoI18nService struct {
	localizer Localizer
}

func (s *GoI18nService) Translate(param TranslatorDto) string {
	if s.localizer == nil {
		return param.MessageId
	}

	translation, err := s.localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    param.MessageId,
		TemplateData: param.Template,
	})
	if err != nil {
		return param.MessageId
	}
	return translation
}

func NewGoI18nService(localizer Localizer) *GoI18nService {
	return &GoI18nService{
		localizer: localizer,
	}
}
