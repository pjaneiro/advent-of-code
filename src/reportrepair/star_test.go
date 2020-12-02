package reportrepair_test

import (
	"fmt"
	. "reportrepair"
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

func ExampleChallenge1() {
	result, err := Challenge1([]int{1010, 1010})
	fmt.Printf("Result is %d, error is %v\n", result, err)
	// Output: Result is 1020100, error is <nil>
}

func ExampleChallenge1_error() {
	result, err := Challenge1([]int{2020})
	fmt.Printf("Result is %d, error is '%v'\n", result, err)
	// Output: Result is 0, error is 'couldn't find a satisfying solution'
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

func ExampleChallenge2() {
	result, err := Challenge2([]int{1000, 1000, 20})
	fmt.Printf("Result is %d, error is %v\n", result, err)
	// Output: Result is 20000000, error is <nil>
}

func ExampleChallenge2_error() {
	result, err := Challenge2([]int{2020})
	fmt.Printf("Result is %d, error is '%v'\n", result, err)
	// Output: Result is 0, error is 'couldn't find a satisfying solution'
}
