package repository

import "time"

type FileRepository interface {
	FindByIdentifier(identifier string) (*FileModel, error)
	Save(p SaveFileParam) error
	Delete(identifier string) error
}

type SaveFileParam struct {
	UniqueId     string
	OriginalName string
	Name         string
	Extension    string
	Size         int64
	Mimetype     string
	FileLocation string
	FileName     string
	CreatedAt    *time.Time
}
