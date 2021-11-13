package retrieving

import (
	"idaman.id/storage/internal/file"
	"idaman.id/storage/internal/repository"
	"idaman.id/storage/internal/storage"
	"idaman.id/storage/pkg/config"
)

type retrieveService struct {
	fileRepo     repository.FileRepository
	configGetter config.Getter
	fileService  file.FileService
}

func (s *retrieveService) GetFile(identifier string) (*FileEntity, error) {

	uniqueId := s.fileService.RemoveFileExtension(identifier)
	file, err := s.fileRepo.FindByUniqueId(uniqueId)
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

func (s *retrieveService) RetrieveFile(identifier string) (*RetrieveFileResult, error) {

	uniqueId := s.fileService.RemoveFileExtension(identifier)
	fileRecord, err := s.fileRepo.FindByUniqueId(uniqueId)
	isRecordAvailable := err == nil
	if !isRecordAvailable {
		return nil, err
	}

	/*
		@todo
		1. set provider from file.provider.id
		2. set config from file.provider.[`${file.provider.type}_config`]
		3. test
	*/
	provider := "local"

	storageService, err := storage.NewStorage(provider, s.configGetter, s.fileService)
	isStorageUnsupported := err != nil
	if isStorageUnsupported {
		return nil, err
	}

	storageFile := &storage.FileEntity{
		UniqueId:     fileRecord.UniqueId,  //local only use this field
		Extension:    fileRecord.Extension, //local only use this field
		OriginalName: fileRecord.OriginalName,
		Name:         fileRecord.Name,
		Size:         fileRecord.Size,
		Mimetype:     fileRecord.Mimetype,
		Url:          fileRecord.Url,
		Path:         fileRecord.Path,
		CreatedAt:    fileRecord.CreatedAt,
		UpdatedAt:    fileRecord.UpdatedAt,
	}

	/*
		@todo
		1. refactor function param using dto if necessary (consistency)
	*/
	fileData, err := storageService.RetrieveFile(storageFile)
	isRetrieveSuccess := err == nil
	if !isRetrieveSuccess {
		return nil, err
	}

	fileResult := FileEntity{
		UniqueId:      fileRecord.UniqueId,
		OriginalName:  fileRecord.OriginalName,
		Name:          fileRecord.Name,
		Extension:     fileRecord.Extension,
		Mimetype:      fileRecord.Mimetype,
		Size:          fileRecord.Size,
		Url:           fileRecord.Url,
		Path:          fileRecord.Path,
		ProviderId:    fileRecord.ProviderId,
		ApplicationId: fileRecord.ApplicationId,
		CreatedAt:     fileRecord.CreatedAt,
		UpdatedAt:     fileRecord.UpdatedAt,
	}
	result := &RetrieveFileResult{
		FileData: fileData,
		File:     fileResult,
	}

	return result, nil
}

func NewRetrieveService(fileRepo repository.FileRepository, fileService file.FileService) RetrieveService {
	s := &retrieveService{
		fileRepo:    fileRepo,
		fileService: fileService,
	}
	return s
}
