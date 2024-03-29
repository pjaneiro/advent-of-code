package dive_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2021/golang/dive"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Command
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Command{{Cmd: "forward", Val: 5}, {Cmd: "down", Val: 5}, {Cmd: "forward", Val: 8}, {Cmd: "up", Val: 3}, {Cmd: "down", Val: 8}, {Cmd: "forward", Val: 2}},
			expected: 150,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []Command{},
			expected: 0,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v) threw '%v', want %d", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v) = %d, want to throw", tc.input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v) = %d, want %d", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Command
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Command{{Cmd: "forward", Val: 5}, {Cmd: "down", Val: 5}, {Cmd: "forward", Val: 8}, {Cmd: "up", Val: 3}, {Cmd: "down", Val: 8}, {Cmd: "forward", Val: 2}},
			expected: 900,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []Command{},
			expected: 0,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v) threw '%v', want %d", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v) = %d, want to throw", tc.input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v) = %d, want %d", tc.input, actual, tc.expected)
			}
		})
	}
}
