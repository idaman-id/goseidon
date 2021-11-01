package rest_fiber_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	rest_fiber "idaman.id/storage/internal/rest-fiber"
	response "idaman.id/storage/internal/rest-response"
	test "idaman.id/storage/internal/rest-test"
	"idaman.id/storage/internal/retrieving"
	"idaman.id/storage/pkg/app"
)

var _ = Describe("File Handler", func() {
	var (
		fiberApp *fiber.App
	)

	BeforeEach(func() {
		fiberApp = rest_fiber.NewApp()
	})

	Describe("FileGetDetail Handler", func() {
		var (
			identifier        string
			fileGetterService retrieving.FileGetter
		)

		BeforeEach(func() {
			fileGetterService = &retrieving.StubFileGetterService{}
			fiberApp.Get("/v1/file/:identifier", rest_fiber.NewFileGetDetailHandler(fileGetterService))
		})

		BeforeEach(func() {
			identifier = "fake-identifier"
		})

		When("file not found", func() {
			It("should return not found response", func() {
				identifier = "not-found"
				req := httptest.NewRequest(http.MethodGet, "/v1/file/"+identifier, nil)
				res, _ := fiberApp.Test(req)

				resEntity := test.UnmarshallResponseBody(res.Body)

				expected := response.NewErrorResponse(&response.ResponseParam{
					Message: app.STATUS_NOT_FOUND,
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

				resEntity := test.UnmarshallResponseBody(res.Body)

				expected := response.NewErrorResponse(&response.ResponseParam{
					Message: app.STATUS_ERROR,
				})

				Expect(res.StatusCode).To(Equal(fiber.StatusBadRequest))
				Expect(resEntity).To(Equal(expected))
			})
		})

		When("file available", func() {
			It("should return success response", func() {
				req := httptest.NewRequest(http.MethodGet, "/v1/file/"+identifier, nil)
				res, _ := fiberApp.Test(req)

				resEntity := test.UnmarshallResponseBody(res.Body)

				file := retrieving.FileEntity{}
				expected := response.NewSuccessResponse(&response.ResponseParam{
					Message: app.STATUS_OK,
					Data:    file,
				})

				Expect(res.StatusCode).To(Equal(fiber.StatusOK))
				Expect(resEntity.Message).To(Equal(expected.Message))
				Expect(resEntity.Data).ToNot(BeNil())
				Expect(resEntity.Error).To(BeNil())
			})
		})
	})

})
