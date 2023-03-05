// Package set provides a simple set wrapper for maps
package set

type (
	// MapBased is a simple map-based set via map[T]struct{}
	MapBased[T comparable] struct {
		data map[T]struct{}
	}
)

// Add will add an item to the set
func (s *MapBased[T]) Add(item T) {
	if s.data == nil {
		s.data = make(map[T]struct{})
	}
	s.data[item] = struct{}{}
}

// Clear will reset the set
func (s *MapBased[T]) Clear() {
	s.data = nil
}

// Remove will remove an item from the set
func (s *MapBased[T]) Remove(item T) {
	if !s.Has(item) {
		return
	}
	delete(s.data, item)
}

// Has will indicate if the set contains a value
func (s *MapBased[T]) Has(item T) bool {
	if s.data == nil {
		return false
	}
	_, ok := s.data[item]
	return ok
}
