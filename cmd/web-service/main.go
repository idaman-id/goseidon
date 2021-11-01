package main

import (
	"idaman.id/storage/internal/bootstraping"
	rest "idaman.id/storage/internal/rest-fiber"
	"idaman.id/storage/pkg/config"
)

func main() {
	err := bootstraping.Setup()
	if err != nil {
		panic(err.Error())
	}

	app := rest.NewApp()
	rest.RegisterRoute(app)

	address := config.Service.GetString("APP_HOST") + ":" + config.Service.GetString("APP_PORT")
	app.Listen(address)
}
