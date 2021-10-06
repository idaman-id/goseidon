package uploading

import (
	"mime/multipart"

	"idaman.id/storage/pkg/file"
)

/* max_size: 128MB = 134217728 bytes */
type FileRule struct {
	Type string `json:"files.type" validate:"valid_file_type"`
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
		var file file.FileEntity
		file.New(fileHeader)

		fileRule := &FileRule{
			Type: file.Type,
			Size: file.Size,
		}
		fileRules = append(fileRules, fileRule)
	}

	rule.Provider = provider
	rule.Files = fileRules
}
