package rest

import (
	"github.com/gofiber/fiber/v2"
)

func createErrorHandler() ErrorHandler {
	return func(ctx Context, err error) Result {
		code := fiber.StatusInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		response := createFailedResponse(ResponseDto{
			Message: err.Error(),
		})

		return ctx.Status(code).JSON(response)
	}
}
