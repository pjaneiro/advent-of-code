package nomatterhowyousliceit_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2018/nomatterhowyousliceit"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Claim
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Claim{Claim{Id: 1, X: 1, Y: 3, W: 4, H: 4}, Claim{Id: 2, X: 3, Y: 1, W: 4, H: 4}, Claim{Id: 3, X: 5, Y: 5, W: 2, H: 2}},
			expected: 4,
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
		input    []Claim
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Claim{Claim{Id: 1, X: 1, Y: 3, W: 4, H: 4}, Claim{Id: 2, X: 3, Y: 1, W: 4, H: 4}, Claim{Id: 3, X: 5, Y: 5, W: 2, H: 2}},
			expected: 3,
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
