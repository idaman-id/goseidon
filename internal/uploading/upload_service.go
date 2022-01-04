package uploading

import (
	"fmt"

	"idaman.id/storage/internal/config"
	"idaman.id/storage/internal/file"
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
	fileService     file.FileService
}

func (s *uploadService) UploadFile(p UploadFileParam) (*FileEntity, error) {
	ur := NewUploadRule(p.File)
	err := s.validator.Validate(*ur)

	if err != nil {
		return nil, err
	}

	uniqueId := s.stringGenerator.GenerateUuid()
	ext := s.fileService.ParseExtension(&p.File)
	mime := s.fileService.ParseMimeType(&p.File)
	name := s.fileService.ParseName(&p.File)
	oriName := s.fileService.ParseOriginalName(&p.File)
	size := s.fileService.ParseSize(&p.File)

	res, err := s.storageSaver.SaveFile(storage.SaveFileParam{
		FileHeader: p.File,
		FileName:   uniqueId + "." + ext,
	})
	if err != nil {
		return nil, err
	}

	appUrl := s.configGetter.GetString("APP_URL")
	publicUrl := fmt.Sprintf("%s/%s", appUrl, res.LocalPath)

	err = s.fileRepo.Save(repository.SaveFileParam{
		UniqueId:     uniqueId,
		OriginalName: oriName,
		Name:         name,
		Size:         size,
		PublicUrl:    publicUrl,
		LocalPath:    res.LocalPath,
		CreatedAt:    &res.CreatedAt,
		Extension:    ext,
		Mimetype:     mime,
	})
	if err != nil {
		return nil, err
	}

	file := FileEntity{
		UniqueId:     uniqueId,
		Name:         name,
		OriginalName: oriName,
		Size:         size,
		Extension:    ext,
		Mimetype:     mime,
		Path:         res.LocalPath,
		Url:          publicUrl,
		CreatedAt:    &res.CreatedAt,
		UpdatedAt:    nil,
		DeletedAt:    nil,
	}

	return &file, nil
}

func NewUploadService(v validation.Validator, cg config.Getter, ss storage.Saver, sg text.Generator, fr repository.FileRepository, fs file.FileService) UploadService {
	s := &uploadService{
		validator:       v,
		configGetter:    cg,
		storageSaver:    ss,
		stringGenerator: sg,
		fileRepo:        fr,
		fileService:     fs,
	}
	return s
}
