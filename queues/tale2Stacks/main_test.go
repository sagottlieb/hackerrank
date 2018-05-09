package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"fmt"

	"github.com/sagottlieb/hackerrank/queues/tale2Stacks/constantDequeue"
	"github.com/sagottlieb/hackerrank/queues/tale2Stacks/constantEnqueue"
	"github.com/sagottlieb/hackerrank/queues/tale2Stacks/core"
	"github.com/sagottlieb/hackerrank/queues/tale2Stacks/lazyQueue"
	"github.com/sagottlieb/hackerrank/queues/tale2Stacks/linkedListStack"
	"github.com/sagottlieb/hackerrank/queues/tale2Stacks/naiveStack"
	"github.com/sagottlieb/hackerrank/queues/tale2Stacks/otsStack"
	"github.com/sagottlieb/hackerrank/queues/tale2Stacks/pointerStack"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	file        string
	stackMaker  core.StackIniter
	queueMaker  core.QueueIniter
	description string
}

func TestQueue(t *testing.T) {
	testCase := getTestCases()

	for _, tc := range testCase {
		t.Run(tc.description, func(t *testing.T) {

			queries, err := getTestInputQueries(tc.file)
			if err != nil {
				t.Fatalf("Could not open test data file for input file %s: %v", tc.file, err)
			}

			queue := tc.queueMaker(tc.stackMaker)

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

		b.Run(tc.description, func(b *testing.B) {

			queries, err := getTestInputQueries(tc.file)
			if err != nil {
				b.Fatalf("Could not open test data file for input file %s: %v", tc.file, err)
			}

			for i := 0; i < b.N; i++ {
				queue := tc.queueMaker(tc.stackMaker)

				_ = doQuerySequence(queue, queries)
			}

		})
	}
}

func getTestInputQueries(filename string) ([]core.QueryDoer, error) {
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

	testCases := []testCase{}

	for _, in := range inputFiles {
		newCases := []testCase{
			{file: in,
				stackMaker:  naiveStack.New,
				queueMaker:  constantEnqueue.New,
				description: fmt.Sprintf("naiveStack_constantEnqueue_%s", in),
			},
			{file: in,
				stackMaker:  naiveStack.New,
				queueMaker:  constantDequeue.New,
				description: fmt.Sprintf("naiveStack_constantDequeue_%s", in),
			},
			{file: in,
				stackMaker:  naiveStack.New,
				queueMaker:  lazyQueue.New,
				description: fmt.Sprintf("naiveStack_lazyQueue_%s", in),
			},
			{file: in,
				stackMaker:  pointerStack.New,
				queueMaker:  constantEnqueue.New,
				description: fmt.Sprintf("pointerStack_constantEnqueue_%s", in),
			},
			{file: in,
				stackMaker:  pointerStack.New,
				queueMaker:  constantDequeue.New,
				description: fmt.Sprintf("pointerStack_constantDequeue_%s", in),
			},
			{file: in,
				stackMaker:  pointerStack.New,
				queueMaker:  lazyQueue.New,
				description: fmt.Sprintf("pointerStack_lazyQueue_%s", in),
			},
			{file: in,
				stackMaker:  linkedListStack.New,
				queueMaker:  constantEnqueue.New,
				description: fmt.Sprintf("linkedListStack_constantEnqueue_%s", in),
			},
			{file: in,
				stackMaker:  linkedListStack.New,
				queueMaker:  constantDequeue.New,
				description: fmt.Sprintf("linkedListStack_constantDequeue_%s", in),
			},
			{file: in,
				stackMaker:  linkedListStack.New,
				queueMaker:  lazyQueue.New,
				description: fmt.Sprintf("linkedListStack_lazyQueue_%s", in),
			},
			{file: in,
				stackMaker:  otsStack.New,
				queueMaker:  constantEnqueue.New,
				description: fmt.Sprintf("OTSStack_constantEnqueue_%s", in),
			},
			{file: in,
				stackMaker:  otsStack.New,
				queueMaker:  constantDequeue.New,
				description: fmt.Sprintf("OTSStack_constantDequeue_%s", in),
			},
			{file: in,
				stackMaker:  otsStack.New,
				queueMaker:  lazyQueue.New,
				description: fmt.Sprintf("OTSStack_lazyQueue_%s", in),
			},
		}
		testCases = append(testCases, newCases...)
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
