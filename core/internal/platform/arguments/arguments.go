package arguments

import (
	"os"
	"strings"
)

// readArgs converts command-line arguments into a map.
// Supported formats: --flag="value" --flag=value --flag value -f value
// (!) AI generated
func Read() map[string]string {
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
