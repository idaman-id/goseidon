package text

type TextService interface {
	Slugify(text string) string
}