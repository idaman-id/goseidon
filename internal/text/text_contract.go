package text

type StringParser interface {
	ParseString(d interface{}) string
}

type Slugger interface {
	Slugify(t string) string
}

type Generator interface {
	GenerateUuid() string
}

type TextService interface {
	Slugger
	StringParser
	Generator
}
