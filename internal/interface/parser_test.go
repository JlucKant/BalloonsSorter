package parser_test

import (
	"strings"
	"testing"

	parser "github.com/JlucKant/BalloonsSorter/internal/interface"
)

type TestParsing struct {
	name          string
	input         string
	expectedValue [][]uint32
	expectedError bool
}

func checkEqual(actual, expected [][]uint32) bool {
	if len(actual) != len(expected) {
		return false
	}

	for i, _ := range actual {
		for j, _ := range expected {
			if actual[i][j] != expected[i][j] {
				return false
			}
		}
	}
	return true
}

func TestInputParser(t *testing.T) {
	tests := []TestParsing{
		{
			name: "Example 1: Valid",
			input: `2
					12 0
					0 21`,
			expectedValue: [][]uint32{{12, 0}, {0, 21}},
		},
		{
			name: "Example 2: Invalid",
			input: `2
					12 
					21`,
			expectedValue: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			reader := strings.NewReader(testCase.input)
			actual, _ := parser.InputParser(reader)

			if !checkEqual(actual, testCase.expectedValue) {
				t.Errorf("expected %v, got %v", testCase.expectedValue, actual)
			}
		})
	}
}
