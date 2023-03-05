// Package paths is responsible for pathing operations/commands
package paths

import (
	"errors"
	"os"
)

// Exists indicates whether a path exists (true) or not (false)
func Exists(file string) bool {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
