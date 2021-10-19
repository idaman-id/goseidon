package rest

import (
	"idaman.id/storage/pkg/translation"
)

func createHomeHandler(dependency *Dependency) Handler {
	return func(ctx Context) Result {
		localizer := dependency.getLocalizer(ctx)
		translator := translation.NewGoI18nService(localizer).Translate

		response := createSuccessResponse(ResponseDto{
			Translator: translator,
		})
		return ctx.JSON(response)
	}
}
