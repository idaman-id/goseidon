package main

import (
	rest "idaman.id/storage/internal/rest-fiber"
)

func main() {
	app, err := rest.NewApp()
	if err != nil {
		panic(err.Error())
	}

	err = app.Run()
	if err != nil {
		panic(err.Error())
	}
}
