// Package os can write to stderr and exit
package exit

import (
	"fmt"
	"os"
)

const (
	// DefaultDieExitCode is the default code that Die/Dief will use
	DefaultDieExitCode = 1
)

// Dief provides formatting outputs to write prior to exit
func Dief(format string, a ...any) {
	DieAndExitf(DefaultDieExitCode, format, a...)
}

// Die will write to stderr and exit (1)
func Die(a any) {
	DieAndExit(DefaultDieExitCode, a)
}

// DieAndExitf will format a message and exit with the given code
func DieAndExitf(code int, format string, a ...any) {
	DieAndExit(code, fmt.Sprintf(format, a...))
}

// DieAndExit will write to stderr and exit with the given code
func DieAndExit(code int, a any) {
	if a != nil {
		fmt.Fprintf(os.Stderr, "%v\n", a)
	}
	os.Exit(code)
}
