package file

const (
	TYPE_AUDIO    = "audio"
	TYPE_VIDEO    = "video"
	TYPE_DOCUMENT = "document"
	TYPE_IMAGE    = "image"
	TYPE_ARCHIEVE = "archieve"
)

var (
	SUPPORTED_TYPES = map[string]bool{
		TYPE_AUDIO:    true,
		TYPE_VIDEO:    true,
		TYPE_DOCUMENT: true,
		TYPE_IMAGE:    true,
		TYPE_ARCHIEVE: true,
	}
)
