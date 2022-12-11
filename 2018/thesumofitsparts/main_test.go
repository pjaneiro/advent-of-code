package thesumofitsparts_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2018/thesumofitsparts"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Dependency
		expected string
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Dependency{Dependency{From: 'C', To: 'A'}, Dependency{From: 'C', To: 'F'}, Dependency{From: 'A', To: 'B'}, Dependency{From: 'A', To: 'D'}, Dependency{From: 'B', To: 'E'}, Dependency{From: 'D', To: 'E'}, Dependency{From: 'F', To: 'E'}},
			expected: "CABDFE",
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.input, 6)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v) threw '%v', want %s", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v) = %s, want to throw", tc.input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v) = %s, want %s", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Dependency
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Dependency{Dependency{From: 'C', To: 'A'}, Dependency{From: 'C', To: 'F'}, Dependency{From: 'A', To: 'B'}, Dependency{From: 'A', To: 'D'}, Dependency{From: 'B', To: 'E'}, Dependency{From: 'D', To: 'E'}, Dependency{From: 'F', To: 'E'}},
			expected: 15,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.input, 6, 2, 0)
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
