package retrieving

type FileGetter interface {
	GetFile(identifier string) (*FileEntity, error)
}

type FileRetriever interface {
	RetrieveFile(identifier string) (*RetrieveFileResult, error)
}

type RetrieveService interface {
	FileGetter
	FileRetriever
}

type RetrieveFileResult struct {
	File     FileEntity
	FileData []byte
}
