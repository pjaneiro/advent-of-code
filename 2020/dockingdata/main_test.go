package dockingdata_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2020/dockingdata"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Operation
		expected int64
	}{
		{
			name:     "Example 1",
			input:    []Operation{},
			expected: 0,
		}, {
			name: "Example 2",
			input: []Operation{
				{Mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", Address: 8, Value: 11},
				{Mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", Address: 7, Value: 101},
				{Mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", Address: 8, Value: 0},
			},
			expected: 165,
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
		input    []Operation
		expected int64
	}{
		{
			name:     "Example 1",
			input:    []Operation{},
			expected: 0,
		}, {
			name: "Example 2",
			input: []Operation{
				{Mask: "000000000000000000000000000000X1001X", Address: 42, Value: 100},
				{Mask: "00000000000000000000000000000000X0XX", Address: 26, Value: 1},
			},
			expected: 208,
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
