package app

type Result = error

type App interface {
	Listen(address string) Result
}
