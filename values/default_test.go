package values_test

import (
	"testing"

	"github.com/enckse/pgl/values"
)

func TestIfNotSet(t *testing.T) {
	if values.IfNotSet("", "test") != "test" {
		t.Error("invalid default")
	}
	if values.IfNotSet("given", "test") != "given" {
		t.Error("invalid default")
	}
	if values.IfNotSet(0, 100) != 100 {
		t.Error("invalid default")
	}
	if values.IfNotSet(10, 100) != 10 {
		t.Error("invalid default")
	}
}
