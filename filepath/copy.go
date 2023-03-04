// Package filepath has the ability to read and copy by path
package filepath

import (
	"fmt"
	"io/fs"
	"os"
)

// CopyOverwrite will overwrite a file if it exists
func CopyOverwrite(src, dst string, mode fs.FileMode) error {
	return doCopy(src, dst, mode, true)
}

// Copy will copy a file from source to destination, will not overwrite
func Copy(src, dst string, mode fs.FileMode) error {
	return doCopy(src, dst, mode, true)
}

func doCopy(src, dst string, mode fs.FileMode, force bool) error {
	if !PathExists(src) {
		return fmt.Errorf("source file '%s' does not exist", src)
	}

	in, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	if err := os.WriteFile(dst, in, mode); err != nil {
		return err
	}

	return nil
}
