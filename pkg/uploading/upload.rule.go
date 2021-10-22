package uploading

import (
	"idaman.id/storage/pkg/file"
)

type FileRule struct {
	Size uint64 `json:"files.size" validate:"required,valid_file_size"`
}

type FileRules = []*FileRule

type UploadRule struct {
	Files    FileRules `json:"files" validate:"required,valid_file_amount,dive,required"`
	Provider string    `json:"provider" validate:"required,valid_provider"`
}

func (rule *UploadRule) SetData(param UploadRuleParam) {
	var fileRules FileRules

	for _, fileHeader := range param.FileHeaders {

		fileRule := &FileRule{
			Size: file.ParseSize(fileHeader),
		}
		fileRules = append(fileRules, fileRule)
	}

	rule.Provider = param.Provider
	rule.Files = fileRules
}
