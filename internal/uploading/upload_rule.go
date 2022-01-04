package uploading

import "mime/multipart"

type fileRule struct {
	Size int64 `json:"size" validate:"required,valid_file_size"`
}

func NewUploadRule(fh multipart.FileHeader) *fileRule {
	f := fileRule{
		Size: fh.Size,
	}
	return &f
}
