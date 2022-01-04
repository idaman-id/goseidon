package builtin_app

import "github.com/gofiber/fiber/v2"

type App = fiber.App
type Context = fiber.Ctx
type Handler = func(*Context) error
type ErrorHandler = func(*Context, error) error
