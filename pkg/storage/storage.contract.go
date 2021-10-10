package storage

import "mime/multipart"

type FileDto = *multipart.FileHeader
type BinaryFile = []byte

type Uploader interface {
	SaveFile(fileHeader FileDto) (result *FileEntity, err error)
}

type Retriever interface {
	RetrieveFile(file *FileEntity) (result BinaryFile, err error)
}

type Storage interface {
	Uploader
	Retriever
}
