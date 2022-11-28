package crabcombat_test

import (
	. "github.com/pjaneiro/advent-of-code/2020/crabcombat"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		deck1    Deck
		deck2    Deck
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			deck1:    Deck{9, 2, 6, 3, 1},
			deck2:    Deck{5, 8, 4, 7, 10},
			expected: 306,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.deck1, tc.deck2)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %v) threw '%v', want %d", tc.deck1, tc.deck2, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %v) = %d, want to throw", tc.deck1, tc.deck2, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v) = %d, want %d", tc.deck1, tc.deck2, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		deck1    Deck
		deck2    Deck
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			deck1:    Deck{9, 2, 6, 3, 1},
			deck2:    Deck{5, 8, 4, 7, 10},
			expected: 291,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.deck1, tc.deck2)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v, %v) threw '%v', want %d", tc.deck1, tc.deck2, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v, %v) = %d, want to throw", tc.deck1, tc.deck2, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v, %v) = %d, want %d", tc.deck1, tc.deck2, actual, tc.expected)
			}
		})
	}
}
