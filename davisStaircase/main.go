package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := parseInput()

	c := count_generator()

	var counts []int

	for n := 0; n < max(inputs); n++ {
		v := <-c
		counts = append(counts, v)
	}

	for _, x := range inputs {
		fmt.Printf("%d\n", counts[x-1])
	}
}

func count_generator() chan int {
	c := make(chan int)

	go func() {
		c <- 1
		c <- 2
		for i, j, k := 4, 2, 1; ; i, j, k = i+j+k, i, j {
			c <- i
		}
	}()

	return c
}

func countWaysRecursive(numStairs int32) int32 {
	if numStairs == 0 {
		return 1
	}
	if numStairs == 1 {
		return 1
	}
	if numStairs == 2 {
		return 2
	}
	return countWaysRecursive(numStairs-1) + countWaysRecursive(numStairs-2) + countWaysRecursive(numStairs-3)
}

func parseInput() []int32 {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	s := int32(sTemp)

	inputs := make([]int32, s)

	for sItr := 0; sItr < int(s); sItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		inputs[sItr] = int32(nTemp)
	}

	return inputs
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func max(x []int32) int {
	maxSoFar := 0
	for _, y := range x {
		if int(y) > maxSoFar {
			maxSoFar = int(y)
		}
	}

	return maxSoFar
}
