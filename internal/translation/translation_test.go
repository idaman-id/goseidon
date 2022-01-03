package translation_test

import (
	"errors"
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
)

func TestTranslation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Translation Package")
}

type StubGoI18n struct {
	success bool
}

func (mock *StubGoI18n) Localize(lc *i18n.LocalizeConfig) (string, error) {
	if mock.success {
		return "localized", nil
	}
	return "", errors.New("Failed translate")
}
