package repository

type FileRepository interface {
	FindByIdentifier(identifier string) (*FileModel, error)
}
