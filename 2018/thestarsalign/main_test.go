package thestarsalign_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2018/thestarsalign"
)

func equalStringSlices(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Star
		expected []string
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Star{Star{X: 9, Y: 1, Vx: 0, Vy: 2}, Star{X: 7, Y: 0, Vx: -1, Vy: 0}, Star{X: 3, Y: -2, Vx: -1, Vy: 1}, Star{X: 6, Y: 10, Vx: -2, Vy: -1}, Star{X: 2, Y: -4, Vx: 2, Vy: 2}, Star{X: -6, Y: 10, Vx: 2, Vy: -2}, Star{X: 1, Y: 8, Vx: 1, Vy: -1}, Star{X: 1, Y: 7, Vx: 1, Vy: 0}, Star{X: -3, Y: 11, Vx: 1, Vy: -2}, Star{X: 7, Y: 6, Vx: -1, Vy: -1}, Star{X: -2, Y: 3, Vx: 1, Vy: 0}, Star{X: -4, Y: 3, Vx: 2, Vy: 0}, Star{X: 10, Y: -3, Vx: -1, Vy: 1}, Star{X: 5, Y: 11, Vx: 1, Vy: -2}, Star{X: 4, Y: 7, Vx: 0, Vy: -1}, Star{X: 8, Y: -2, Vx: 0, Vy: 1}, Star{X: 15, Y: 0, Vx: -2, Vy: 0}, Star{X: 1, Y: 6, Vx: 1, Vy: 0}, Star{X: 8, Y: 9, Vx: 0, Vy: -1}, Star{X: 3, Y: 3, Vx: -1, Vy: 1}, Star{X: 0, Y: 5, Vx: 0, Vy: -1}, Star{X: -2, Y: 2, Vx: 2, Vy: 0}, Star{X: 5, Y: -2, Vx: 1, Vy: 2}, Star{X: 1, Y: 4, Vx: 2, Vy: 1}, Star{X: -2, Y: 7, Vx: 2, Vy: -2}, Star{X: 3, Y: 6, Vx: -1, Vy: -1}, Star{X: 5, Y: 0, Vx: 1, Vy: 0}, Star{X: -6, Y: 0, Vx: 2, Vy: 0}, Star{X: 5, Y: 9, Vx: 1, Vy: -2}, Star{X: 14, Y: 7, Vx: -2, Vy: 0}, Star{X: -3, Y: 6, Vx: 2, Vy: -1}},
			expected: []string{"#...#..###", "#...#...#.", "#...#...#.", "#####...#.", "#...#...#.", "#...#...#.", "#...#...#.", "#...#..###"},
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v) threw '%v', want %v", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v) = %v, want to throw", tc.input, actual)
			} else if !equalStringSlices(actual, tc.expected) {
				t.Errorf("Challenge1(%v) = %v, want %v", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Star
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Star{Star{X: 9, Y: 1, Vx: 0, Vy: 2}, Star{X: 7, Y: 0, Vx: -1, Vy: 0}, Star{X: 3, Y: -2, Vx: -1, Vy: 1}, Star{X: 6, Y: 10, Vx: -2, Vy: -1}, Star{X: 2, Y: -4, Vx: 2, Vy: 2}, Star{X: -6, Y: 10, Vx: 2, Vy: -2}, Star{X: 1, Y: 8, Vx: 1, Vy: -1}, Star{X: 1, Y: 7, Vx: 1, Vy: 0}, Star{X: -3, Y: 11, Vx: 1, Vy: -2}, Star{X: 7, Y: 6, Vx: -1, Vy: -1}, Star{X: -2, Y: 3, Vx: 1, Vy: 0}, Star{X: -4, Y: 3, Vx: 2, Vy: 0}, Star{X: 10, Y: -3, Vx: -1, Vy: 1}, Star{X: 5, Y: 11, Vx: 1, Vy: -2}, Star{X: 4, Y: 7, Vx: 0, Vy: -1}, Star{X: 8, Y: -2, Vx: 0, Vy: 1}, Star{X: 15, Y: 0, Vx: -2, Vy: 0}, Star{X: 1, Y: 6, Vx: 1, Vy: 0}, Star{X: 8, Y: 9, Vx: 0, Vy: -1}, Star{X: 3, Y: 3, Vx: -1, Vy: 1}, Star{X: 0, Y: 5, Vx: 0, Vy: -1}, Star{X: -2, Y: 2, Vx: 2, Vy: 0}, Star{X: 5, Y: -2, Vx: 1, Vy: 2}, Star{X: 1, Y: 4, Vx: 2, Vy: 1}, Star{X: -2, Y: 7, Vx: 2, Vy: -2}, Star{X: 3, Y: 6, Vx: -1, Vy: -1}, Star{X: 5, Y: 0, Vx: 1, Vy: 0}, Star{X: -6, Y: 0, Vx: 2, Vy: 0}, Star{X: 5, Y: 9, Vx: 1, Vy: -2}, Star{X: 14, Y: 7, Vx: -2, Vy: 0}, Star{X: -3, Y: 6, Vx: 2, Vy: -1}},
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
