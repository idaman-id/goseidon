package app

type Result = error

type App interface {
	Listen(address string) Result
}

const (
	STATUS_OK               = "OK"
	STATUS_ERROR            = "ERROR"
	STATUS_INVALID_DATA     = "INVALID_DATA"
	STATUS_TOO_MANY_REQUEST = "TOO_MANY_REQUEST"
)
