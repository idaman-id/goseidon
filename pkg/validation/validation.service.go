package validation

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"idaman.id/storage/pkg/app"
)

func ValidateRule(locale string, data interface{}) *app.ValidationError {
	translator := createTranslator(locale)
	validate, validatorErr := createValidator(translator, locale)
	if validatorErr != nil {
		panic("Failed to get validator")
	}

	errors := app.ValidationError{
		Message: app.STATUS_INVALID_DATA,
	}

	err := validate.Struct(data)

	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	for _, err := range validationErrors {
		element := app.ValidationItem{
			Field:   err.Field(),
			Message: err.Translate(translator),
			Value:   err.Param(),
		}
		errors.Items = append(errors.Items, &element)
	}

	return &errors
}

func createValidator(translator ut.Translator, locale string) (*validator.Validate, error) {

	validate := validator.New()

	err := registerTranslation(validate, translator, locale)
	if err != nil {
		return nil, err
	}

	registerTag(validate)

	return validate, nil
}
