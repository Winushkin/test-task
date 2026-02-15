// Package config package used for upload app configuration
package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresName     string `env:"POSTRGES_NAME"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresPort     int    `env:"POSTGRES_PORT"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	ReportsDirPath   string `env:"REPORT_DIR_PATH"`
	TSVDirPath       string `env:"TSV_DIR_PATH"`
	PollinInterval   int    `env:"POLLING_SECONDS_INTERVAL"`
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
