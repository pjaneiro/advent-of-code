package extendedpolymerization_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2021/golang/extendedpolymerization"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		starter  string
		rules    map[string]map[string]string
		expected int64
		error    bool
	}{
		{
			name:     "Example 1",
			starter:  "NNCB",
			rules:    map[string]map[string]string{"B": {"B": "N", "C": "B", "H": "H", "N": "B"}, "C": {"B": "H", "C": "N", "H": "B", "N": "C"}, "H": {"B": "C", "C": "B", "H": "N", "N": "C"}, "N": {"B": "B", "C": "B", "H": "C", "N": "C"}},
			expected: 1588,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.starter, tc.rules)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %v) threw '%v', want %d", tc.starter, tc.rules, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %v) = %d, want to throw", tc.starter, tc.rules, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v) = %d, want %d", tc.starter, tc.rules, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		starter  string
		rules    map[string]map[string]string
		expected int64
		error    bool
	}{
		{
			name:     "Example 1",
			starter:  "NNCB",
			rules:    map[string]map[string]string{"B": {"B": "N", "C": "B", "H": "H", "N": "B"}, "C": {"B": "H", "C": "N", "H": "B", "N": "C"}, "H": {"B": "C", "C": "B", "H": "N", "N": "C"}, "N": {"B": "B", "C": "B", "H": "C", "N": "C"}},
			expected: 2188189693529,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.starter, tc.rules)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v, %v) threw '%v', want %d", tc.starter, tc.rules, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v, %v) = %d, want to throw", tc.starter, tc.rules, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v, %v) = %d, want %d", tc.starter, tc.rules, actual, tc.expected)
			}
		})
	}
}
