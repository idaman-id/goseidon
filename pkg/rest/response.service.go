package rest

func createSuccessResponse(dto ResponseDto) ResponseEntity {

	response := ResponseEntity{
		Message: "ok",
	}

	if dto.Data != nil {
		response.Data = dto.Data
	}

	if dto.Message != "" {
		response.Message = dto.Message
	}

	return response
}

func createFailedResponse(dto ResponseDto) ResponseEntity {

	response := ResponseEntity{
		Message: "error occured",
	}

	if dto.Error != nil {
		response.Error = dto.Error
	}

	if dto.Message != "" {
		response.Message = dto.Message
	}

	return response
}
