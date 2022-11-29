package chronalcalibration_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2018/chronalcalibration"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Operation
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Operation{},
			expected: 0,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []Operation{Operation{Op: '+', Val: 1}, Operation{Op: '-', Val: 2}, Operation{Op: '+', Val: 3}, Operation{Op: '+', Val: 1}},
			expected: 3,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    []Operation{Operation{Op: '+', Val: 1}, Operation{Op: '+', Val: 1}, Operation{Op: '+', Val: 1}},
			expected: 3,
			error:    false,
		},
		{
			name:     "Example 4",
			input:    []Operation{Operation{Op: '+', Val: 1}, Operation{Op: '+', Val: 1}, Operation{Op: '-', Val: 2}},
			expected: 0,
			error:    false,
		},
		{
			name:     "Example 5",
			input:    []Operation{Operation{Op: '-', Val: 1}, Operation{Op: '-', Val: 2}, Operation{Op: '-', Val: 3}},
			expected: -6,
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
		input    []Operation
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Operation{},
			expected: 0,
			error:    true,
		},
		{
			name:     "Example 2",
			input:    []Operation{Operation{Op: '+', Val: 1}, Operation{Op: '-', Val: 2}, Operation{Op: '+', Val: 3}, Operation{Op: '+', Val: 1}},
			expected: 2,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    []Operation{Operation{Op: '+', Val: 1}, Operation{Op: '-', Val: 1}},
			expected: 0,
			error:    false,
		},
		{
			name:     "Example 4",
			input:    []Operation{Operation{Op: '+', Val: 3}, Operation{Op: '+', Val: 3}, Operation{Op: '+', Val: 4}, Operation{Op: '-', Val: 2}, Operation{Op: '-', Val: 4}},
			expected: 10,
			error:    false,
		},
		{
			name:     "Example 5",
			input:    []Operation{Operation{Op: '-', Val: 6}, Operation{Op: '+', Val: 3}, Operation{Op: '+', Val: 8}, Operation{Op: '+', Val: 5}, Operation{Op: '-', Val: 6}},
			expected: 5,
			error:    false,
		},
		{
			name:     "Example 6",
			input:    []Operation{Operation{Op: '+', Val: 7}, Operation{Op: '+', Val: 7}, Operation{Op: '-', Val: 2}, Operation{Op: '-', Val: 7}, Operation{Op: '-', Val: 4}},
			expected: 14,
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
