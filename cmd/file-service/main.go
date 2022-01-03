package main

import (
	fiber "idaman.id/storage/internal/rest-fiber"
)

func main() {
	app, err := fiber.NewApp()
	if err != nil {
		panic(err.Error())
	}

	err = app.Run()
	if err != nil {
		panic(err.Error())
	}
}
