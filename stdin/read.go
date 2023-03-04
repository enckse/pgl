// Package stdin can read one (or more) lines from stdin
package stdin

import (
	"bufio"
	"bytes"
	"os"
)

// ReadAllStdin will read all text from stdin
func ReadAll() ([]byte, error) {
	return read(false)
}

// ReadStdinLine will read one line of stdin input
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
