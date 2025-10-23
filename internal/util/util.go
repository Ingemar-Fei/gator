package util

import (
	"os"
	"slices"
)

/* print a debug message when run with --verbose */
func DebugMode() bool {
	return slices.Contains(os.Args, "--verbose")
}
