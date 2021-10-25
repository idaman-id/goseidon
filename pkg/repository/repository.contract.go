package repository

import "idaman.id/storage/pkg/file"

type FileRepository interface {
	FindByUniqueId(uniqueId string) (*file.FileEntity, error)
}
