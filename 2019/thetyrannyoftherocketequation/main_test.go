package thetyrannyoftherocketequation_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2019/thetyrannyoftherocketequation"
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
			input:    []int{12},
			expected: 2,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []int{14},
			expected: 2,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    []int{1969},
			expected: 654,
			error:    false,
		},
		{
			name:     "Example 4",
			input:    []int{100756},
			expected: 33583,
			error:    false,
		},
		{
			name:     "Example 5",
			input:    []int{12, 14, 1969, 100756},
			expected: 34241,
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
		input    []int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []int{14},
			expected: 2,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []int{1969},
			expected: 966,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    []int{100756},
			expected: 50346,
			error:    false,
		},
		{
			name:     "Example 4",
			input:    []int{14, 1969, 100756},
			expected: 51314,
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
