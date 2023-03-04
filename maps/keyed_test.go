package maps_test

import (
	"fmt"
	"testing"

	"github.com/enckse/pgl/maps"
)

func TestAddKeyValue(t *testing.T) {
	k := &maps.KeyedMap[string]{}
	maps.AddKeyValue(nil, "", nil)
	maps.AddKeyValue(k, "test", nil)
	if fmt.Sprintf("%v", maps.GetKeys(k)) != "[test]" {
		t.Error(maps.GetKeys(k))
		t.Error("invalid map")
	}
	v, ok := maps.GetKeyValue(k, "test")
	if !ok || v != nil {
		t.Error("invalid get")
	}
	maps.AddKeyValue(k, "test", 1)
	if fmt.Sprintf("%v", maps.GetKeys(k)) != "[test]" {
		t.Error(maps.GetKeys(k))
		t.Error("invalid map")
	}
	v, ok = maps.GetKeyValue(k, "test")
	if !ok || v != 1 {
		t.Error("invalid get")
	}
}

func TestDeleteKeyValue(t *testing.T) {
	k := &maps.KeyedMap[string]{}
	maps.DeleteKey(nil, struct{}{})
	maps.DeleteKey(k, "")
	maps.AddKeyValue(k, "test", nil)
	maps.AddKeyValue(k, "test2", 2)
	maps.DeleteKey(k, "test")
	if fmt.Sprintf("%v", maps.GetKeys(k)) != "[test2]" {
		t.Error("invalid map")
	}
	_, ok := maps.GetKeyValue(k, "test")
	if ok {
		t.Error("invalid key")
	}
	val, ok := maps.GetKeyValue(k, "test2")
	if !ok || val != 2 {
		t.Error("invalid key")
	}
}

func TestGetKeyValue(t *testing.T) {
	maps.GetKeyValue(nil, 1)
	k := &maps.KeyedMap[string]{}
	maps.AddKeyValue(k, "test", "TEST")
	maps.AddKeyValue(k, "test2", 2)
	val, ok := maps.GetKeyValue(k, "test")
	if !ok || val != "TEST" {
		t.Error("invalid key")
	}
	val, ok = maps.GetKeyValue(k, "test2")
	if !ok || val != 2 {
		t.Error("invalid key")
	}
	if _, ok := maps.GetKeyValue(k, "invalid"); ok {
		t.Error("invalid key")
	}
}

func TestGetKeys(t *testing.T) {
	if len(maps.GetKeys[string](nil)) != 0 {
		t.Error("invalid get")
	}
	if len(maps.GetKeys(&maps.KeyedMap[string]{})) != 0 {
		t.Error("invalid get")
	}
	k := &maps.KeyedMap[string]{}
	maps.AddKeyValue(k, "test2", 2)
	maps.AddKeyValue(k, "test", "TEST")
	if fmt.Sprintf("%v", maps.GetKeys(k)) != "[test2 test]" {
		t.Error("invalid map")
	}
}
