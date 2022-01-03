package rest_fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"idaman.id/storage/internal/app"
	"idaman.id/storage/internal/config"
	"idaman.id/storage/internal/file"
	"idaman.id/storage/internal/repository"
	"idaman.id/storage/internal/retrieving"
	"idaman.id/storage/internal/text"
	"idaman.id/storage/internal/uploading"
	"idaman.id/storage/internal/validation"
)

func NewApp() (app.App, error) {
	configService, err := config.NewConfig(config.CONFIG_VIPER)
	if err != nil {
		return nil, err
	}
	err = config.InitConfig(configService)
	if err != nil {
		return nil, err
	}

	validator := validation.NewValidator(validation.VALIDATOR_GO_I18N)

	repo, err := repository.NewRepository(repository.DATABASE_MONGO)
	if err != nil {
		return nil, err
	}

	textService := text.NewTextService()
	fileService := file.NewFileService(textService)
	retrieveService := retrieving.NewRetrieveService(repo.FileRepo, fileService)
	uploadService := uploading.NewUploadService(validator, configService, fileService)

	app := fiber.New(fiber.Config{
		ErrorHandler: NewErrorHandler(),
	})
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(etag.New())
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", NewHomeHandler())
	app.Get("/file/:identifier", NewGetResourceHandler(retrieveService))
	app.Post("/v1/file", NewUploadFileHandler(uploadService))
	app.Get("/v1/file/:identifier", NewFileGetDetailHandler(retrieveService))
	app.Get("*", NewNotFoundHandler())

	fiberApp := &FiberApp{
		fiber:        app,
		configGetter: configService,
	}
	return fiberApp, nil
}
