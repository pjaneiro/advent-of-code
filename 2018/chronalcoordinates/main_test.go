package chronalcoordinates_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2018/chronalcoordinates"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Coordinate
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Coordinate{Coordinate{X: 1, Y: 1}, Coordinate{X: 1, Y: 6}, Coordinate{X: 8, Y: 3}, Coordinate{X: 3, Y: 4}, Coordinate{X: 5, Y: 5}, Coordinate{X: 8, Y: 9}},
			expected: 17,
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
		input    []Coordinate
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Coordinate{Coordinate{X: 1, Y: 1}, Coordinate{X: 1, Y: 6}, Coordinate{X: 8, Y: 3}, Coordinate{X: 3, Y: 4}, Coordinate{X: 5, Y: 5}, Coordinate{X: 8, Y: 9}},
			expected: 16,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.input, 32)
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
