package restroomredoubt_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2024/restroomredoubt"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []string
		width    int
		height   int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []string{"p=0,4 v=3,-3", "p=6,3 v=-1,-3", "p=10,3 v=-1,2", "p=2,0 v=2,-1", "p=0,0 v=1,3", "p=3,0 v=-2,-2", "p=7,6 v=-1,-3", "p=3,0 v=-1,-2", "p=9,3 v=2,3", "p=7,3 v=-1,2", "p=2,4 v=2,-3", "p=9,5 v=-3,-3"},
			width:    11,
			height:   7,
			expected: 12,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.input, tc.width, tc.height)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %d, %d) threw '%v', want %d", tc.input, tc.width, tc.height, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %d, %d) = %d, want to throw", tc.input, tc.width, tc.height, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %d, %d) = %d, want %d", tc.input, tc.width, tc.height, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []string
		width    int
		height   int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []string{"p=0,4 v=3,-3", "p=6,3 v=-1,-3", "p=10,3 v=-1,2", "p=2,0 v=2,-1", "p=0,0 v=1,3", "p=3,0 v=-2,-2", "p=7,6 v=-1,-3", "p=3,0 v=-1,-2", "p=9,3 v=2,3", "p=7,3 v=-1,2", "p=2,4 v=2,-3", "p=9,5 v=-3,-3"},
			width:    11,
			height:   7,
			expected: 12,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.input, tc.width, tc.height)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v, %d, %d) threw '%v', want %d", tc.input, tc.width, tc.height, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v, %d, %d) = %d, want to throw", tc.input, tc.width, tc.height, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v, %d, %d) = %d, want %d", tc.input, tc.width, tc.height, actual, tc.expected)
			}
		})
	}
}
