package builtin_app_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	builtin_app "idaman.id/storage/internal/builtin-app"
	"idaman.id/storage/internal/deleting"
	app_error "idaman.id/storage/internal/error"
	response "idaman.id/storage/internal/response"
	"idaman.id/storage/internal/retrieving"
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
			fileGetterService = &FakeFileGetterService{}
			fiberApp.Get("/v1/file/:identifier", builtin_app.NewFileGetDetailHandler(fileGetterService))
		})

		When("file not found", func() {
			It("should return not found response", func() {
				identifier = "not-found"
				req := httptest.NewRequest(http.MethodGet, "/v1/file/"+identifier, nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)

				expected := response.NewErrorResponse(&response.ResponseParam{
					Message: "File is not found",
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

	Context("GetFileResource Handler", func() {
		var (
			identifier           string
			fileRetrieverService retrieving.FileRetriever
		)

		BeforeEach(func() {
			identifier = "fake-identifier"
			fileRetrieverService = &FakeFileRetrieverService{}
			fiberApp.Get("/file/:identifier", builtin_app.NewGetResourceHandler(fileRetrieverService))
		})

		When("file not found", func() {
			It("should return not found response", func() {
				identifier = "not-found"
				req := httptest.NewRequest(http.MethodGet, "/file/"+identifier, nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)

				expected := response.NewErrorResponse(&response.ResponseParam{
					Message: "File is not found",
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

	Context("DeleteFile Handler", func() {

		var (
			identifier       string
			deleteServiceSpy *FileDeleteServiceSpy
		)

		markDeleteSucceeded := func() {
			deleteServiceSpy.ErrorResultOfDeleteFile = nil
		}

		markDeleteFailed := func(theError error) {
			deleteServiceSpy.ErrorResultOfDeleteFile = theError
		}

		BeforeEach(func() {
			identifier = "fake-identifier"
			deleteServiceSpy = &FileDeleteServiceSpy{}
			markDeleteSucceeded()
		})

		sendRequest := func(service deleting.DeleteService) (*http.Request, *http.Response) {
			fiberApp.Delete("/v1/file/:identifier", builtin_app.DeleteFileHandler(service))

			req := httptest.NewRequest(http.MethodDelete, "/v1/file/"+identifier, nil)
			res, _ := fiberApp.Test(req)
			return req, res
		}

		It("calls service 'Delete File' method with correct argument", func() {
			sendRequest(deleteServiceSpy)
			Expect(deleteServiceSpy.LastIdentifierOfDeleteFile).To(Equal(identifier))
		})

		It("returns http success response after file deletion was successful", func() {
			_, res := sendRequest(deleteServiceSpy)
			Expect(res.StatusCode).To(Equal(fiber.StatusOK))
			Expect(res.Body.Close()).To(BeNil())
		})

		It("returns Not Found http response if file to be deleted was not found", func() {
			markDeleteFailed(app_error.NewNotfoundError("abc.jpeg"))
			expectedResponseEntity := response.NewErrorResponse(&response.ResponseParam{
				Message: "abc.jpeg is not found",
			})

			_, res := sendRequest(deleteServiceSpy)
			responseEntity := UnmarshallResponseBody(res.Body)

			Expect(res.StatusCode).To(Equal(fiber.StatusNotFound))
			Expect(responseEntity).To(Equal(expectedResponseEntity))
		})

		It("returns Bad Request http response if unexpected error happened", func() {
			markDeleteFailed(app_error.NewUnknownError("unknown error message"))
			expectedResponseEntity := response.NewErrorResponse(&response.ResponseParam{
				Message: "unknown error message",
			})

			_, res := sendRequest(deleteServiceSpy)
			responseEntity := UnmarshallResponseBody(res.Body)

			Expect(res.StatusCode).To(Equal(fiber.StatusBadRequest))
			Expect(responseEntity).To(Equal(expectedResponseEntity))
		})
	})

})
