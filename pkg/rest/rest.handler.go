package rest

import (
	"github.com/gofiber/fiber/v2"
	"idaman.id/storage/pkg/retrieving"
	"idaman.id/storage/pkg/translation"
	"idaman.id/storage/pkg/uploading"
	"idaman.id/storage/pkg/validation"
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

		identifier := ctx.Params("identifier")
		fileEntity, err := retrieving.GetFile(identifier)
		isFileAvailable := err == nil
		if isFileAvailable {
			response := createSuccessResponse(ResponseDto{
				Translator: translator,
				Data:       fileEntity,
			})
			return ctx.JSON(response)
		}

		response := createFailedResponse(ResponseDto{
			Message:    err.Error(),
			Translator: translator,
		})
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
}

func createGetResourceHandler(dependency *Dependency) Handler {
	return func(ctx Context) Result {
		localizer := dependency.getLocalizer(ctx)
		translator := translation.CreateSimpleTranslator(localizer)

		identifier := ctx.Params("identifier")
		result, err := retrieving.RetrieveFile(identifier)
		isFileAvailable := err == nil

		if isFileAvailable {
			ctx.Set("Content-Type", result.File.Mimetype)
			return ctx.Send(result.FileData)
		}

		var response ResponseEntity
		var statusCode int

		switch err.(type) {
		case *retrieving.FileNotFoundError:
			statusCode = fiber.StatusNotFound
			response = createFailedResponse(ResponseDto{
				Message:    err.Error(),
				Translator: translator,
			})
		default:
			statusCode = fiber.StatusBadRequest
			response = createFailedResponse(ResponseDto{
				Message:    err.Error(),
				Translator: translator,
			})
		}

		return ctx.Status(statusCode).JSON(response)
	}
}

func createUploadFileHandler(dependency *Dependency) Handler {
	return func(ctx Context) Result {
		locale := dependency.getLocale(ctx)
		localizer := dependency.getLocalizer(ctx)
		translator := translation.CreateSimpleTranslator(localizer)

		form, err := ctx.MultipartForm()

		isFormInvalid := err != nil
		if isFormInvalid {
			response := createFailedResponse(ResponseDto{
				Message:    err.Error(),
				Translator: translator,
			})
			return ctx.Status(fiber.StatusBadRequest).JSON(response)
		}

		uploadData := uploading.UploadFileDto{
			Files:    form.File["files"],
			Provider: ctx.FormValue("provider"),
			Locale:   locale,
		}
		result, err := uploading.UploadFile(uploadData)

		isUploadSuccess := err == nil
		if isUploadSuccess {
			response := createSuccessResponse(ResponseDto{
				Data:       result.Items,
				Translator: translator,
			})
			return ctx.JSON(response)
		}

		var response ResponseEntity
		var status int

		switch err.(type) {
		case *validation.ValidationError:
			status = fiber.StatusUnprocessableEntity
			validationError := err.(*validation.ValidationError)
			response = createFailedResponse(ResponseDto{
				Message:    validationError.Error(),
				Error:      validationError.Items,
				Translator: translator,
			})
		default:
			status = fiber.StatusBadRequest
			response = createFailedResponse(ResponseDto{
				Message:    err.Error(),
				Translator: translator,
			})
		}

		return ctx.Status(status).JSON(response)
	}
}
