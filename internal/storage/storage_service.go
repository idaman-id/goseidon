package storage

import (
	"idaman.id/storage/internal/file"
	"idaman.id/storage/pkg/app"
	"idaman.id/storage/pkg/config"
)

func NewStorage(provider string, configGetter config.Getter, fileService file.FileService) (Storage, error) {

	if provider != "local" {
		err := app.NewNotSupportedError("Storage")
		return nil, err
	}

	storage := NewStorageLocal(configGetter, fileService)

	return storage, nil
}
