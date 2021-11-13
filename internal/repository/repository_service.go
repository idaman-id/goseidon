package repository

import (
	mongo "idaman.id/storage/internal/repository-mongo"
	"idaman.id/storage/pkg/app"
)

func NewRepository(provider string) (*Repository, error) {
	if provider != DATABASE_MONGO {
		err := app.NewNotSupportedError("Database")
		return nil, err
	}

	repo := &Repository{
		FileRepo: &mongo.FileRepository{},
	}

	return repo, nil
}
