// Package collections contains a number of collection types
// that are useful for managing various data
package collections

type (
	// Set is a simple map-based set via map[T]struct{}
	Set[T comparable] struct {
		data map[T]struct{}
	}
)

// Add will add an item to the set
func (s *Set[T]) Add(item T) {
	if s.data == nil {
		s.data = make(map[T]struct{})
	}
	s.data[item] = struct{}{}
}

// Clear will reset the set
func (s *Set[T]) Clear() {
	s.data = nil
}

// Remove will remove an item from the set
func (s *Set[T]) Remove(item T) {
	if !s.Contains(item) {
		return
	}
	delete(s.data, item)
}

// Contains will indicate if the set contains a value
func (s *Set[T]) Contains(item T) bool {
	if s.data == nil {
		return false
	}
	_, ok := s.data[item]
	return ok
}
