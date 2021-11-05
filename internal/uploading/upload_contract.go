package uploading

import "mime/multipart"

const (
	UPLOAD_SUCCESS = "success"
	UPLOAD_FAILED  = "failed"
)

type UploadService interface {
	UploadFile(param UploadFileParam) (*UploadResult, error)
}

type UploadFileParam struct {
	Files    []*multipart.FileHeader
	Provider string
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
