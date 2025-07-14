package testutils

import (
	"testing"

	"github.com/ahobsonsayers/utilopia/envutils"
)

// SkipIfCI skips the test if the current environment is a CI/CD environment.
func SkipIfCI(t *testing.T, args ...any) {
	if envutils.IsCI() {
		t.Skip(args...)
	}
}
