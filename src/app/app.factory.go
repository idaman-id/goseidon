package app

import (
	"github.com/gofiber/fiber/v2"
	"idaman.id/storage/src/file"
)

func CreateApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})
	return app
}

func CreateRouter(App *fiber.App) {
	Router(App)
	file.Router(App)
}
