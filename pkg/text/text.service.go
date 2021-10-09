package text

import "github.com/gosimple/slug"

func Slugify(text string) string {
	result := slug.Make(text)
	return result
}
