package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Router(App *fiber.App) {
	App.Use(recover.New())
	App.Get("/", HomeHandler)
}
