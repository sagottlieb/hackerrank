package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	inputs := parseFromStdin()

	for _, n := range inputs {
		prettyPrintPrimality(isPrime(n))
	}
}

func prettyPrintPrimality(b bool) {
	if b {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not prime")
	}
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}

	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// not pretty but does the trick
func parseFromStdin() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// "The first line contains an integer, p, denoting the number of integers to check for primality."
	scanner.Scan()
	x := scanner.Text()
	x = strings.TrimSuffix(x, "\n")
	p, _ := strconv.Atoi(x) // TODO check error

	var inputs = make([]int, p)

	// "Each of the subsequent lines contains an integer, n, you must test for primality."
	for i := 0; i < p; i++ {
		scanner.Scan()
		s := scanner.Text()
		s = strings.TrimSuffix(s, "\n")
		y, _ := strconv.Atoi(s) // TODO check error
		inputs[i] = y
	}

	return inputs
}
