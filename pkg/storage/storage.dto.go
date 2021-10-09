package storage

import (
	"mime/multipart"
)

type FileDto = *multipart.FileHeader
type BinaryFile = []byte
