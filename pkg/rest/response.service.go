package rest

import (
	"idaman.id/storage/pkg/app"
)

func createSuccessResponse(dto ResponseDto) ResponseEntity {

	response := ResponseEntity{
		Message: app.STATUS_OK,
	}

	if dto.Data != nil {
		response.Data = dto.Data
	}

	if dto.Message != "" {
		response.Message = dto.Message
	}

	translation := dto.translator(response.Message)
	response.Message = translation

	return response
}

func createFailedResponse(dto ResponseDto) ResponseEntity {

	response := ResponseEntity{
		Message: app.STATUS_ERROR,
	}

	if dto.Error != nil {
		response.Error = dto.Error
	}

	if dto.Message != "" {
		response.Message = dto.Message
	}

	translation := dto.translator(response.Message)
	response.Message = translation

	return response
}
