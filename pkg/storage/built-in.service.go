package storage

type StorageBuiltIn struct {
}

func (storage *StorageBuiltIn) SaveFile(fileHeader FileDto) (result *SaveFileResult, err error) {

	return &SaveFileResult{}, nil
}
