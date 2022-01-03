package validation_test

type StubStringParser struct {
}

func (s *StubStringParser) ParseString(param interface{}) string {
	return ""
}

type StubConfigGetter struct {
}

func (s *StubConfigGetter) GetString(key string) string {
	return ""
}

func (s *StubConfigGetter) GetInt(key string) int {
	return 0
}

func (s *StubConfigGetter) Get(key string) interface{} {
	return ""
}
