package week1

import (
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []int{},
			expected: 0,
			error:    true,
		}, {
			name:     "Example 2",
			input:    []int{2020},
			expected: 0,
			error:    true,
		}, {
			name:     "Example 3",
			input:    []int{2020, 0},
			expected: 0,
			error:    false,
		}, {
			name:     "Example 4",
			input:    []int{1010, 1010},
			expected: 1020100,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := challenge1(tc.input)
			if err != nil && tc.error == false {
				t.Errorf("expected %d, threw '%v'", tc.expected, err)
			} else if err == nil && tc.error == true {
				t.Errorf("expected to throw, got %d", actual)
			} else if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []int{},
			expected: 0,
			error:    true,
		}, {
			name:     "Example 2",
			input:    []int{2020},
			expected: 0,
			error:    true,
		}, {
			name:     "Example 3",
			input:    []int{2020, 0, 0},
			expected: 0,
			error:    false,
		}, {
			name:     "Example 4",
			input:    []int{1000, 1000, 20},
			expected: 20000000,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := challenge2(tc.input)
			if err != nil && tc.error == false {
				t.Errorf("expected %d, threw '%v'", tc.expected, err)
			} else if err == nil && tc.error == true {
				t.Errorf("expected to throw, got %d", actual)
			} else if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
