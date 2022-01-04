package file

import "mime/multipart"

type FileService interface {
	FileParser
	FileRemover
}

type FileParser interface {
	ParseOriginalName(fh *multipart.FileHeader) string
	ParseName(fh *multipart.FileHeader) string
	ParseSize(fh *multipart.FileHeader) int64
	ParseMimeType(fh *multipart.FileHeader) string
	ParseExtension(fh *multipart.FileHeader) string
}

type FileRemover interface {
	RemoveFileExtension(fn string) string
}
