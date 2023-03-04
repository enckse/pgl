package paths_test

import (
	"os"
	"path/filepath"
	"testing"

	fp "github.com/enckse/pgl/paths"
)

func TestPathExists(t *testing.T) {
	testDir := filepath.Join("testdata", "exists")
	os.RemoveAll(testDir)
	if fp.Exists(testDir) {
		t.Error("test dir SHOULD NOT exist")
	}
	os.Mkdir(testDir, 0o755)
	if !fp.Exists(testDir) {
		t.Error("test dir SHOULD exist")
	}
}
