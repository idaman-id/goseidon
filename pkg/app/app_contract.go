package app

const (
	STATUS_OK               = "OK"
	STATUS_ERROR            = "ERROR"
	STATUS_INVALID_DATA     = "INVALID_DATA"
	STATUS_TOO_MANY_REQUEST = "TOO_MANY_REQUEST"
	STATUS_NOT_FOUND        = "NOT_FOUND"
	STATUS_NOT_SUPPORTED    = "NOT_SUPPORTED"
)

type App interface {
	Run() error
}
