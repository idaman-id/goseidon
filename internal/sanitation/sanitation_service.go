package sanitation

import "github.com/go-sanitize/sanitize"

type Sanitator interface {
	Sanitize(o interface{}) error
}

type sanitationService struct {
	sa *sanitize.Sanitizer
}

func (s *sanitationService) Sanitize(o interface{}) error {
	return s.sa.Sanitize(o)
}

func NewSanitator() (*sanitationService, error) {
	sa, err := sanitize.New()
	if err != nil {
		return nil, err
	}

	s := &sanitationService{sa}
	return s, nil
}
