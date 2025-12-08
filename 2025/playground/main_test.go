package playground_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2025/playground"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []string
		steps    int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []string{"162,817,812", "57,618,57", "906,360,560", "592,479,940", "352,342,300", "466,668,158", "542,29,236", "431,825,988", "739,650,466", "52,470,668", "216,146,977", "819,987,18", "117,168,530", "805,96,715", "346,949,466", "970,615,88", "941,993,340", "862,61,35", "984,92,344", "425,690,689"},
			steps:    10,
			expected: 40,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.input, tc.steps)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %d) threw '%v', want %d", tc.input, tc.steps, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %d) = %d, want to throw", tc.input, tc.steps, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %d) = %d, want %d", tc.input, tc.steps, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []string{"162,817,812", "57,618,57", "906,360,560", "592,479,940", "352,342,300", "466,668,158", "542,29,236", "431,825,988", "739,650,466", "52,470,668", "216,146,977", "819,987,18", "117,168,530", "805,96,715", "346,949,466", "970,615,88", "941,993,340", "862,61,35", "984,92,344", "425,690,689"},
			expected: 25272,
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
