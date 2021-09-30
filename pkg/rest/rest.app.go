package rest

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"idaman.id/storage/pkg/app"
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

func CreateApp() app.App {
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

	app.Use(recover.New())
	app.Get("/", createHomeHandler())
	app.Get("/file/:id", createGetResourceHandler())
	app.Get("/v1/file/:id", createGetDetailHandler())
	app.Post("/v1/file", createUploadFileHandler(dependency))

	return app
}
