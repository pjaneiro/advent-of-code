package shuttlesearch_test

import (
	. "github.com/pjaneiro/advent-of-code-2020/shuttlesearch"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		goal     int
		ids      []string
		expected int
	}{
		{
			name:     "Example 1",
			goal:     939,
			ids:      []string{"7", "13", "x", "x", "59", "x", "31", "19"},
			expected: 295,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := Challenge1(tc.goal, tc.ids)
			if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v) = %d, want %d", tc.goal, tc.ids, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		ids      []string
		expected int64
	}{
		{
			name:     "Example 1",
			ids:      []string{"7", "13", "x", "x", "59", "x", "31", "19"},
			expected: 1068781,
		}, {
			name:     "Example 2",
			ids:      []string{"17", "x", "13", "19"},
			expected: 3417,
		}, {
			name:     "Example 3",
			ids:      []string{"67", "7", "59", "61"},
			expected: 754018,
		}, {
			name:     "Example 4",
			ids:      []string{"67", "x", "7", "59", "61"},
			expected: 779210,
		}, {
			name:     "Example 5",
			ids:      []string{"67", "7", "x", "59", "61"},
			expected: 1261476,
		}, {
			name:     "Example 6",
			ids:      []string{"1789", "37", "47", "1889"},
			expected: 1202161486,
		}, {
			name:     "Example 7",
			ids:      []string{"13", "x", "x", "41", "x", "x", "x", "x", "x", "x", "x", "x", "x", "467", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "19", "x", "x", "x", "x", "17", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "29", "x", "353", "x", "x", "x", "x", "x", "37", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "23"},
			expected: 672754131923874,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := Challenge2(tc.ids)
			if actual != tc.expected {
				t.Errorf("Challenge2(%v) = %d, want %d", tc.ids, actual, tc.expected)
			}
		})
	}
}
