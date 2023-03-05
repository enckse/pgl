// Package paths is responsible for pathing operations/commands
package paths

import (
	"errors"
	"io/fs"
	"os"
)

var (
	// ErrNoSourceFile indicates the source file for copy does NOT exist
	ErrNoSourceFile = errors.New("source file does not exist")
	// ErrDestExists indicates that the destination already exists
	ErrDestExists = errors.New("destination file exists")
	// ErrModeMismatch means that the destination overwrite would be changing mode
	ErrModeMismatch = errors.New("overwrite will change mode")
)

// CopyOverwrite is the same as Copy however it will overwrite the
// destination UNLESS the file mode is different than what was given
func CopyOverwrite(src, dst string, mode fs.FileMode) error {
	return doCopy(src, dst, mode, true)
}

// Copy will copy a file from source to destination but
// will not overwrite
func Copy(src, dst string, mode fs.FileMode) error {
	return doCopy(src, dst, mode, false)
}

func doCopy(src, dst string, mode fs.FileMode, force bool) error {
	if !Exist(src) {
		return ErrNoSourceFile
	}

	in, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	if Exist(dst) {
		if !force {
			return ErrDestExists
		}
		stat, err := os.Stat(dst)
		if err != nil {
			return err
		}
		if stat.Mode() != mode {
			return ErrModeMismatch
		}
	}

	if err := os.WriteFile(dst, in, mode); err != nil {
		return err
	}

	return nil
}
