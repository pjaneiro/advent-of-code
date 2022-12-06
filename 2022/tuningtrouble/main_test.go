package tuningtrouble_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2022/tuningtrouble"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 7,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 5,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 6,
			error:    false,
		},
		{
			name:     "Example 4",
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 10,
			error:    false,
		},
		{
			name:     "Example 5",
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
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
		input    string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 19,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 23,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 23,
			error:    false,
		},
		{
			name:     "Example 4",
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 29,
			error:    false,
		},
		{
			name:     "Example 5",
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 26,
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
