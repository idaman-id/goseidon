package error

type ValidationItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value"`
}

type ValidationError struct {
	Message string
	Items   []ValidationItem
}

func (error *ValidationError) Error() string {
	return error.Message
}

func NewValidationError(items []ValidationItem) *ValidationError {
	err := &ValidationError{
		Message: STATUS_INVALID_DATA,
		Items:   items,
	}
	return err
}

type UnsupportedError struct {
	Message string
	Context string
}

func (error *UnsupportedError) Error() string {
	return error.Message
}

func NewUnsupportedError(context string) *UnsupportedError {
	err := &UnsupportedError{
		Message: STATUS_NOT_SUPPORTED,
		Context: context,
	}
	return err
}

type NotfoundError struct {
	Message string
	Context string
}

func (error *NotfoundError) Error() string {
	return error.Message
}

func NewNotfoundError(context string) *NotfoundError {
	err := &NotfoundError{
		Message: STATUS_NOT_FOUND,
		Context: context,
	}
	return err
}
