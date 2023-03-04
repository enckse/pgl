// Package filepath has the ability to read and copy by path
package filepath

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

// CopyOverwrite will overwrite a file if it exists
func CopyOverwrite(src, dst string, mode fs.FileMode) error {
	return doCopy(src, dst, mode, true)
}

// Copy will copy a file from source to destination, will not overwrite
func Copy(src, dst string, mode fs.FileMode) error {
	return doCopy(src, dst, mode, false)
}

func doCopy(src, dst string, mode fs.FileMode, force bool) error {
	if !PathExists(src) {
		return ErrNoSourceFile
	}

	in, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	if PathExists(dst) {
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
