	package main

	import (
		"bufio"
		"os"
		"strconv"
		"strings"
		"fmt"
	)

	// not pretty but does the trick
	func parseInput(file *os.File) []queryDoer {
		r := bufio.NewReader(file)

		// "The first line contains a single integer, q, denoting the number of queries."
		x, _ := r.ReadString('\n') // TODO check error
		x = strings.TrimSuffix(x, "\n")
		q, _ := strconv.Atoi(x) // TODO check error

		var queries = make([]queryDoer, q)
		var err error

		for i := 0; i < q; i++ {
			s, _ := r.ReadString('\n') // TODO check error
			queries[i], err = convertToQuery(s)
			if err != nil {
				panic(fmt.Sprintf("input line %d: %s", i+1, err.Error()))
			}
		}

		return queries
	}

	func convertToQuery(s string) (queryDoer, error) {

		s = strings.TrimSuffix(s, "\n")

		parts := strings.Split(s, " ")

		if len(parts) == 2 { // type 1 query
			y := parts[0]
			z, _ := strconv.Atoi(y) // TODO check error
			if z != 1 {
				return nil, fmt.Errorf("input format violates exercise instructions: line has two space separated parts but first part is not 1 (%s)", s)
			}

			u := parts[1]
			v, _ := strconv.Atoi(u) // TODO check error

			return enqueueQuery{val: v}, nil
		}

		if len(parts) == 1 { // type 2 or 3 query
			y := parts[0]
			z, _ := strconv.Atoi(y) // TODO check error

			if z == 2 {
				return dequeueQuery{}, nil
			}

			if z == 3 {
				return peekQuery{}, nil
			}

			return nil, fmt.Errorf("input format violates exercise instructions: line does not begin with 2 or 3 (%s)", s)
		}

		return nil, fmt.Errorf("input format violates exercise instructions: line has %d space separated parts (%s)", len(parts), s)
	}

	func prettyFormat(x []string) string {
		out := ""
		for i := 0; i < len(x); i++ {
			out += string(x[i])
		}
		return out
	}
