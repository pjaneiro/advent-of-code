package combobreaker_test

import (
	. "github.com/pjaneiro/advent-of-code/2020/combobreaker"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		cardPK   int
		doorPK   int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			cardPK:   5764801,
			doorPK:   17807724,
			expected: 14897079,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.cardPK, tc.doorPK)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %v) threw '%v', want %d", tc.cardPK, tc.doorPK, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %v) = %d, want to throw", tc.cardPK, tc.doorPK, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v) = %d, want %d", tc.cardPK, tc.doorPK, actual, tc.expected)
			}
		})
	}
}
