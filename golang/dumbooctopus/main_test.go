package dumbooctopus_test

import (
	. "github.com/pjaneiro/advent-of-code-2021/dumbooctopus"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    [][]int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    [][]int{{5, 4, 8, 3, 1, 4, 3, 2, 2, 3}, {2, 7, 4, 5, 8, 5, 4, 7, 1, 1}, {5, 2, 6, 4, 5, 5, 6, 1, 7, 3}, {6, 1, 4, 1, 3, 3, 6, 1, 4, 6}, {6, 3, 5, 7, 3, 8, 5, 4, 7, 8}, {4, 1, 6, 7, 5, 2, 4, 6, 4, 5}, {2, 1, 7, 6, 8, 4, 1, 7, 2, 1}, {6, 8, 8, 2, 8, 8, 1, 1, 3, 4}, {4, 8, 4, 6, 8, 4, 8, 5, 5, 4}, {5, 2, 8, 3, 7, 5, 1, 5, 2, 6}},
			expected: 1656,
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
		input    [][]int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    [][]int{{5, 4, 8, 3, 1, 4, 3, 2, 2, 3}, {2, 7, 4, 5, 8, 5, 4, 7, 1, 1}, {5, 2, 6, 4, 5, 5, 6, 1, 7, 3}, {6, 1, 4, 1, 3, 3, 6, 1, 4, 6}, {6, 3, 5, 7, 3, 8, 5, 4, 7, 8}, {4, 1, 6, 7, 5, 2, 4, 6, 4, 5}, {2, 1, 7, 6, 8, 4, 1, 7, 2, 1}, {6, 8, 8, 2, 8, 8, 1, 1, 3, 4}, {4, 8, 4, 6, 8, 4, 8, 5, 5, 4}, {5, 2, 8, 3, 7, 5, 1, 5, 2, 6}},
			expected: 195,
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
