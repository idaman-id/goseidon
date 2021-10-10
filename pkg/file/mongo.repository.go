package file

import (
	"time"

	"idaman.id/storage/pkg/app"
)

type MongoRepository struct {
}

/*
	@todo
	1. mongo db implementation
*/
func (repo *MongoRepository) FindByUniqueId(uniqueId string) (*FileEntity, error) {

	if uniqueId == "not_found" {
		return nil, &app.NotFoundError{
			Message: app.STATUS_NOT_FOUND,
			Context: "File",
		}
	}

	file := &FileEntity{
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
