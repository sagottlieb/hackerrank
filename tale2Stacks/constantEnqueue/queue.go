package constantEnqueue

import "github.com/sagottlieb/hackerrank/tale2Stacks/core"

type queue struct {
	newestOnTop core.Stack // source of truth
	oldestOnTop core.Stack // empty except during peek/dequeue ops
}

func New(newStack core.StackIniter) core.Queue {
	return &queue{
		newestOnTop: newStack(),
		oldestOnTop: newStack(),
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

	v := q.oldestOnTop.Peek()

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

	v := q.oldestOnTop.Pop()

	// move everything back to newestOnTop
	for q.oldestOnTop.Len() > 0 {
		q.newestOnTop.Push(q.oldestOnTop.Pop())
	}

	return v
}
