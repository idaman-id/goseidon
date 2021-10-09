package retrieving

type FileNotFoundError struct {
	Message string
}

func (error *FileNotFoundError) Error() string {
	return error.Message
}
