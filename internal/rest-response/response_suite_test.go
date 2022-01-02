package rest_response_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	response "idaman.id/storage/internal/rest-response"
)

func TestResponse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Response Suite")
}

var _ = Describe("Response Contract", func() {
	Context("Contract constant", func() {
		It("should contain valid constant", func() {
			Expect(response.STATUS_OK).To(Equal("OK"))
			Expect(response.STATUS_ERROR).To(Equal("ERROR"))
		})
	})
})

var _ = Describe("Response Service", func() {

	Context("NewSuccessResponse function", func() {
		var (
			param *response.ResponseParam
		)

		BeforeEach(func() {
			param = &response.ResponseParam{}
		})

		When("parameter is nil", func() {
			It("should return default success response", func() {
				param = nil
				res := response.NewSuccessResponse(param)

				expected := &response.ResponseEntity{
					Message: response.STATUS_OK,
				}
				Expect(res).To(Equal(expected))
			})
		})

		When("parameter is empty", func() {
			It("should return default success response", func() {
				res := response.NewSuccessResponse(param)

				expected := &response.ResponseEntity{
					Message: response.STATUS_OK,
				}
				Expect(res).To(Equal(expected))
			})
		})

		When("message is specified in parameter", func() {
			It("should return message in response", func() {
				message := "custom message"
				param = &response.ResponseParam{
					Message: message,
				}
				res := response.NewSuccessResponse(param)

				expected := &response.ResponseEntity{
					Message: message,
				}
				Expect(res).To(Equal(expected))
			})
		})

		When("message is blank space", func() {
			It("should return default success message", func() {
				message := "   "
				param = &response.ResponseParam{
					Message: message,
				}
				res := response.NewSuccessResponse(param)

				expected := &response.ResponseEntity{
					Message: response.STATUS_OK,
				}
				Expect(res).To(Equal(expected))
			})
		})

		When("data is specified in parameter", func() {
			It("should return data in response", func() {
				data := []string{}
				param = &response.ResponseParam{
					Data: data,
				}
				res := response.NewSuccessResponse(param)

				expected := &response.ResponseEntity{
					Message: response.STATUS_OK,
					Data:    data,
				}
				Expect(res).To(Equal(expected))
			})
		})

	})

	Context("NewErrorResponse function", func() {
		var (
			param *response.ResponseParam
		)

		BeforeEach(func() {
			param = &response.ResponseParam{}
		})

		When("parameter is nil", func() {
			It("should return default error response", func() {
				param = nil
				res := response.NewErrorResponse(param)

				expected := &response.ResponseEntity{
					Message: response.STATUS_ERROR,
				}
				Expect(res).To(Equal(expected))
			})
		})

		When("parameter is empty", func() {
			It("should return default error response", func() {
				res := response.NewErrorResponse(param)

				expected := &response.ResponseEntity{
					Message: response.STATUS_ERROR,
				}
				Expect(res).To(Equal(expected))
			})
		})

		When("message is specified in parameter", func() {
			It("should return message in response", func() {
				message := "custom message"
				param = &response.ResponseParam{
					Message: message,
				}
				res := response.NewErrorResponse(param)

				expected := &response.ResponseEntity{
					Message: message,
				}
				Expect(res).To(Equal(expected))
			})
		})

		When("message is blank space", func() {
			It("should return default error message", func() {
				message := "   "
				param = &response.ResponseParam{
					Message: message,
				}
				res := response.NewErrorResponse(param)

				expected := &response.ResponseEntity{
					Message: response.STATUS_ERROR,
				}
				Expect(res).To(Equal(expected))
			})
		})

		When("error is specified in parameter", func() {
			It("should return error in response", func() {
				error := []string{}
				param = &response.ResponseParam{
					Error: error,
				}
				res := response.NewErrorResponse(param)

				expected := &response.ResponseEntity{
					Message: response.STATUS_ERROR,
					Error:   error,
				}
				Expect(res).To(Equal(expected))
			})
		})

	})

})
