package file

import "time"

type File struct {
	UUID      string    `json:"uuid"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Extension string    `json:"extension"`
	Size      uint32    `json:"size"`
	Mimetype  string    `json:"mimetype"`
	Url       string    `json:"url"`
	Path      string    `json:"path"`
	ServerId  string    `json:"server_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
