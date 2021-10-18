package config_test

type MockConfig struct {
	values          map[string]interface{}
	loadConfigError error
}

func NewMockconfig() *MockConfig {
	config := &MockConfig{}
	config.values = map[string]interface{}{}
	return config
}

func (config *MockConfig) LoadConfiguration() error {
	return config.loadConfigError
}

func (config *MockConfig) GetString(key string) string {
	value, _ := config.values[key]
	return value.(string)
}

func (config *MockConfig) GetInt(key string) int {
	value, _ := config.values[key]
	return value.(int)
}

func (config *MockConfig) Get(key string) interface{} {
	value, _ := config.values[key]
	return value
}

func (config *MockConfig) Set(key string, value interface{}) {
	config.values[key] = value
}

func (config *MockConfig) SetDefault(key string, value interface{}) {
	config.Set(key, value)
}
