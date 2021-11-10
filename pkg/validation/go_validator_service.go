package validation

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"idaman.id/storage/pkg/app"
	"idaman.id/storage/pkg/text"
)

type GoValidatorService struct {
	validate Validator
}

func NewGoValidator(validate Validator) (*GoValidatorService, error) {
	service := &GoValidatorService{
		validate: validate,
	}
	err := service.registerCustomization()
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (s *GoValidatorService) ValidateStruct(param interface{}) error {

	isDataTypeStruct := reflect.ValueOf(param).Kind() == reflect.Struct
	if !isDataTypeStruct {
		return app.NewNotSupportedError("Validation")
	}

	vResult := s.validate.Struct(param)
	isDataValid := vResult == nil
	if isDataValid {
		return nil
	}

	var items []*app.ValidationItem

	errors := vResult.(GoValidationErrors)
	for _, err := range errors {
		value := text.Service.ParseString(err.Value())
		element := app.ValidationItem{
			Field:   err.Field(),
			Message: err.Error(),
			Value:   value,
		}
		items = append(items, &element)
	}

	vErr := app.NewValidationError(items)

	return vErr
}

func (s *GoValidatorService) registerCustomization() error {
	var err error

	s.validate.RegisterTagNameFunc(tagNameFunc)

	customValidations := []struct {
		name string
		fn   validator.Func
	}{
		{
			name: "valid_provider",
			fn:   validProviderRule,
		},
		{
			name: "valid_file_amount",
			fn:   validFileAmountRule,
		},
		{
			name: "valid_file_size",
			fn:   validFileSizeRule,
		},
	}

	for _, cValidation := range customValidations {
		err = s.validate.RegisterValidation(cValidation.name, cValidation.fn)
		if err != nil {
			break
		}
	}

	return err
}
