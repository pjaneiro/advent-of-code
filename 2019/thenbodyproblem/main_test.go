package thenbodyproblem_test

import (
	. "github.com/pjaneiro/advent-of-code/2019/thenbodyproblem"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Moon
		steps    int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Moon{Moon{X: -1, Y: 0, Z: 2, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 2, Y: -10, Z: -7, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 4, Y: -8, Z: 8, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 3, Y: 5, Z: -1, Vx: 0, Vy: 0, Vz: 0}},
			steps:    10,
			expected: 179,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []Moon{Moon{X: -8, Y: -10, Z: 0, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 5, Y: 5, Z: 10, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 2, Y: -7, Z: 3, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 9, Y: -8, Z: -3, Vx: 0, Vy: 0, Vz: 0}},
			steps:    100,
			expected: 1940,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.input, tc.steps)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %v) threw '%v', want %d", tc.input, tc.steps, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %v) = %d, want to throw", tc.input, tc.steps, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v) = %d, want %d", tc.input, tc.steps, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Moon
		expected int64
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Moon{Moon{X: -1, Y: 0, Z: 2, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 2, Y: -10, Z: -7, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 4, Y: -8, Z: 8, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 3, Y: 5, Z: -1, Vx: 0, Vy: 0, Vz: 0}},
			expected: 2772,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []Moon{Moon{X: -8, Y: -10, Z: 0, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 5, Y: 5, Z: 10, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 2, Y: -7, Z: 3, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 9, Y: -8, Z: -3, Vx: 0, Vy: 0, Vz: 0}},
			expected: 4686774924,
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
