package repository

import "idaman.id/storage/internal/file"

const (
	DATABASE_MONGO = "mongo"
)

type FileRepository interface {
	FindByUniqueId(uniqueId string) (*file.FileEntity, error)
}

type Repository struct {
	FileRepo FileRepository
}
