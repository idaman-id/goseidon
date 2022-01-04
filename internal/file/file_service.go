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

func (s *fileService) ParseOriginalName(fh *multipart.FileHeader) string {
	if fh == nil {
		return ""
	}
	fn := strings.ToLower(fh.Filename)
	return fn
}

func (s *fileService) ParseName(fh *multipart.FileHeader) string {
	fn := s.ParseOriginalName(fh)
	fileNameWithoutExtension := s.RemoveFileExtension(fn)
	fn = s.slugger.Slugify(fileNameWithoutExtension)
	return fn
}

func (s *fileService) ParseSize(fh *multipart.FileHeader) int64 {
	if fh == nil {
		return 0
	}
	size := fh.Size
	if size < 0 {
		return 0
	}
	return size
}

func (s *fileService) ParseMimeType(fh *multipart.FileHeader) string {
	if fh == nil {
		return ""
	}
	contentType, isAvailable := fh.Header["Content-Type"]
	isMimeAvailable := isAvailable && len(contentType) > 0

	if !isMimeAvailable {
		return ""
	}
	return contentType[0]
}

func (s *fileService) ParseExtension(fh *multipart.FileHeader) string {
	if fh == nil {
		return ""
	}
	ext := filepath.Ext(fh.Filename)
	extWithoutDot := strings.ReplaceAll(ext, ".", "")
	lExt := strings.ToLower(extWithoutDot)
	return lExt
}

func (s *fileService) RemoveFileExtension(fn string) string {
	ext := filepath.Ext(fn)
	name := strings.TrimSuffix(fn, ext)
	lName := strings.ToLower(name)
	return lName
}

func NewFileService(sl text.Slugger) FileService {
	return &fileService{
		slugger: sl,
	}
}
