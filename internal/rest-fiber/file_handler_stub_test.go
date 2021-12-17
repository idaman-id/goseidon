package rest_fiber_test

import (
	"errors"

	"idaman.id/storage/internal/retrieving"
	"idaman.id/storage/pkg/app"
)

type StubDeleteService struct {
}

func (s *StubDeleteService) DeleteFile(identifier string) error {
	if identifier == "not-found" {
		return app.NewNotFoundError("File")
	} else if identifier == "error" {
		return errors.New(app.STATUS_ERROR)
	}
	return nil
}

type StubFileGetterService struct {
}

func (stub *StubFileGetterService) GetFile(identifier string) (*retrieving.FileEntity, error) {
	if identifier == "not-found" {
		return nil, app.NewNotFoundError("File")
	} else if identifier == "error" {
		return nil, errors.New(app.STATUS_ERROR)
	}
	file := &retrieving.FileEntity{}
	return file, nil
}

type StubFileRetrieverService struct {
}

func (stub *StubFileRetrieverService) RetrieveFile(identifier string) (*retrieving.RetrieveFileResult, error) {
	if identifier == "not-found" {
		return nil, app.NewNotFoundError("File")
	} else if identifier == "error" {
		return nil, errors.New(app.STATUS_ERROR)
	}
	file := retrieving.FileEntity{}
	fileData := make([]byte, 0)
	result := &retrieving.RetrieveFileResult{
		File:     file,
		FileData: fileData,
	}
	return result, nil
}
