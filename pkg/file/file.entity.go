package file

import (
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

type FileEntity struct {
	UniqueId      string    `json:"unique_id"`
	Type          string    `json:"type"`
	Name          string    `json:"name"`
	Extension     string    `json:"extension"`
	Size          uint64    `json:"size"`
	Mimetype      string    `json:"mimetype"`
	Url           string    `json:"url"`
	Path          string    `json:"path"`
	ProviderId    string    `json:"provider_id"`
	ApplicationId string    `json:"application_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (file *FileEntity) New(fileHeader *multipart.FileHeader) {
	file.detectMetaData(fileHeader)
}

func (file *FileEntity) detectMetaData(fileHeader *multipart.FileHeader) {
	file.detectName(fileHeader)
	file.detectSize(fileHeader)
	file.detectMimeType(fileHeader)
	file.detectExtension()
	file.detectType()
}

func (file *FileEntity) detectName(fileHeader *multipart.FileHeader) {
	file.Name = strings.ToLower(fileHeader.Filename)
}

func (file *FileEntity) detectSize(fileHeader *multipart.FileHeader) {
	file.Size = uint64(fileHeader.Size)
}

func (file *FileEntity) detectMimeType(fileHeader *multipart.FileHeader) {
	contentType, isAvailable := fileHeader.Header["Content-Type"]
	isContentTypeAvailable := isAvailable && len(contentType) > 0

	if isContentTypeAvailable {
		file.Mimetype = strings.ToLower(contentType[0])
	}
}

func (file *FileEntity) detectExtension() {
	extension := filepath.Ext(file.Name)
	extensionWithoutDot := strings.ReplaceAll(extension, ".", "")

	file.Extension = strings.ToLower(extensionWithoutDot)
}

func (file *FileEntity) detectType() {
	extension := file.Extension
	fileType, isExtensionSupported := SUPPORTED_EXTENSIONS[extension]

	if isExtensionSupported {
		file.Type = strings.ToLower(fileType)
	}
}
