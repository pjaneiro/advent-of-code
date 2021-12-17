package spaceimageformat_test

import (
	. "github.com/pjaneiro/advent-of-code-2019/spaceimageformat"
	"testing"
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
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
			expected: 1,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.input, 2, 3)
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
