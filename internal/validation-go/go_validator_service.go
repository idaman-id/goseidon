package validation_go

import (
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"idaman.id/storage/internal/config"
	app_error "idaman.id/storage/internal/error"
)

type goValidationService struct {
	v  *validator.Validate
	t  ut.Translator
	cg config.Getter
}

func (s *goValidationService) Validate(i interface{}) error {
	isTypeValid := reflect.ValueOf(i).Kind() == reflect.Struct
	if !isTypeValid {
		return app_error.NewUnsupportedError("Validation")
	}

	err := s.v.Struct(i)
	if err == nil {
		return nil
	}

	var items []app_error.ValidationItem
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		element := app_error.ValidationItem{
			Field:   err.Field(),
			Message: err.Error(),
		}
		items = append(items, element)
	}

	vErr := app_error.NewValidationError(items)
	return vErr
}

func NewGoValidator(cg config.Getter) (*goValidationService, error) {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	v := validator.New()
	err := en_translations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		return nil, err
	}
	tFn := NewTagNameFunc()
	v.RegisterTagNameFunc(tFn)

	cValidations := []struct {
		name string
		fn   validator.Func
	}{
		{
			name: "valid_file_size",
			fn:   NewValidFileSizeRule(cg),
		},
	}

	for _, cv := range cValidations {
		err = v.RegisterValidation(cv.name, cv.fn)
		if err != nil {
			break
		}
	}

	s := &goValidationService{
		v:  v,
		t:  trans,
		cg: cg,
	}
	return s, nil
}
