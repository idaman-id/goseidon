package builtin_app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"idaman.id/storage/internal/app"
	"idaman.id/storage/internal/config"
	"idaman.id/storage/internal/database"
	"idaman.id/storage/internal/file"
	repository_mysql "idaman.id/storage/internal/repository-mysql"
	"idaman.id/storage/internal/retrieving"
	storage_local "idaman.id/storage/internal/storage-local"
	"idaman.id/storage/internal/text"
	"idaman.id/storage/internal/uploading"
	"idaman.id/storage/internal/validation"
)

func NewApp() (app.App, error) {
	configService, err := config.NewConfigService()
	if err != nil {
		return nil, err
	}
	validatorService, err := validation.NewValidator(configService)
	if err != nil {
		return nil, err
	}

	textService := text.NewTextService()
	fileService := file.NewFileService(textService)

	mysqlClient, err := database.NewMySQLClient(configService)
	if err != nil {
		return nil, err
	}
	fileRepo := repository_mysql.NewFileRepository(mysqlClient, fileService)

	localStorage := storage_local.NewStorageLocal("storage/file")

	retrieveService := retrieving.NewRetrieveService(fileRepo, configService, fileService, localStorage)
	uploadService := uploading.NewUploadService(validatorService, configService, localStorage, textService, fileRepo)

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
	app.Post("/v1/file", NewUploadFileHandler(uploadService, fileService))
	app.Get("/v1/file/:identifier", NewFileGetDetailHandler(retrieveService))
	app.Get("*", NewNotFoundHandler())

	fiberApp := &FiberApp{
		fiber:        app,
		configGetter: configService,
	}
	return fiberApp, nil
}
