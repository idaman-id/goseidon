package rest_response

const (
	STATUS_OK    = "OK"
	STATUS_ERROR = "ERROR"
)

type ResponseParam struct {
	Message string
	Data    interface{}
	Error   interface{}
}
