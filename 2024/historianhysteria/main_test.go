package historianhysteria_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2024/historianhysteria"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input1   []int
		input2   []int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input1:   []int{3, 4, 2, 1, 3, 3},
			input2:   []int{4, 3, 5, 3, 9, 3},
			expected: 11,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.input1, tc.input2)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %v) threw '%v', want %d", tc.input1, tc.input2, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %v) = %d, want to throw", tc.input1, tc.input2, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v) = %d, want %d", tc.input1, tc.input2, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input1   []int
		input2   []int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input1:   []int{3, 4, 2, 1, 3, 3},
			input2:   []int{4, 3, 5, 3, 9, 3},
			expected: 31,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.input1, tc.input2)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v, %v) threw '%v', want %d", tc.input1, tc.input2, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v, %v) = %d, want to throw", tc.input1, tc.input2, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v, %v) = %d, want %d", tc.input1, tc.input2, actual, tc.expected)
			}
		})
	}
}
