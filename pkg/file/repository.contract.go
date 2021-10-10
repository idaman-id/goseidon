package file

type FileRepository interface {
	FindByUniqueId(uniqueId string) (*FileEntity, error)
}
