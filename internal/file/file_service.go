package file

import (
	"mime/multipart"
	"path/filepath"
	"strings"

	"idaman.id/storage/pkg/text"
)

func ParseOriginalName(fileHeader *multipart.FileHeader) string {
	fileName := strings.ToLower(fileHeader.Filename)
	return fileName
}

func ParseName(fileHeader *multipart.FileHeader) string {
	fileName := ParseOriginalName(fileHeader)
	fileNameWithoutExtension := RemoveFileExtension(fileName)
	fileName = text.Service.Slugify(fileNameWithoutExtension)
	return fileName
}

func ParseSize(fileHeader *multipart.FileHeader) uint64 {
	fileSize := uint64(fileHeader.Size)
	return fileSize
}

func ParseMimeType(fileHeader *multipart.FileHeader) string {
	contentType, isAvailable := fileHeader.Header["Content-Type"]
	isMimeTypeAvailable := isAvailable && len(contentType) > 0

	if !isMimeTypeAvailable {
		return ""
	}
	return contentType[0]
}

func ParseExtension(fileHeader *multipart.FileHeader) string {
	extension := filepath.Ext(fileHeader.Filename)
	extensionWithoutDot := strings.ReplaceAll(extension, ".", "")
	fileExtension := strings.ToLower(extensionWithoutDot)
	return fileExtension
}

func RemoveFileExtension(fileName string) string {
	name := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return name
}
