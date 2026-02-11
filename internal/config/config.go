//Package config package used for upload app configuration
package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct{
	PostgresUser string `env:"POSTGRES_USER"`
	PostgresName string `env:"POSTRGES_NAME"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresPort int `env:"POSTGRES_PORT"`
	PostgresHost string `env:"POSTGRES_HOST"`
}


func initEnv() error {
	if err := godotenv.Load(); err != nil{
		return fmt.Errorf("failed to get Envs: %w", err)
	}
	return nil
}

func NewConfig() (*Config, error) {
	if err := initEnv(); err != nil{
		return nil, fmt.Errorf("initEnv: %w", err)
	}

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil{
		return nil, fmt.Errorf("ReadEnv: %w", err)
	}

	return &cfg, nil
}