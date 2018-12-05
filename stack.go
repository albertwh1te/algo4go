package algo4go

import (
	"sync"
)

//Stack is thread safe FILO datastructure
type Stack struct {
	lock sync.Mutex
	s    []interface{}
}

// NewStack return a new thread safe Stack
func NewStack() *Stack {
	return &Stack{s: []interface{}{}}
}

// Push one element into stack
func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

// Pop last element from stack
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.Empty() {
		return nil
	}

	res := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return res
}

// Empty checks if stack is empty
func (s *Stack) Empty() bool {
	return len(s.s) == 0
}
