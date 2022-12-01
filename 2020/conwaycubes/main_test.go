package conwaycubes_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2020/conwaycubes"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    [][][]rune
		expected int
	}{
		{
			name: "Example 1",
			input: [][][]rune{
				{
					{'.', '#', '.'},
					{'.', '.', '#'},
					{'#', '#', '#'},
				},
			},
			expected: 112,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := Challenge1(tc.input)
			if actual != tc.expected {
				t.Errorf("Challenge1(%v) = %d, want %d", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input    [][][][]rune
		expected int
	}{
		{
			name: "Example 1",
			input: [][][][]rune{
				{
					{
						{'.', '#', '.'},
						{'.', '.', '#'},
						{'#', '#', '#'},
					},
				},
			},
			expected: 848,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := Challenge2(tc.input)
			if actual != tc.expected {
				t.Errorf("Challenge2(%v) = %d, want %d", tc.input, actual, tc.expected)
			}
		})
	}
}
