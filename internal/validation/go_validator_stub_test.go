package validation_test

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"idaman.id/storage/internal/validation"
)

type StubGoValidator struct {
	StructShouldError             bool
	RegisterTagNameFuncCounter    int
	RegisterValidationCounter     int
	RegisterValidationShouldError bool
}

func (mock *StubGoValidator) Struct(data interface{}) error {
	if mock.StructShouldError {
		var errors validation.GoValidationErrors
		err := &StubFieldError{}
		errors = append(errors, err)
		return errors
	}
	return nil
}

func (mock *StubGoValidator) RegisterTagNameFunc(fn validator.TagNameFunc) {
	mock.RegisterTagNameFuncCounter++
}

func (mock *StubGoValidator) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	if mock.RegisterValidationShouldError {
		return errors.New("Stubed error")
	}
	mock.RegisterValidationCounter++
	return nil
}

type StubFieldError struct {
}

func (mock *StubFieldError) Field() string {
	return "Value"
}

func (mock *StubFieldError) Value() interface{} {
	return ""
}

func (mock *StubFieldError) Error() string {
	return "Field value is required"
}
