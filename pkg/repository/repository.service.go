package repository

import (
	"idaman.id/storage/pkg/app"
	mongo "idaman.id/storage/pkg/repository/mongo"
)

var (
	FileRepo FileRepository
)

func Init(provider string) error {
	if provider != app.DATABASE_MONGO {
		err := app.NewNotSupportedError("Database")
		return err
	}

	FileRepo = &mongo.FileRepository{}
	return nil
}
