package lazyQueue

import "github.com/sagottlieb/hackerrank/tale2Stacks/core"

type queue struct {
	// except during operations, only one of these will contain entries
	// move entries between them as necessary to complete ops
	newestOnTop core.Stack
	oldestOnTop core.Stack
}

func New(newStack core.StackIniter) core.Queue {
	return &queue{
		newestOnTop: newStack(),
		oldestOnTop: newStack(),
	}
}

func (q *queue) Enqueue(v int) { // add item
	// before we can push, we need the data to be in newestOnTop

	for q.oldestOnTop.Len() > 0 {
		q.newestOnTop.Push(q.oldestOnTop.Pop())
	}

	q.newestOnTop.Push(v)
}

func (q *queue) Peek() int { // get "oldest" item
	// before we can peek, we need the data to be in oldestOnTop

	for q.newestOnTop.Len() > 0 {
		q.oldestOnTop.Push(q.newestOnTop.Pop())
	}

	return q.oldestOnTop.Peek()
}

func (q *queue) Dequeue() int { // get "oldest" and remove it
	// before we can pop, we need the data to be in oldestOnTop

	for q.newestOnTop.Len() > 0 {
		q.oldestOnTop.Push(q.newestOnTop.Pop())
	}

	return q.oldestOnTop.Pop()
}
