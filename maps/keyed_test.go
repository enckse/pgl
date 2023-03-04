package maps_test

import (
	"fmt"
	"testing"

	"github.com/enckse/pgl/maps"
)

func TestAddKeyValue(t *testing.T) {
	k := &maps.KeyedMap[string, int]{}
	k.Add("test", 0)
	if fmt.Sprintf("%v", k.Keys()) != "[test]" {
		t.Error("invalid map")
	}
	v, ok := k.Get("test")
	if !ok || v != 0 {
		t.Error("invalid get")
	}
	k.Add("test", 1)
	if fmt.Sprintf("%v", k.Keys()) != "[test]" {
		t.Error("invalid map")
	}
	v, ok = k.Get("test")
	if !ok || v != 1 {
		t.Error("invalid get")
	}
}

func TestDeleteKeyValue(t *testing.T) {
	k := &maps.KeyedMap[string, string]{}
	k.Delete("")
	k.Add("test", "abc")
	k.Add("test2", "cde")
	k.Delete("test")
	if fmt.Sprintf("%v", k.Keys()) != "[test2]" {
		t.Error("invalid map")
	}
	_, ok := k.Get("test")
	if ok {
		t.Error("invalid key")
	}
	val, ok := k.Get("test2")
	if !ok || val != "cde" {
		t.Error("invalid key")
	}
}

func TestGetKeyValue(t *testing.T) {
	k := &maps.KeyedMap[string, bool]{}
	k.Add("test", true)
	k.Add("test2", false)
	val, ok := k.Get("test")
	if !ok || !val {
		t.Error("invalid key")
	}
	val, ok = k.Get("test2")
	if !ok || val {
		t.Error("invalid key")
	}
	if _, ok := k.Get("invalid"); ok {
		t.Error("invalid key")
	}
}

func TestGetKeys(t *testing.T) {
	k := &maps.KeyedMap[string, uint]{}
	if len(k.Keys()) != 0 {
		t.Error("invalid get")
	}
	k.Add("test2", 2)
	k.Add("test", 5)
	if fmt.Sprintf("%v", k.Keys()) != "[test2 test]" {
		t.Error("invalid map")
	}
}
