package index

import (
	"sync"
)

// Set holds a group of indexes
type Set struct {
	indexes map[string]Backend
	sync.RWMutex
}

// CreateSet creates a set of indexes
func CreateSet() *Set {
	return &Set{indexes: make(map[string]Backend)}
}

// Get retrieves a specific index
func (s *Set) Get(index string) Backend {
	s.RLock()
	defer s.RUnlock()

	return s.indexes[index]
}

// Add adds an index to the set
func (s *Set) Add(index string, backend Backend) {
	s.Lock()
	defer s.Unlock()

	s.indexes[index] = backend
}
