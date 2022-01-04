package uploading

import (
	"fmt"

	"idaman.id/storage/internal/config"
	"idaman.id/storage/internal/file"
	"idaman.id/storage/internal/validation"
)

type uploadService struct {
	validationService validation.ValidationService
	configGetter      config.Getter
	fileService       file.FileService
}

func (s *uploadService) UploadFile(param UploadFileParam) (*UploadResult, error) {
	var rule UploadRule
	rule.SetData(UploadRuleParam{
		FileHeaders: param.Files,
		Provider:    param.Provider,
	})

	err := s.validationService.Validate(rule)

	if err != nil {
		return nil, err
	}

	// storage, err := storage.NewStorage(param.Provider, s.configGetter, s.fileService)
	// if err != nil {
	// 	return nil, err
	// }

	uploadResult := UploadResult{}
	for _, fileHeader := range param.Files {
		fmt.Println("fileHeader", fileHeader)

		// fileResult, err := storage.SaveFile(fileHeader)
		// isSaveSuccess := err == nil

		// if isSaveSuccess {
		// 	file := FileEntity{
		// 		UniqueId:      fileResult.UniqueId,
		// 		OriginalName:  fileResult.OriginalName,
		// 		Name:          fileResult.Name,
		// 		Extension:     fileResult.Extension,
		// 		Size:          fileResult.Size,
		// 		Mimetype:      fileResult.Mimetype,
		// 		Url:           fileResult.Url,
		// 		Path:          fileResult.Path,
		// 		CreatedAt:     fileResult.CreatedAt,
		// 		ProviderId:    "", //@todo: update this field
		// 		ApplicationId: "", //@todo: update this field
		// 	}
		// 	uploadResult.Items = append(uploadResult.Items, UploadResultItem{
		// 		Status: UPLOAD_SUCCESS,
		// 		File:   &file,
		// 	})
		// } else {
		uploadResult.Items = append(uploadResult.Items, UploadResultItem{
			Status:  UPLOAD_FAILED,
			Message: err.Error(),
		})
		// }

	}

	/*
		@todo:
		1. insert record for all success upload files (including provider_id, application_id)
		2. test
	*/
	return &uploadResult, nil
}

func NewUploadService(validationService validation.ValidationService, configGetter config.Getter, fileService file.FileService) UploadService {
	s := &uploadService{
		validationService: validationService,
		configGetter:      configGetter,
		fileService:       fileService,
	}
	return s
}
