package rest_fiber

import (
	"github.com/gofiber/fiber/v2"
	"idaman.id/storage/internal/deleting"
	response "idaman.id/storage/internal/rest-response"
	"idaman.id/storage/internal/retrieving"
	"idaman.id/storage/internal/uploading"
	app_error "idaman.id/storage/pkg/error"
)

func NewFileGetDetailHandler(rService retrieving.FileGetter) Handler {
	return func(ctx *Context) error {

		fileEntity, err := rService.GetFile(ctx.Params("identifier"))
		isFileAvailable := err == nil
		if isFileAvailable {
			responseEntity := response.NewSuccessResponse(&response.ResponseParam{
				Data: fileEntity,
			})
			return ctx.JSON(responseEntity)
		}

		var statusCode int
		var responseEntity *response.ResponseEntity

		switch err.(type) {
		case *app_error.NotFoundError:
			notFoundError := err.(*app_error.NotFoundError)
			statusCode = fiber.StatusNotFound
			responseEntity = response.NewErrorResponse(&response.ResponseParam{
				Message: notFoundError.Error(),
			})
		default:
			statusCode = fiber.StatusBadRequest
			responseEntity = response.NewErrorResponse(&response.ResponseParam{
				Message: err.Error(),
			})
		}

		return ctx.Status(statusCode).JSON(responseEntity)
	}
}

func NewDeleteFileHandler(dService deleting.DeleteService) Handler {
	return func(ctx *Context) error {
		err := dService.DeleteFile(ctx.Params("identifier"))
		isSuccessDelete := err == nil
		if isSuccessDelete {
			responseEntity := response.NewSuccessResponse(nil)
			return ctx.JSON(responseEntity)
		}

		var statusCode int
		var responseEntity *response.ResponseEntity

		switch err.(type) {
		case *app_error.NotFoundError:
			notFoundError := err.(*app_error.NotFoundError)
			statusCode = fiber.StatusNotFound
			responseEntity = response.NewErrorResponse(&response.ResponseParam{
				Message: notFoundError.Error(),
			})
		default:
			statusCode = fiber.StatusBadRequest
			responseEntity = response.NewErrorResponse(&response.ResponseParam{
				Message: err.Error(),
			})
		}

		return ctx.Status(statusCode).JSON(responseEntity)
	}
}

func NewGetResourceHandler(rService retrieving.FileRetriever) Handler {
	return func(ctx *Context) error {
		result, err := rService.RetrieveFile(ctx.Params("identifier"))
		isFileAvailable := err == nil

		if isFileAvailable {
			ctx.Set("Content-Type", result.File.Mimetype)
			return ctx.Send(result.FileData)
		}

		var responseEntity *response.ResponseEntity
		var statusCode int

		switch err.(type) {
		case *app_error.NotFoundError:
			notFoundError := err.(*app_error.NotFoundError)
			statusCode = fiber.StatusNotFound
			responseEntity = response.NewErrorResponse(&response.ResponseParam{
				Message: notFoundError.Error(),
			})
		default:
			statusCode = fiber.StatusBadRequest
			responseEntity = response.NewErrorResponse(&response.ResponseParam{
				Message: err.Error(),
			})
		}

		return ctx.Status(statusCode).JSON(responseEntity)
	}
}

func NewUploadFileHandler(uService uploading.UploadService) Handler {
	return func(ctx *Context) error {

		form, err := ctx.MultipartForm()

		isFormInvalid := err != nil
		if isFormInvalid {
			responseEntity := response.NewErrorResponse(&response.ResponseParam{
				Message: err.Error(),
			})
			return ctx.Status(fiber.StatusBadRequest).JSON(responseEntity)
		}

		uploadData := uploading.UploadFileParam{
			Files:    form.File["files"],
			Provider: ctx.FormValue("provider"),
		}
		result, err := uService.UploadFile(uploadData)

		isUploadSuccess := err == nil
		if isUploadSuccess {
			responseEntity := response.NewSuccessResponse(&response.ResponseParam{
				Data: result.Items,
			})
			return ctx.JSON(responseEntity)
		}

		var responseEntity *response.ResponseEntity
		var status int

		switch err.(type) {
		case *app_error.ValidationError:
			status = fiber.StatusUnprocessableEntity
			validationError := err.(*app_error.ValidationError)
			responseEntity = response.NewErrorResponse(&response.ResponseParam{
				Message: validationError.Error(),
				Error:   validationError.Items,
			})
		default:
			status = fiber.StatusBadRequest
			responseEntity = response.NewErrorResponse(&response.ResponseParam{
				Message: err.Error(),
			})
		}

		return ctx.Status(status).JSON(responseEntity)
	}
}
