package validation_go_test

import (
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"idaman.id/storage/internal/config"
	"idaman.id/storage/internal/error"
	"idaman.id/storage/internal/text"
	validation_go "idaman.id/storage/internal/validation-go"
)

var _ = Describe("GoValidator Service", func() {
	Context("NewGoValidator function", func() {
		var (
			stubValidator    *FakeGoValidator
			stubStringParser text.StringParser
			stubConfigGetter config.Getter
			service          *validation_go.GoValidatorService
		)

		BeforeEach(func() {
			stubValidator = &FakeGoValidator{}
			stubStringParser = &FakeStringParser{}
			stubConfigGetter = &FakeConfigGetter{}
			service, _ = validation_go.NewGoValidator(stubValidator, stubStringParser, stubConfigGetter)
		})

		When("called", func() {
			It("should return GoValidator instance", func() {
				resDataType := reflect.TypeOf(service)
				expectedDataType := reflect.TypeOf(&validation_go.GoValidatorService{})

				Expect(resDataType).To(Equal(expectedDataType))
			})

			It("should register tag name function", func() {
				Expect(stubValidator.RegisterTagNameFuncCounter).To(Equal(1))
			})

			It("should register custom validation function", func() {
				Expect(stubValidator.RegisterValidationCounter).To(Equal(3))
			})

			It("should return error when failed create validator", func() {
				stubValidator = &FakeGoValidator{
					RegisterValidationShouldError: true,
				}
				service, err := validation_go.NewGoValidator(stubValidator, stubStringParser, stubConfigGetter)

				Expect(err).ToNot(BeNil())
				Expect(service).To(BeNil())
			})
		})
	})

	Context("Validate method", func() {
		var (
			stubValidator    *FakeGoValidator
			stubStringParser text.StringParser
			stubConfigGetter config.Getter
			service          *validation_go.GoValidatorService
		)

		BeforeEach(func() {
			stubValidator = &FakeGoValidator{}
			stubStringParser = &FakeStringParser{}
			stubConfigGetter = &FakeConfigGetter{}
			service, _ = validation_go.NewGoValidator(stubValidator, stubStringParser, stubConfigGetter)
		})

		When("param is not a struct", func() {
			It("should return UnsupportedError", func() {
				expected := error.NewUnsupportedError("Validation")
				res := service.Validate("")

				Expect(res).To(MatchError(expected))
			})
		})

		When("param is struct and contain no rule", func() {
			It("should return nil", func() {
				type Rule struct {
				}
				param := Rule{}
				res := service.Validate(param)

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
				res := service.Validate(param)

				Expect(res).To(BeNil())
			})
		})

		When("param contain invalid data", func() {
			It("should return ValidationError", func() {
				stubValidator = &FakeGoValidator{
					StructShouldError: true,
				}
				service, _ = validation_go.NewGoValidator(stubValidator, stubStringParser, stubConfigGetter)

				type Rule struct {
					Value int `validate:"required"`
				}
				param := Rule{}
				res := service.Validate(param)

				var items []error.ValidationItem
				item := error.ValidationItem{
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
