package main

import (
	"github.com/golang-collections/collections/stack"
)

// implements queue
type queue struct {
	newestOnTop *stack.Stack
	oldestOnTop *stack.Stack
}

func newQueue() *queue {
	return &queue{
		newestOnTop: stack.New(),
		oldestOnTop: stack.New(),
	}
}

func (q *queue) Enqueue(v int) { // add item
	q.newestOnTop.Push(v)
}

func (q *queue) Peek() int { // get "oldest" item

	// move everything from newestOnTop to oldestOnTop
	for q.newestOnTop.Len() > 0 {
		q.oldestOnTop.Push(q.newestOnTop.Pop())
	}

	v := q.oldestOnTop.Peek().(int)

	// move everything back to newestOnTop
	for q.oldestOnTop.Len() > 0 {
		q.newestOnTop.Push(q.oldestOnTop.Pop())
	}

	return v
}

func (q *queue) Dequeue() int { // get "oldest" and remove it

	// move everything from newestOnTop to oldestOnTop
	for q.newestOnTop.Len() > 0 {
		q.oldestOnTop.Push(q.newestOnTop.Pop())
	}

	v := q.oldestOnTop.Pop().(int)

	// move everything back to newestOnTop
	for q.oldestOnTop.Len() > 0 {
		q.newestOnTop.Push(q.oldestOnTop.Pop())
	}

	return v
}
