package repository

import (
	"time"
)

type FileModel struct {
	Id           int64
	UniqueId     string
	OriginalName string
	Name         string
	Extension    string
	Size         int64
	Mimetype     string
	PublicUrl    string
	LocalPath    string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}

func (m *FileModel) SetCreatedAtFromUnixTime(t int64) *FileModel {
	if t > 0 {
		ts := time.Unix(t, 0)
		m.CreatedAt = &ts
	}
	return m
}

func (m *FileModel) SetUpdatedAtFromUnixTime(t int64) *FileModel {
	if t > 0 {
		ts := time.Unix(t, 0)
		m.UpdatedAt = &ts
	}
	return m
}

func (m *FileModel) SetDeletedAtFromUnixTime(t int64) *FileModel {
	if t > 0 {
		ts := time.Unix(t, 0)
		m.DeletedAt = &ts
	}
	return m
}
