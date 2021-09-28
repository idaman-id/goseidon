package storage

type Uploader interface {
	SaveFile(fileHeader FileDto) (result *SaveFileResult, err error)
}

type Storage interface {
	Uploader
}
