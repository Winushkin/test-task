// Package config package used for upload app configuration
package config

import (
	"fmt"

	"file-manager/internal/postgres"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	PostgresCfg     postgres.Config `env-prefix:"POSTGRES_"`
	ReportsDirPath  string          `env:"REPORT_DIR_PATH"`
	TSVDirPath      string          `env:"TSV_DIR_PATH"`
	LogDIRPath      string          `env:"LOG_DIR_PATH"`
	PollingInterval int             `env:"POLLING_SECONDS_INTERVAL"`
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Load: %w", err)
	}

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("ReadEnv: %w", err)
	}

	return &cfg, nil
}
