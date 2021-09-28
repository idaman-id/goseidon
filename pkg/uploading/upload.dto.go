package uploading

import (
	"mime/multipart"

	"idaman.id/storage/pkg/file"
	"idaman.id/storage/pkg/storage"
)

type UploadFileDto struct {
	Files   []*multipart.FileHeader
	Storage storage.Uploader
}

type UploadResultItem struct {
	Status  string          `json:"status"`
	Message string          `json:"message,omitempty"`
	File    file.FileEntity `json:"file,omitempty"`
}

type UploadResult struct {
	Items []UploadResultItem
}
