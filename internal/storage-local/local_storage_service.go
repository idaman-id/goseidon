package storage_local

import (
	"bufio"
	"io/ioutil"
	"os"

	"idaman.id/storage/internal/config"
	"idaman.id/storage/internal/file"
	"idaman.id/storage/internal/storage"
)

type storageLocal struct {
	storageDir   string
	configGetter config.Getter
	fileService  file.FileService
}

func (s *storageLocal) RetrieveFile(localPath string) (result storage.BinaryFile, err error) {
	osFile, err := os.Open(localPath)
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

// func (s *storageLocal) SaveFile(fileHeader *multipart.FileHeader) (result *FileEntity, err error) {
// 	file := NewStorageFile(fileHeader, s.fileService)

// 	path := s.StorageDir + "/" + file.UniqueId + "." + file.Extension

// 	appUrl := s.configGetter.GetString("APP_URL")
// 	file.Path = path
// 	file.Url = appUrl + "/" + path

// 	err = fasthttp.SaveMultipartFile(fileHeader, path)

// 	isSaveFailed := err != nil
// 	if isSaveFailed {
// 		return nil, err
// 	}

// 	return file, nil
// }

// func (s *storageLocal) DeleteFile(file *FileEntity) error {
// 	path := s.StorageDir + "/" + file.UniqueId + "." + file.Extension
// 	err := os.Remove(path)

// 	switch err.(type) {
// 	case *fs.PathError:
// 		err = app_error.NewNotfoundError("File")
// 	}

// 	return err
// }

func NewStorageLocal(configGetter config.Getter, fileService file.FileService) *storageLocal {
	storage := &storageLocal{
		storageDir:   "storage/file",
		configGetter: configGetter,
		fileService:  fileService,
	}
	return storage
}
