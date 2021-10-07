package config

var config Config

func InitConfig(provider string) {
	config = CreateConfig(provider)
}

func CreateConfig(provider string) Config {
	if provider != PROVIDER_VIPER {
		panic("Config provider is not supported")
	}

	return &ViperConfig{
		fileName: ".env",
	}
}

func LoadConfiguration() error {
	err := config.LoadConfiguration()
	if err != nil {
		return err
	}

	config.SetDefault("MIN_UPLOADED_FILE", 1)
	config.SetDefault("MAX_UPLOADED_FILE", 5)
	return nil
}

func GetAsString(key string) string {
	return config.GetAsString(key)
}

func GetAsUint(key string) uint {
	return config.GetAsUint(key)
}

func Get(key string) interface{} {
	return config.Get(key)
}

func Set(key string, value interface{}) {
	config.Set(key, value)
}

func SetDefault(key string, value interface{}) {
	config.SetDefault(key, value)
}
