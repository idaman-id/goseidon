package uploading

import (
	"mime/multipart"

	"idaman.id/storage/internal/file"
)

type UploadService interface {
	UploadFile(p UploadFileParam) (*FileEntity, error)
}

type UploadFileParam struct {
	File *file.FileEntity
}

type UploadRuleParam struct {
	FileHeader *multipart.FileHeader
	Provider   string
}
