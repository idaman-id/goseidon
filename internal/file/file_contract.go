package file

import "mime/multipart"

type FileService interface {
	ParseOriginalName(fileHeader *multipart.FileHeader) string
	ParseName(fileHeader *multipart.FileHeader) string
	ParseSize(fileHeader *multipart.FileHeader) uint64
	ParseMimeType(fileHeader *multipart.FileHeader) string
	ParseExtension(fileHeader *multipart.FileHeader) string
	RemoveFileExtension(fileName string) string
}
