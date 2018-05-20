package otsStack

import (
	"github.com/golang-collections/collections/stack"
	"github.com/sagottlieb/hackerrank/tale2Stacks/core"
)

// For benchmarking fun; wrapper was a quick way to deal with different return types of stack implementations

type wrappedStack struct {
	s *stack.Stack
}

func New() core.Stack {
	return wrappedStack{
		s: stack.New(),
	}
}

func (ws wrappedStack) Len() int {
	return ws.s.Len()
}

func (ws wrappedStack) Push(v int) {
	ws.s.Push(v)
}

func (ws wrappedStack) Pop() int {
	return ws.s.Pop().(int)
}

func (ws wrappedStack) Peek() int {
	return ws.s.Peek().(int)
}
