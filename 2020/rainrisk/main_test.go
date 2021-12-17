package rainrisk_test

import (
	. "github.com/pjaneiro/advent-of-code-2020/rainrisk"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Instruction
		expected int
	}{
		{
			name: "Example 1",
			input: []Instruction{
				{Dir: 'F', Val: 10},
				{Dir: 'N', Val: 3},
				{Dir: 'F', Val: 7},
				{Dir: 'R', Val: 90},
				{Dir: 'F', Val: 11},
			},
			expected: 25,
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
		input    []Instruction
		expected int
	}{
		{
			name: "Example 1",
			input: []Instruction{
				{Dir: 'F', Val: 10},
				{Dir: 'N', Val: 3},
				{Dir: 'F', Val: 7},
				{Dir: 'R', Val: 90},
				{Dir: 'F', Val: 11},
			},
			expected: 286,
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
