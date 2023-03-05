package paths_test

import (
	"os"
	"path/filepath"
	"testing"

	fp "github.com/enckse/pgl/os/paths"
)

func TestPathExist(t *testing.T) {
	testDir := filepath.Join("testdata", "exists")
	os.RemoveAll(testDir)
	if fp.Exist(testDir) {
		t.Error("test dir SHOULD NOT exist")
	}
	os.Mkdir(testDir, 0o755)
	if !fp.Exist(testDir) {
		t.Error("test dir SHOULD exist")
	}
}
