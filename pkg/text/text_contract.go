package text

type StringParser interface {
	ParseString(data interface{}) string
}

type Slugger interface {
	Slugify(text string) string
}

type TextService interface {
	Slugger
	StringParser
}
