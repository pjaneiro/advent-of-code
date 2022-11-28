package crabcups_test

import (
	. "github.com/pjaneiro/advent-of-code/2020/crabcups"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		expected string
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expected: "67384529",
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v) threw '%v', want %s", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v) = %s, want to throw", tc.input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v) = %s, want %s", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		expected string
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expected: "149245887792",
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v) threw '%v', want %s", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v) = %s, want to throw", tc.input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v) = %s, want %s", tc.input, actual, tc.expected)
			}
		})
	}
}
