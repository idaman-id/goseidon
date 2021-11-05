package validation_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"idaman.id/storage/pkg/app"
	"idaman.id/storage/pkg/validation"
)

var _ = Describe("GoValidator Service", func() {
	Describe("NewGoValidator function", func() {
		var (
			mock    *StubGoValidator
			service *validation.GoValidatorService
		)

		BeforeEach(func() {
			mock = &StubGoValidator{}
			service, _ = validation.NewGoValidator(mock)
		})

		When("called", func() {
			It("should return GoValidator instance", func() {
				resDataType := reflect.TypeOf(service)
				expectedDataType := reflect.TypeOf(&validation.GoValidatorService{})

				Expect(resDataType).To(Equal(expectedDataType))
			})

			It("should register tag name function", func() {
				Expect(mock.RegisterTagNameFuncCounter).To(Equal(1))
			})

			It("should register custom validation function", func() {
				Expect(mock.RegisterValidationCounter).To(Equal(4))
			})

			It("should return error when failed create validator", func() {
				mock = &StubGoValidator{
					RegisterValidationShouldError: true,
				}
				service, err := validation.NewGoValidator(mock)

				Expect(err).ToNot(BeNil())
				Expect(service).To(BeNil())
			})
		})
	})

	Describe("ValidateStruct method", func() {
		var (
			mock    *StubGoValidator
			service *validation.GoValidatorService
		)

		BeforeEach(func() {
			mock = &StubGoValidator{}
			service, _ = validation.NewGoValidator(mock)
		})

		When("param is not a struct", func() {
			It("should return NotSupportedError", func() {
				expected := app.NewNotSupportedError("Validation")
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
				mock = &StubGoValidator{
					StructShouldError: true,
				}
				service, _ = validation.NewGoValidator(mock)

				type Rule struct {
					Value int `validate:"required"`
				}
				param := Rule{}
				res := service.ValidateStruct(param)

				var items []*app.ValidationItem
				item := &app.ValidationItem{
					Field:   "Value",
					Message: "Field value is required",
					Value:   "",
				}
				items = append(items, item)
				expected := app.NewValidationError(items)
				Expect(res).To(MatchError(expected))
			})
		})
	})

})
