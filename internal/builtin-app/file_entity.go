package builtin_app

import (
	"time"
)

type FileDetailEntity struct {
	UniqueId  string     `json:"unique_id"`
	Name      string     `json:"name"`
	Extension string     `json:"extension"`
	Size      int64      `json:"size"`
	Mimetype  string     `json:"mimetype"`
	Url       string     `json:"url"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
