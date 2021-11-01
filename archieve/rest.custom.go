package archieve

// type Dependency struct {
// 	getLocalizer Localizer
// 	getLocale    LocaleParser
// }

// type LocaleParser = func(ctx Context) string
// type Localizer = func(ctx Context) *i18n.Localizer

// locale := dependency.getLocale(ctx)
// localizer := dependency.getLocalizer(ctx)
// translator := translation.NewGoI18nService(localizer).Translate

// Translator: translator,
// TranslationData: map[string]interface{}{
// 	"context": notFoundError.Context,
// },

// func createLimiterConfig() func() limiter.Config {
// 	return func() limiter.Config {
// 		config := limiter.Config{
// 			Max:        20,
// 			Expiration: 30 * time.Second,
// 			KeyGenerator: func(ctx *fiber.Ctx) string {
// 				return ctx.Get("x-forwarded-for")
// 			},
// 			LimitReached: func(ctx *fiber.Ctx) error {
// 				responseEntity := response.NewErrorResponse(&response.ResponseParam{
// 					Message: app.STATUS_TOO_MANY_REQUEST,
// 				})
// 				return ctx.Status(fiber.StatusTooManyRequests).JSON(responseEntity)
// 			},
// 		}
// 		return config
// 	}
// }

// createConfig := createLimiterConfig()
// app.Use(limiter.New(createConfig()))

// func localeParser(ctx handler.Context) string {
// 	locale := ctx.Query("lang")
// 	if locale == "" {
// 		defaultLocale := config.Service.GetString("APP_DEFAULT_LOCALE")
// 		locale = ctx.Get("Accept-Language", defaultLocale, "en")
// 	}
// 	return locale
// }

// func createLocalizer(i18nBundle *i18n.Bundle) func(ctx handler.Context) *i18n.Localizer {
// 	return func(ctx handler.Context) *i18n.Localizer {
// 		locale := localeParser(ctx)
// 		localizer := i18n.NewLocalizer(i18nBundle, locale)
// 		return localizer
// 	}
// }
