package main

import (
	"idaman.id/storage/internal/rest"
	"idaman.id/storage/pkg/config"
)

func main() {
	app := rest.CreateApp()
	address := config.Service.GetString("APP_HOST") + ":" + config.Service.GetString("APP_PORT")
	app.Listen(address)
}
