package storage

import "mime/multipart"

type BinaryFile = []byte
type Retriever interface {
	RetrieveFile(localPath string) (result BinaryFile, err error)
}

type Deleter interface {
	DeleteFile(localPath string) error
}

type Saver interface {
	SaveFile(param SaveFileParam) (file *FileEntity, err error)
}

type SaveFileParam struct {
	FileHeader multipart.FileHeader
	FileName   string
}

type Storage interface {
	Saver
	Retriever
	Deleter
}
