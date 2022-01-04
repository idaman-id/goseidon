package builtin_app

import (
	"github.com/gofiber/fiber/v2"
	"idaman.id/storage/internal/config"
)

type FiberApp struct {
	fiber        *fiber.App
	configGetter config.Getter
}

func (app *FiberApp) Run() error {
	addr := app.configGetter.GetString("APP_HOST") + ":" + app.configGetter.GetString("APP_PORT")
	return app.fiber.Listen(addr)
}
