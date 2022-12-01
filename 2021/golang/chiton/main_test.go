package chiton_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2021/golang/chiton"
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
			input:    [][]int{{1, 1, 6, 3, 7, 5, 1, 7, 4, 2}, {1, 3, 8, 1, 3, 7, 3, 6, 7, 2}, {2, 1, 3, 6, 5, 1, 1, 3, 2, 8}, {3, 6, 9, 4, 9, 3, 1, 5, 6, 9}, {7, 4, 6, 3, 4, 1, 7, 1, 1, 1}, {1, 3, 1, 9, 1, 2, 8, 1, 3, 7}, {1, 3, 5, 9, 9, 1, 2, 4, 2, 1}, {3, 1, 2, 5, 4, 2, 1, 6, 3, 9}, {1, 2, 9, 3, 1, 3, 8, 5, 2, 1}, {2, 3, 1, 1, 9, 4, 4, 5, 8, 1}},
			expected: 40,
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
			input:    [][]int{{1, 1, 6, 3, 7, 5, 1, 7, 4, 2}, {1, 3, 8, 1, 3, 7, 3, 6, 7, 2}, {2, 1, 3, 6, 5, 1, 1, 3, 2, 8}, {3, 6, 9, 4, 9, 3, 1, 5, 6, 9}, {7, 4, 6, 3, 4, 1, 7, 1, 1, 1}, {1, 3, 1, 9, 1, 2, 8, 1, 3, 7}, {1, 3, 5, 9, 9, 1, 2, 4, 2, 1}, {3, 1, 2, 5, 4, 2, 1, 6, 3, 9}, {1, 2, 9, 3, 1, 3, 8, 5, 2, 1}, {2, 3, 1, 1, 9, 4, 4, 5, 8, 1}},
			expected: 315,
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
