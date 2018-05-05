package main

import "fmt"

// type 1
// implements queryDoer
type enqueueQuery struct {
	val int
}

func (e enqueueQuery) prettyPrint() {
	fmt.Printf("ENQUEUE %d\n", e.val)
}

func (e enqueueQuery) doQuery(q *queue) {
	q.Enqueue(e.val)
}

// type 2
// implements queryDoer
type dequeueQuery struct{}

func (d dequeueQuery) prettyPrint() {
	fmt.Println("DEQUEUE HEAD")
}

func (d dequeueQuery) doQuery(q *queue) {
	_ = q.Dequeue()
}

// type 3
// implements queryDoer
type peekQuery struct{}

func (p peekQuery) prettyPrint() {
	fmt.Println("PRINT HEAD")
}

func (p peekQuery) doQuery(q *queue) {
	v := q.Peek()
	fmt.Printf("%d\n", v)
}
