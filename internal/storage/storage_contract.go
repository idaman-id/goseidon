package storage

import "mime/multipart"

type BinaryFile = []byte

type Uploader interface {
	SaveFile(fh *multipart.FileHeader) (result *FileEntity, err error)
}

type Retriever interface {
	RetrieveFile(localPath string) (result BinaryFile, err error)
}

type Deleter interface {
	DeleteFile(file *FileEntity) error
}

type Saver interface {
	SaveFile(fh *multipart.FileHeader) (result *FileEntity, err error)
}

type Storage interface {
	Uploader
	Retriever
	Deleter
}
