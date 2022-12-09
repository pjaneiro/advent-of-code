package treetoptreehouse_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2022/treetoptreehouse"
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
			input:    [][]int{[]int{3, 0, 3, 7, 3}, []int{2, 5, 5, 1, 2}, []int{6, 5, 3, 3, 2}, []int{3, 3, 5, 4, 9}, []int{3, 5, 3, 9, 0}},
			expected: 21,
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
			input:    [][]int{[]int{3, 0, 3, 7, 3}, []int{2, 5, 5, 1, 2}, []int{6, 5, 3, 3, 2}, []int{3, 3, 5, 4, 9}, []int{3, 5, 3, 9, 0}},
			expected: 8,
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
