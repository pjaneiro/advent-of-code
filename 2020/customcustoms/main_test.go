package customcustoms_test

import (
	. "github.com/pjaneiro/advent-of-code-2020/customcustoms"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Form
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Form{{Npeople: 1, Answers: map[rune]int{'a': 1, 'b': 1, 'c': 1}}},
			expected: 3,
			error:    false,
		}, {
			name:     "Example 2",
			input:    []Form{{Npeople: 3, Answers: map[rune]int{'a': 1, 'b': 1, 'c': 1}}},
			expected: 3,
			error:    false,
		}, {
			name:     "Example 3",
			input:    []Form{{Npeople: 2, Answers: map[rune]int{'a': 2, 'b': 1, 'c': 1}}},
			expected: 3,
			error:    false,
		}, {
			name:     "Example 4",
			input:    []Form{{Npeople: 4, Answers: map[rune]int{'a': 4}}},
			expected: 1,
			error:    false,
		}, {
			name:     "Example 5",
			input:    []Form{{Npeople: 1, Answers: map[rune]int{'b': 1}}},
			expected: 1,
			error:    false,
		}, {
			name:     "Example 6",
			input:    []Form{{Npeople: 1, Answers: map[rune]int{'a': 1, 'b': 1, 'c': 1}}, {Npeople: 3, Answers: map[rune]int{'a': 1, 'b': 1, 'c': 1}}, {Npeople: 2, Answers: map[rune]int{'a': 2, 'b': 1, 'c': 1}}, {Npeople: 4, Answers: map[rune]int{'a': 4}}, {Npeople: 1, Answers: map[rune]int{'b': 1}}},
			expected: 11,
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
		input    []Form
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Form{{Npeople: 1, Answers: map[rune]int{'a': 1, 'b': 1, 'c': 1}}},
			expected: 3,
			error:    false,
		}, {
			name:     "Example 2",
			input:    []Form{{Npeople: 3, Answers: map[rune]int{'a': 1, 'b': 1, 'c': 1}}},
			expected: 0,
			error:    false,
		}, {
			name:     "Example 3",
			input:    []Form{{Npeople: 2, Answers: map[rune]int{'a': 2, 'b': 1, 'c': 1}}},
			expected: 1,
			error:    false,
		}, {
			name:     "Example 4",
			input:    []Form{{Npeople: 4, Answers: map[rune]int{'a': 4}}},
			expected: 1,
			error:    false,
		}, {
			name:     "Example 5",
			input:    []Form{{Npeople: 1, Answers: map[rune]int{'b': 1}}},
			expected: 1,
			error:    false,
		}, {
			name:     "Example 6",
			input:    []Form{{Npeople: 1, Answers: map[rune]int{'a': 1, 'b': 1, 'c': 1}}, {Npeople: 3, Answers: map[rune]int{'a': 1, 'b': 1, 'c': 1}}, {Npeople: 2, Answers: map[rune]int{'a': 2, 'b': 1, 'c': 1}}, {Npeople: 4, Answers: map[rune]int{'a': 4}}, {Npeople: 1, Answers: map[rune]int{'b': 1}}},
			expected: 6,
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
