package validation

import "github.com/go-playground/validator/v10"

var (
	Service ValidationService
)

func Init() error {
	if Service != nil {
		return nil
	}

	service, err := NewGoValidator(validator.New())
	if err != nil {
		return err
	}

	Service = service
	return nil
}
