// Package config containing necessary env set and init for the related service
package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
)

// Config for the service
type Config struct {
	BaseURL         string `env:"BASE_URL" json:"base_url"`                   // URL where server will be started
	ServerAddress   string `env:"SERVER_ADDRESS" json:"server_address"`       // Server port
	FileStoragePath string `env:"FILE_STORAGE_PATH" json:"file_storage_path"` // Path to a file which will be used as a storage
	SecretKey       string `env:"SECRET_KEY" envDefault:"hello"`              // Secret for hashing ops
	DatabaseDSN     string `env:"DATABASE_DSN" json:"database_dsn"`           // Database connection string for DB-style storage
	EnableHTTPS     bool   `env:"ENABLE_HTTPS" json:"enable_https"`           // flag enables https
	TrustedSubnet   string `env:"TRUSTED_SUBNET" json:"trusted_subnet"`       // trusted subnet cidr

}

// Init tries to parse os.env and flags passed to the service run command.
// Flags have priority over os envs.
// Then in returns a config ready for usage by other service layers.
func Init() (*Config, error) {
	cfg := &Config{BaseURL: "http://localhost:8080", ServerAddress: ":8080", EnableHTTPS: false}

	cfgPath := os.Getenv("CONFIG")
	configFlag := flag.NewFlagSet("Only config", flag.ContinueOnError)
	configFlag.StringVar(&cfgPath, "c", cfgPath, "Config file")
	configFlag.Parse(os.Args[1:])

	if cfgPath != "" {
		fileCfg, err := parseFileConfig(cfgPath)

		if err == nil {
			cfg = fileCfg
		}
	}

	err := env.Parse(cfg)

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "Server address")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "base url to use in strings")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "where to save db")
	flag.StringVar(&cfg.SecretKey, "k", cfg.SecretKey, "secret key")
	flag.StringVar(&cfg.DatabaseDSN, "d", cfg.DatabaseDSN, "secret key")
	flag.BoolVar(&cfg.EnableHTTPS, "s", cfg.EnableHTTPS, "flag enable HTTPS")
	flag.StringVar(&cfgPath, "c", cfgPath, "Config file")
	flag.StringVar(&cfg.TrustedSubnet, "t", cfg.TrustedSubnet, "Trusted subnet")

	flag.Parse()

	return cfg, nil
}

func parseFileConfig(name string) (*Config, error) {
	file, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	fileCfg := &Config{}
	err = json.NewDecoder(file).Decode(&fileCfg)

	if err != nil {
		return nil, err
	}

	return fileCfg, nil
}
