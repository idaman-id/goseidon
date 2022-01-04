package main

import (
	builtin_app "idaman.id/storage/internal/builtin-app"
)

func main() {
	app, err := builtin_app.NewApp()
	if err != nil {
		panic(err.Error())
	}

	err = app.Run()
	if err != nil {
		panic(err.Error())
	}
}
