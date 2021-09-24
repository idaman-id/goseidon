package http

type ResponseParameter struct {
	Message string
	Data    interface{}
	Error   interface{}
}

func CreateSuccessResponse(param ResponseParameter) Response {

	response := Response{
		Message: "ok",
	}

	if param.Data != nil {
		response.Data = param.Data
	}

	if param.Message != "" {
		response.Message = param.Message
	}

	return response
}

func CreateFailedResponse(param ResponseParameter) Response {

	response := Response{
		Message: "error occured",
	}

	if param.Error != nil {
		response.Error = param.Error
	}

	if param.Message != "" {
		response.Message = param.Message
	}

	return response
}
