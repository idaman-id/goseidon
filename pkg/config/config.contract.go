package config

type Config interface {
	loadConfiguration() error
	GetString(key string) string
	GetInt(key string) int
	Get(key string) interface{}
	Set(key string, value interface{})
	SetDefault(key string, value interface{})
}

const (
	CONFIG_VIPER = "viper"
)
