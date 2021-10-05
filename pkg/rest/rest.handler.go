package rest

import (
	"github.com/gofiber/fiber/v2"
	"idaman.id/storage/pkg/storage"
	"idaman.id/storage/pkg/translation"
	"idaman.id/storage/pkg/uploading"
)

func createHomeHandler(dependency *Dependency) Handler {
	return func(ctx Context) Result {
		localizer := dependency.getLocalizer(ctx)
		translator := translation.CreateSimpleTranslator(localizer)
		response := createSuccessResponse(ResponseDto{
			Translator: translator,
		})
		return ctx.JSON(response)
	}
}

func createGetDetailHandler(dependency *Dependency) Handler {
	return func(ctx Context) Result {
		localizer := dependency.getLocalizer(ctx)
		translator := translation.CreateSimpleTranslator(localizer)
		response := createSuccessResponse(ResponseDto{
			Translator: translator,
		})
		return ctx.JSON(response)
	}
}

func createGetResourceHandler(dependency *Dependency) Handler {
	return func(ctx Context) Result {
		localizer := dependency.getLocalizer(ctx)
		translator := translation.CreateSimpleTranslator(localizer)
		response := createSuccessResponse(ResponseDto{
			Translator: translator,
		})
		return ctx.JSON(response)
	}
}

func createUploadFileHandler(dependency *Dependency) Handler {
	return func(ctx Context) Result {
		// locale := dependency.getLocale(ctx)
		localizer := dependency.getLocalizer(ctx)
		translator := translation.CreateSimpleTranslator(localizer)

		form, err := ctx.MultipartForm()

		if err != nil {
			response := createFailedResponse(ResponseDto{
				Message:    err.Error(),
				Translator: translator,
			})
			return ctx.Status(fiber.StatusBadRequest).JSON(response)
		}

		// type UploadFileRule struct {
		// 	Provider string `json:"provider" validate:"required"`
		// }
		// rule := &UploadFileRule{
		// 	Example: "1",
		// }

		// validationError := validation.ValidateRule(locale, rule)

		// if validationError != nil {
		// 	response := createFailedResponse(ResponseDto{
		// 		Message:    validationError.Error(),
		// 		Error:      validationError.Items,
		// 		Translator: translator,
		// 	})
		// 	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response)
		// }

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
			Translator: translator,
		})
		return ctx.JSON(response)
	}
}
