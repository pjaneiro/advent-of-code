package proboscideavolcanium_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2022/proboscideavolcanium"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name       string
		input_path string
		expected   int
		error      bool
	}{
		{
			name:       "Example 1",
			input_path: "input_test.txt",
			expected:   1651,
			error:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input, err := ReadLines(tc.input_path)
			if err != nil {
				t.Errorf("Error reading input")
				t.FailNow()
			}
			actual, err := Challenge1(input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v) threw '%v', want %d", input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v) = %d, want to throw", input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v) = %d, want %d", input, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name       string
		input_path string
		expected   int
		error      bool
	}{
		{
			name:       "Example 1",
			input_path: "input_test.txt",
			expected:   1707,
			error:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input, err := ReadLines(tc.input_path)
			if err != nil {
				t.Errorf("Error reading input")
				t.FailNow()
			}
			actual, err := Challenge2(input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v) threw '%v', want %d", input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v) = %d, want to throw", input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v) = %d, want %d", input, actual, tc.expected)
			}
		})
	}
}
