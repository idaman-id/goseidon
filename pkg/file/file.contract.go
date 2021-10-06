package file

const (
	TYPE_AUDIO    = "audio"
	TYPE_VIDEO    = "video"
	TYPE_DOCUMENT = "document"
	TYPE_IMAGE    = "image"
	TYPE_ARCHIEVE = "archieve"
)

var (
	SUPPORTED_EXTENSIONS = map[string]string{
		"xls":  TYPE_DOCUMENT,
		"xlsx": TYPE_DOCUMENT,
		"doc":  TYPE_DOCUMENT,
		"docx": TYPE_DOCUMENT,
		"pdf":  TYPE_DOCUMENT,
		"ppt":  TYPE_DOCUMENT,
		"pptx": TYPE_DOCUMENT,
		"txt":  TYPE_DOCUMENT,

		"jpg":  TYPE_IMAGE,
		"jpeg": TYPE_IMAGE,
		"png":  TYPE_IMAGE,
		"gif":  TYPE_IMAGE,
		"svg":  TYPE_IMAGE,

		"mp3":  TYPE_AUDIO,
		"mpeg": TYPE_AUDIO,
		"aac":  TYPE_AUDIO,
		"wav":  TYPE_AUDIO,
		"ogg":  TYPE_AUDIO,

		"mp4": TYPE_VIDEO,
		"mkv": TYPE_VIDEO,
		"3gp": TYPE_VIDEO,

		"zip": TYPE_ARCHIEVE,
		"rar": TYPE_ARCHIEVE,
	}
	SUPPORTED_TYPES = map[string]bool{
		TYPE_AUDIO:    true,
		TYPE_VIDEO:    true,
		TYPE_DOCUMENT: true,
		TYPE_IMAGE:    true,
		TYPE_ARCHIEVE: true,
	}
)
