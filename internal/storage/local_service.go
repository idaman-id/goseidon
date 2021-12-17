package storage

import (
	"bufio"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/valyala/fasthttp"
	"idaman.id/storage/internal/file"
	"idaman.id/storage/pkg/app"
	"idaman.id/storage/pkg/config"
)

type StorageLocal struct {
	StorageDir   string
	configGetter config.Getter
	fileService  file.FileService
}

func (s *StorageLocal) SaveFile(fileHeader *multipart.FileHeader) (result *FileEntity, err error) {
	file := NewFile(fileHeader, s.fileService)

	path := s.StorageDir + "/" + file.UniqueId + "." + file.Extension

	appUrl := s.configGetter.GetString("APP_URL")
	file.Path = path
	file.Url = appUrl + "/" + path

	err = fasthttp.SaveMultipartFile(fileHeader, path)

	isSaveFailed := err != nil
	if isSaveFailed {
		return nil, err
	}

	return file, nil
}

func (s *StorageLocal) RetrieveFile(file *FileEntity) (result BinaryFile, err error) {
	path := s.StorageDir + "/" + file.UniqueId + "." + file.Extension
	osFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer osFile.Close()

	reader := bufio.NewReader(osFile)
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (s *StorageLocal) DeleteFile(file *FileEntity) error {
	path := s.StorageDir + "/" + file.UniqueId + "." + file.Extension
	err := os.Remove(path)

	switch err.(type) {
	case *fs.PathError:
		err = &app.NotFoundError{
			Message: app.STATUS_NOT_FOUND,
			Context: "File",
		}
	}

	return err
}

func NewStorageLocal(configGetter config.Getter, fileService file.FileService) *StorageLocal {
	storage := StorageLocal{
		StorageDir:   "storage/file",
		configGetter: configGetter,
		fileService:  fileService,
	}
	return &storage
}
