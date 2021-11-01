package deleting

import (
	"idaman.id/storage/internal/file"
	"idaman.id/storage/internal/repository"
	"idaman.id/storage/internal/storage"
)

type BasicService struct {
	fileRepo repository.FileRepository
}

func (s *BasicService) DeleteFile(identifier string) error {

	uniqueId := file.RemoveFileExtension(identifier)
	fileRecord, err := s.fileRepo.FindByUniqueId(uniqueId)
	isRecordAvailable := err == nil
	if !isRecordAvailable {
		return err
	}

	/*
		@todo
		1. choose provider according to fileRecord.provider.type
		2. test
	*/
	provider := "local"
	storageService, err := storage.NewStorage(provider)
	if err != nil {
		return err
	}

	storageFile := &storage.FileEntity{
		UniqueId:     fileRecord.UniqueId,  //note: local only use this field
		Extension:    fileRecord.Extension, //note: local only use this field
		OriginalName: fileRecord.OriginalName,
		Name:         fileRecord.Name,
		Size:         fileRecord.Size,
		Mimetype:     fileRecord.Mimetype,
		Url:          fileRecord.Url,
		Path:         fileRecord.Path,
		CreatedAt:    fileRecord.CreatedAt,
		UpdatedAt:    fileRecord.UpdatedAt,
	}

	err = storageService.DeleteFile(storageFile)
	if err != nil {
		return err
	}

	/*
		@todo
		1. flag file as deleted (deleted_at) in database
	*/
	return nil
}
