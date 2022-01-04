package retrieving

import (
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

	fileEntity := &FileEntity{
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
	return fileEntity, nil
}

func (s *retrieveService) RetrieveFile(identifier string) (*RetrieveFileResult, error) {

	fileRecord, err := s.fileRepo.FindByIdentifier(identifier)
	if err != nil {
		return nil, err
	}

	fileData, err := s.storageRetriever.RetrieveFile(fileRecord.LocalPath)
	if err != nil {
		return nil, err
	}

	fileResult := &FileEntity{
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

func NewRetrieveService(fr repository.FileRepository, cg config.Getter, fs file.FileService, sr storage.Retriever) RetrieveService {
	return &retrieveService{
		configGetter:     cg,
		fileRepo:         fr,
		fileService:      fs,
		storageRetriever: sr,
	}
}
