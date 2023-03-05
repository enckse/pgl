// Package exit handles wrappers around os.Exit via output formatting
// and exit codes
package exit

import (
	"fmt"
	"os"
)

const (
	// DefaultDieExitCode is the default code that Die/Dief will use
	// when calling os.Exit
	DefaultDieExitCode = 1
)

// Dief will write to stderr a formatted message string and
// exit using the default exit code
func Dief(format string, a ...any) {
	DieAndExitf(DefaultDieExitCode, format, a...)
}

// Die will write to stderr a non-nil 0-N input set and then exit using
// the default exit code
func Die(a ...any) {
	DieAndExit(DefaultDieExitCode, a...)
}

// DieAndExitf will write to stderr a formatted message string and
// exit using the given exit code
func DieAndExitf(code int, format string, a ...any) {
	DieAndExit(code, fmt.Sprintf(format, a...))
}

// DieAndExit will write to stderr a non-nil 0-N input set and then exit
// using the given exit code
func DieAndExit(code int, a ...any) {
	for _, item := range a {
		if item == nil {
			continue
		}
		fmt.Fprintf(os.Stderr, "%v\n", item)
	}
	os.Exit(code)
}
