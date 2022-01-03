package rest_fiber_test

import (
	"errors"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	app_error "idaman.id/storage/internal/error"
	response "idaman.id/storage/internal/response"
	"idaman.id/storage/internal/retrieving"
)

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Package")
}

type StubDeleteService struct {
}

func (s *StubDeleteService) DeleteFile(identifier string) error {
	if identifier == "not-found" {
		return app_error.NewNotFoundError("File")
	} else if identifier == "error" {
		return errors.New(response.STATUS_ERROR)
	}
	return nil
}

type StubFileGetterService struct {
}

func (stub *StubFileGetterService) GetFile(identifier string) (*retrieving.FileEntity, error) {
	if identifier == "not-found" {
		return nil, app_error.NewNotFoundError("File")
	} else if identifier == "error" {
		return nil, errors.New(response.STATUS_ERROR)
	}
	file := &retrieving.FileEntity{}
	return file, nil
}

type StubFileRetrieverService struct {
}

func (stub *StubFileRetrieverService) RetrieveFile(identifier string) (*retrieving.RetrieveFileResult, error) {
	if identifier == "not-found" {
		return nil, app_error.NewNotFoundError("File")
	} else if identifier == "error" {
		return nil, errors.New(response.STATUS_ERROR)
	}
	file := retrieving.FileEntity{}
	fileData := make([]byte, 0)
	result := &retrieving.RetrieveFileResult{
		File:     file,
		FileData: fileData,
	}
	return result, nil
}
