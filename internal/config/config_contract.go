package config

const (
	CONFIG_VIPER = "viper"
)

type Defaultable interface {
	SetDefault(key string, value interface{})
}

type Getter interface {
	GetString(key string) string
	GetInt(key string) int
	Get(key string) interface{}
}

type Setter interface {
	Set(key string, value interface{})
}

type ConfigService interface {
	LoadConfiguration() error
	Defaultable
	Setter
	Getter
}
