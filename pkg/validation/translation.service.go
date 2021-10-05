package validation

import (
	"reflect"
	"strings"

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

func registerTranslation(validate *validator.Validate, translator ut.Translator, locale string) error {
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

func registerTag(validate *validator.Validate) {

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
