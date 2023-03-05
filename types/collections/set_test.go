package collections_test

import (
	"testing"

	set "github.com/enckse/pgl/types/collections"
)

func TestAdd(t *testing.T) {
	s := &set.Set[string]{}
	s.Add("test")
	if !s.Has("test") {
		t.Error("invalid set")
	}
	s.Add("test")
	if !s.Has("test") {
		t.Error("invalid set")
	}
	s.Add("test2")
	if !s.Has("test2") {
		t.Error("invalid set")
	}
}

func TestClear(t *testing.T) {
	s := &set.Set[string]{}
	s.Clear()
	s.Add("test")
	if !s.Has("test") {
		t.Error("invalid clear")
	}
	s.Add("test2")
	if !s.Has("test2") {
		t.Error("invalid clear")
	}
	s.Clear()
	if s.Has("test") || s.Has("test2") {
		t.Error("invalid clear")
	}
	s.Add("test")
	if !s.Has("test") || s.Has("test2") {
		t.Error("invalid clear")
	}
}

func TestRemove(t *testing.T) {
	s := &set.Set[string]{}
	s.Remove("test1")
	s.Add("test")
	if !s.Has("test") {
		t.Error("invalid remove")
	}
	s.Add("test2")
	if !s.Has("test2") {
		t.Error("invalid remove")
	}
	s.Remove("a")
	if !s.Has("test2") || !s.Has("test") {
		t.Error("invalid remove")
	}
	s.Remove("test")
	if !s.Has("test2") || s.Has("test") {
		t.Error("invalid remove")
	}
	s.Remove("test2")
	if s.Has("test2") {
		t.Error("invalid remove")
	}
}

func TestHas(t *testing.T) {
	s := &set.Set[int]{}
	if s.Has(1) {
		t.Error("invalid has")
	}
	s.Add(1)
	if !s.Has(1) {
		t.Error("invalid has")
	}
	s.Clear()
	if s.Has(1) {
		t.Error("invalid has")
	}
}