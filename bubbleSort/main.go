package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	inputs := parseFromStdin()

	out, swaps := bubbleSort(inputs)

	fmt.Printf("Array is sorted in %d swaps.\n", swaps)
	fmt.Printf("First Element: %d\n", out[0])
	fmt.Printf("Last Element: %d\n", out[len(out)-1])
}

func bubbleSort(in []int) ([]int, int) {
	swaps := 0

	for _, _ = range in {
		for j, _ := range in[:len(in)-1] {
			// Swap adjacent elements if they are in decreasing order
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
				swaps++
			}
		}
	}

	return in, swaps

}

// not pretty but does the trick
func parseFromStdin() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	x := scanner.Text()
	x = strings.TrimSuffix(x, "\n")
	n, _ := strconv.Atoi(x) // TODO check error

	var inputs = make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		s = strings.TrimSuffix(s, "\n")
		y, _ := strconv.Atoi(s) // TODO check error
		inputs[i] = y
	}

	return inputs
}
