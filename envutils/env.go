package envutils

import (
	"os"
	"strconv"
)

// IsCI checks if the current environment is a CI/CD environment.
// It does this by looking for the "CI" env var
func IsCI() bool {
	env, ok := os.LookupEnv("CI")
	if !ok {
		return false
	}

	isCI, err := strconv.ParseBool(env)
	if err != nil {
		return false
	}

	return isCI
}
