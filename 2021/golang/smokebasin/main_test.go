package smokebasin_test

import (
	. "github.com/pjaneiro/advent-of-code-2021/smokebasin"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    [][]int
		expected int
		error    bool
	}{
		{
			name: "Example 1",
			input: [][]int{{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
				{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
				{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
				{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
				{9, 8, 9, 9, 9, 6, 5, 6, 7, 8}},
			expected: 15,
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
		input    [][]int
		expected int
		error    bool
	}{
		{
			name: "Example 1",
			input: [][]int{{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
				{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
				{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
				{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
				{9, 8, 9, 9, 9, 6, 5, 6, 7, 8}},
			expected: 1134,
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
