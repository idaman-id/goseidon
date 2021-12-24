package rest_fiber_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"idaman.id/storage/internal/deleting"
	rest_fiber "idaman.id/storage/internal/rest-fiber"
	response "idaman.id/storage/internal/rest-response"
	"idaman.id/storage/internal/retrieving"
	app_error "idaman.id/storage/pkg/error"
)

var _ = Describe("File Handler", func() {
	var (
		fiberApp *fiber.App
	)

	BeforeEach(func() {
		fiberApp = fiber.New()
	})

	Context("FileGetDetail Handler", func() {
		var (
			identifier        string
			fileGetterService retrieving.FileGetter
		)

		BeforeEach(func() {
			identifier = "fake-identifier"
			fileGetterService = &StubFileGetterService{}
			fiberApp.Get("/v1/file/:identifier", rest_fiber.NewFileGetDetailHandler(fileGetterService))
		})

		When("file not found", func() {
			It("should return not found response", func() {
				identifier = "not-found"
				req := httptest.NewRequest(http.MethodGet, "/v1/file/"+identifier, nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)

				expected := response.NewErrorResponse(&response.ResponseParam{
					Message: app_error.ERROR_NOT_FOUND,
				})

				Expect(res.StatusCode).To(Equal(fiber.StatusNotFound))
				Expect(resEntity).To(Equal(expected))
			})
		})

		When("unexpected error happened", func() {
			It("should return error response", func() {
				identifier = "error"
				req := httptest.NewRequest(http.MethodGet, "/v1/file/"+identifier, nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)

				expected := response.NewErrorResponse(&response.ResponseParam{
					Message: response.STATUS_ERROR,
				})

				Expect(res.StatusCode).To(Equal(fiber.StatusBadRequest))
				Expect(resEntity).To(Equal(expected))
			})
		})

		When("file available", func() {
			It("should return success response", func() {
				req := httptest.NewRequest(http.MethodGet, "/v1/file/"+identifier, nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)

				file := retrieving.FileEntity{}
				expected := response.NewSuccessResponse(&response.ResponseParam{
					Message: response.STATUS_OK,
					Data:    file,
				})

				Expect(res.StatusCode).To(Equal(fiber.StatusOK))
				Expect(resEntity.Message).To(Equal(expected.Message))
				Expect(resEntity.Data).ToNot(BeNil())
				Expect(resEntity.Error).To(BeNil())
			})
		})
	})

	Context("DeleteFile Handler", func() {
		var (
			identifier         string
			fileDeleterService deleting.DeleteService
		)

		BeforeEach(func() {
			identifier = "fake-identifier"
			fileDeleterService = &StubDeleteService{}
			fiberApp.Delete("/v1/file/:identifier", rest_fiber.NewDeleteFileHandler(fileDeleterService))
		})

		When("file not found", func() {
			It("should return not found response", func() {
				identifier = "not-found"
				req := httptest.NewRequest(http.MethodDelete, "/v1/file/"+identifier, nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)

				expected := response.NewErrorResponse(&response.ResponseParam{
					Message: app_error.ERROR_NOT_FOUND,
				})

				Expect(res.StatusCode).To(Equal(fiber.StatusNotFound))
				Expect(resEntity).To(Equal(expected))
			})
		})

		When("unexpected error happened", func() {
			It("should return error response", func() {
				identifier = "error"
				req := httptest.NewRequest(http.MethodDelete, "/v1/file/"+identifier, nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)

				expected := response.NewErrorResponse(&response.ResponseParam{
					Message: response.STATUS_ERROR,
				})

				Expect(res.StatusCode).To(Equal(fiber.StatusBadRequest))
				Expect(resEntity).To(Equal(expected))
			})
		})

		When("file available", func() {
			It("should return success response", func() {
				req := httptest.NewRequest(http.MethodDelete, "/v1/file/"+identifier, nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)
				expected := response.NewSuccessResponse(nil)

				Expect(res.StatusCode).To(Equal(fiber.StatusOK))
				Expect(resEntity).To(Equal(expected))
			})
		})

	})

	Context("GetFileResource Handler", func() {
		var (
			identifier           string
			fileRetrieverService retrieving.FileRetriever
		)

		BeforeEach(func() {
			identifier = "fake-identifier"
			fileRetrieverService = &StubFileRetrieverService{}
			fiberApp.Get("/file/:identifier", rest_fiber.NewGetResourceHandler(fileRetrieverService))
		})

		When("file not found", func() {
			It("should return not found response", func() {
				identifier = "not-found"
				req := httptest.NewRequest(http.MethodGet, "/file/"+identifier, nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)

				expected := response.NewErrorResponse(&response.ResponseParam{
					Message: app_error.ERROR_NOT_FOUND,
				})

				Expect(res.StatusCode).To(Equal(fiber.StatusNotFound))
				Expect(resEntity).To(Equal(expected))
			})
		})

		When("unexpected error happened", func() {
			It("should return error response", func() {
				identifier = "error"
				req := httptest.NewRequest(http.MethodGet, "/file/"+identifier, nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)

				expected := response.NewErrorResponse(&response.ResponseParam{
					Message: response.STATUS_ERROR,
				})

				Expect(res.StatusCode).To(Equal(fiber.StatusBadRequest))
				Expect(resEntity).To(Equal(expected))
			})
		})

		When("file available", func() {
			It("should return success response", func() {
				req := httptest.NewRequest(http.MethodGet, "/file/"+identifier, nil)
				res, _ := fiberApp.Test(req)

				Expect(res.StatusCode).To(Equal(fiber.StatusOK))
				Expect(res.Body.Close()).To(BeNil())
			})
		})

	})

})
