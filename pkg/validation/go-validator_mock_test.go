package validation_test

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"idaman.id/storage/pkg/validation"
)

type MockGoValidator struct {
	StructShouldError             bool
	RegisterTagNameFuncCounter    int
	RegisterValidationCounter     int
	RegisterValidationShouldError bool
}

func (mock *MockGoValidator) Struct(data interface{}) error {
	if mock.StructShouldError {
		var errors validation.GoValidationErrors
		err := &MockFieldError{}
		errors = append(errors, err)
		return errors
	}
	return nil
}

func (mock *MockGoValidator) RegisterTagNameFunc(fn validator.TagNameFunc) {
	mock.RegisterTagNameFuncCounter++
}

func (mock *MockGoValidator) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	if mock.RegisterValidationShouldError {
		return errors.New("Mocked error")
	}
	mock.RegisterValidationCounter++
	return nil
}

type MockFieldError struct {
}

func (mock *MockFieldError) Field() string {
	return "Value"
}

func (mock *MockFieldError) Value() interface{} {
	return ""
}

func (mock *MockFieldError) Error() string {
	return "Field value is required"
}
