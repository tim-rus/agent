package main

import (
	"core/internal/arguments"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

//

const (
	defaultModel = "qwen/qwen3.7-flash"
)

//

type Args struct {
	EnvPath string
}

//

func main() {

	args := loadArgs()

	if err := godotenv.Load(args.EnvPath); err != nil {
		slog.Error("failed to load env file", "err", err)
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

	return args
}
