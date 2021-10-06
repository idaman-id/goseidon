package storage

import "errors"

func CreateStorage(provider string) (Storage, error) {

	if provider != "local" {
		return nil, errors.New("unsupported storage provider")
	}

	storage := CreateStorageLocal()

	return storage, nil
}
