package file

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"idaman.id/storage/src/http"
)

func GetDetailHandler(ctx *fiber.Ctx) error {
	file1 := File{
		UUID: uuid.NewString(),
	}
	response := http.CreateSuccessResponse(http.ResponseParameter{
		Data: file1,
	})
	return ctx.JSON(response)
}

func GetResourceHandler(ctx *fiber.Ctx) error {

	response := http.CreateSuccessResponse(http.ResponseParameter{})
	return ctx.JSON(response)
}

func UploadFileHandler(ctx *fiber.Ctx) error {

	response := http.CreateSuccessResponse(http.ResponseParameter{})
	return ctx.JSON(response)
}
