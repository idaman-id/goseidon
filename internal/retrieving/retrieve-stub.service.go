package retrieving

import (
	"errors"

	"idaman.id/storage/pkg/app"
)

type StubFileGetterService struct {
}

func (stub *StubFileGetterService) GetFile(identifier string) (*FileEntity, error) {
	if identifier == "not-found" {
		return nil, app.NewNotFoundError("File")
	} else if identifier == "error" {
		return nil, errors.New(app.STATUS_ERROR)
	}
	file := &FileEntity{}
	return file, nil
}
