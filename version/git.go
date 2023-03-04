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

	"github.com/enckse/pgl/exit"
	"github.com/enckse/pgl/paths"
)

type (
	// GitManager is used to setup version information
	GitManager interface {
		Tag() string
		Error(error)
		Noop(string)
		Write(string)
		GitRoot() string
	}
	// DefaultGitManager is the default manager for most applications
	DefaultGitManager struct{}
)

// Error will die/exit
func (d DefaultGitManager) Error(err error) {
	exit.Die(err)
}

// Tag will get the tag via git describe
func (d DefaultGitManager) Tag() string {
	b, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	if err != nil {
		d.Error(err)
		return ""
	}
	return string(b)
}

// Write will simply write the tag to stdout
func (d DefaultGitManager) Write(tag string) {
	args := os.Args
	if len(args) != 2 {
		d.Error(errors.New("invalid arguments, must be: '<cmd> <file>'"))
		return
	}
	if err := os.WriteFile(args[1], []byte(tag), 0o644); err != nil {
		d.Error(err)
		return
	}
}

// Noop writes the message to stderr
func (d DefaultGitManager) Noop(message string) {
	fmt.Fprintln(os.Stderr, message)
}

// GitRoot gets the root directory where the git directory is
func (d DefaultGitManager) GitRoot() string {
	return ".git"
}

// DefaultGitVersion is a wrapper around default git versioning settings
func DefaultGitVersion() {
	Git(DefaultGitManager{})
}

var (
	errInvalidTag = errors.New("current version tag is malformed: ")
	// ErrInvalidMinorTagZero means the parsed tag had a minor component below 0
	ErrInvalidMinorTagZero = fmt.Errorf("%wminor component less than 0", errInvalidTag)
	// ErrInvalidMinorTagExceeds means the parsed minor tag +1 is 100 (or more)
	ErrInvalidMinorTagExceeds = fmt.Errorf("%wmaximum minor tag value exceeded", errInvalidTag)
	// ErrInvalidTagSpace means the tag was only whitespace
	ErrInvalidTagSpace = fmt.Errorf("%wfound only whitespace", errInvalidTag)
	// ErrInvalidInputTag means the input tag was invalid (wrong length, values, etc.)
	ErrInvalidInputTag = fmt.Errorf("%wtag format is invalid", errInvalidTag)
	// Version holds the internal version of pgl
	//go:embed "vers.txt"
	Version string
)

const (
	timeFormat = "v06.01."
)

// Git handles versioning for git-based repos
func Git(v GitManager) {
	if !paths.Exists(v.GitRoot()) {
		v.Noop("no git root found")
		return
	}
	tag := strings.TrimSpace(v.Tag())
	if tag == "" {
		v.Error(ErrInvalidTagSpace)
		return
	}
	currentVersion := time.Now().Format(timeFormat)
	if len(tag) != len(timeFormat)+2 {
		v.Error(ErrInvalidInputTag)
		return
	}
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
	if !strings.HasPrefix(tag, currentVersion) {
		v.Write(fmt.Sprintf("%s00", currentVersion))
		return
	}
	if converted < 0 {
		v.Error(ErrInvalidMinorTagZero)
		return
	}
	minor := uint(converted) + 1
	if minor > 99 {
		v.Error(ErrInvalidMinorTagExceeds)
		return
	}
	v.Write(fmt.Sprintf("%s%02d", currentVersion, minor))
}
