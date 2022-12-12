package chronalcharge_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2018/chronalcharge"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    int
		expected []int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    18,
			expected: []int{33, 45},
			error:    false,
		},
		{
			name:     "Example 2",
			input:    42,
			expected: []int{21, 61},
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v) threw '%v', want %v", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v) = %v, want to throw", tc.input, actual)
			} else if actual[0] != tc.expected[0] || actual[1] != tc.expected[1] {
				t.Errorf("Challenge1(%v) = %v, want %d", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input    int
		expected []int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    18,
			expected: []int{90, 269, 16},
			error:    false,
		},
		{
			name:     "Example 2",
			input:    42,
			expected: []int{232, 251, 12},
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v) threw '%v', want %v", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v) = %v, want to throw", tc.input, actual)
			} else if actual[0] != tc.expected[0] || actual[1] != tc.expected[1] || actual[2] != tc.expected[2] {
				t.Errorf("Challenge2(%v) = %v, want %d", tc.input, actual, tc.expected)
			}
		})
	}
}
