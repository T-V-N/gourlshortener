// Package config is responsible for taking the runtime configuration from
// multiple sources of parameters and providing a structured configuration
// data to the service at the time of launch. It is also provides sensible
// defaults.
//
// Environment variables are considered the primary source of configuration.
// It supports the 12-factors app approach.
// For developers' convenience configuration can be overridden
// with CLI parameters.
package config

import (
	"fmt"

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

	return cfg, nil
}
