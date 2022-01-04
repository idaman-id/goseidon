package repository_mysql

import (
	"database/sql"
)

type FileModel struct {
	Id           int64
	UniqueId     string
	OriginalName string
	Name         string
	Extension    string
	Size         uint64
	Mimetype     string
	PublicUrl    string
	LocalPath    string
	CreatedAt    int64
	UpdatedAt    sql.NullInt64
	DeletedAt    sql.NullInt64
}