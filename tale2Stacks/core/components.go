package core

type Stack interface {
	Len() int
	Push(int)
	Pop() int
	Peek() int
}

type StackIniter func() Stack

type Queue interface {
	Enqueue(int)
	Dequeue() int
	Peek() int
}

type QueueIniter func(StackIniter) Queue

type QueryDoer interface {
	PrettyPrint()
	DoQuery(Queue) string
}
