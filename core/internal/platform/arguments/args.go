package arguments

import (
	"log/slog"
	"os"
	"strings"
)

type Args struct {
	LogLevel   slog.Level
	ConfigPath string
	EnvPath    string
	DevMode    bool
}

func Parse() Args {

	argsRaw := readArgs()

	args := Args{}

	// log
	if logLevel, ok := argsRaw["log-level"]; ok {
		args.LogLevel = parseLogLevel(logLevel)
	} else if logLevel, ok := argsRaw["l"]; ok {
		args.LogLevel = parseLogLevel(logLevel)
	} else {
		args.LogLevel = slog.LevelInfo
	}

	// config path
	if configPath, ok := argsRaw["config"]; ok {
		args.ConfigPath = configPath
	} else if configPath, ok := argsRaw["c"]; ok {
		args.ConfigPath = configPath
	} else {
		args.ConfigPath = "./config.yaml"
	}

	// env path
	if envPath, ok := argsRaw["env"]; ok {
		args.EnvPath = envPath
	} else if envPath, ok := argsRaw["e"]; ok {
		args.EnvPath = envPath
	} else {
		args.EnvPath = "./.env"
	}

	// dev mode
	if _, ok := argsRaw["dev"]; ok {
		args.DevMode = true
	} else if _, ok := argsRaw["d"]; ok {
		args.DevMode = true
	}

	return args
}

//

// readArgs converts command-line arguments into a map.
// Supported formats: --flag="value" --flag=value --flag value -f value
// (!) AI generated
func readArgs() map[string]string {
	args := os.Args[1:]
	result := make(map[string]string)

	for i := 0; i < len(args); i++ {
		arg := args[i]

		// Handle --key=value
		if strings.HasPrefix(arg, "--") && strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			key := strings.TrimPrefix(parts[0], "--")
			result[key] = strings.Trim(parts[1], `"'`)
			continue
		}

		// Handle --key value or -k value
		if strings.HasPrefix(arg, "-") {
			key := strings.TrimPrefix(arg, "-")
			key = strings.TrimPrefix(key, "-") // Handles the second dash if it was --

			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				result[key] = strings.Trim(args[i+1], `"'`)
				i++ // Skip the next element since it is the value
			} else {
				result[key] = "" // Flag present without value
			}
		}
	}

	return result
}

//

// parseLogLevel maps string values ["debug", "info", "warn", "error"] to slog.Level
// Defaults to INFO
func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError

	}
	return slog.LevelInfo
}
