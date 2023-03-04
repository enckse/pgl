// Package version can help manage git-based versions
package version

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/enckse/pgl/filepath"
	o "github.com/enckse/pgl/os"
)

type (
	// VersionManager is used to setup version information
	VersionManager interface {
		Tag() (string, error)
		Error(error)
		Noop(string)
		Write(string) error
	}
	DefaultVersionManager struct{}
)

//go:embed vers.txt
var Version string

// Error will die/exit
func (d DefaultVersionManager) Error(err error) {
	o.Die(err)
}

// Tag will get the tag via git describe
func (d DefaultVersionManager) Tag() (string, error) {
	b, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	if err != nil {
		return "", err
	}
	return string(b), err
}

// Write will simply write the tag to stdout
func (d DefaultVersionManager) Write(tag string) error {
	args := os.Args
	if len(args) != 2 {
		return errors.New("invalid arguments, must be: '<cmd> <file>'")
	}
	return os.WriteFile(args[1], []byte(tag), 0o644)
}

// Noop writes the message to stderr
func (d DefaultVersionManager) Noop(message string) {
	fmt.Fprintln(os.Stderr, message)
}

// DefaultGitVersion is a wrapper around default git versioning settings
func DefaultGitVersion() {
	Git(DefaultVersionManager{})
}

var (
	// ErrInvalidInputTag indicates the tag is malformed
	ErrInvalidInputTag = errors.New("current version tag is malformed")
	// ErrInvalidInputTagMinor means the minor component is < 0
	ErrInvalidInputTagMinor = fmt.Errorf("%w: minor component less than 0", ErrInvalidInputTag)
	// ErrMaxMinorTag indicates that somehow the minor version is > 99
	ErrMaxMinorTag = fmt.Errorf("%w: maximum minor tag value exceeded", ErrInvalidInputTag)
	// ErrInvalidInputTagSpace indicates the tag was only whitespace
	ErrInvalidInputTagSpace = fmt.Errorf("%w: found only whitespace", ErrInvalidInputTag)
)

// Git handles versioning for git-based repos
func Git(v VersionManager) {
	if !filepath.PathExists(".git") {
		v.Noop("not git controlled")
		return
	}
	tag, err := v.Tag()
	if err != nil {
		v.Error(err)
		return
	}
	trimmed := strings.TrimSpace(tag)
	if trimmed == "" {
		v.Error(ErrInvalidInputTagSpace)
		return
	}
	tag = trimmed
	currentVersion := time.Now().Format("v06.01.")
	var minor uint
	if strings.HasPrefix(tag, currentVersion) {
		parts := strings.Split(tag, ".")
		if len(parts) != 3 {
			v.Error(ErrInvalidInputTag)
			return
		}
		minorVer := strings.TrimPrefix(parts[2], "0")
		converted, err := strconv.Atoi(minorVer)
		if err != nil {
			v.Error(err)
			return
		}
		if converted < 0 {
			v.Error(ErrInvalidInputTagMinor)
			return
		}
		minor = uint(converted) + 1
	}
	if minor > 99 {
		v.Error(ErrMaxMinorTag)
		return
	}
	if err := v.Write(fmt.Sprintf("%s%02d", currentVersion, minor)); err != nil {
		v.Error(err)
	}
}
