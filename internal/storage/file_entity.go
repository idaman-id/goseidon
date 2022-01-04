package storage

import (
	"time"
)

type FileEntity struct {
	Name      string
	Size      int64
	LocalPath string
	CreatedAt time.Time
}
