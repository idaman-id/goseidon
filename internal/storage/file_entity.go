package storage

import (
	"time"
)

type FileEntity struct {
	UniqueId     string
	OriginalName string
	Name         string
	Extension    string
	Size         uint64
	Mimetype     string
	Url          string
	Path         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// func NewStorageFile(fileHeader *multipart.FileHeader, fileService file.FileService) *FileEntity {
// 	fileEntity := &FileEntity{}
// 	// fileEntity.detectMetaData(fileHeader)
// 	fileEntity.UniqueId = uuid.New().String()
// 	fileEntity.CreatedAt = time.Now()
// 	return fileEntity
// }

// func (fileEntity *FileEntity) detectMetaData(fileHeader *multipart.FileHeader) {
// 	fileEntity.detectName(fileHeader)
// 	fileEntity.detectSize(fileHeader)
// 	fileEntity.detectMimeType(fileHeader)
// 	fileEntity.detectExtension(fileHeader)
// }

// func (fileEntity *FileEntity) detectName(fileHeader *multipart.FileHeader) {
// 	fileEntity.OriginalName = file.ParseOriginalName(fileHeader)
// 	fileEntity.Name = file.ParseName(fileHeader)
// }

// func (fileEntity *FileEntity) detectSize(fileHeader *multipart.FileHeader) {
// 	fileEntity.Size = file.ParseSize(fileHeader)
// }

// func (fileEntity *FileEntity) detectMimeType(fileHeader *multipart.FileHeader) {
// 	fileEntity.Mimetype = file.ParseMimeType(fileHeader)
// }

// func (fileEntity *FileEntity) detectExtension(fileHeader *multipart.FileHeader) {
// 	fileEntity.Extension = file.ParseExtension(fileHeader)
// }
