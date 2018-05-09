package constantDequeue

import (
	"github.com/sagottlieb/hackerrank/queues/tale2Stacks/core"
)

type queue struct {
	newestOnTop core.Stack // empty except during enqueue ops
	oldestOnTop core.Stack // source of truth
}

func New(newStack core.StackIniter) core.Queue {
	return &queue{
		newestOnTop: newStack(),
		oldestOnTop: newStack(),
	}
}

func (q *queue) Enqueue(v int) { // add item
	// move everything from oldestOnTop into newestOnTop
	for q.oldestOnTop.Len() > 0 {
		q.newestOnTop.Push(q.oldestOnTop.Pop())
	}

	// add to the top of newestOnTop
	q.newestOnTop.Push(v)

	// move everything back to oldestOnTop
	for q.newestOnTop.Len() > 0 {
		q.oldestOnTop.Push(q.newestOnTop.Pop())
	}
}

func (q *queue) Peek() int { // get "oldest" item
	return q.oldestOnTop.Peek()
}

func (q *queue) Dequeue() int { // get "oldest" and remove it
	return q.oldestOnTop.Pop()
}
