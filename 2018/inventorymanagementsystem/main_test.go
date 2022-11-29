package inventorymanagementsystem_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2018/inventorymanagementsystem"
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
			input:    []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"},
			expected: 12,
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
		expected string
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"},
			expected: "fgij",
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v) threw '%v', want %s", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v) = %s, want to throw", tc.input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v) = %s, want %s", tc.input, actual, tc.expected)
			}
		})
	}
}
