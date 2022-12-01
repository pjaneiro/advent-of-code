package operationorder_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2020/operationorder"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "Example 1",
			input:    []string{"1 + 2 * 3 + 4 * 5 + 6"},
			expected: 71,
		}, {
			name:     "Example 2",
			input:    []string{"1 + (2 * 3) + (4 * (5 + 6))"},
			expected: 51,
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
		input    []string
		expected int
	}{
		{
			name:     "Example 1",
			input:    []string{"1 + 2 * 3 + 4 * 5 + 6"},
			expected: 231,
		}, {
			name:     "Example 2",
			input:    []string{"1 + (2 * 3) + (4 * (5 + 6))"},
			expected: 51,
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
