package query

import (
	"fmt"

	"github.com/sagottlieb/hackerrank/queues/tale2Stacks/core"
)

// type 1
// implements queryDoer
type Enqueue struct {
	Val int
}

func (e Enqueue) PrettyPrint() {
	fmt.Printf("ENQUEUE %d\n", e.Val)
}

func (e Enqueue) DoQuery(q core.Queue) string {
	q.Enqueue(e.Val)
	return ""
}

// type 2
// implements queryDoer
type Dequeue struct{}

func (d Dequeue) PrettyPrint() {
	fmt.Println("DEQUEUE HEAD")
}

func (d Dequeue) DoQuery(q core.Queue) string {
	_ = q.Dequeue()
	return ""
}

// type 3
// implements queryDoer
type Peek struct{}

func (p Peek) PrettyPrint() {
	fmt.Println("PRINT HEAD")
}

func (p Peek) DoQuery(q core.Queue) string {
	v := q.Peek()
	return fmt.Sprintf("%d\n", v)
}
