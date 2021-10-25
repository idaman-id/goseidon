package uploading

import (
	"mime/multipart"
)

type UploadFileParam struct {
	Files    []*multipart.FileHeader
	Provider string
	Locale   string
}

type UploadResultItem struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	File    *FileEntity `json:"file,omitempty"`
}

type UploadResult struct {
	Items []UploadResultItem
}

type UploadRuleParam struct {
	FileHeaders []*multipart.FileHeader
	Provider    string
}
