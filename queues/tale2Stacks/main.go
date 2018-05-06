package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	inputs := parseInput(os.Stdin) // []queryDoer

	q := newQueue()

	output := doQuerySequence(q, inputs)

	fmt.Printf(output)

}

type stack interface {
	Len() int
	Push(int)
	Pop() int
	Peek() int
}

type queue struct {
	newestOnTop stack
	oldestOnTop stack
}

type queryDoer interface {
	prettyPrint()
	doQuery(*queue) string
}

// This setup allows for testing while continuing to print to STDOUT in the way
// expected by hackerrank.
func doQuerySequence(q *queue, inputs []queryDoer) string {

	// much more efficient than naive string concatenation
	var buffer bytes.Buffer

	for _, x := range inputs {

		// x.prettyPrint() // uncomment to debug

		if out := x.doQuery(q); out != "" {
			buffer.WriteString(out)
		}

	}

	return buffer.String()
}
