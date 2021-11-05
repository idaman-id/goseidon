package repository

import (
	mongo "idaman.id/storage/internal/repository/mongo"
	"idaman.id/storage/pkg/app"
)

var (
	FileRepo FileRepository
)

func Init(provider string) error {
	if provider != DATABASE_MONGO {
		err := app.NewNotSupportedError("Database")
		return err
	}

	FileRepo = &mongo.FileRepository{}
	return nil
}
