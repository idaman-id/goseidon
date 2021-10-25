package validation

import (
	"github.com/go-playground/validator/v10"
	"idaman.id/storage/pkg/app"
)

type Map = map[string]interface{}

type Validator interface {
	Struct(data interface{}) error
	RegisterTagNameFunc(fn validator.TagNameFunc)
	RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error
}

type GoFieldError interface {
	Value() interface{}
	Field() string
	Error() string
}

type GoValidationErrors []GoFieldError

func (errs GoValidationErrors) Error() string {
	return app.STATUS_INVALID_DATA
}
