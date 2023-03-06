// Package config containing necessary env set and init for the related service
package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

// Config for the service
type Config struct {
	SimpleAnalyzerName     string `env:"SIMPLE_ANALYZER_NAME" envDefault:""`     // Simple analyzer to be used
	QuickfixAnalyzerName   string `env:"QUICKFIX_ANALYZER_NAME" envDefault:""`   // Qucickfixes analyzer to be used
	StylecheckAnalyzerName string `env:"STYLECHECK_ANALYZER_NAME" envDefault:""` // Stylecheck analyzer to be used
}

// Init tries to parse os.env passed to the service run command.
// Then in returns a config ready for usage by checker
func Init() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return cfg, nil
}
