package config

type Config interface {
	LoadConfiguration() error
	GetAsString(key string) string
	GetAsUint(key string) uint
	Get(key string) interface{}
	Set(key string, value interface{})
	SetDefault(key string, value interface{})
}

const (
	PROVIDER_VIPER = "viper"
)
