package rambunctiousrecitation_test

import (
	. "github.com/pjaneiro/advent-of-code/2020/rambunctiousrecitation"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1",
			input:    []int{0, 3, 6},
			expected: 436,
		}, {
			name:     "Example 2",
			input:    []int{1, 3, 2},
			expected: 1,
		}, {
			name:     "Example 3",
			input:    []int{2, 1, 3},
			expected: 10,
		}, {
			name:     "Example 4",
			input:    []int{1, 2, 3},
			expected: 27,
		}, {
			name:     "Example 5",
			input:    []int{2, 3, 1},
			expected: 78,
		}, {
			name:     "Example 6",
			input:    []int{3, 2, 1},
			expected: 438,
		}, {
			name:     "Example 7",
			input:    []int{3, 1, 2},
			expected: 1836,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := Challenge1(tc.input)
			if actual != tc.expected {
				t.Errorf("Challenge1(%v) = %d, want %d", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1",
			input:    []int{0, 3, 6},
			expected: 175594,
		}, {
			name:     "Example 2",
			input:    []int{1, 3, 2},
			expected: 2578,
		}, {
			name:     "Example 3",
			input:    []int{2, 1, 3},
			expected: 3544142,
		}, {
			name:     "Example 4",
			input:    []int{1, 2, 3},
			expected: 261214,
		}, {
			name:     "Example 5",
			input:    []int{2, 3, 1},
			expected: 6895259,
		}, {
			name:     "Example 6",
			input:    []int{3, 2, 1},
			expected: 18,
		}, {
			name:     "Example 7",
			input:    []int{3, 1, 2},
			expected: 362,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := Challenge2(tc.input)
			if actual != tc.expected {
				t.Errorf("Challenge2(%v) = %d, want %d", tc.input, actual, tc.expected)
			}
		})
	}
}
