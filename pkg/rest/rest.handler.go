package rest

import (
	"github.com/gofiber/fiber/v2"
	"idaman.id/storage/pkg/storage"
	"idaman.id/storage/pkg/uploading"
)

func createHomeHandler() Handler {
	return func(ctx Context) Result {
		response := createSuccessResponse(ResponseDto{})
		return ctx.JSON(response)
	}
}

func createGetDetailHandler() Handler {
	return func(ctx Context) Result {
		response := createSuccessResponse(ResponseDto{})
		return ctx.JSON(response)
	}
}

func createGetResourceHandler() Handler {
	return func(ctx Context) Result {
		response := createSuccessResponse(ResponseDto{})
		return ctx.JSON(response)
	}
}

func createUploadFileHandler() Handler {
	return func(ctx Context) Result {
		/**
		@todo
		1. validation
		*/
		form, err := ctx.MultipartForm()

		if err != nil {
			response := createFailedResponse(ResponseDto{
				Message: err.Error(),
			})
			return ctx.Status(fiber.StatusBadRequest).JSON(response)
		}

		// Provider: ctx.FormValue("provider"),
		storageProvider := &storage.StorageLocal{
			StorageDir: "storage",
		}

		data := uploading.UploadFileDto{
			Files:   form.File["files"],
			Storage: storageProvider,
		}
		result := uploading.UploadFile(data)

		response := createSuccessResponse(ResponseDto{
			Data: result.Items,
		})
		return ctx.JSON(response)
	}
}
