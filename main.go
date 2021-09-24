package main

import (
	App "idaman.id/storage/src/app"
)

func main() {
	app := App.CreateApp()
	App.CreateRouter(app)
	app.Listen(":3000")
}
