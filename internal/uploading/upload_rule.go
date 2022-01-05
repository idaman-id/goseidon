package uploading

import (
	"idaman.id/storage/internal/file"
)

type fileRule struct {
	Name      string `json:"name" validate:"required"`
	Extension string `json:"ext" validate:"required"`
	Mimetype  string `json:"mimetype" validate:"required"`
	Size      int64  `json:"size" validate:"required,valid_file_size"`
}

func NewUploadRule(f *file.FileEntity) *fileRule {
	fr := fileRule{
		Name:      f.Name,
		Extension: f.Extension,
		Mimetype:  f.Mimetype,
		Size:      f.Size,
	}
	return &fr
}
