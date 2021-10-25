package storage

import "mime/multipart"

type BinaryFile = []byte

type Uploader interface {
	SaveFile(fileHeader *multipart.FileHeader) (result *FileEntity, err error)
}

type Retriever interface {
	RetrieveFile(file *FileEntity) (result BinaryFile, err error)
}

type Deleter interface {
	DeleteFile(file *FileEntity) error
}

type Storage interface {
	Uploader
	Retriever
	Deleter
}
