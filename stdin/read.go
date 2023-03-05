// Package stdin is responsible for reading the stdin buffer
// and reading 0-N lines from it
package stdin

import (
	"bufio"
	"bytes"
	"os"
)

// ReadAll will read all stdin lines with an appended
// newline at the end of each read line
func ReadAll() ([]byte, error) {
	return read(false)
}

// ReadLine will read a single line from stdin with
// an appended newline at the end of the read line
func ReadLine() ([]byte, error) {
	return read(true)
}

func read(one bool) ([]byte, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var b bytes.Buffer
	for scanner.Scan() {
		b.WriteString(scanner.Text())
		b.WriteString("\n")
		if one {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
