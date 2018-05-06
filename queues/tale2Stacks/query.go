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

func (e enqueueQuery) doQuery(q *queue) string {
	q.Enqueue(e.val)
	return ""
}

// type 2
// implements queryDoer
type dequeueQuery struct{}

func (d dequeueQuery) prettyPrint() {
	fmt.Println("DEQUEUE HEAD")
}

func (d dequeueQuery) doQuery(q *queue) string {
	_ = q.Dequeue()
	return ""
}

// type 3
// implements queryDoer
type peekQuery struct{}

func (p peekQuery) prettyPrint() {
	fmt.Println("PRINT HEAD")
}

func (p peekQuery) doQuery(q *queue) string {
	v := q.Peek()
	return fmt.Sprintf("%d\n", v)
}
