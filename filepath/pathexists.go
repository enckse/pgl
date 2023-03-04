// Package filepath can tell if paths exist
package filepath

import (
	"errors"
	"os"
)

// PathExists indicates whether a file exists (true) or not (false)
func PathExists(file string) bool {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
