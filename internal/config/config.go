package config

import (
	"fmt"

	"flag"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerURL       string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"./../../file_storage"`
}

func Init() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "Server address")
	flag.StringVar(&cfg.ServerURL, "b", cfg.ServerURL, "base url to use in strings")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "where to save db")
	flag.Parse()

	return cfg, nil
}