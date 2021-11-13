package validation

type ValidationService interface {
	ValidateStruct(param interface{}) error
}

const (
	VALIDATOR_GO_I18N = "go_i18n"
)
