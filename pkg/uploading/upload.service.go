package uploading

func UploadFile(param UploadFileDto) UploadResult {

	uploadResult := UploadResult{}

	for _, file := range param.Files {

		saveResult, err := param.Storage.SaveFile(file)

		isSaveSuccess := err == nil
		if isSaveSuccess {
			uploadResult.Items = append(uploadResult.Items, UploadResultItem{
				Status: STATUS_SUCCESS,
				File:   saveResult.File,
			})
		} else {
			uploadResult.Items = append(uploadResult.Items, UploadResultItem{
				Status:  STATUS_FAILED,
				Message: err.Error(),
			})
		}

	}

	/**
	@todo:
	1. insert record for all files uploaded to provider
	2. set provider_id, application_id
	*/
	return uploadResult
}
