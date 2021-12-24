package storage

import (
	"idaman.id/storage/internal/file"
	"idaman.id/storage/pkg/config"
	app_error "idaman.id/storage/pkg/error"
)

func NewStorage(provider string, configGetter config.Getter, fileService file.FileService) (Storage, error) {

	if provider != "local" {
		err := app_error.NewNotSupportedError("Storage")
		return nil, err
	}

	storage := NewStorageLocal(configGetter, fileService)

	return storage, nil
}
