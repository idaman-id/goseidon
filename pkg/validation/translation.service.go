package validation

import (
	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"idaman.id/storage/pkg/config"
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

	err = validate.RegisterTranslation("valid_file_amount", translator, func(ut ut.Translator) error {
		minLength := config.Service.GetInt("MIN_UPLOADED_FILE")
		maxLength := config.Service.GetInt("MAX_UPLOADED_FILE")
		message := fmt.Sprintf("{0} must be greater than or equal %d and less than or equal %d", minLength, maxLength)
		return ut.Add("valid_file_amount", message, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("valid_file_amount", fe.Field())
		return t
	})
	if err != nil {
		return err
	}

	err = validate.RegisterTranslation("valid_file_size", translator, func(ut ut.Translator) error {
		minSize := config.Service.GetInt("MIN_FILE_SIZE")
		maxSize := config.Service.GetInt("MAX_FILE_SIZE")
		message := fmt.Sprintf("{0} must be greater than or equal %d and less than or equal %d", minSize, maxSize)
		return ut.Add("valid_file_size", message, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("valid_file_size", fe.Field())
		return t
	})
	if err != nil {
		return err
	}

	return nil
}
