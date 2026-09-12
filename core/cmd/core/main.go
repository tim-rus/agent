package main

import (
	"core/internal/dialog"
	"core/internal/llm"
	"core/internal/platform/arguments"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"go.yaml.in/yaml/v2"
)

//

type Args struct {
	EnvPath     string
	CfgPath     string
	PromptsPath string
}

type Env struct {
	OpenAIKey     string `env:"OPENAI_KEY"`
	OpenAIBaseURL string `env:"OPENAI_BASE_URL"`
}

type Config struct {
	Models struct {
		Default string `yaml:"default"`
	} `yaml:"models"`
}

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
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		slog.Error("parse config file", "err", err)
		os.Exit(1)
	}

	slog.Info("running app")

	if err := run(args, cfg); err != nil {
		slog.Error("app failed", "err", err)
		os.Exit(1)
	}

	slog.Info("app exited")
}

//

func run(args Args, cfg Config) error {
	oai := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_KEY")),
		option.WithBaseURL(os.Getenv("OPENAI_BASE_URL")),
	)

	prompts, err := readFilesToMap(args.PromptsPath)
	if err != nil {
		return fmt.Errorf("failed to read prompt files: %w", err)
	}

	llm := llm.New(oai, cfg.Models.Default)
	dialog := dialog.New(llm, prompts["system"])

	return cliLoop(dialog)
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

	// config path
	if cfgPath, ok := argsRaw["config"]; ok {
		args.CfgPath = cfgPath
	} else if cfgPath, ok := argsRaw["c"]; ok {
		args.CfgPath = cfgPath
	} else {
		args.CfgPath = "./config.yaml" // default
	}

	// prompts path
	if promtsPath, ok := argsRaw["prompts"]; ok {
		args.PromptsPath = promtsPath
	} else if promtsPath, ok := argsRaw["p"]; ok {
		args.PromptsPath = promtsPath
	} else {
		args.PromptsPath = "./prompts" // default
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

//

// ai generated (!)
func readFilesToMap(dirPath string) (map[string]string, error) {
	// 1. Read all entries within the designated folder
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	// 2. Initialize the map (pre-allocating size boosts performance)
	fileMap := make(map[string]string, len(entries))

	for _, entry := range entries {
		// 3. Skip subdirectories to only evaluate standard files
		if entry.IsDir() {
			continue
		}

		// 4. Resolve the absolute/relative layout path for the file
		filePath := filepath.Join(dirPath, entry.Name())

		// 5. Load file contents entirely into memory
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", entry.Name(), err)
		}

		// 6. Convert file name to lowercase for the map key
		lowerName := strings.ReplaceAll(strings.ToLower(entry.Name()), ".md", "")

		// 7. Store filename and content string in map
		fileMap[lowerName] = string(content)
	}

	return fileMap, nil
}
