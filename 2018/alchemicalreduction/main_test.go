package alchemicalreduction_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2018/alchemicalreduction"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    "aA",
			expected: 0,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    "abBA",
			expected: 0,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    "abAB",
			expected: 4,
			error:    false,
		},
		{
			name:     "Example 4",
			input:    "aabAAB",
			expected: 6,
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
		input    string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    "dabAcCaCBAcCcaDA",
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
