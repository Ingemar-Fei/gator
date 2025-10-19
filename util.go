package main

import (
	"fmt"
	"os"
	"slices"
)

/* print a debug message when run with --verbose */
func DebugMessage(s string, a ...any) {
	if slices.Contains(os.Args, "--verbose") {
		fmt.Printf(s, a...)
	}
}
