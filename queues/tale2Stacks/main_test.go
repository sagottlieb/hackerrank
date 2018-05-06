package main

import (
	"testing"
	"os"
	"log"
	"path/filepath"
	"strings"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	file            string
}

func TestQueue(t *testing.T) {
	testCase:= getTestCases()

	for _, tc := range testCase {
		t.Run(tc.file, func(t *testing.T) {
			f, err := os.Open(tc.file)
			if err != nil {
				t.Fatalf("Could not open test data file %s: %v", tc.file, err)
			}

			queries := parseInput(f)
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