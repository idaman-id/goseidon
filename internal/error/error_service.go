package error

type ValidationItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value"`
}

type ValidationError struct {
	Message string
	Items   []*ValidationItem
}

func (error *ValidationError) Error() string {
	return error.Message
}

func NewValidationError(items []*ValidationItem) *ValidationError {
	err := &ValidationError{
		Message: ERROR_INVALID_DATA,
		Items:   items,
	}
	return err
}

type NotSupportedError struct {
	Message string
	Context string
}

func (error *NotSupportedError) Error() string {
	return error.Message
}

func NewNotSupportedError(context string) *NotSupportedError {
	err := &NotSupportedError{
		Message: ERROR_NOT_SUPPORTED,
		Context: context,
	}
	return err
}

type NotFoundError struct {
	Message string
	Context string
}

func (error *NotFoundError) Error() string {
	return error.Message
}

func NewNotFoundError(context string) *NotFoundError {
	err := &NotFoundError{
		Message: ERROR_NOT_FOUND,
		Context: context,
	}
	return err
}
