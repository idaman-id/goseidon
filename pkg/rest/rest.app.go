package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"idaman.id/storage/pkg/app"
)

func CreateApp() app.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: createErrorHandler(),
	})

	app.Use(recover.New())
	app.Get("/", createHomeHandler())
	app.Get("/file/:id", createGetResourceHandler())
	app.Get("/v1/file/:id", createGetDetailHandler())
	app.Post("/v1/file", createUploadFileHandler())

	return app
}
