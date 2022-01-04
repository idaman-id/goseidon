package validation

const (
	VALIDATOR_GO_I18N = "go_i18n"
)

type ValidationService interface {
	Validate(i interface{}) error
}
