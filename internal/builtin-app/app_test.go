package builtin_app_test

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	app_error "idaman.id/storage/internal/error"
	response "idaman.id/storage/internal/response"
	"idaman.id/storage/internal/retrieving"
)

func TestBuiltinApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BuiltinApp Package")
}

func StringifyResponse(r io.Reader) string {
	body, _ := ioutil.ReadAll(r)
	resBody := string(body)

	return resBody
}

func UnmarshallResponseBody(r io.Reader) *response.ResponseEntity {
	resBody := StringifyResponse(r)

	var resEntity *response.ResponseEntity
	json.Unmarshal([]byte(resBody), &resEntity)

	return resEntity
}

type FakeDeleteService struct {
}

func (s *FakeDeleteService) DeleteFile(identifier string) error {
	if identifier == "not-found" {
		return app_error.NewNotfoundError("File")
	} else if identifier == "error" {
		return errors.New(response.STATUS_ERROR)
	}
	return nil
}

type FakeFileGetterService struct {
}

func (stub *FakeFileGetterService) GetFile(identifier string) (*retrieving.FileEntity, error) {
	if identifier == "not-found" {
		return nil, app_error.NewNotfoundError("File")
	} else if identifier == "error" {
		return nil, errors.New(response.STATUS_ERROR)
	}
	file := &retrieving.FileEntity{}
	return file, nil
}

type FakeFileRetrieverService struct {
}

func (stub *FakeFileRetrieverService) RetrieveFile(identifier string) (*retrieving.RetrieveFileResult, error) {
	if identifier == "not-found" {
		return nil, app_error.NewNotfoundError("File")
	} else if identifier == "error" {
		return nil, errors.New(response.STATUS_ERROR)
	}
	file := retrieving.FileEntity{}
	fileData := make([]byte, 0)
	result := &retrieving.RetrieveFileResult{
		File:     &file,
		FileData: fileData,
	}
	return result, nil
}
