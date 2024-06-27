package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbSource      string `env:"DB_SOURCE"`
	DbType        string `env:"DB_TYPE"`
	DbUsername    string `env:"DB_USERNAME"`
	DbPassword    string `env:"DB_PASSWORD"`
	DbName        string `env:"DB_NAME"`
	DbTestName    string `env:"DB_TEST_NAME"`
	DbHost        string `env:"DB_HOST" env-default:"localhost"`
	DbPort        string `env:"DB_PORT" env-default:"5432"`
	DbSSLMode     string `env:"DB_SSL_MODE" env-default:"disable"`
	ServerAddress string `env:"SERVER_ADDRESS" env-default:"0.0.0.0:8888"`
}

func NewConfig(path string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path, cfg)

	if err != nil {
		return nil, err
	}

	return cfg, err
}
