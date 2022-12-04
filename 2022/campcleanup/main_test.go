package campcleanup_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2022/campcleanup"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Range
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Range{Range{Min1: 2, Max1: 4, Min2: 6, Max2: 8}, Range{Min1: 2, Max1: 3, Min2: 4, Max2: 5}, Range{Min1: 5, Max1: 7, Min2: 7, Max2: 9}, Range{Min1: 2, Max1: 8, Min2: 3, Max2: 7}, Range{Min1: 6, Max1: 6, Min2: 4, Max2: 6}, Range{Min1: 2, Max1: 6, Min2: 4, Max2: 8}},
			expected: 2,
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
		input    []Range
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Range{Range{Min1: 2, Max1: 4, Min2: 6, Max2: 8}, Range{Min1: 2, Max1: 3, Min2: 4, Max2: 5}, Range{Min1: 5, Max1: 7, Min2: 7, Max2: 9}, Range{Min1: 2, Max1: 8, Min2: 3, Max2: 7}, Range{Min1: 6, Max1: 6, Min2: 4, Max2: 6}, Range{Min1: 2, Max1: 6, Min2: 4, Max2: 8}},
			expected: 4,
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
