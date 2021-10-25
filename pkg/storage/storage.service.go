package storage

import "idaman.id/storage/pkg/app"

func NewStorage(provider string) (Storage, error) {

	if provider != "local" {
		return nil, &app.NotSupportedError{
			Message: app.STATUS_NOT_SUPPORTED,
			Context: "Storage",
		}
	}

	storage := NewStorageLocal()

	return storage, nil
}
