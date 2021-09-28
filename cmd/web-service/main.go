package main

import (
	"idaman.id/storage/pkg/rest"
)

func main() {
	app := rest.CreateApp()
	app.Listen(":3000")
}
