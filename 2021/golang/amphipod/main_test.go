package amphipod_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2021/golang/amphipod"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    State
		expected int
		error    bool
	}{
		{
			name: "Example 1",
			input: State{
				Cols: []Column{
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'A', 'B'}, Size: 2, Type: 'A'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'D', 'C'}, Size: 2, Type: 'B'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'C', 'B'}, Size: 2, Type: 'C'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'A', 'D'}, Size: 2, Type: 'D'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
				},
			},
			expected: 12521,
			error:    false,
		},
		{
			name: "Example 2",
			input: State{
				Cols: []Column{
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'D', 'B'}, Size: 2, Type: 'A'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'D', 'C'}, Size: 2, Type: 'B'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'A', 'C'}, Size: 2, Type: 'C'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'A', 'B'}, Size: 2, Type: 'D'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
				},
			},
			expected: 18051,
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
		input    State
		expected int
		error    bool
	}{
		{
			name: "Example 1",
			input: State{
				Cols: []Column{
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'A', 'D', 'D', 'B'}, Size: 4, Type: 'A'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'D', 'B', 'C', 'C'}, Size: 4, Type: 'B'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'C', 'A', 'B', 'B'}, Size: 4, Type: 'C'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'A', 'C', 'A', 'D'}, Size: 4, Type: 'D'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
				},
			},
			expected: 44169,
			error:    false,
		},
		{
			name: "Example 2",
			input: State{
				Cols: []Column{
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'D', 'D', 'D', 'B'}, Size: 4, Type: 'A'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'D', 'B', 'C', 'C'}, Size: 4, Type: 'B'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'A', 'A', 'B', 'C'}, Size: 4, Type: 'C'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{'A', 'C', 'A', 'B'}, Size: 4, Type: 'D'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
					Column{Data: []byte{}, Size: 1, Type: '.'},
				},
			},
			expected: 50245,
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
