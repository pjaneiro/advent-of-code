package diracdice_test

import (
	. "github.com/pjaneiro/advent-of-code/2021/golang/diracdice"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		player1  int
		player2  int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			player1:  4,
			player2:  8,
			expected: 739785,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.player1, tc.player2)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %v) threw '%v', want %d", tc.player1, tc.player2, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %v) = %d, want to throw", tc.player1, tc.player2, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v) = %d, want %d", tc.player1, tc.player2, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		player1  int
		player2  int
		expected int64
		error    bool
	}{
		{
			name:     "Example 1",
			player1:  4,
			player2:  8,
			expected: 444356092776315,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.player1, tc.player2)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v, %v) threw '%v', want %d", tc.player1, tc.player2, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v, %v) = %d, want to throw", tc.player1, tc.player2, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v, %v) = %d, want %d", tc.player1, tc.player2, actual, tc.expected)
			}
		})
	}
}
