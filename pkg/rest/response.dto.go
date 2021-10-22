package rest

import (
	"idaman.id/storage/pkg/translation"
)

type ResponseParam struct {
	Message         string
	Data            interface{}
	Error           interface{}
	Translator      translation.Translator
	TranslationData map[string]interface{}
}
