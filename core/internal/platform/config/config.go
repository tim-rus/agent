package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

func Read(cfg interface{}, cfgFile, envFile string) error {

	if cfgFile == "" {
		return errors.New("failed to read confg file")
	}

	// yaml
	f, err := os.ReadFile(cfgFile)
	if err != nil {
		return fmt.Errorf("failed to load yaml file: %w", err)
	}
	if err := yaml.Unmarshal(f, cfg); err != nil {
		return fmt.Errorf("failed to unmarshal yaml file: %w", err)
	}

	// env
	if envFile != "" {
		if err := godotenv.Load(envFile); err != nil {
			slog.Warn("CONFIG", "failed to load env file", envFile)
		}
	}
	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("failed to parse env vars: %w", err)
	}

	// ok
	return nil
}
