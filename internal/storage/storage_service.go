package storage

import "idaman.id/storage/pkg/app"

func NewStorage(provider string) (Storage, error) {

	if provider != "local" {
		err := app.NewNotSupportedError("Storage")
		return nil, err
	}

	storage := NewStorageLocal()

	return storage, nil
}
