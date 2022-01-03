package repository

import (
	app_error "idaman.id/storage/internal/error"
	mongo "idaman.id/storage/internal/repository-mongo"
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
