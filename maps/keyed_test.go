package maps_test

import (
	"fmt"
	"testing"

	"github.com/enckse/pgl/maps"
)

func TestAddKeyValue(t *testing.T) {
	k := &maps.KeyedMap[string]{}
	k.Add("test", nil)
	if fmt.Sprintf("%v", k.Keys()) != "[test]" {
		t.Error("invalid map")
	}
	v, ok := k.Get("test")
	if !ok || v != nil {
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
	k := &maps.KeyedMap[string]{}
	k.Delete("")
	k.Add("test", nil)
	k.Add("test2", 2)
	k.Delete("test")
	if fmt.Sprintf("%v", k.Keys()) != "[test2]" {
		t.Error("invalid map")
	}
	_, ok := k.Get("test")
	if ok {
		t.Error("invalid key")
	}
	val, ok := k.Get("test2")
	if !ok || val != 2 {
		t.Error("invalid key")
	}
}

func TestGetKeyValue(t *testing.T) {
	k := &maps.KeyedMap[string]{}
	k.Add("test", "TEST")
	k.Add("test2", 2)
	val, ok := k.Get("test")
	if !ok || val != "TEST" {
		t.Error("invalid key")
	}
	val, ok = k.Get("test2")
	if !ok || val != 2 {
		t.Error("invalid key")
	}
	if _, ok := k.Get("invalid"); ok {
		t.Error("invalid key")
	}
}

func TestGetKeys(t *testing.T) {
	k := &maps.KeyedMap[string]{}
	if len(k.Keys()) != 0 {
		t.Error("invalid get")
	}
	k.Add("test2", 2)
	k.Add("test", "TEST")
	if fmt.Sprintf("%v", k.Keys()) != "[test2 test]" {
		t.Error("invalid map")
	}
}
