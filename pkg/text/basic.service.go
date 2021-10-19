package text

import "github.com/gosimple/slug"

type BasicService struct {
}

func (service *BasicService) Slugify(text string) string {
	result := slug.Make(text)
	return result
}
