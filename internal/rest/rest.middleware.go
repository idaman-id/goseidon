package rest

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	application "idaman.id/storage/pkg/app"
	"idaman.id/storage/pkg/translation"
)

func createErrorHandler(dependency *Dependency) ErrorHandler {
	return func(ctx Context, err error) Result {
		code := fiber.StatusInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		localizer := dependency.getLocalizer(ctx)
		translator := translation.NewGoI18nService(localizer).Translate

		response := createFailedResponse(ResponseParam{
			Message:    err.Error(),
			Translator: translator,
		})

		return ctx.Status(code).JSON(response)
	}
}

func createLimiterConfig(dependency *Dependency) func() limiter.Config {
	return func() limiter.Config {
		config := limiter.Config{
			Max:        20,
			Expiration: 30 * time.Second,
			KeyGenerator: func(ctx *fiber.Ctx) string {
				return ctx.Get("x-forwarded-for")
			},
			LimitReached: func(ctx *fiber.Ctx) error {
				localizer := dependency.getLocalizer(ctx)
				translator := translation.NewGoI18nService(localizer).Translate
				response := createFailedResponse(ResponseParam{
					Message:    application.STATUS_TOO_MANY_REQUEST,
					Translator: translator,
				})
				return ctx.Status(fiber.StatusTooManyRequests).JSON(response)
			},
		}
		return config
	}
}
