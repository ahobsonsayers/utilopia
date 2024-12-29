package envutils

import (
	"os"
	"strconv"
	"testing"
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

// SkipIfCI skips the test if the current environment is a CI/CD environment.
func SkipIfCI(t *testing.T, args ...any) {
	if IsCI() {
		t.Skip(args...)
	}
}
