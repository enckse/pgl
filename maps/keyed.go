// Package maps has a comparable keyed map
package maps

type (
	// KeyedMap contains keys synced with the underlying map keys
	KeyedMap[T comparable] struct {
		data map[T]any
		keys []T
	}
)

// Add will add a new key/value to the map
func (m *KeyedMap[T]) Add(key T, value any) {
	if m == nil {
		return
	}
	needAdd := false
	if _, ok := m.Get(key); !ok {
		needAdd = true
	}
	if m.data == nil {
		m.data = make(map[T]any)
	}
	if needAdd {
		m.keys = append(m.keys, key)
	}
	m.data[key] = value
}

// Delete will remove a key from the map
func (m *KeyedMap[T]) Delete(key T) {
	if !validKeyedMap(m) {
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
func (m *KeyedMap[T]) Get(key T) (any, bool) {
	if !validKeyedMap(m) {
		return nil, false
	}
	d, ok := m.data[key]
	return d, ok
}

// Keys will retrieve the map keys
func (m *KeyedMap[T]) Keys() []T {
	if !validKeyedMap(m) {
		return []T{}
	}
	return m.keys
}

func validKeyedMap[T comparable](m *KeyedMap[T]) bool {
	return m != nil && m.data != nil
}
