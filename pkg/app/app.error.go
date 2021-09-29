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
