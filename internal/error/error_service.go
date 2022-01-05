package error

import "fmt"

type ValidationItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Message string
	Items   []ValidationItem
}

func (error *ValidationError) Error() string {
	return error.Message
}

func NewValidationError(items []ValidationItem) *ValidationError {
	return &ValidationError{
		Message: STATUS_INVALID_DATA,
		Items:   items,
	}
}

type UnsupportedError struct {
	Message string
	Context string
}

func (error *UnsupportedError) Error() string {
	return error.Message
}

func NewUnsupportedError(context string) *UnsupportedError {
	return &UnsupportedError{
		Message: STATUS_NOT_SUPPORTED,
		Context: context,
	}
}

type NotfoundError struct {
	Message string
	Context string
}

func (error *NotfoundError) Error() string {
	return fmt.Sprintf("%s is not found", error.Context)
}

func NewNotfoundError(context string) *NotfoundError {
	return &NotfoundError{
		Message: STATUS_NOT_FOUND,
		Context: context,
	}
}

type AlreadyExistsError struct {
	Message string
	Context string
}

func (error *AlreadyExistsError) Error() string {
	return error.Message
}

func NewAlreadyExistsError(context string) *AlreadyExistsError {
	return &AlreadyExistsError{
		Message: STATUS_ALREADY_EXISTS,
		Context: context,
	}
}
