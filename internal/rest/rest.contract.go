package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Result = error
type Context = *fiber.Ctx
type Handler = func(ctx Context) Result
type ErrorHandler = func(ctx Context, err error) Result

type LocaleParser = func(ctx Context) string
type Localizer = func(ctx Context) *i18n.Localizer

type App interface {
	Listen(address string) Result
}

type Dependency struct {
	getLocalizer Localizer
	getLocale    LocaleParser
}
