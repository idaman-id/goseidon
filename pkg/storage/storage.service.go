package storage

import "idaman.id/storage/pkg/app"

func CreateStorage(provider string) (Storage, error) {

	if provider != "local" {
		return nil, &app.NotSupportedError{
			Message: app.STATUS_NOT_SUPPORTED,
			Context: "Storage",
		}
	}

	storage := CreateStorageLocal()

	return storage, nil
}
