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

type FakeStringParser struct {
}

func (s *FakeStringParser) ParseString(param interface{}) string {
	return ""
}

type FakeConfigGetter struct {
}

func (s *FakeConfigGetter) GetString(key string) string {
	return ""
}

func (s *FakeConfigGetter) GetInt(key string) int {
	return 0
}

func (s *FakeConfigGetter) Get(key string) interface{} {
	return ""
}

type FakeGoValidator struct {
	StructShouldError             bool
	RegisterTagNameFuncCounter    int
	RegisterValidationCounter     int
	RegisterValidationShouldError bool
}

func (mock *FakeGoValidator) Struct(data interface{}) error {
	if mock.StructShouldError {
		var errors validation_go.GoValidationErrors
		err := &FakeFieldError{}
		errors = append(errors, err)
		return errors
	}
	return nil
}

func (mock *FakeGoValidator) RegisterTagNameFunc(fn validator.TagNameFunc) {
	mock.RegisterTagNameFuncCounter++
}

func (mock *FakeGoValidator) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	if mock.RegisterValidationShouldError {
		return errors.New("Fakeed error")
	}
	mock.RegisterValidationCounter++
	return nil
}

type FakeFieldError struct {
}

func (mock *FakeFieldError) Field() string {
	return "Value"
}

func (mock *FakeFieldError) Value() interface{} {
	return ""
}

func (mock *FakeFieldError) Error() string {
	return "Field value is required"
}
