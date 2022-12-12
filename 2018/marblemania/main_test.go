package marblemania_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2018/marblemania"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input1   int
		input2   int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input1:   9,
			input2:   25,
			expected: 32,
			error:    false,
		},
		{
			name:     "Example 2",
			input1:   10,
			input2:   1618,
			expected: 8317,
			error:    false,
		},
		{
			name:     "Example 3",
			input1:   13,
			input2:   7999,
			expected: 146373,
			error:    false,
		},
		{
			name:     "Example 4",
			input1:   17,
			input2:   1104,
			expected: 2764,
			error:    false,
		},
		{
			name:     "Example 5",
			input1:   21,
			input2:   6111,
			expected: 54718,
			error:    false,
		},
		{
			name:     "Example 6",
			input1:   30,
			input2:   5807,
			expected: 37305,
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
