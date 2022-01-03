package file

import (
	"mime/multipart"
	"path/filepath"
	"strings"

	"idaman.id/storage/internal/text"
)

type fileService struct {
	slugger text.Slugger
}

func (s *fileService) ParseOriginalName(fileHeader *multipart.FileHeader) string {
	if fileHeader == nil {
		return ""
	}
	fileName := strings.ToLower(fileHeader.Filename)
	return fileName
}

func (s *fileService) ParseName(fileHeader *multipart.FileHeader) string {
	fileName := s.ParseOriginalName(fileHeader)
	fileNameWithoutExtension := s.RemoveFileExtension(fileName)
	fileName = s.slugger.Slugify(fileNameWithoutExtension)
	return fileName
}

func (s *fileService) ParseSize(fileHeader *multipart.FileHeader) uint64 {
	if fileHeader == nil {
		return 0
	}
	size := fileHeader.Size
	if size < 0 {
		return 0
	}
	return uint64(size)
}

func (s *fileService) ParseMimeType(fileHeader *multipart.FileHeader) string {
	if fileHeader == nil {
		return ""
	}
	contentType, isAvailable := fileHeader.Header["Content-Type"]
	isMimeAvailable := isAvailable && len(contentType) > 0

	if !isMimeAvailable {
		return ""
	}
	return contentType[0]
}

func (s *fileService) ParseExtension(fileHeader *multipart.FileHeader) string {
	if fileHeader == nil {
		return ""
	}
	extension := filepath.Ext(fileHeader.Filename)
	extensionWithoutDot := strings.ReplaceAll(extension, ".", "")
	lowerExt := strings.ToLower(extensionWithoutDot)
	return lowerExt
}

func (s *fileService) RemoveFileExtension(fileName string) string {
	extenstion := filepath.Ext(fileName)
	name := strings.TrimSuffix(fileName, extenstion)
	lowerName := strings.ToLower(name)
	return lowerName
}

func NewFileService(slugger text.Slugger) FileService {
	return &fileService{
		slugger: slugger,
	}
}
