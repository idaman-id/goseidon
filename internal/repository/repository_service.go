package repository

import (
	mongo "idaman.id/storage/internal/repository-mongo"
	app_error "idaman.id/storage/pkg/error"
)

func NewRepository(provider string) (*Repository, error) {
	if provider != DATABASE_MONGO {
		err := app_error.NewNotSupportedError("Database")
		return nil, err
	}

	repo := &Repository{
		FileRepo: &mongo.FileRepository{},
	}

	return repo, nil
}
