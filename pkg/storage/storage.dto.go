package storage

import (
	"mime/multipart"

	"idaman.id/storage/pkg/file"
)

type FileDto = *multipart.FileHeader

type SaveFileResult struct {
	File file.FileEntity
}

type GetFileResult struct {
}
