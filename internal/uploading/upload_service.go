package uploading

import (
	"fmt"
	"time"

	"idaman.id/storage/internal/config"
	"idaman.id/storage/internal/repository"
	"idaman.id/storage/internal/storage"
	"idaman.id/storage/internal/text"
	"idaman.id/storage/internal/validation"
)

type uploadService struct {
	validator       validation.Validator
	configGetter    config.Getter
	storageSaver    storage.Saver
	stringGenerator text.Generator
	fileRepo        repository.FileRepository
}

func (s *uploadService) UploadFile(p UploadFileParam) (*FileEntity, error) {
	ur := NewUploadRule(p.File)
	err := s.validator.Validate(*ur)

	if err != nil {
		return nil, err
	}

	uniqueId := s.stringGenerator.GenerateUuid()
	createdAt := time.Now()
	fileName := uniqueId + "." + p.File.Extension
	res, err := s.storageSaver.SaveFile(storage.SaveFileParam{
		FileName: fileName,
		FileData: p.File.Data,
	})
	if err != nil {
		return nil, err
	}

	appUrl := s.configGetter.GetString("APP_URL")
	publicUrl := fmt.Sprintf("%s/%s/%s", appUrl, "file", res.FileName)

	err = s.fileRepo.Save(repository.SaveFileParam{
		UniqueId:     uniqueId,
		OriginalName: p.File.OriginalName,
		Name:         p.File.Name,
		Size:         p.File.Size,
		CreatedAt:    &createdAt,
		Extension:    p.File.Extension,
		Mimetype:     p.File.Mimetype,
		FileLocation: res.FileLocation,
		FileName:     res.FileName,
	})
	if err != nil {
		return nil, err
	}

	file := FileEntity{
		UniqueId:     uniqueId,
		Name:         p.File.Name,
		OriginalName: p.File.OriginalName,
		Size:         p.File.Size,
		Extension:    p.File.Extension,
		Mimetype:     p.File.Mimetype,
		Url:          publicUrl,
		CreatedAt:    &createdAt,
		UpdatedAt:    nil,
		DeletedAt:    nil,
	}
	return &file, nil
}

func NewUploadService(v validation.Validator, cg config.Getter, ss storage.Saver, sg text.Generator, fr repository.FileRepository) UploadService {
	return &uploadService{
		validator:       v,
		configGetter:    cg,
		storageSaver:    ss,
		stringGenerator: sg,
		fileRepo:        fr,
	}
}
