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
	Size         int64
	Mimetype     string
	FileLocation string
	FileName     string
	CreatedAt    int64
	UpdatedAt    sql.NullInt64
	DeletedAt    sql.NullInt64
}
