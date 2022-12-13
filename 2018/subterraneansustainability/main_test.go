package subterraneansustainability_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2018/subterraneansustainability"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		state    []bool
		rules    []Rule
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			state:    []bool{true, false, false, true, false, true, false, false, true, true, false, false, false, false, false, false, true, true, true, false, false, false, true, true, true},
			rules:    []Rule{Rule{Plants: [5]bool{false, false, false, true, true}, AddPlant: true}, Rule{Plants: [5]bool{false, false, true, false, false}, AddPlant: true}, Rule{Plants: [5]bool{false, true, false, false, false}, AddPlant: true}, Rule{Plants: [5]bool{false, true, false, true, false}, AddPlant: true}, Rule{Plants: [5]bool{false, true, false, true, true}, AddPlant: true}, Rule{Plants: [5]bool{false, true, true, false, false}, AddPlant: true}, Rule{Plants: [5]bool{false, true, true, true, true}, AddPlant: true}, Rule{Plants: [5]bool{true, false, true, false, true}, AddPlant: true}, Rule{Plants: [5]bool{true, false, true, true, true}, AddPlant: true}, Rule{Plants: [5]bool{true, true, false, true, false}, AddPlant: true}, Rule{Plants: [5]bool{true, true, false, true, true}, AddPlant: true}, Rule{Plants: [5]bool{true, true, true, false, false}, AddPlant: true}, Rule{Plants: [5]bool{true, true, true, false, true}, AddPlant: true}, Rule{Plants: [5]bool{true, true, true, true, false}, AddPlant: true}},
			expected: 325,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.state, tc.rules)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %v) threw '%v', want %d", tc.state, tc.rules, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %v) = %d, want to throw", tc.state, tc.rules, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v) = %d, want %d", tc.state, tc.rules, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		state    []bool
		rules    []Rule
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			state:    []bool{true, false, false, true, false, true, false, false, true, true, false, false, false, false, false, false, true, true, true, false, false, false, true, true, true},
			rules:    []Rule{Rule{Plants: [5]bool{false, false, false, true, true}, AddPlant: true}, Rule{Plants: [5]bool{false, false, true, false, false}, AddPlant: true}, Rule{Plants: [5]bool{false, true, false, false, false}, AddPlant: true}, Rule{Plants: [5]bool{false, true, false, true, false}, AddPlant: true}, Rule{Plants: [5]bool{false, true, false, true, true}, AddPlant: true}, Rule{Plants: [5]bool{false, true, true, false, false}, AddPlant: true}, Rule{Plants: [5]bool{false, true, true, true, true}, AddPlant: true}, Rule{Plants: [5]bool{true, false, true, false, true}, AddPlant: true}, Rule{Plants: [5]bool{true, false, true, true, true}, AddPlant: true}, Rule{Plants: [5]bool{true, true, false, true, false}, AddPlant: true}, Rule{Plants: [5]bool{true, true, false, true, true}, AddPlant: true}, Rule{Plants: [5]bool{true, true, true, false, false}, AddPlant: true}, Rule{Plants: [5]bool{true, true, true, false, true}, AddPlant: true}, Rule{Plants: [5]bool{true, true, true, true, false}, AddPlant: true}},
			expected: 999999999374,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.state, tc.rules)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v, %v) threw '%v', want %d", tc.state, tc.rules, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v, %v) = %d, want to throw", tc.state, tc.rules, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v, %v) = %d, want %d", tc.state, tc.rules, actual, tc.expected)
			}
		})
	}
}
