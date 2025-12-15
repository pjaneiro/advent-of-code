package christmastreefarm_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2025/christmastreefarm"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []string{"0:", "###", "##.", "##.", "", "1:", "###", "##.", ".##", "", "2:", ".##", "###", "##.", "", "3:", "##.", "###", "##.", "", "4:", "###", "#..", "###", "", "5:", "###", ".#.", "###", "", "4x4: 0 0 0 0 2 0", "12x5: 1 0 1 0 2 2", "12x5: 1 0 1 0 3 2"},
			expected: 2,
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
