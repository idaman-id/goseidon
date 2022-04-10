package storage_local

import (
	"bufio"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"

	app_error "idaman.id/storage/internal/error"
	"idaman.id/storage/internal/storage"
)

type storageLocal struct {
	storageDir string
}

func (s *storageLocal) RetrieveFile(fileLocation string) (storage.BinaryFile, error) {
	osFile, err := os.Open(fileLocation)
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

func (s *storageLocal) SaveFile(param storage.SaveFileParam) (*storage.SaveFileResult, error) {
	fl := s.storageDir
	fn := param.FileName
	path := fl + "/" + fn

	_, err := os.Stat(path)
	if !errors.Is(err, os.ErrNotExist) {
		return nil, app_error.NewAlreadyExistsError("File")
	}

	err = ioutil.WriteFile(path, param.FileData, 0644)
	if err != nil {
		return nil, err
	}

	res := storage.SaveFileResult{
		FileLocation: fl,
		FileName:     fn,
	}
	return &res, nil
}

func (s *storageLocal) DeleteFile(fileLocation string) error {
	err := os.Remove(s.storageDir + "/" + fileLocation)

	switch err.(type) {
	case *fs.PathError:
		err = app_error.NewNotfoundError("File")
	}

	return err
}

func NewStorageLocal(sDir string) *storageLocal {
	return &storageLocal{storageDir: sDir}
}
