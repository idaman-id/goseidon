package translation_test

import (
	"errors"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type StubGoI18n struct {
	success bool
}

func (mock *StubGoI18n) Localize(lc *i18n.LocalizeConfig) (string, error) {
	if mock.success {
		return "localized", nil
	}
	return "", errors.New("Failed translate")
}
