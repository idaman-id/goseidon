package validation

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

func createTranslator(locale string) ut.Translator {
	en := en.New()
	id := id.New()
	uni := ut.New(en, id, en)
	translator, _ := uni.GetTranslator(locale)
	return translator
}

func registerDefaultTranslation(validate *validator.Validate, translator ut.Translator, locale string) error {
	var err error

	switch locale {
	case "id":
		err = id_translations.RegisterDefaultTranslations(validate, translator)
	case "en":
		err = en_translations.RegisterDefaultTranslations(validate, translator)
	default:
		err = en_translations.RegisterDefaultTranslations(validate, translator)
	}

	return err
}

/*
	@todo
	1. translate multilingual (indonesian)
	2. use .env variable
*/
func registerTranslation(validate *validator.Validate, translator ut.Translator) error {
	var err error

	err = validate.RegisterTranslation("valid_provider", translator, func(ut ut.Translator) error {
		return ut.Add("valid_provider", "{0} must be a valid provider", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("valid_provider", fe.Field())
		return t
	})
	if err != nil {
		return err
	}

	err = validate.RegisterTranslation("valid_file_type", translator, func(ut ut.Translator) error {
		return ut.Add("valid_file_type", "{0} must be a valid file type", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("valid_file_type", fe.Field())
		return t
	})
	if err != nil {
		return err
	}

	err = validate.RegisterTranslation("valid_file_amounts", translator, func(ut ut.Translator) error {
		return ut.Add("valid_file_amounts", "{0} must be greater than or equal 1 and less than or equal 5", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("valid_file_amounts", fe.Field())
		return t
	})
	if err != nil {
		return err
	}

	return nil
}
