package app

import (
	"github.com/gofiber/fiber/v2"
	"idaman.id/storage/src/http"
)

func HomeHandler(ctx *fiber.Ctx) error {

	response := http.CreateSuccessResponse(http.ResponseParameter{})
	return ctx.JSON(response)
}

var ErrorHandler = func(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	response := http.CreateFailedResponse(http.ResponseParameter{
		Message: err.Error(),
	})

	return ctx.Status(code).JSON(response)
}
