package validation

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type validationError struct {
	msg string
}

func (e *validationError) Error() string {
	return e.msg
}

type validationService struct {
	v *validator.Validate
	t ut.Translator
}

func (s *validationService) Validate(i interface{}) error {
	err := s.v.Struct(i)
	if err == nil {
		return nil
	}

	msg := ""
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		msg += e.Translate(s.t)
	}

	return &validationError{msg}
}

func NewValidator() (Validator, error) {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	v := validator.New()
	err := en_translations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		return nil, err
	}
	s := &validationService{
		v: v,
		t: trans,
	}
	return s, nil
}
