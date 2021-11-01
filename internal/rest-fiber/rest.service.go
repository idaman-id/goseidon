package rest_fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"idaman.id/storage/internal/retrieving"
)

func NewApp() *App {
	app := fiber.New(fiber.Config{
		ErrorHandler: NewErrorHandler(),
	})

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(etag.New())
	app.Use(cors.New())
	app.Use(logger.New())

	return app
}

func RegisterRoute(app *App) {
	app.Get("/", NewHomeHandler())
	// app.Get("/file/:identifier", NewGetResourceHandler())
	// app.Post("/v1/file", NewUploadFileHandler())
	app.Get("/v1/file/:identifier", NewFileGetDetailHandler(retrieving.Service))
	// app.Delete("/v1/file/:identifier", NewDeleteFileHandler())
	app.Get("*", NewNotFoundHandler())
}
