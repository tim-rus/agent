package main

import (
	"core/internal/platform/arguments"
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

//

type Args struct {
	EnvPath string
}

type Env struct {
	OpenAIKey string `env:"OPENAI_KEY"`
}

//

func main() {
	slog.Info("hi")

	args := loadArgs()

	slog.Info("read args", "args", args)

	env := Env{}
	if err := loadEnv(args.EnvPath, &env); err != nil {
		slog.Error("failed to load env file", "err", err)
		os.Exit(1)
	}

	slog.Info("read env", "env", env)
}

//

func loadArgs() Args {
	argsRaw := arguments.Read()

	args := Args{}

	// env path
	if envPath, ok := argsRaw["env"]; ok {
		args.EnvPath = envPath
	} else if envPath, ok := argsRaw["e"]; ok {
		args.EnvPath = envPath
	} else {
		args.EnvPath = "./.env" // default
	}

	return args
}

func loadEnv(envFilePath string, e *Env) error {
	if envFilePath == "" {
		return fmt.Errorf("env file path is empty")
	}
	if err := godotenv.Load(envFilePath); err != nil {
		return fmt.Errorf("failed to load env file: %w", err)
	}
	if err := env.Parse(e); err != nil {
		slog.Error("CONFIG: failed to parse env vars", "err", err)
		return fmt.Errorf("failed to parse env vars: %w", err)
	}
	return nil
}
