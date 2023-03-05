// Package paths is responsible for pathing operations/commands
package paths

import (
	"errors"
	"os"
)

// Exist indicates whether a path exists (true) or not (false)
func Exist(file string) bool {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
