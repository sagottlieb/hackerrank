package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/sagottlieb/hackerrank/queues/tale2Stacks/core"
	"github.com/sagottlieb/hackerrank/queues/tale2Stacks/lazyQueue"
	"github.com/sagottlieb/hackerrank/queues/tale2Stacks/linkedListStack"
)

func main() {

	inputs := parseInput(os.Stdin) // []queryDoer

	q := lazyQueue.New(linkedListStack.New)

	output := doQuerySequence(q, inputs)

	fmt.Printf(output)

}

// This setup allows for testing while continuing to print to STDOUT in the way
// expected by hackerrank.
func doQuerySequence(q core.Queue, inputs []core.QueryDoer) string {

	// much more efficient than naive string concatenation
	var buffer bytes.Buffer

	for _, x := range inputs {

		// x.PrettyPrint() // uncomment to debug

		if out := x.DoQuery(q); out != "" {
			buffer.WriteString(out)
		}

	}

	return buffer.String()
}
