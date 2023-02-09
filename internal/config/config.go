// Package config containing necessary env set and init for the related service
package config

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env/v6"
)

// Config for the service
type Config struct {
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"` // URL where server will be started
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:":8080"`           // Server port
	FileStoragePath string `env:"FILE_STORAGE_PATH"`                           // Path to a file which will be used as a storage
	SecretKey       string `env:"SECRET_KEY" envDefault:"hello"`               // Secret for hashing ops
	DatabaseDSN     string `env:"DATABASE_DSN"`                                // Database connection string for DB-style storage
}

// Init tries to parse os.env and flags passed to the service run command.
// Flags have priority over os envs.
// Then in returns a config ready for usage by other service layers.
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
