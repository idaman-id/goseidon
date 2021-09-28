package rest

type ResponseDto struct {
	Message string
	Data    interface{}
	Error   interface{}
}
