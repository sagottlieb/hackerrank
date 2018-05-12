package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)

// not pretty but does the trick
func parseFromStdin() []string {
	r := bufio.NewReader(os.Stdin)

	// "The first line contains a single integer, n, denoting the number of strings."
	x, _ := r.ReadString('\n') // TODO check error
	x = strings.TrimSuffix(x, "\n")
	n, _ := strconv.Atoi(x) // TODO check error

	var inputs = make([]string, n)

	// "Each line i of the n subsequent lines consists of a single string, s, denoting a sequence of brackets."
	for i :=  0; i < n; i++ {
		s, _ := r.ReadString('\n') // TODO check error
		inputs[i] = strings.TrimSuffix(s, "\n")
	}

	return inputs
}

func prettyFormat(x []string) string {
	out := ""
	for i:= 0 ; i < len(x); i++ {
		out += string(x[i])
	}
	return out
}
