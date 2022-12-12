package monkeyinthemiddle_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2022/monkeyinthemiddle"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Monkey
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Monkey{Monkey{Items: []int{79, 98}, Operator: '*', Factor: 19, Test: 23, True: 2, False: 3}, Monkey{Items: []int{54, 65, 75, 74}, Operator: '+', Factor: 6, Test: 19, True: 2, False: 0}, Monkey{Items: []int{79, 60, 97}, Operator: '*', Factor: -1, Test: 13, True: 1, False: 3}, Monkey{Items: []int{74}, Operator: '+', Factor: 3, Test: 17, True: 0, False: 1}},
			expected: 10605,
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
		input    []Monkey
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Monkey{Monkey{Items: []int{79, 98}, Operator: '*', Factor: 19, Test: 23, True: 2, False: 3}, Monkey{Items: []int{54, 65, 75, 74}, Operator: '+', Factor: 6, Test: 19, True: 2, False: 0}, Monkey{Items: []int{79, 60, 97}, Operator: '*', Factor: -1, Test: 13, True: 1, False: 3}, Monkey{Items: []int{74}, Operator: '+', Factor: 3, Test: 17, True: 0, False: 1}},
			expected: 2713310158,
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
