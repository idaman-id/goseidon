package uploading

import (
	"idaman.id/storage/internal/storage"
	"idaman.id/storage/pkg/validation"
)

func UploadFile(param UploadFileParam) (*UploadResult, error) {
	var rule UploadRule
	rule.SetData(UploadRuleParam{
		FileHeaders: param.Files,
		Provider:    param.Provider,
	})

	err := validation.Service.ValidateStruct(rule)

	isDataInvalid := err != nil
	if isDataInvalid {
		return nil, err
	}

	storage, err := storage.NewStorage(param.Provider)
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
		2. test
	*/
	return &uploadResult, nil
}
