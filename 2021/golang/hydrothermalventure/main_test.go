package hydrothermalventure_test

import (
	. "github.com/pjaneiro/advent-of-code/2021/golang/hydrothermalventure"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Vector
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Vector{},
			expected: 0,
			error:    false,
		},
		{
			name: "Example 2",
			input: []Vector{
				{Point{0, 9}, Point{5, 9}},
				{Point{8, 0}, Point{0, 8}},
				{Point{9, 4}, Point{3, 4}},
				{Point{2, 2}, Point{2, 1}},
				{Point{7, 0}, Point{7, 4}},
				{Point{6, 4}, Point{2, 0}},
				{Point{0, 9}, Point{2, 9}},
				{Point{3, 4}, Point{1, 4}},
				{Point{0, 0}, Point{8, 8}},
				{Point{5, 5}, Point{8, 2}},
			},
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
		input    []Vector
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Vector{},
			expected: 0,
			error:    false,
		},
		{
			name: "Example 2",
			input: []Vector{
				{Point{0, 9}, Point{5, 9}},
				{Point{8, 0}, Point{0, 8}},
				{Point{9, 4}, Point{3, 4}},
				{Point{2, 2}, Point{2, 1}},
				{Point{7, 0}, Point{7, 4}},
				{Point{6, 4}, Point{2, 0}},
				{Point{0, 9}, Point{2, 9}},
				{Point{3, 4}, Point{1, 4}},
				{Point{0, 0}, Point{8, 8}},
				{Point{5, 5}, Point{8, 2}},
			},
			expected: 12,
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
