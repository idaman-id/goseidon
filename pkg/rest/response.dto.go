package rest

import (
	"idaman.id/storage/pkg/translation"
)

type ResponseDto struct {
	Message         string
	Data            interface{}
	Error           interface{}
	Translator      translation.Translator
	TranslationData map[string]interface{}
}
