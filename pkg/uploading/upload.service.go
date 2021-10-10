package uploading

import (
	"idaman.id/storage/pkg/storage"
	"idaman.id/storage/pkg/validation"
)

func UploadFile(param UploadFileDto) (*UploadResult, error) {
	var rule UploadRule
	rule.SetData(param.Files, param.Provider)

	validationError := validation.ValidateStruct(validation.ValidationStructDto{
		Locale: param.Locale,
		Struct: rule,
	})

	isDataInvalid := validationError != nil
	if isDataInvalid {
		return nil, validationError
	}

	storage, err := storage.CreateStorage(param.Provider)
	isProviderUnsupported := err != nil

	if isProviderUnsupported {
		return nil, err
	}

	uploadResult := UploadResult{}
	for _, fileHeader := range param.Files {

		fileResult, err := storage.SaveFile(fileHeader)
		isSaveSuccess := err == nil

		if isSaveSuccess {
			file := FileEntity{
				UniqueId:      fileResult.UniqueId,
				OriginalName:  fileResult.OriginalName,
				Name:          fileResult.Name,
				Extension:     fileResult.Extension,
				Size:          fileResult.Size,
				Mimetype:      fileResult.Mimetype,
				Url:           fileResult.Url,
				Path:          fileResult.Path,
				CreatedAt:     fileResult.CreatedAt,
				ProviderId:    "", //@todo: update this field
				ApplicationId: "", //@todo: update this field
			}
			uploadResult.Items = append(uploadResult.Items, UploadResultItem{
				Status: UPLOAD_SUCCESS,
				File:   &file,
			})
		} else {
			uploadResult.Items = append(uploadResult.Items, UploadResultItem{
				Status:  UPLOAD_FAILED,
				Message: err.Error(),
			})
		}

	}

	/*
		@todo:
		1. insert record for all success upload files (including provider_id, application_id)
	*/
	return &uploadResult, nil
}
