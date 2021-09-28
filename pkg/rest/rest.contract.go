package rest

import "github.com/gofiber/fiber/v2"

type Result = error
type Context = *fiber.Ctx
type Handler = func(ctx Context) Result
type ErrorHandler = func(ctx Context, err error) Result
