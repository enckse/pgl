// Package maps has a comparable keyed map
package maps

type (
	// KeyedMap contains keys synced with the underlying map keys
	KeyedMap[T comparable] struct {
		data map[T]any
		keys []T
	}
)

// AddKeyValue will add a new key/value to the map
func AddKeyValue[T comparable](m *KeyedMap[T], key T, value any) {
	if m == nil {
		return
	}
	needAdd := false
	if _, ok := GetKeyValue(m, key); !ok {
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

// DeleteKey will remove a key from the map
func DeleteKey[T comparable](m *KeyedMap[T], key T) {
	if !validKeyedMap(m) {
		return
	}
	if _, ok := GetKeyValue(m, key); ok {
		delete(m.data, key)
		rekey := []T{}
		for _, k := range GetKeys(m) {
			if k == key {
				continue
			}
			rekey = append(rekey, k)
		}
		m.keys = rekey
	}
}

// GetKeyValue will get the value of a key
func GetKeyValue[T comparable](m *KeyedMap[T], key T) (any, bool) {
	if !validKeyedMap(m) {
		return nil, false
	}
	d, ok := m.data[key]
	return d, ok
}

// GetKeys will retrieve the map keys
func GetKeys[T comparable](m *KeyedMap[T]) []T {
	if !validKeyedMap(m) {
		return []T{}
	}
	return m.keys
}

func validKeyedMap[T comparable](m *KeyedMap[T]) bool {
	return m != nil && m.data != nil
}
