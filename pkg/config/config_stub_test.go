package config_test

type StubConfig struct {
	values          map[string]interface{}
	loadConfigError error
}

func NewStubconfig() *StubConfig {
	config := &StubConfig{}
	config.values = map[string]interface{}{}
	return config
}

func (config *StubConfig) LoadConfiguration() error {
	return config.loadConfigError
}

func (config *StubConfig) GetString(key string) string {
	value, _ := config.values[key]
	return value.(string)
}

func (config *StubConfig) GetInt(key string) int {
	value, _ := config.values[key]
	return value.(int)
}

func (config *StubConfig) Get(key string) interface{} {
	value, _ := config.values[key]
	return value
}

func (config *StubConfig) Set(key string, value interface{}) {
	config.values[key] = value
}

func (config *StubConfig) SetDefault(key string, value interface{}) {
	config.Set(key, value)
}
