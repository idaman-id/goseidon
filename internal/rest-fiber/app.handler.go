package rest_fiber

import (
	"github.com/gofiber/fiber/v2"
	response "idaman.id/storage/internal/rest-response"
	"idaman.id/storage/pkg/app"
)

func NewHomeHandler() Handler {
	return func(ctx *Context) error {
		responseEntity := response.NewSuccessResponse(nil)
		return ctx.JSON(responseEntity)
	}
}

func NewNotFoundHandler() Handler {
	return func(ctx *Context) error {
		responseEntity := response.NewErrorResponse(&response.ResponseParam{
			Message: app.STATUS_NOT_FOUND,
		})
		return ctx.Status(fiber.StatusNotFound).JSON(responseEntity)
	}
}

func NewErrorHandler() ErrorHandler {
	return func(ctx *Context, err error) error {

		statusCode := fiber.StatusInternalServerError

		fiberErr, isFiberError := err.(*fiber.Error)
		if isFiberError {
			statusCode = fiberErr.Code
		}

		responseEntity := response.NewErrorResponse(&response.ResponseParam{
			Message: err.Error(),
		})

		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return ctx.Status(statusCode).JSON(responseEntity)
	}
}
