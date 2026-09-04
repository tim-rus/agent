package main

import (
	"core/internal/platform/arguments"
	"log/slog"
)

//

type Args struct {
	EnvPath string
}

//

func main() {
	slog.Info("hi")

	args := loadArgs()

	slog.Info("read args", "args", args)
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
