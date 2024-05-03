package config

var GlobalConfig *Config

type Config struct {
	MessageApp string
}

func NewConfig(messageApp string) *Config {
	return &Config{
		MessageApp: messageApp,
	}
}
