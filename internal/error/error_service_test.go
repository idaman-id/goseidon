package error_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"idaman.id/storage/internal/error"
)

var _ = Describe("App Contract", func() {
	Context("Contract constant", func() {
		It("should contain valid constant", func() {
			Expect(error.STATUS_INVALID_DATA).To(Equal("INVALID_DATA"))
			Expect(error.STATUS_TOO_MANY_REQUEST).To(Equal("TOO_MANY_REQUEST"))
			Expect(error.STATUS_NOT_FOUND).To(Equal("NOT_FOUND"))
			Expect(error.STATUS_NOT_SUPPORTED).To(Equal("NOT_SUPPORTED"))
			Expect(error.STATUS_ALREADY_EXISTS).To(Equal("ALREADY_EXISTS"))
		})
	})
})

var _ = Describe("Error Service", func() {
	Describe("Validation Error", func() {
		Context("ValidationError struct", func() {
			var (
				err *error.ValidationError
			)

			BeforeEach(func() {
				err = &error.ValidationError{
					Message: error.STATUS_INVALID_DATA,
				}
			})

			When("Error method called", func() {
				It("should return error message", func() {

					Expect(err.Error()).To(Equal(error.STATUS_INVALID_DATA))
				})
			})
		})

		Context("NewValidationError function", func() {
			var (
				items []error.ValidationItem
			)

			BeforeEach(func() {
				item1 := error.ValidationItem{
					Field:   "name",
					Message: "invalid value",
					Value:   "",
				}
				items = append(items, item1)
			})
			When("function called", func() {
				It("should return ValidationError instance", func() {
					expected := &error.ValidationError{
						Message: error.STATUS_INVALID_DATA,
						Items:   items,
					}
					err := error.NewValidationError(items)

					Expect(err).To(MatchError(expected))
				})
			})
		})
	})

	Describe("Unsupported Error", func() {
		Context("UnsupportedError struct", func() {
			var (
				err *error.UnsupportedError
			)

			BeforeEach(func() {
				err = &error.UnsupportedError{
					Message: error.STATUS_NOT_SUPPORTED,
				}
			})

			When("Error method called", func() {
				It("should return error message", func() {

					Expect(err.Error()).To(Equal(error.STATUS_NOT_SUPPORTED))
				})
			})
		})

		Context("NewUnsupportedError function", func() {
			var (
				context string
			)

			BeforeEach(func() {
				context = "Config"
			})

			When("function called", func() {
				It("should return UnsupportedError instance", func() {
					expected := &error.UnsupportedError{
						Message: error.STATUS_NOT_SUPPORTED,
						Context: context,
					}
					err := error.NewUnsupportedError(context)

					Expect(err).To(MatchError(expected))
				})
			})
		})
	})

	Describe("Notfound Error", func() {
		Context("NotfoundError struct", func() {
			var (
				err *error.NotfoundError
			)

			BeforeEach(func() {
				err = &error.NotfoundError{
					Context: "File",
					Message: error.STATUS_NOT_FOUND,
				}
			})

			When("Error method called", func() {
				It("should return error message", func() {

					Expect(err.Error()).To(Equal("File is not found"))
				})
			})
		})

		Context("NewNotfoundError function", func() {
			var (
				context string
			)

			BeforeEach(func() {
				context = "Config"
			})

			When("function called", func() {
				It("should return NotfoundError instance", func() {
					expected := &error.NotfoundError{
						Message: error.STATUS_NOT_FOUND,
						Context: context,
					}
					err := error.NewNotfoundError(context)

					Expect(err).To(MatchError(expected))
				})
			})
		})
	})

	Describe("AlreadyExists Error", func() {
		Context("AlreadyExistsError struct", func() {
			var (
				err *error.AlreadyExistsError
			)

			BeforeEach(func() {
				err = &error.AlreadyExistsError{
					Message: error.STATUS_ALREADY_EXISTS,
				}
			})

			When("Error method called", func() {
				It("should return error message", func() {

					Expect(err.Error()).To(Equal(error.STATUS_ALREADY_EXISTS))
				})
			})
		})

		Context("NewAlreadyExistsError function", func() {
			var (
				context string
			)

			BeforeEach(func() {
				context = "Config"
			})

			When("function called", func() {
				It("should return AlreadyExistsError instance", func() {
					expected := &error.AlreadyExistsError{
						Message: error.STATUS_ALREADY_EXISTS,
						Context: context,
					}
					err := error.NewAlreadyExistsError(context)

					Expect(err).To(MatchError(expected))
				})
			})
		})
	})

})
