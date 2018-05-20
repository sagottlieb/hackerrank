package linkedListStack

import "github.com/sagottlieb/hackerrank/tale2Stacks/core"

type stack struct {
	top  *entry
	size int
}

type entry struct {
	val      int
	previous *entry
}

func New() core.Stack {
	return new(stack)
}

func (s *stack) Len() int {
	return s.size
}

func (s *stack) Push(v int) {
	newTop := entry{
		val:      v,
		previous: s.top,
	}

	s.top = &newTop
	s.size += 1
}

func (s *stack) Pop() int {
	if s.size == 0 {
		panic("Attempted to pop from empty stack")
	}

	v := s.top.val

	s.top = s.top.previous
	s.size -= 1

	return v
}

func (s *stack) Peek() int {
	if s.size == 0 {
		panic("Attempted to peek on empty stack")
	}

	return s.top.val
}
