package uploading

import (
	"mime/multipart"

	"idaman.id/storage/pkg/file"
)

type UploadFileDto struct {
	Files    []*multipart.FileHeader
	Provider string
	Locale   string
}

type UploadResultItem struct {
	Status  string           `json:"status"`
	Message string           `json:"message,omitempty"`
	File    *file.FileEntity `json:"file,omitempty"`
}

type UploadResult struct {
	Items []UploadResultItem
}
