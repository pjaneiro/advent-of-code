package sensorboost_test

import (
	. "github.com/pjaneiro/advent-of-code-2019/sensorboost"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		data     []int64
		input    []int64
		expected int64
		error    bool
	}{
		{
			name:     "Example 1",
			data:     []int64{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99},
			input:    []int64{1},
			expected: 43210,
			error:    false,
		}, {
			name:     "Example 2",
			data:     []int64{1102,34915192,34915192,7,4,7,99,0},
			input:    []int64{1},
			expected: 54321,
			error:    false,
		}, {
			name:     "Example 3",
			data:     []int64{104,1125899906842624,99},
			input:    []int64{1},
			expected: 65210,
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
