package trickshot_test

import (
	. "github.com/pjaneiro/advent-of-code-2021/trickshot"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		minX     int
		maxX     int
		minY     int
		maxY     int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			minX:     20,
			maxX:     30,
			minY:     -10,
			maxY:     -5,
			expected: 45,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.minX, tc.maxX, tc.minY, tc.maxY)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %v, %v, %v) threw '%v', want %d", tc.minX, tc.maxX, tc.minY, tc.maxY, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %v, %v, %v) = %d, want to throw", tc.minX, tc.maxX, tc.minY, tc.maxY, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v, %v, %v) = %d, want %d", tc.minX, tc.maxX, tc.minY, tc.maxY, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		minX     int
		maxX     int
		minY     int
		maxY     int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			minX:     20,
			maxX:     30,
			minY:     -10,
			maxY:     -5,
			expected: 112,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.minX, tc.maxX, tc.minY, tc.maxY)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v, %v, %v, %v) threw '%v', want %d", tc.minX, tc.maxX, tc.minY, tc.maxY, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v, %v, %v, %v) = %d, want to throw", tc.minX, tc.maxX, tc.minY, tc.maxY, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v, %v, %v, %v) = %d, want %d", tc.minX, tc.maxX, tc.minY, tc.maxY, actual, tc.expected)
			}
		})
	}
}
