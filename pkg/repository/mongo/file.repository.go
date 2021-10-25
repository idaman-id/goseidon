package repository_mongo

import (
	"time"

	"idaman.id/storage/pkg/app"
	"idaman.id/storage/pkg/file"
)

type FileRepository struct {
}

/*
	@todo
	1. mongo db implementation
	2. test
*/
func (repo *FileRepository) FindByUniqueId(uniqueId string) (*file.FileEntity, error) {

	if uniqueId == "not_found" {
		err := app.NewNotFoundError("File")
		return nil, err
	}

	file := &file.FileEntity{
		UniqueId:      uniqueId,
		OriginalName:  "hio.jpeg",
		Name:          "hio",
		Extension:     "jpeg",
		Size:          6720,
		Mimetype:      "image/jpeg",
		Url:           "http://storage.domain.tld/storage/file/4980a441-b747-4226-ada0-63a5b2cac26d.jpeg",
		Path:          "",
		ProviderId:    "",
		ApplicationId: "",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return file, nil
}
