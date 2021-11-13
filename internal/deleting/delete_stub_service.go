package deleting

import (
	"errors"

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
