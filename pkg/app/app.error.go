package app

type ValidationItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Err              error
	ValidationErrors []ValidationItem
}

func (validation *ValidationError) Error() string {
	return validation.Err.Error()
}
