package builtin_app_test

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	builtin_app "idaman.id/storage/internal/builtin-app"
	app_error "idaman.id/storage/internal/error"
	response "idaman.id/storage/internal/response"
)

var _ = Describe("App Handler", func() {
	var (
		fiberApp *fiber.App
	)

	BeforeEach(func() {
		fiberApp = fiber.New(fiber.Config{
			ErrorHandler: builtin_app.NewErrorHandler(),
		})
	})

	Context("Home Handler", func() {
		BeforeEach(func() {
			fiberApp.Get("/", builtin_app.NewHomeHandler())
		})
		When("home endpoint accessed", func() {
			It("should return success response", func() {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)

				expected := response.NewSuccessResponse(nil)

				Expect(res.StatusCode).To(Equal(fiber.StatusOK))
				Expect(resEntity).To(Equal(expected))
			})
		})
	})

	Context("NotFound Handler", func() {
		BeforeEach(func() {
			fiberApp.Get("*", builtin_app.NewNotFoundHandler())
		})
		When("unavailable endpoint accessed", func() {
			It("should return not found response", func() {
				req := httptest.NewRequest(http.MethodGet, "/unavailable-endpoint", nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)

				expected := response.NewErrorResponse(&response.ResponseParam{
					Message: app_error.STATUS_NOT_FOUND,
				})

				Expect(res.StatusCode).To(Equal(fiber.StatusNotFound))
				Expect(resEntity).To(Equal(expected))
			})
		})
	})

	Context("Error Handler", func() {
		When("error is not fiber error", func() {
			BeforeEach(func() {
				fiberApp.Get("error-handler", func(c *fiber.Ctx) error {
					return errors.New("custom error handler")
				})
			})

			It("should return error response with internal server error status code", func() {
				req := httptest.NewRequest(http.MethodGet, "/error-handler", nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)

				expected := response.NewErrorResponse(&response.ResponseParam{
					Message: "custom error handler",
				})

				Expect(res.StatusCode).To(Equal(fiber.StatusInternalServerError))
				Expect(resEntity).To(Equal(expected))
			})

			It("should return application/json header content type", func() {
				fiberApp.Get("error-handler", func(c *fiber.Ctx) error {
					return errors.New("custom error handler")
				})

				req := httptest.NewRequest(http.MethodGet, "/error-handler", nil)
				res, _ := fiberApp.Test(req)

				Expect(res.Header.Get(fiber.HeaderContentType)).To(Equal(fiber.MIMEApplicationJSON))
			})
		})

		When("error is fiber error", func() {
			BeforeEach(func() {
				fiberApp.Get("error-handler", func(c *fiber.Ctx) error {
					return fiber.ErrBadGateway
				})
			})
			It("should return error response with custom status code", func() {
				req := httptest.NewRequest(http.MethodGet, "/error-handler", nil)
				res, _ := fiberApp.Test(req)

				resEntity := UnmarshallResponseBody(res.Body)

				expected := response.NewErrorResponse(&response.ResponseParam{
					Message: "Bad Gateway",
				})

				Expect(res.StatusCode).To(Equal(fiber.StatusBadGateway))
				Expect(resEntity).To(Equal(expected))
			})

			It("should return application/json header content type", func() {

				req := httptest.NewRequest(http.MethodGet, "/error-handler", nil)
				res, _ := fiberApp.Test(req)

				Expect(res.Header.Get(fiber.HeaderContentType)).To(Equal(fiber.MIMEApplicationJSON))
			})
		})
	})

})
