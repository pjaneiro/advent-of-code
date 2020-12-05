package binaryboarding_test

import (
	. "binaryboarding"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []string{"BFFFBBFRRR"},
			expected: 567,
			error:    false,
		}, {
			name:     "Example 2",
			input:    []string{"FFFBBBFRRR"},
			expected: 119,
			error:    false,
		}, {
			name:     "Example 3",
			input:    []string{"BBFFBBFRLL"},
			expected: 820,
			error:    false,
		}, {
			name:     "Example 4",
			input:    []string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"},
			expected: 820,
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
		input    []string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []string{"BFFFBBFRRR", "BFFFBBFRLR"},
			expected: 566,
			error:    false,
		}, {
			name:     "Example 2",
			input:    []string{"FFFBBBFRRR", "FFFBBBFRLR"},
			expected: 118,
			error:    false,
		}, {
			name:     "Example 3",
			input:    []string{"BBFFBBFRLL", "BBFFBBFRRL"},
			expected: 821,
			error:    false,
		}, {
			name:     "Example 4",
			input:    []string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"},
			expected: 0,
			error:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.input)
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
