// Package io can read one (or more) lines from stdin
package io

import (
	"bufio"
	"bytes"
	"os"
)

// ReadAllStdin will read all text from stdin
func ReadAllStdin() ([]byte, error) {
	return readStdin(false)
}

// ReadStdinLine will read one line of stdin input
func ReadStdinLine() ([]byte, error) {
	return readStdin(true)
}

func readStdin(one bool) ([]byte, error) {
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
