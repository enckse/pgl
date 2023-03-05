// Package maps provides a simple key-based map in which a map
// is maintained along with a set of keys
package maps

type (
	// KeyedMap is wrapper around a map with a synchronized
	// set of keys
	KeyedMap[T comparable, V any] struct {
		data map[T]V
		keys []T
	}
)

// Convert will convert an existing 0-N maps to a KeyedMap
func Convert[T comparable, V any](m ...map[T]V) *KeyedMap[T, V] {
	obj := &KeyedMap[T, V]{}
	for _, item := range m {
		for k, v := range item {
			obj.Add(k, v)
		}
	}
	return obj
}

// Add will add a new key/value to the map
func (m *KeyedMap[T, V]) Add(key T, value V) {
	needAdd := false
	if _, ok := m.Get(key); !ok {
		needAdd = true
	}
	if m.data == nil {
		m.data = make(map[T]V)
	}
	if needAdd {
		m.keys = append(m.keys, key)
	}
	m.data[key] = value
}

// Delete will remove a key from the map
func (m *KeyedMap[T, V]) Delete(key T) {
	if m.data == nil {
		return
	}
	if _, ok := m.Get(key); ok {
		delete(m.data, key)
		rekey := []T{}
		for _, k := range m.Keys() {
			if k == key {
				continue
			}
			rekey = append(rekey, k)
		}
		m.keys = rekey
	}
}

// Get will get the value of a key
func (m *KeyedMap[T, V]) Get(key T) (V, bool) {
	if m.data == nil {
		return *new(V), false
	}
	d, ok := m.data[key]
	return d, ok
}

// Keys will retrieve the map keys
func (m *KeyedMap[T, V]) Keys() []T {
	if m.data == nil {
		return []T{}
	}
	return m.keys
}
