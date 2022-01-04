package uploading

import "mime/multipart"

type UploadService interface {
	UploadFile(p UploadFileParam) (*FileEntity, error)
}

type UploadFileParam struct {
	File multipart.FileHeader
}

type UploadRuleParam struct {
	FileHeader *multipart.FileHeader
	Provider   string
}
