package retrieving

import (
	"time"
)

type FileEntity struct {
	UniqueId     string
	OriginalName string
	Name         string
	Extension    string
	Size         int64
	Mimetype     string
	Url          string
	Path         string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}
