package rest

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	application "idaman.id/storage/pkg/app"
	"idaman.id/storage/pkg/translation"
)

func localeParser(ctx Context) string {
	locale := ctx.Query("lang")
	if locale == "" {
		locale = ctx.Get("Accept-Language", "en")
	}
	return locale
}

func createLocalizerFactory(i18nBundle *i18n.Bundle) func(ctx Context) *i18n.Localizer {
	return func(ctx Context) *i18n.Localizer {
		locale := localeParser(ctx)
		localizer := i18n.NewLocalizer(i18nBundle, locale)
		return localizer
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
				localizer := dependency.localizer(ctx)
				translator := translation.CreateSimpleTranslator(localizer)
				response := createFailedResponse(ResponseDto{
					Message:    application.STATUS_TOO_MANY_REQUEST,
					Translator: translator,
				})
				return ctx.Status(fiber.StatusTooManyRequests).JSON(response)
			},
		}
		return config
	}
}

func CreateApp() application.App {
	i18nBundle := i18n.NewBundle(language.English)

	i18nBundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	i18nBundle.MustLoadMessageFile("pkg/translation/status.en.json")
	i18nBundle.MustLoadMessageFile("pkg/translation/status.id.json")

	localizer := createLocalizerFactory(i18nBundle)

	dependency := &Dependency{
		localizer:    localizer,
		localeParser: localeParser,
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: createErrorHandler(),
	})

	createConfig := createLimiterConfig(dependency)

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(etag.New())
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(limiter.New(createConfig()))
	app.Get("/", createHomeHandler(dependency))
	app.Get("/file/:id", createGetResourceHandler(dependency))
	app.Get("/v1/file/:id", createGetDetailHandler(dependency))
	app.Post("/v1/file", createUploadFileHandler(dependency))

	return app
}
