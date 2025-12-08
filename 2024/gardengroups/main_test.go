package gardengroups_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2024/gardengroups"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []string{"AAAA", "BBCD", "BBCC", "EEEC"},
			expected: 140,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []string{"OOOOO", "OXOXO", "OOOOO", "OXOXO", "OOOOO"},
			expected: 772,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    []string{"RRRRIICCFF", "RRRRIICCCF", "VVRRRCCFFF", "VVRCCCJFFF", "VVVVCJJCFE", "VVIVCCJJEE", "VVIIICJJEE", "MIIIIIJJEE", "MIIISIJEEE", "MMMISSJEEE"},
			expected: 1930,
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
		input    []string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []string{"AAAA", "BBCD", "BBCC", "EEEC"},
			expected: 80,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []string{"OOOOO", "OXOXO", "OOOOO", "OXOXO", "OOOOO"},
			expected: 436,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    []string{"EEEEE", "EXXXX", "EEEEE", "EXXXX", "EEEEE"},
			expected: 236,
			error:    false,
		},
		{
			name:     "Example 4",
			input:    []string{"AAAAAA", "AAABBA", "AAABBA", "ABBAAA", "ABBAAA", "AAAAAA"},
			expected: 368,
			error:    false,
		},
		{
			name:     "Example 5",
			input:    []string{"RRRRIICCFF", "RRRRIICCCF", "VVRRRCCFFF", "VVRCCCJFFF", "VVVVCJJCFE", "VVIVCCJJEE", "VVIIICJJEE", "MIIIIIJJEE", "MIIISIJEEE", "MMMISSJEEE"},
			expected: 1206,
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
