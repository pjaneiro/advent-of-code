package universalorbitmap_test

import (
	. "github.com/pjaneiro/advent-of-code/2019/universalorbitmap"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    map[string]string
		expected int
		error    bool
	}{
		{
			name: "Example 1",
			input: map[string]string{
				"B": "COM",
				"C": "B",
				"D": "C",
				"E": "D",
				"F": "E",
				"G": "B",
				"H": "G",
				"I": "D",
				"J": "E",
				"K": "J",
				"L": "K",
			},
			expected: 42,
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
		input    map[string]string
		expected int
		error    bool
	}{
		{
			name: "Example 1",
			input: map[string]string{
				"B":   "COM",
				"C":   "B",
				"D":   "C",
				"E":   "D",
				"F":   "E",
				"G":   "B",
				"H":   "G",
				"I":   "D",
				"J":   "E",
				"K":   "J",
				"L":   "K",
				"YOU": "K",
				"SAN": "I",
			},
			expected: 4,
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
