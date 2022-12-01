package seacucumber_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2021/golang/seacucumber"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    [][]byte
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    [][]byte{[]byte{'v', '.', '.', '.', '>', '>', '.', 'v', 'v', '>'}, []byte{'.', 'v', 'v', '>', '>', '.', 'v', 'v', '.', '.'}, []byte{'>', '>', '.', '>', 'v', '>', '.', '.', '.', 'v'}, []byte{'>', '>', 'v', '>', '>', '.', '>', '.', 'v', '.'}, []byte{'v', '>', 'v', '.', 'v', 'v', '.', 'v', '.', '.'}, []byte{'>', '.', '>', '>', '.', '.', 'v', '.', '.', '.'}, []byte{'.', 'v', 'v', '.', '.', '>', '.', '>', 'v', '.'}, []byte{'v', '.', 'v', '.', '.', '>', '>', 'v', '.', 'v'}, []byte{'.', '.', '.', '.', 'v', '.', '.', 'v', '.', '>'}},
			expected: 58,
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
