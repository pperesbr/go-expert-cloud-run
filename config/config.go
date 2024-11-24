package config

type AppConfig struct {
	ServerPort    string
	GoogleKey     string
	WeatherAPIKey string
}

func NewAppConfig(serverPort string, weatherAPIKey string) *AppConfig {
	return &AppConfig{
		ServerPort:    serverPort,
		WeatherAPIKey: weatherAPIKey,
	}
}
