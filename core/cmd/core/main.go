package main

import (
	"core/internal/arguments"
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"go.yaml.in/yaml/v2"
)

//

const (
	defaultModel = "qwen/qwen3.7-flash"
)

//

type Args struct {
	EnvPath string
	CfgPath string
}

type Env struct {
	OpenAIKey     string `env:"OPENAI_KEY"`
	OpenAIBaseURL string `env:"OPENAI_BASE_URL"`
}

type Config struct{}

//

func main() {

	args := loadArgs()

	env := Env{}
	if err := loadEnv(&env, args.EnvPath); err != nil {
		slog.Error("load env", "err", err)
		os.Exit(1)
	}

	cfg := Config{}
	f, err := os.ReadFile(args.CfgPath)
	if err != nil {
		slog.Error("read config file", "err", err)
		os.Exit(1)
	}
	if err := yaml.Unmarshal(f, cfg); err != nil {
		slog.Error("parse config file", "err", err)
		os.Exit(1)
	}

	slog.Info("running app")

	if err := run(); err != nil {
		slog.Error("app failed", "err", err)
		os.Exit(1)
	}

	slog.Info("app exited")
}

//

func run() error {
	oai := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_KEY")),
		option.WithBaseURL(os.Getenv("OPENAI_BASE_URL")),
	)

	return cliLoop(useRequest(oai))
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

	// env path
	if cfgPath, ok := argsRaw["config"]; ok {
		args.CfgPath = cfgPath
	} else if cfgPath, ok := argsRaw["c"]; ok {
		args.CfgPath = cfgPath
	} else {
		args.CfgPath = "./config.yaml" // default
	}

	return args
}

func loadEnv(e *Env, envFilePath ...string) error {
	if err := godotenv.Load(envFilePath...); err != nil {
		return fmt.Errorf("failed to load env files: %w", err)
	}
	if err := env.Parse(e); err != nil {
		return fmt.Errorf("failed to parse env vars: %w", err)
	}
	return nil
}
