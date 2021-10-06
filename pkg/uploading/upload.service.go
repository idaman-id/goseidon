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
	isProviderUnsupported := err == nil
	if isProviderUnsupported {
		return nil, err
	}

	uploadResult := UploadResult{}
	for _, file := range param.Files {

		saveResult, err := storage.SaveFile(file)
		isSaveSuccess := err == nil

		if isSaveSuccess {
			uploadResult.Items = append(uploadResult.Items, UploadResultItem{
				Status: UPLOAD_SUCCESS,
				File:   &saveResult.File,
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

		@nicetohave:
		1. concurrent/pararell processing
	*/
	return &uploadResult, nil
}
