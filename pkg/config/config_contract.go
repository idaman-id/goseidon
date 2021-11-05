package config

type Defaultable interface {
	SetDefault(key string, value interface{})
}

type ConfigService interface {
	LoadConfiguration() error
	GetString(key string) string
	GetInt(key string) int
	Get(key string) interface{}
	Set(key string, value interface{})
	Defaultable
}

const (
	CONFIG_VIPER = "viper"
)
