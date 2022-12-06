package supplystacks_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2022/supplystacks"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		state    [][]byte
		input    []Instruction
		expected string
		error    bool
	}{
		{
			name:     "Example 1",
			state:    [][]byte{[]byte{'Z', 'N'}, []byte{'M', 'C', 'D'}, []byte{'P'}},
			input:    []Instruction{Instruction{Amt: 1, From: 2, To: 1}, Instruction{Amt: 3, From: 1, To: 3}, Instruction{Amt: 2, From: 2, To: 1}, Instruction{Amt: 1, From: 1, To: 2}},
			expected: "CMZ",
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.state, tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v) threw '%v', want %s", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v) = %s, want to throw", tc.input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v) = %s, want %s", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		state    [][]byte
		input    []Instruction
		expected string
		error    bool
	}{
		{
			name:     "Example 1",
			state:    [][]byte{[]byte{'Z', 'N'}, []byte{'M', 'C', 'D'}, []byte{'P'}},
			input:    []Instruction{Instruction{Amt: 1, From: 2, To: 1}, Instruction{Amt: 3, From: 1, To: 3}, Instruction{Amt: 2, From: 2, To: 1}, Instruction{Amt: 1, From: 1, To: 2}},
			expected: "MCD",
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.state, tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v) threw '%v', want %s", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v) = %s, want to throw", tc.input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v) = %s, want %s", tc.input, actual, tc.expected)
			}
		})
	}
}
