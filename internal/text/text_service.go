package text

import (
	"fmt"
	"strconv"

	"github.com/gosimple/slug"
)

type textService struct {
}

func (s *textService) Slugify(text string) string {
	result := slug.Make(text)
	return result
}

func (s *textService) ParseString(param interface{}) string {
	var value string
	switch param.(type) {
	case string:
		value = param.(string)
	case bool:
		bValue := param.(bool)
		value = strconv.FormatBool(bValue)
	case float64, float32:
		value = fmt.Sprint(param)
	case uint64, uint32, uint16, uint8:
		value = fmt.Sprint(param)
	case int64, int32, int16, int8:
		value = fmt.Sprint(param)
	default:
		value = ""
	}
	return value
}

func NewTextService() TextService {
	return &textService{}
}
