package collections_test

import (
	"fmt"
	"testing"

	maps "github.com/enckse/pgl/types/collections"
)

func TestSetKeyValue(t *testing.T) {
	k := &maps.Map[string, int]{}
	k.Set("test", 0)
	if fmt.Sprintf("%v", k.Keys()) != "[test]" {
		t.Error("invalid map")
	}
	v, ok := k.Get("test")
	if !ok || v != 0 {
		t.Error("invalid get")
	}
	k.Set("test", 1)
	if fmt.Sprintf("%v", k.Keys()) != "[test]" {
		t.Error("invalid map")
	}
	v, ok = k.Get("test")
	if !ok || v != 1 {
		t.Error("invalid get")
	}
}

func TestDeleteKeyValue(t *testing.T) {
	k := &maps.Map[string, string]{}
	k.Delete("")
	k.Set("test", "abc")
	k.Set("test2", "cde")
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
	k.Set("test3", "abc")
	k.Set("test4", "cde")
	for _, key := range []string{"test4", "test3", "test2"} {
		k.Delete(key)
	}
}

func TestGetKeyValue(t *testing.T) {
	k := &maps.Map[string, bool]{}
	k.Set("test", true)
	k.Set("test2", false)
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
	k := &maps.Map[string, uint]{}
	if len(k.Keys()) != 0 {
		t.Error("invalid get")
	}
	k.Set("test2", 2)
	k.Set("test", 5)
	if fmt.Sprintf("%v", k.Keys()) != "[test2 test]" {
		t.Error("invalid map")
	}
}

func TestConvert(t *testing.T) {
	k := maps.Convert[string, int]()
	if k == nil || len(k.Keys()) != 0 {
		t.Error("invalid new keyed map")
	}
	nm := maps.Convert(map[string]int{})
	if nm == nil || len(nm.Keys()) != 0 {
		t.Error("invalid new keyed map")
	}
	nm = maps.Convert(map[string]int{}, nil)
	if nm == nil || len(nm.Keys()) != 0 {
		t.Error("invalid new keyed map")
	}
	nm = maps.Convert(map[string]int{"test": 1}, nil, map[string]int{"test": 2, "test2": 3})
	if nm == nil || len(nm.Keys()) != 2 {
		t.Error("invalid new keyed map")
	}
	if v, ok := nm.Get("test"); !ok || v != 2 {
		t.Error("invalid map")
	}
	if v, ok := nm.Get("test2"); !ok || v != 3 {
		t.Error("invalid map")
	}
}

func TestMapCount(t *testing.T) {
	k := &maps.Map[string, uint]{}
	if k.Count() != 0 {
		t.Error("invalid count")
	}
	k.Set("test2", 2)
	k.Set("test", 5)
	if k.Count() != 2 {
		t.Error("invalid count")
	}
	k.Set("test2", 2)
	if k.Count() != 2 {
		t.Error("invalid count")
	}
	k.Set("test3", 2)
	if k.Count() != 3 {
		t.Error("invalid count")
	}
	k.Delete("test")
	k.Delete("test2")
	if k.Count() != 1 {
		t.Error("invalid count")
	}
}
