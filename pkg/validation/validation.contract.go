package validation

type ValidationService interface {
	ValidateStruct(param interface{}) error
}
