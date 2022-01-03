package error_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"idaman.id/storage/internal/error"
)

func TestApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App Suite")
}

var _ = Describe("App Contract", func() {
	Context("Contract constant", func() {
		It("should contain valid constant", func() {
			Expect(error.ERROR_INVALID_DATA).To(Equal("INVALID_DATA"))
			Expect(error.ERROR_TOO_MANY_REQUEST).To(Equal("TOO_MANY_REQUEST"))
			Expect(error.ERROR_NOT_FOUND).To(Equal("NOT_FOUND"))
			Expect(error.ERROR_NOT_SUPPORTED).To(Equal("NOT_SUPPORTED"))
		})
	})
})

var _ = Describe("App Error", func() {
	Context("ValidationError struct", func() {
		var (
			err *error.ValidationError
		)

		BeforeEach(func() {
			err = &error.ValidationError{
				Message: error.ERROR_INVALID_DATA,
			}
		})

		When("Error method called", func() {
			It("should return error message", func() {

				Expect(err.Error()).To(Equal(error.ERROR_INVALID_DATA))
			})
		})
	})

	Context("NewValidationError function", func() {
		var (
			items []*error.ValidationItem
		)

		BeforeEach(func() {
			item1 := &error.ValidationItem{
				Field:   "name",
				Message: "invalid value",
				Value:   "",
			}
			items = append(items, item1)
		})
		When("function called", func() {
			It("should return ValidationError instance", func() {
				expected := &error.ValidationError{
					Message: error.ERROR_INVALID_DATA,
					Items:   items,
				}
				err := error.NewValidationError(items)

				Expect(err).To(MatchError(expected))
			})
		})
	})

	Context("NotSupportedError struct", func() {
		var (
			err *error.NotSupportedError
		)

		BeforeEach(func() {
			err = &error.NotSupportedError{
				Message: error.ERROR_NOT_SUPPORTED,
			}
		})

		When("Error method called", func() {
			It("should return error message", func() {

				Expect(err.Error()).To(Equal(error.ERROR_NOT_SUPPORTED))
			})
		})
	})

	Context("NewNotSupportedError function", func() {
		var (
			context string
		)

		BeforeEach(func() {
			context = "Config"
		})

		When("function called", func() {
			It("should return NotSupportedError instance", func() {
				expected := &error.NotSupportedError{
					Message: error.ERROR_NOT_SUPPORTED,
					Context: context,
				}
				err := error.NewNotSupportedError(context)

				Expect(err).To(MatchError(expected))
			})
		})
	})

	Context("NotFoundError struct", func() {
		var (
			err *error.NotFoundError
		)

		BeforeEach(func() {
			err = &error.NotFoundError{
				Message: error.ERROR_NOT_FOUND,
			}
		})

		When("Error method called", func() {
			It("should return error message", func() {

				Expect(err.Error()).To(Equal(error.ERROR_NOT_FOUND))
			})
		})
	})

	Context("NewNotFoundError function", func() {
		var (
			context string
		)

		BeforeEach(func() {
			context = "Config"
		})

		When("function called", func() {
			It("should return NotFoundError instance", func() {
				expected := &error.NotFoundError{
					Message: error.ERROR_NOT_FOUND,
					Context: context,
				}
				err := error.NewNotFoundError(context)

				Expect(err).To(MatchError(expected))
			})
		})
	})

})
