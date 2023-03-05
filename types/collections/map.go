// Package collections contains a number of collection types
// that are useful for managing various data
package collections

type (
	// Map is wrapper around a map with a synchronized
	// set of keys
	Map[T comparable, V any] struct {
		data map[T]V
		keys []T
	}
)

// Convert will convert an existing 0-N maps to a Map
func Convert[T comparable, V any](m ...map[T]V) *Map[T, V] {
	obj := &Map[T, V]{}
	for _, item := range m {
		for k, v := range item {
			obj.Set(k, v)
		}
	}
	return obj
}

// Set will add a new key/value to the map
func (m *Map[T, V]) Set(key T, value V) {
	needSet := false
	if _, ok := m.Get(key); !ok {
		needSet = true
	}
	if m.data == nil {
		m.data = make(map[T]V)
	}
	if needSet {
		m.keys = append(m.keys, key)
	}
	m.data[key] = value
}

// Delete will remove a key from the map
func (m *Map[T, V]) Delete(key T) {
	if m.data == nil {
		return
	}
	if _, ok := m.Get(key); ok {
		delete(m.data, key)
		deleting := -1
		for idx, k := range m.Keys() {
			if k == key {
				deleting = idx
				break
			}
		}
		if deleting >= 0 {
			m.keys = append(m.keys[:deleting], m.keys[deleting+1:]...)
		}
	}
}

// Get will get the value of a key
func (m *Map[T, V]) Get(key T) (V, bool) {
	if m.data == nil {
		return *new(V), false
	}
	d, ok := m.data[key]
	return d, ok
}

// Keys will retrieve the map keys
func (m *Map[T, V]) Keys() []T {
	if m.data == nil {
		return []T{}
	}
	return m.keys
}

// Count will get the count of items in the map
func (m *Map[T, V]) Count() int {
	return len(m.keys)
}
