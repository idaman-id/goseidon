package deleting

type DeleteService interface {
	DeleteFile(identifier string) error
}
