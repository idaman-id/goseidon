package uploading

import (
	"mime/multipart"

	"idaman.id/storage/pkg/file"
)

/* max_size: 128MB = 134217728 bytes */
type FileRule struct {
	Size uint64 `json:"files.size" validate:"required,min=1,max=134217728"`
}

type FileRules = []*FileRule

type UploadRule struct {
	Files    FileRules `json:"files" validate:"required,valid_file_amounts,dive,required"`
	Provider string    `json:"provider" validate:"required,valid_provider"`
}

func (rule *UploadRule) SetData(fileHeaders []*multipart.FileHeader, provider string) {
	var fileRules FileRules

	for _, fileHeader := range fileHeaders {

		fileRule := &FileRule{
			Size: file.ParseSize(fileHeader),
		}
		fileRules = append(fileRules, fileRule)
	}

	rule.Provider = provider
	rule.Files = fileRules
}
