package cafeteria_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2025/cafeteria"
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
			input:    []string{"3-5", "10-14", "16-20", "12-18", "", "1", "5", "8", "11", "17", "32"},
			expected: 3,
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
			input:    []string{"3-5", "10-14", "16-20", "12-18", "", "1", "5", "8", "11", "17", "32"},
			expected: 14,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []string{"200-300", "100-101", "1-1", "2-2", "3-3", "1-3", "1-3", "2-2", "50-70", "10-10", "98-99", "99-99", "99-99", "99-100", "1-1", "2-1", "100-100", "100-100", "100-101", "200-300", "201-300", "202-300", "250-251", "98-99", "100-100", "100-101", "1-101", "", "202"},
			expected: 202,
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
