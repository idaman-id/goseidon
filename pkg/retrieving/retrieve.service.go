package retrieving

import (
	"idaman.id/storage/pkg/file"
	"idaman.id/storage/pkg/storage"
)

func GetFile(identifier string) (*FileEntity, error) {

	/*
		@todo
		1. check file existance
		2. if not found return NotFoundError
	*/

	uniqueId := file.RemoveFileExtension(identifier)
	return &FileEntity{
		UniqueId:      uniqueId,
		OriginalName:  "hio.jpeg",
		Name:          "hio",
		Extension:     "jpeg",
		Size:          6720,
		Mimetype:      "image/jpeg",
		Url:           "http://storage.domain.tld/storage/file/4980a441-b747-4226-ada0-63a5b2cac26d.jpeg",
		Path:          "",
		ProviderId:    "",
		ApplicationId: "",
	}, nil
}

func RetrieveFile(identifier string) (*RetrieveFileResult, error) {
	/*
		@todo
		1. check file existance
		2. if not found return NotFoundError
		3.set fileEntity using data from repo
	*/
	uniqueId := file.RemoveFileExtension(identifier)

	provider := "local"

	storageService, err := storage.CreateStorage(provider)
	isStorageUnsupported := err != nil
	if isStorageUnsupported {
		return nil, err
	}

	storageFile := storage.FileEntity{
		UniqueId:  uniqueId,
		Extension: "jpeg",
	}

	fileData, err := storageService.RetrieveFile(&storageFile)
	isFileAvailable := err == nil
	if !isFileAvailable {
		return nil, err
	}

	file := FileEntity{
		Mimetype: "image/jpeg",
	}
	result := &RetrieveFileResult{
		FileData: fileData,
		File:     file,
	}

	return result, nil
}
