package storage

import (
	"idaman.id/storage/internal/config"
	app_error "idaman.id/storage/internal/error"
	"idaman.id/storage/internal/file"
)

func NewStorage(provider string, configGetter config.Getter, fileService file.FileService) (Storage, error) {

	if provider != "local" {
		err := app_error.NewNotSupportedError("Storage")
		return nil, err
	}

	storage := NewStorageLocal(configGetter, fileService)

	return storage, nil
}
