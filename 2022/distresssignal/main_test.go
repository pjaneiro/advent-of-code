package distresssignal_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2022/distresssignal"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    [][]string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    [][]string{[]string{"[1,1,3,1,1]", "[1,1,5,1,1]"}, []string{"[[1],[2,3,4]]", "[[1],4]"}, []string{"[9]", "[[8,7,6]]"}, []string{"[[4,4],4,4]", "[[4,4],4,4,4]"}, []string{"[7,7,7,7]", "[7,7,7]"}, []string{"[]", "[3]"}, []string{"[[[]]]", "[[]]"}, []string{"[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]"}},
			expected: 13,
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
		input    [][]string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    [][]string{[]string{"[1,1,3,1,1]", "[1,1,5,1,1]"}, []string{"[[1],[2,3,4]]", "[[1],4]"}, []string{"[9]", "[[8,7,6]]"}, []string{"[[4,4],4,4]", "[[4,4],4,4,4]"}, []string{"[7,7,7,7]", "[7,7,7]"}, []string{"[]", "[3]"}, []string{"[[[]]]", "[[]]"}, []string{"[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]"}},
			expected: 140,
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
