package paths_test

import (
	"os"
	"path/filepath"
	"testing"

	fp "github.com/enckse/pgl/paths"
)

var testCopyDir = filepath.Join("testdata", "cp")

func setupCopy() {
	os.RemoveAll(testCopyDir)
	os.MkdirAll(testCopyDir, 0o755)
}

func TestCopy(t *testing.T) {
	setupCopy()
	src := filepath.Join(testCopyDir, "copysrc")
	dst := filepath.Join(testCopyDir, "copydst")
	if err := fp.Copy(src, dst, 0o644); err != fp.ErrNoSourceFile {
		t.Errorf("copy error should be no source: %v", err)
	}
	os.WriteFile(src, []byte("test"), 0o644)
	if err := fp.Copy(src, dst, 0o600); err != nil {
		t.Errorf("no copy error should have happened: %v", err)
	}
	stat, _ := os.Stat(dst)
	if stat.Mode() != 0o600 {
		t.Error("invalid copy write mods")
	}
	b, _ := os.ReadFile(dst)
	if string(b) != "test" {
		t.Error("invalid resulting dest file")
	}
	os.WriteFile(src, []byte("test2"), 0o644)
	if err := fp.Copy(src, dst, 0o644); err != fp.ErrDestExists {
		t.Errorf("copy error should be no source: %v", err)
	}
	b, _ = os.ReadFile(dst)
	if string(b) != "test" {
		t.Error("invalid resulting dest file")
	}
}

func TestCopyOverwrite(t *testing.T) {
	setupCopy()
	src := filepath.Join(testCopyDir, "copysrcow")
	dst := filepath.Join(testCopyDir, "copydstow")
	if err := fp.CopyOverwrite(src, dst, 0o644); err != fp.ErrNoSourceFile {
		t.Errorf("copy error should be no source: %v", err)
	}
	os.WriteFile(src, []byte("test"), 0o644)
	if err := fp.CopyOverwrite(src, dst, 0o600); err != nil {
		t.Errorf("no copy error should have happened: %v", err)
	}
	stat, _ := os.Stat(dst)
	if stat.Mode() != 0o600 {
		t.Error("invalid copy write mods")
	}
	b, _ := os.ReadFile(dst)
	if string(b) != "test" {
		t.Error("invalid resulting dest file")
	}
	os.WriteFile(src, []byte("test2"), 0o644)
	if err := fp.CopyOverwrite(src, dst, 0o644); err != fp.ErrModeMismatch {
		t.Errorf("copy error should have happened, mode: %v", err)
	}
	if err := fp.CopyOverwrite(src, dst, 0o600); err != nil {
		t.Errorf("no copy error should have happened: %v", err)
	}
	b, _ = os.ReadFile(dst)
	if string(b) != "test2" {
		t.Error("invalid resulting dest file")
	}
}
