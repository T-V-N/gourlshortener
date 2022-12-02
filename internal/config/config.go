package config

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:":8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	SecretKey       string `env:"SECRET_KEY" envDefault:"hello"`
	DatabaseDSN     string `env:"DATABASE_DSN"`
}

func Init() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "Server address")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "base url to use in strings")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "where to save db")
	flag.StringVar(&cfg.SecretKey, "s", cfg.SecretKey, "secret key")
	flag.StringVar(&cfg.DatabaseDSN, "d", cfg.DatabaseDSN, "secret key")
	flag.Parse()

	return cfg, nil
}
