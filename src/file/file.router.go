package file

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {

	app.Get("/file/:id", GetResourceHandler)

	v1 := app.Group("v1")
	v1.Post("/file", UploadFileHandler)
	v1.Get("/file/:id", GetDetailHandler)

}
