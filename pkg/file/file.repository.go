package file

import (
	"idaman.id/storage/pkg/app"
)

var repository FileRepository

func InitRepo(provider string) error {
	repo, err := NewRepo(provider)
	if err != nil {
		return err
	}

	repository = repo
	return nil
}

func NewRepo(provider string) (FileRepository, error) {
	if provider != app.DATABASE_MONGO {
		return nil, &app.NotSupportedError{
			Message: app.STATUS_NOT_SUPPORTED,
			Context: "Database",
		}
	}

	repo := &MongoRepository{}
	return repo, nil
}

func FindByUniqueId(uniqueId string) (*FileEntity, error) {
	return repository.FindByUniqueId(uniqueId)
}
