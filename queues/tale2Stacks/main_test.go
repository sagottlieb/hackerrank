package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	file string
}

func TestQueue(t *testing.T) {
	testCase := getTestCases()

	for _, tc := range testCase {
		t.Run(tc.file, func(t *testing.T) {

			queries, err := getTestInputQueries(tc.file)
			if err != nil {
				t.Fatalf("Could not open test data file for input file %s: %v", tc.file, err)
			}

			queue := newQueue()

			output := doQuerySequence(queue, queries)

			expected, err := getExpectations(tc.file)
			if err != nil {
				t.Fatalf("Could not open test expectation file for input file %s: %v", tc.file, err)
			}

			assert.Equal(t, expected, output)
		})
	}
}

func BenchmarkQueue(b *testing.B) {
	testCase := getTestCases()

	for _, tc := range testCase {

		b.Run(tc.file, func(b *testing.B) {

			queries, err := getTestInputQueries(tc.file)
			if err != nil {
				b.Fatalf("Could not open test data file for input file %s: %v", tc.file, err)
			}

			for i := 0; i < b.N; i++ {
				queue := newQueue()

				_ = doQuerySequence(queue, queries)
			}

		})
	}
}

func getTestInputQueries(filename string) ([]queryDoer, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	queries := parseInput(f)

	f.Close()

	return queries, nil
}

func getTestCases() []testCase {

	inputFiles := []string{}

	err := filepath.Walk("testcases", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}

		if info.Name() == "input.txt" {
			inputFiles = append(inputFiles, path)
		}

		return nil
	})

	if err != nil {
		log.Fatalf("walk error [%v]", err)
	}

	testCases := make([]testCase, len(inputFiles))

	for i, in := range inputFiles {
		testCases[i] = testCase{file: in}
	}

	return testCases
}

// hack together output file's path based on the input file path
func constructOutputFilename(in string) string {
	return strings.Replace(in, "input.txt", "output.txt", -1)
}

func getExpectations(inFile string) (string, error) {
	out, err := ioutil.ReadFile(constructOutputFilename(inFile))
	if err != nil {
		return "", err
	}

	return string(out), nil
}
