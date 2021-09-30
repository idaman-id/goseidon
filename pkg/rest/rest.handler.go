package rest

import (
	"github.com/gofiber/fiber/v2"
	"idaman.id/storage/pkg/storage"
	"idaman.id/storage/pkg/translation"
	"idaman.id/storage/pkg/uploading"
	"idaman.id/storage/pkg/validation"
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

func createUploadFileHandler(dependency *Dependency) Handler {
	return func(ctx Context) Result {

		form, err := ctx.MultipartForm()

		type UploadFileParam struct {
			Example string `json:"example" validate:"required,email,min=3"`
		}
		rule := &UploadFileParam{
			Example: "1",
		}

		locale := dependency.localeParser(ctx)
		localizer := dependency.localizer(ctx)
		translator := translation.CreateSimpleTranslator(localizer)

		validationError := validation.ValidateRule(locale, rule)

		if validationError != nil {
			response := createFailedResponse(ResponseDto{
				Message:    validationError.Error(),
				Error:      validationError.Items,
				translator: translator,
			})
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response)
		}

		if err != nil {
			response := createFailedResponse(ResponseDto{
				Message:    err.Error(),
				translator: translator,
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
			Data:       result.Items,
			translator: translator,
		})
		return ctx.JSON(response)
	}
}
