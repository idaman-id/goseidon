package validation_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"idaman.id/storage/pkg/config"
	"idaman.id/storage/pkg/error"
	"idaman.id/storage/pkg/text"
	"idaman.id/storage/pkg/validation"
)

var _ = Describe("GoValidator Service", func() {
	Context("NewGoValidator function", func() {
		var (
			stubValidator    *StubGoValidator
			stubStringParser text.StringParser
			stubConfigGetter config.Getter
			service          *validation.GoValidatorService
		)

		BeforeEach(func() {
			stubValidator = &StubGoValidator{}
			stubStringParser = &StubStringParser{}
			stubConfigGetter = &StubConfigGetter{}
			service, _ = validation.NewGoValidator(stubValidator, stubStringParser, stubConfigGetter)
		})

		When("called", func() {
			It("should return GoValidator instance", func() {
				resDataType := reflect.TypeOf(service)
				expectedDataType := reflect.TypeOf(&validation.GoValidatorService{})

				Expect(resDataType).To(Equal(expectedDataType))
			})

			It("should register tag name function", func() {
				Expect(stubValidator.RegisterTagNameFuncCounter).To(Equal(1))
			})

			It("should register custom validation function", func() {
				Expect(stubValidator.RegisterValidationCounter).To(Equal(3))
			})

			It("should return error when failed create validator", func() {
				stubValidator = &StubGoValidator{
					RegisterValidationShouldError: true,
				}
				service, err := validation.NewGoValidator(stubValidator, stubStringParser, stubConfigGetter)

				Expect(err).ToNot(BeNil())
				Expect(service).To(BeNil())
			})
		})
	})

	Context("ValidateStruct method", func() {
		var (
			stubValidator    *StubGoValidator
			stubStringParser text.StringParser
			stubConfigGetter config.Getter
			service          *validation.GoValidatorService
		)

		BeforeEach(func() {
			stubValidator = &StubGoValidator{}
			stubStringParser = &StubStringParser{}
			stubConfigGetter = &StubConfigGetter{}
			service, _ = validation.NewGoValidator(stubValidator, stubStringParser, stubConfigGetter)
		})

		When("param is not a struct", func() {
			It("should return NotSupportedError", func() {
				expected := error.NewNotSupportedError("Validation")
				res := service.ValidateStruct("")

				Expect(res).To(MatchError(expected))
			})
		})

		When("param is struct and contain no rule", func() {
			It("should return nil", func() {
				type Rule struct {
				}
				param := Rule{}
				res := service.ValidateStruct(param)

				Expect(res).To(BeNil())
			})
		})

		When("param contain valid data", func() {
			It("should return nil", func() {
				type Rule struct {
					Value int `validate:"required"`
				}
				param := Rule{
					Value: 1,
				}
				res := service.ValidateStruct(param)

				Expect(res).To(BeNil())
			})
		})

		When("param contain invalid data", func() {
			It("should return ValidationError", func() {
				stubValidator = &StubGoValidator{
					StructShouldError: true,
				}
				service, _ = validation.NewGoValidator(stubValidator, stubStringParser, stubConfigGetter)

				type Rule struct {
					Value int `validate:"required"`
				}
				param := Rule{}
				res := service.ValidateStruct(param)

				var items []*error.ValidationItem
				item := &error.ValidationItem{
					Field:   "Value",
					Message: "Field value is required",
					Value:   "",
				}
				items = append(items, item)
				expected := error.NewValidationError(items)
				Expect(res).To(MatchError(expected))
			})
		})
	})

})
