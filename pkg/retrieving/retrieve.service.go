package retrieving

import (
	"idaman.id/storage/pkg/file"
	"idaman.id/storage/pkg/storage"
)

func GetFile(identifier string) (*FileEntity, error) {

	uniqueId := file.RemoveFileExtension(identifier)
	file, err := file.FindByUniqueId(uniqueId)
	isRecordAvailable := err == nil
	if !isRecordAvailable {
		return nil, err
	}

	result := &FileEntity{
		UniqueId:      file.UniqueId,
		OriginalName:  file.OriginalName,
		Name:          file.Name,
		Extension:     file.Extension,
		Size:          file.Size,
		Mimetype:      file.Mimetype,
		Url:           file.Url,
		Path:          file.Path,
		ProviderId:    file.ProviderId,
		ApplicationId: file.ApplicationId,
		CreatedAt:     file.CreatedAt,
		UpdatedAt:     file.UpdatedAt,
	}
	return result, nil
}

func RetrieveFile(identifier string) (*RetrieveFileResult, error) {

	uniqueId := file.RemoveFileExtension(identifier)
	file, err := file.FindByUniqueId(uniqueId)
	isRecordAvailable := err == nil
	if !isRecordAvailable {
		return nil, err
	}

	/*
		@todo
		1. set provider from file.provider.id
		2. set config from file.provider.[`${file.provider.type}_config`]
	*/
	provider := "local"

	storageService, err := storage.CreateStorage(provider)
	isStorageUnsupported := err != nil
	if isStorageUnsupported {
		return nil, err
	}

	storageFile := storage.FileEntity{
		UniqueId:     file.UniqueId, //local only use this field
		Extension:    "jpeg",        //local only use this field
		OriginalName: file.OriginalName,
		Name:         file.Name,
		Size:         file.Size,
		Mimetype:     file.Mimetype,
		Url:          file.Url,
		Path:         file.Path,
		CreatedAt:    file.CreatedAt,
		UpdatedAt:    file.UpdatedAt,
	}

	/*
		@todo
		1. refactor function param using dto if necessary (consistency)
	*/
	fileData, err := storageService.RetrieveFile(&storageFile)
	isRetrieveSuccess := err == nil
	if !isRetrieveSuccess {
		return nil, err
	}

	fileResult := FileEntity{
		UniqueId:      file.UniqueId,
		OriginalName:  file.OriginalName,
		Name:          file.Name,
		Extension:     file.Extension,
		Mimetype:      file.Mimetype,
		Size:          file.Size,
		Url:           file.Url,
		Path:          file.Path,
		ProviderId:    file.ProviderId,
		ApplicationId: file.ApplicationId,
		CreatedAt:     file.CreatedAt,
		UpdatedAt:     file.UpdatedAt,
	}
	result := &RetrieveFileResult{
		FileData: fileData,
		File:     fileResult,
	}

	return result, nil
}
