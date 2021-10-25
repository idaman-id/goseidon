package app_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"idaman.id/storage/pkg/app"
)

func TestApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App Suite")
}

var _ = Describe("App Contract", func() {
	Describe("Contract constant", func() {
		It("should contain valid constant", func() {
			Expect(app.STATUS_OK).To(Equal("OK"))
			Expect(app.STATUS_ERROR).To(Equal("ERROR"))
			Expect(app.STATUS_INVALID_DATA).To(Equal("INVALID_DATA"))
			Expect(app.STATUS_TOO_MANY_REQUEST).To(Equal("TOO_MANY_REQUEST"))
			Expect(app.STATUS_NOT_FOUND).To(Equal("NOT_FOUND"))
			Expect(app.STATUS_NOT_SUPPORTED).To(Equal("NOT_SUPPORTED"))
		})
	})
})

var _ = Describe("App Error", func() {
	Describe("ValidationError struct", func() {
		var (
			err *app.ValidationError
		)

		BeforeEach(func() {
			err = &app.ValidationError{
				Message: app.STATUS_INVALID_DATA,
			}
		})

		When("Error method called", func() {
			It("should return error message", func() {

				Expect(err.Error()).To(Equal(app.STATUS_INVALID_DATA))
			})
		})
	})

	Describe("NewValidationError function", func() {
		var (
			items []*app.ValidationItem
		)

		BeforeEach(func() {
			item1 := &app.ValidationItem{
				Field:   "name",
				Message: "invalid value",
				Value:   "",
			}
			items = append(items, item1)
		})
		When("function called", func() {
			It("should return ValidationError instance", func() {
				expected := &app.ValidationError{
					Message: app.STATUS_INVALID_DATA,
					Items:   items,
				}
				err := app.NewValidationError(items)

				Expect(err).To(MatchError(expected))
			})
		})
	})

	Describe("NotSupportedError struct", func() {
		var (
			err *app.NotSupportedError
		)

		BeforeEach(func() {
			err = &app.NotSupportedError{
				Message: app.STATUS_NOT_SUPPORTED,
			}
		})

		When("Error method called", func() {
			It("should return error message", func() {

				Expect(err.Error()).To(Equal(app.STATUS_NOT_SUPPORTED))
			})
		})
	})

	Describe("NewNotSupportedError function", func() {
		var (
			context string
		)

		BeforeEach(func() {
			context = "Config"
		})

		When("function called", func() {
			It("should return NotSupportedError instance", func() {
				expected := &app.NotSupportedError{
					Message: app.STATUS_NOT_SUPPORTED,
					Context: context,
				}
				err := app.NewNotSupportedError(context)

				Expect(err).To(MatchError(expected))
			})
		})
	})

	Describe("NotFoundError struct", func() {
		var (
			err *app.NotFoundError
		)

		BeforeEach(func() {
			err = &app.NotFoundError{
				Message: app.STATUS_NOT_FOUND,
			}
		})

		When("Error method called", func() {
			It("should return error message", func() {

				Expect(err.Error()).To(Equal(app.STATUS_NOT_FOUND))
			})
		})
	})

	Describe("NewNotFoundError function", func() {
		var (
			context string
		)

		BeforeEach(func() {
			context = "Config"
		})

		When("function called", func() {
			It("should return NotFoundError instance", func() {
				expected := &app.NotFoundError{
					Message: app.STATUS_NOT_FOUND,
					Context: context,
				}
				err := app.NewNotFoundError(context)

				Expect(err).To(MatchError(expected))
			})
		})
	})

})
