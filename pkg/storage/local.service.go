package storage

import (
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"idaman.id/storage/pkg/file"
)

type StorageLocal struct {
	StorageDir string
}

func (storage *StorageLocal) SaveFile(fileHeader FileDto) (result *SaveFileResult, err error) {
	uuid := uuid.New().String()
	path := storage.StorageDir + "/" + uuid + "-" + fileHeader.Filename
	err = fasthttp.SaveMultipartFile(fileHeader, path)

	isSaveFailed := err != nil
	if isSaveFailed {
		return nil, err
	}

	fileData := file.FileEntity{
		UniqueId: uuid,
		Url:      "http://storage.domain.tld/" + path,
		Path:     path,
	}
	fileData.FetchMetaData(fileHeader)

	result = &SaveFileResult{
		File: fileData,
	}
	return result, nil
}
