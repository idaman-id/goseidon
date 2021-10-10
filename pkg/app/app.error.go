package app

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

type NotSupportedError struct {
	Message string
	Context string
}

func (error *NotSupportedError) Error() string {
	return error.Message
}

type NotFoundError struct {
	Message string
	Context string
}

func (error *NotFoundError) Error() string {
	return error.Message
}
