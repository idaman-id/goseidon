package storage

type BinaryFile = []byte

type Retriever interface {
	RetrieveFile(localPath string) (result BinaryFile, err error)
}

type Deleter interface {
	DeleteFile(localPath string) error
}

type Saver interface {
	SaveFile(param SaveFileParam) (result *SaveFileResult, err error)
}

type SaveFileParam struct {
	FileName string
	FileData BinaryFile
}

type SaveFileResult struct {
	FileLocation string
	FileName     string
}

type Storage interface {
	Saver
	Retriever
	Deleter
}
