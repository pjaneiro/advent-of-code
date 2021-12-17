package handheldhalting_test

import (
	. "github.com/pjaneiro/advent-of-code-2020/handheldhalting"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Command
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Command{{Cmd: "nop", Val: 0}, {Cmd: "acc", Val: 1}, {Cmd: "jmp", Val: 4}, {Cmd: "acc", Val: 3}, {Cmd: "jmp", Val: -3}, {Cmd: "acc", Val: -99}, {Cmd: "acc", Val: 1}, {Cmd: "jmp", Val: -4}, {Cmd: "acc", Val: 6}},
			expected: 5,
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
		input    []Command
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Command{{Cmd: "nop", Val: 0}, {Cmd: "acc", Val: 1}, {Cmd: "jmp", Val: 4}, {Cmd: "acc", Val: 3}, {Cmd: "jmp", Val: -3}, {Cmd: "acc", Val: -99}, {Cmd: "acc", Val: 1}, {Cmd: "jmp", Val: -4}, {Cmd: "acc", Val: 6}},
			expected: 8,
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
