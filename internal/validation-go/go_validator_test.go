package validation_go_test

import (
	"errors"

	"testing"

	"github.com/go-playground/validator/v10"
	validation_go "idaman.id/storage/internal/validation-go"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
)

func TestGoValidation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoValidation Package")
}

type StubStringParser struct {
}

func (s *StubStringParser) ParseString(param interface{}) string {
	return ""
}

type StubConfigGetter struct {
}

func (s *StubConfigGetter) GetString(key string) string {
	return ""
}

func (s *StubConfigGetter) GetInt(key string) int {
	return 0
}

func (s *StubConfigGetter) Get(key string) interface{} {
	return ""
}

type StubGoValidator struct {
	StructShouldError             bool
	RegisterTagNameFuncCounter    int
	RegisterValidationCounter     int
	RegisterValidationShouldError bool
}

func (mock *StubGoValidator) Struct(data interface{}) error {
	if mock.StructShouldError {
		var errors validation_go.GoValidationErrors
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
