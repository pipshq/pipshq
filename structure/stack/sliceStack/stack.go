/*
 * Implements a stack with slice
 * Thread-safe
 */

package sliceStack

import (
	"sync"
)

type stack struct {
	container []interface{}
	mu        sync.RWMutex
}

func New() *stack {
	return &stack{}
}

func (s *stack) Push(v interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.container = append(s.container, v)
}

func (s *stack) Top() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	top := s.container[len(s.container)-1]
	return top
}

// Pop an item from the end of stack
func (s *stack) Pop() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.container) > 0 {
		tail := s.container[len(s.container)-1]
		s.container = s.container[:len(s.container)-1]
		return tail
	}
	return nil
}

// Remove an item from stack, and keep order unchanged.
func (s *stack) Remove(i int) *stack {
	s.mu.Lock()
	defer s.mu.Unlock()
	// use copy() to move the high indexed element
	// forward to override the location of the removed element

	copy(s.container[1:], s.container[i+1:])
	s.container = s.container[:len(s.container)-1]
	return s
}

func (s *stack) Length() int {
	return len(s.container)
}

func (s *stack) Items() []interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.container
}
