package config

type Config interface {
	LoadConfiguration() error
	GetString(key string) string
	GetInt(key string) int
	Get(key string) interface{}
	Set(key string, value interface{})
	SetDefault(key string, value interface{})
}

const (
	PROVIDER_VIPER = "viper"
)
