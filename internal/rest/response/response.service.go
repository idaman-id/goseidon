package rest_response

import (
	"strings"

	"idaman.id/storage/pkg/app"
)

func NewSuccessResponse(param *ResponseParam) *ResponseEntity {

	response := ResponseEntity{
		Message: app.STATUS_OK,
	}

	if param == nil {
		return &response
	}

	isEmptyString := strings.TrimSpace(param.Message) == ""
	if !isEmptyString {
		response.Message = param.Message
	}

	if param.Data != nil {
		response.Data = param.Data
	}

	return &response
}

func NewErrorResponse(param *ResponseParam) *ResponseEntity {

	response := ResponseEntity{
		Message: app.STATUS_ERROR,
	}

	if param == nil {
		return &response
	}

	isEmptyString := strings.TrimSpace(param.Message) == ""
	if !isEmptyString {
		response.Message = param.Message
	}

	if param.Error != nil {
		response.Error = param.Error
	}

	return &response
}
