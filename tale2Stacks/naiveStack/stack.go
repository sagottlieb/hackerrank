package naiveStack

import "github.com/sagottlieb/hackerrank/tale2Stacks/core"

// implements stack
type stack struct {
	data []int
}

func New() core.Stack {
	return new(stack)
}

func (s *stack) Len() int {
	return len(s.data)
}

func (s *stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *stack) Pop() int {
	if len(s.data) == 0 {
		panic("Attempted to pop from empty stack")
	}

	v := s.data[len(s.data)-1]

	s.data = s.data[:len(s.data)-1]

	return v
}

func (s *stack) Peek() int {
	if len(s.data) == 0 {
		panic("Attempted to peek on empty stack")
	}

	return s.data[len(s.data)-1]
}
