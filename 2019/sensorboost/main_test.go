package sensorboost_test

import (
	. "github.com/pjaneiro/advent-of-code-2019/sensorboost"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    map[int64]int64
		expected int64
		error    bool
	}{
		{
			name:     "Example 1",
			input:    map[int64]int64{0: 109, 1: 1, 2: 204, 3: -1, 4: 1001, 5: 100, 6: 1, 7: 100, 8: 1008, 9: 100, 10: 16, 11: 101, 12: 1006, 13: 101, 14: 0, 15: 99},
			expected: 99,
			error:    false,
		}, {
			name:     "Example 2",
			input:    map[int64]int64{0: 1102, 1: 34915192, 2: 34915192, 3: 7, 4: 4, 5: 7, 6: 99, 7: 0},
			expected: 1219070632396864,
			error:    false,
		}, {
			name:     "Example 3",
			input:    map[int64]int64{0: 104, 1: 1125899906842624, 2: 99},
			expected: 1125899906842624,
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
