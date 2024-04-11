package closer

import (
	"errors"
	"sync"
)

// LifoCloser provides LIFO closing of all closers. (1,2,3 -> 3,2,1).
type LifoCloser struct {
	closers []Closer
	mu      sync.Mutex
}

// NewLifoCloser creates a new LifoCloser and returns a pointer.
func NewLifoCloser() *LifoCloser {
	return new(LifoCloser)
}

// Add adds closers to list of closers.
func (s *LifoCloser) Add(closers ...Closer) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.closers = append(s.closers, closers...)
}

// Close calls all close method (CloseFunc) from closers in reverse sequence (LIFO).
func (s *LifoCloser) Close() (errs error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := len(s.closers) - 1; i >= 0; i-- {
		closer := s.closers[i]
		if err := closer.Close(); err != nil {
			errs = errors.Join(errs, err)
		}
	}

	return
}
