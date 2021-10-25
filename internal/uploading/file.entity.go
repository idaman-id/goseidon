package uploading

import (
	"time"
)

type FileEntity struct {
	UniqueId      string    `json:"unique_id"`
	OriginalName  string    `json:"original_name"`
	Name          string    `json:"name"`
	Extension     string    `json:"extension"`
	Size          uint64    `json:"size"`
	Mimetype      string    `json:"mimetype"`
	Url           string    `json:"url"`
	Path          string    `json:"-"`
	ProviderId    string    `json:"provider_id"`
	ApplicationId string    `json:"application_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
