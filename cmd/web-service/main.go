package main

import (
	"idaman.id/storage/pkg/config"
	"idaman.id/storage/pkg/rest"
)

func main() {
	app := rest.CreateApp()
	address := config.GetString("APP_HOST") + ":" + config.GetString("APP_PORT")
	app.Listen(address)
}
