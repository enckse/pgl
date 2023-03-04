package version_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/enckse/pgl/version"
)

var testDirGitVersion = "testdata"

type (
	mockManager struct {
		tag  string
		err  error
		vers string
		noop string
	}
)

func (m *mockManager) Tag() string {
	return m.tag
}

func (m *mockManager) Write(vers string) {
	m.vers = vers
}

func (m *mockManager) Noop(msg string) {
	m.noop = msg
}

func (m *mockManager) Error(err error) {
	m.err = err
}

func (m *mockManager) GitRoot() string {
	return testDirGitVersion
}

func setupGitVersion() {
	os.RemoveAll(testDirGitVersion)
	os.Mkdir(testDirGitVersion, 0o755)
}

func TestGitVersion(t *testing.T) {
	setupGitVersion()
	m := &mockManager{}
	m.err = errors.New("bad tag")
	version.Git(m)
	if m.err != version.ErrInvalidTagSpace {
		t.Errorf("invalid tag error: %v", m.err)
	}
	m.err = nil
	m.tag = "aa"
	version.Git(m)
	if m.err != version.ErrInvalidInputTag {
		t.Errorf("invalid tag error: %v", m.err)
	}
	m.tag = "v12.34.56"
	version.Git(m)
	if m.err != version.ErrInvalidInputTag {
		t.Errorf("invalid tag error: %v", m.err)
	}
	current := time.Now().Format("v06.01.")
	for _, invalid := range []string{"", "000", "00.", "00.11", "0"} {
		m.err = nil
		m.tag = current + invalid
		version.Git(m)
		if m.err != version.ErrInvalidInputTag {
			t.Errorf("invalid tag error: %v", m.err)
		}
	}
	m.err = nil
	m.tag = current + "a1"
	version.Git(m)
	if m.err.Error() != "strconv.Atoi: parsing \"a1\": invalid syntax" {
		t.Errorf("invalid tag error: %v", m.err)
	}
	m.err = nil
	m.tag = current + "-9"
	version.Git(m)
	if m.err != version.ErrInvalidMinorTagZero {
		t.Errorf("invalid tag error: %v", m.err)
	}
	m.err = nil
	m.tag = current + "99"
	version.Git(m)
	if m.err != version.ErrInvalidMinorTagExceeds {
		t.Errorf("invalid tag error: %v", m.err)
	}
	m.err = nil
	m.tag = current + "01"
	version.Git(m)
	if m.err != nil || m.vers != current+"02" {
		t.Errorf("invalid tag error or tag: %v, %s", m.err, m.vers)
	}
	m.err = nil
	m.tag = current + "09"
	version.Git(m)
	if m.err != nil || m.vers != current+"10" {
		t.Errorf("invalid tag error or tag: %v, %s", m.err, m.vers)
	}
	m.err = nil
	m.tag = current + "98"
	version.Git(m)
	if m.err != nil || m.vers != current+"99" {
		t.Errorf("invalid tag error or tag: %v, %s", m.err, m.vers)
	}
	m.err = nil
	m.tag = "v22.00.98"
	version.Git(m)
	if m.err != nil || m.vers != current+"00" {
		t.Errorf("invalid tag error or tag: %v, %s", m.err, m.vers)
	}
	m.err = nil
	m.tag = time.Now().Format("v06") + ".00.98"
	version.Git(m)
	if m.err != nil || m.vers != current+"00" {
		t.Errorf("invalid tag error or tag: %v, %s", m.err, m.vers)
	}
}

func TestGitNoop(t *testing.T) {
	os.RemoveAll(testDirGitVersion)
	m := &mockManager{}
	version.Git(m)
	if m.noop != "no git root found" {
		t.Error("invalid noop")
	}
	m.err = errors.New("mock")
	m.noop = ""
	setupGitVersion()
	version.Git(m)
	if m.noop != "" {
		t.Error("invalid noop")
	}
}
