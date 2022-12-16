package beaconexclusionzone_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2022/beaconexclusionzone"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name       string
		input_path string
		row        int
		expected   int
		error      bool
	}{
		{
			name:       "Example 1",
			input_path: "input_test.txt",
			row:        10,
			expected:   26,
			error:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input, err := ReadLines(tc.input_path)
			if err != nil {
				t.Errorf("Error reading input")
				t.FailNow()
			}
			actual, err := Challenge1(input, tc.row)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %v) threw '%v', want %d", input, tc.row, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %v) = %d, want to throw", input, tc.row, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v) = %d, want %d", input, tc.row, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name       string
		input_path string
		maxCoord   int
		expected   int
		error      bool
	}{
		{
			name:       "Example 1",
			input_path: "input_test.txt",
			maxCoord:   20,
			expected:   56000011,
			error:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input, err := ReadLines(tc.input_path)
			if err != nil {
				t.Errorf("Error reading input")
				t.FailNow()
			}
			actual, err := Challenge2(input, tc.maxCoord)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v, %v) threw '%v', want %d", input, tc.maxCoord, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v, %v) = %d, want to throw", input, tc.maxCoord, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v, %v) = %d, want %d", input, tc.maxCoord, actual, tc.expected)
			}
		})
	}
}
