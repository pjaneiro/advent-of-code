package a1202programalarm_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2019/a1202programalarm"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			expected: 3500,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []int{1, 0, 0, 0, 99},
			expected: 2,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    []int{2, 3, 0, 3, 99},
			expected: 2,
			error:    false,
		},
		{
			name:     "Example 4",
			input:    []int{2, 4, 4, 5, 99, 0},
			expected: 2,
			error:    false,
		},
		{
			name:     "Example 5",
			input:    []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			expected: 30,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.input, false)
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
