package config

import (
	"fmt"

	"flag"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerURL       string `env:"BASE_URL" envDefault:"example.com"`
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"./../../file_storage"`
}

func Init() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	flag.StringVar(&cfg.ServerAddress, "a", "Address", "Server address")
	flag.StringVar(&cfg.ServerURL, "b", "Base url", "base url to use in strings")
	flag.StringVar(&cfg.FileStoragePath, "f", "File storage path", "where to save db")

	return cfg, nil
}
