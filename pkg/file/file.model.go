package file

import "time"

type FileModel struct {
	UUID          string
	Type          string
	Name          string
	Extension     string
	Size          uint64
	Mimetype      string
	Url           string
	Path          string
	ProviderId    string
	ApplicationId string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
