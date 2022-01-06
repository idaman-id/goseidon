package retrieving

import (
	"fmt"

	"idaman.id/storage/internal/config"
	"idaman.id/storage/internal/file"
	"idaman.id/storage/internal/repository"
	"idaman.id/storage/internal/storage"
)

type retrieveService struct {
	configGetter     config.Getter
	fileRepo         repository.FileRepository
	fileService      file.FileService
	storageRetriever storage.Retriever
}

func (s *retrieveService) GetFile(identifier string) (*FileEntity, error) {

	fileRecord, err := s.fileRepo.FindByIdentifier(identifier)
	if err != nil {
		return nil, err
	}

	appUrl := s.configGetter.GetString("APP_URL")
	url := fmt.Sprintf("%s/%s/%s", appUrl, "file", fileRecord.FileName)

	fileEntity := &FileEntity{
		UniqueId:     fileRecord.UniqueId,
		OriginalName: fileRecord.OriginalName,
		Name:         fileRecord.Name,
		Extension:    fileRecord.Extension,
		Size:         fileRecord.Size,
		Mimetype:     fileRecord.Mimetype,
		Url:          url,
		CreatedAt:    fileRecord.CreatedAt,
		UpdatedAt:    fileRecord.UpdatedAt,
		DeletedAt:    fileRecord.DeletedAt,
	}
	return fileEntity, nil
}

func (s *retrieveService) RetrieveFile(identifier string) (*RetrieveFileResult, error) {

	fileRecord, err := s.fileRepo.FindByIdentifier(identifier)
	if err != nil {
		return nil, err
	}

	localPath := fmt.Sprintf("%s/%s", fileRecord.FileLocation, fileRecord.FileName)
	fileData, err := s.storageRetriever.RetrieveFile(localPath)
	if err != nil {
		return nil, err
	}

	appUrl := s.configGetter.GetString("APP_URL")
	url := fmt.Sprintf("%s/%s/%s", appUrl, "file", fileRecord.FileName)
	fileResult := &FileEntity{
		UniqueId:     fileRecord.UniqueId,
		OriginalName: fileRecord.OriginalName,
		Name:         fileRecord.Name,
		Extension:    fileRecord.Extension,
		Mimetype:     fileRecord.Mimetype,
		Size:         fileRecord.Size,
		Url:          url,
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

func NewRetrieveService(fr repository.FileRepository, cg config.Getter, fs file.FileService, sr storage.Retriever) RetrieveService {
	return &retrieveService{
		configGetter:     cg,
		fileRepo:         fr,
		fileService:      fs,
		storageRetriever: sr,
	}
}
