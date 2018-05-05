package main

func main() {

	inputs := parseFromStdin() // []queryDoer

	q := newQueue()

	for _, x := range inputs {

		// x.prettyPrint() // uncomment to debug

		x.doQuery(q)
	}
}

type queryDoer interface {
	prettyPrint()
	doQuery(*queue)
}
