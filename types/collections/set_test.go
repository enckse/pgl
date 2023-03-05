package collections_test

import (
	"testing"

	set "github.com/enckse/pgl/types/collections"
)

func TestAdd(t *testing.T) {
	s := &set.Set[string]{}
	s.Add("test")
	if !s.Contains("test") {
		t.Error("invalid set")
	}
	s.Add("test")
	if !s.Contains("test") {
		t.Error("invalid set")
	}
	s.Add("test2")
	if !s.Contains("test2") {
		t.Error("invalid set")
	}
}

func TestClear(t *testing.T) {
	s := &set.Set[string]{}
	s.Clear()
	s.Add("test")
	if !s.Contains("test") {
		t.Error("invalid clear")
	}
	s.Add("test2")
	if !s.Contains("test2") {
		t.Error("invalid clear")
	}
	s.Clear()
	if s.Contains("test") || s.Contains("test2") {
		t.Error("invalid clear")
	}
	s.Add("test")
	if !s.Contains("test") || s.Contains("test2") {
		t.Error("invalid clear")
	}
}

func TestRemove(t *testing.T) {
	s := &set.Set[string]{}
	s.Remove("test1")
	s.Add("test")
	if !s.Contains("test") {
		t.Error("invalid remove")
	}
	s.Add("test2")
	if !s.Contains("test2") {
		t.Error("invalid remove")
	}
	s.Remove("a")
	if !s.Contains("test2") || !s.Contains("test") {
		t.Error("invalid remove")
	}
	s.Remove("test")
	if !s.Contains("test2") || s.Contains("test") {
		t.Error("invalid remove")
	}
	s.Remove("test2")
	if s.Contains("test2") {
		t.Error("invalid remove")
	}
}

func TestContains(t *testing.T) {
	s := &set.Set[int]{}
	if s.Contains(1) {
		t.Error("invalid has")
	}
	s.Add(1)
	if !s.Contains(1) {
		t.Error("invalid has")
	}
	s.Clear()
	if s.Contains(1) {
		t.Error("invalid has")
	}
}
