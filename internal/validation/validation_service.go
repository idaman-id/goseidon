package validation

import (
	"idaman.id/storage/internal/config"
	validation_go "idaman.id/storage/internal/validation-go"
)

func NewValidator(cg config.Getter) (Validator, error) {
	return validation_go.NewGoValidator(cg)
}
