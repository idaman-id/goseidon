package deleting

import (
	"idaman.id/storage/internal/file"
	"idaman.id/storage/internal/repository"
	"idaman.id/storage/internal/storage"
)

type deleteService struct {
	fileRepo    repository.FileRepository
	storage     storage.Deleter
	fileService file.FileRemover
}

func (s *deleteService) DeleteFile(identifier string) error {
	fileId := s.fileService.RemoveFileExtension(identifier)

	fileModel, err := s.fileRepo.FindByIdentifier(fileId)
	if err != nil {
		return err
	}

	err = s.fileRepo.Delete(fileId)
	if err != nil {
		return err
	}

	err = s.storage.DeleteFile(fileModel.FileName)
	if err != nil {
		return err
	}

	return nil
}

func NewDeleteService(sr storage.Deleter, fr repository.FileRepository, fs file.FileRemover) DeleteService {
	return &deleteService{
		storage:     sr,
		fileRepo:    fr,
		fileService: fs,
	}
}
