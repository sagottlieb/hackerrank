package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// not pretty but does the trick
func parseFromStdin() []queryDoer {
	r := bufio.NewReader(os.Stdin)

	// "The first line contains a single integer, q, denoting the number of queries."
	x, _ := r.ReadString('\n') // TODO check error
	x = strings.TrimSuffix(x, "\n")
	q, _ := strconv.Atoi(x) // TODO check error

	var queries = make([]queryDoer, q)

	for i := 0; i < q; i++ {
		s, _ := r.ReadString('\n') // TODO check error
		queries[i] = convertToQuery(s)
	}

	return queries
}

func convertToQuery(s string) queryDoer {

	s = strings.TrimSuffix(s, "\n")

	parts := strings.Split(s, " ")

	if len(parts) == 2 { // type 1 query
		y := parts[0]
		z, _ := strconv.Atoi(y) // TODO check error
		if z != 1 {
			panic("input format violates exercise instructions")
		}

		u := parts[1]
		v, _ := strconv.Atoi(u) // TODO check error

		return enqueueQuery{val: v}
	}

	if len(parts) == 1 { // type 2 or 3 query
		y := parts[0]
		z, _ := strconv.Atoi(y) // TODO check error

		if z == 2 {
			return dequeueQuery{}
		}

		if z == 3 {
			return peekQuery{}
		}

		panic("input format violates exercise instructions")
	}

	panic("input format violates exercise instructions")
}

func prettyFormat(x []string) string {
	out := ""
	for i := 0; i < len(x); i++ {
		out += string(x[i])
	}
	return out
}
