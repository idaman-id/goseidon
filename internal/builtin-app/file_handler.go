package builtin_app

import (
	"github.com/gofiber/fiber/v2"
	"idaman.id/storage/internal/deleting"
	app_error "idaman.id/storage/internal/error"
	"idaman.id/storage/internal/file"
	response "idaman.id/storage/internal/response"
	"idaman.id/storage/internal/retrieving"
	"idaman.id/storage/internal/uploading"
)

func NewFileGetDetailHandler(rService retrieving.FileGetter) Handler {
	return func(ctx *Context) error {
		fileDetail, err := rService.GetFile(ctx.Params("identifier"))
		if err != nil {
			var statusCode int
			var resBody *response.ResponseEntity

			switch err.(type) {
			case *app_error.NotfoundError:
				notFoundError := err.(*app_error.NotfoundError)
				statusCode = fiber.StatusNotFound
				resBody = response.NewErrorResponse(&response.ResponseParam{
					Message: notFoundError.Error(),
				})
			default:
				statusCode = fiber.StatusBadRequest
				resBody = response.NewErrorResponse(&response.ResponseParam{
					Message: err.Error(),
				})
			}

			return ctx.Status(statusCode).JSON(resBody)
		}

		fileEntity := &FileDetailEntity{
			UniqueId:  fileDetail.UniqueId,
			Name:      fileDetail.Name,
			Extension: fileDetail.Extension,
			Size:      fileDetail.Size,
			Mimetype:  fileDetail.Mimetype,
			Url:       fileDetail.Url,
			CreatedAt: fileDetail.CreatedAt,
			UpdatedAt: fileDetail.UpdatedAt,
		}
		resBody := response.NewSuccessResponse(&response.ResponseParam{
			Data: fileEntity,
		})
		return ctx.JSON(resBody)
	}
}

func NewGetResourceHandler(rService retrieving.FileRetriever) Handler {
	return func(ctx *Context) error {
		result, err := rService.RetrieveFile(ctx.Params("identifier"))

		if err != nil {
			var responseEntity *response.ResponseEntity
			var statusCode int

			switch err.(type) {
			case *app_error.NotfoundError:
				notFoundError := err.(*app_error.NotfoundError)
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

		ctx.Set("Content-Type", result.File.Mimetype)
		return ctx.Send(result.FileData)
	}
}

func NewUploadFileHandler(uService uploading.UploadService, fService file.FileService) Handler {
	return func(ctx *Context) error {

		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			err = app_error.NewNotfoundError("File")
			responseEntity := response.NewErrorResponse(&response.ResponseParam{
				Message: err.Error(),
			})
			return ctx.Status(fiber.StatusBadRequest).JSON(responseEntity)
		}

		fileEntity, err := file.NewFileFromMultipartHeader(fileHeader, fService)

		if err != nil {
			err = app_error.NewNotfoundError("File")
			responseEntity := response.NewErrorResponse(&response.ResponseParam{
				Message: err.Error(),
			})
			return ctx.Status(fiber.StatusBadRequest).JSON(responseEntity)
		}

		fileDetail, err := uService.UploadFile(uploading.UploadFileParam{
			File: fileEntity,
		})

		if err != nil {
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

		res := &FileDetailEntity{
			UniqueId:  fileDetail.UniqueId,
			Name:      fileDetail.Name,
			Extension: fileDetail.Extension,
			Size:      fileDetail.Size,
			Mimetype:  fileDetail.Mimetype,
			Url:       fileDetail.Url,
			CreatedAt: fileDetail.CreatedAt,
			UpdatedAt: fileDetail.UpdatedAt,
		}

		responseEntity := response.NewSuccessResponse(&response.ResponseParam{
			Data: res,
		})
		return ctx.JSON(responseEntity)
	}
}

func DeleteFileHandler(rService deleting.DeleteService) Handler {
	return func(ctx *Context) error {
		err := rService.DeleteFile(ctx.Params("identifier"))

		if err != nil {
			var statusCode int
			var resBody *response.ResponseEntity

			switch err.(type) {
			case *app_error.NotfoundError:
				notFoundError := err.(*app_error.NotfoundError)
				statusCode = fiber.StatusNotFound
				resBody = response.NewErrorResponse(&response.ResponseParam{
					Message: notFoundError.Error(),
				})
			default:
				statusCode = fiber.StatusBadRequest
				resBody = response.NewErrorResponse(&response.ResponseParam{
					Message: err.Error(),
				})
			}
			return ctx.Status(statusCode).JSON(resBody)
		}

		resBody := response.NewSuccessResponse(nil)
		return ctx.JSON(resBody)
	}
}
