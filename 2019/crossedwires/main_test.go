package crossedwires_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2019/crossedwires"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    [][]Instruction
		expected int
		error    bool
	}{
		{
			name: "Example 1",
			input: [][]Instruction{
				{
					{Dir: 'R', Val: 75}, {Dir: 'D', Val: 30}, {Dir: 'R', Val: 83}, {Dir: 'U', Val: 83}, {Dir: 'L', Val: 12}, {Dir: 'D', Val: 49}, {Dir: 'R', Val: 71}, {Dir: 'U', Val: 7}, {Dir: 'L', Val: 72},
				}, {
					{Dir: 'U', Val: 62}, {Dir: 'R', Val: 66}, {Dir: 'U', Val: 55}, {Dir: 'R', Val: 34}, {Dir: 'D', Val: 71}, {Dir: 'R', Val: 55}, {Dir: 'D', Val: 58}, {Dir: 'R', Val: 83},
				},
			},
			expected: 159,
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
		input    [][]Instruction
		expected int
		error    bool
	}{
		{
			name: "Example 1",
			input: [][]Instruction{
				{
					{Dir: 'R', Val: 75}, {Dir: 'D', Val: 30}, {Dir: 'R', Val: 83}, {Dir: 'U', Val: 83}, {Dir: 'L', Val: 12}, {Dir: 'D', Val: 49}, {Dir: 'R', Val: 71}, {Dir: 'U', Val: 7}, {Dir: 'L', Val: 72},
				}, {
					{Dir: 'U', Val: 62}, {Dir: 'R', Val: 66}, {Dir: 'U', Val: 55}, {Dir: 'R', Val: 34}, {Dir: 'D', Val: 71}, {Dir: 'R', Val: 55}, {Dir: 'D', Val: 58}, {Dir: 'R', Val: 83},
				},
			},
			expected: 610,
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
