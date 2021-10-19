package translation

import "github.com/nicksnyder/go-i18n/v2/i18n"

type Localizer interface {
	Localize(lc *i18n.LocalizeConfig) (string, error)
}
