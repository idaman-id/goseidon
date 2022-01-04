package retrieving

import (
	"idaman.id/storage/internal/config"
	"idaman.id/storage/internal/file"
	"idaman.id/storage/internal/repository"
	"idaman.id/storage/internal/storage"
)

type retrieveService struct {
	fileRepo     repository.FileRepository
	configGetter config.Getter
	fileService  file.FileService
}

func (s *retrieveService) GetFile(identifier string) (*FileEntity, error) {

	fileRecord, err := s.fileRepo.FindByIdentifier(identifier)
	if err != nil {
		return nil, err
	}

	fileEntity := FileEntity{
		UniqueId:     fileRecord.UniqueId,
		OriginalName: fileRecord.OriginalName,
		Name:         fileRecord.Name,
		Extension:    fileRecord.Extension,
		Size:         fileRecord.Size,
		Mimetype:     fileRecord.Mimetype,
		Url:          fileRecord.PublicUrl,
		Path:         fileRecord.LocalPath,
		CreatedAt:    fileRecord.CreatedAt,
		UpdatedAt:    fileRecord.UpdatedAt,
		DeletedAt:    fileRecord.DeletedAt,
	}
	return &fileEntity, nil
}

func (s *retrieveService) RetrieveFile(identifier string) (*RetrieveFileResult, error) {

	fileRecord, err := s.fileRepo.FindByIdentifier(identifier)
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
		Url:          fileRecord.PublicUrl,
		Path:         fileRecord.LocalPath,
		CreatedAt:    *fileRecord.CreatedAt,
		UpdatedAt:    *fileRecord.UpdatedAt,
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
		UniqueId:     fileRecord.UniqueId,
		OriginalName: fileRecord.OriginalName,
		Name:         fileRecord.Name,
		Extension:    fileRecord.Extension,
		Mimetype:     fileRecord.Mimetype,
		Size:         fileRecord.Size,
		Url:          fileRecord.PublicUrl,
		Path:         fileRecord.LocalPath,
		CreatedAt:    fileRecord.CreatedAt,
		UpdatedAt:    fileRecord.UpdatedAt,
		DeletedAt:    fileRecord.DeletedAt,
	}
	result := &RetrieveFileResult{
		FileData: fileData,
		File:     fileResult,
	}

	return result, nil
}

func NewRetrieveService(fileRepo repository.FileRepository, configGetter config.Getter, fileService file.FileService) RetrieveService {
	s := &retrieveService{
		fileRepo:     fileRepo,
		configGetter: configGetter,
		fileService:  fileService,
	}
	return s
}
