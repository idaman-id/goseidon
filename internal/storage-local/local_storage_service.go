package storage_local

import (
	"bufio"
	"io/ioutil"
	"os"
	"time"

	"github.com/valyala/fasthttp"
	"idaman.id/storage/internal/storage"
)

type storageLocal struct {
	storageDir string
}

func (s *storageLocal) RetrieveFile(localPath string) (storage.BinaryFile, error) {
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

func (s *storageLocal) SaveFile(param storage.SaveFileParam) (*storage.FileEntity, error) {
	path := s.storageDir + "/" + param.FileName
	createdAt := time.Now()

	err := fasthttp.SaveMultipartFile(&param.FileHeader, path)
	if err != nil {
		return nil, err
	}

	file := storage.FileEntity{
		Name:      param.FileName,
		Size:      param.FileHeader.Size,
		LocalPath: path,
		CreatedAt: createdAt,
	}
	return &file, nil
}

// func (s *storageLocal) DeleteFile(localPath string) error {
// 	err := os.Remove(localPath)

// 	switch err.(type) {
// 	case *fs.PathError:
// 		err = app_error.NewNotfoundError("File")
// 	}

// 	return err
// }

func NewStorageLocal() *storageLocal {
	storage := &storageLocal{
		storageDir: "storage/file",
	}
	return storage
}
