// Package collections contains a number of collection types
// that are useful for managing various data
package collections

type (
	void struct{}
	// Set is a simple map-based set via map[T]struct{}
	Set[T comparable] struct {
		Map[T, void]
	}
)

// NewSet will create a new set with 0-N values
func NewSet[T comparable](v ...T) *Set[T] {
	s := &Set[T]{}
	s.Add(v...)
	return s
}

// Add will add an item to the set
func (s *Set[T]) Add(items ...T) {
	for _, i := range items {
		s.Set(i, void{})
	}
}

// Clear will reset the set
func (s *Set[T]) Clear() {
	s.data = nil
}

// Remove will remove an item from the set
func (s *Set[T]) Remove(item T) {
	s.Delete(item)
}

// Contains will indicate if the set contains a value
func (s *Set[T]) Contains(item T) bool {
	_, ok := s.Get(item)
	return ok
}
