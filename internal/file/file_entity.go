package file

import (
	"io"
	"mime/multipart"
)

type FileEntity struct {
	OriginalName string
	Size         int64
	Data         []byte

	Name      string
	Extension string
	Mimetype  string
}

func NewFileFromMultipartHeader(fh *multipart.FileHeader, fs FileService) (*FileEntity, error) {
	mFile, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer mFile.Close()

	fileData, err := io.ReadAll(mFile)
	if err != nil {
		return nil, err
	}

	name := fs.ParseName(fh)
	ext := fs.ParseExtension(fh)
	mime := fs.ParseMimeType(fh)

	file := &FileEntity{
		OriginalName: fh.Filename,
		Size:         fh.Size,
		Data:         fileData,

		Name:      name,
		Extension: ext,
		Mimetype:  mime,
	}
	return file, nil
}
