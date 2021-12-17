package handyhaversacks_test

import (
	. "github.com/pjaneiro/advent-of-code-2020/handyhaversacks"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    map[string]map[string]int
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    map[string]map[string]int{"light red": {"bright white": 1, "muted yellow": 2}},
			expected: 0,
			error:    false,
		}, {
			name:     "Example 2",
			input:    map[string]map[string]int{"bright white": {"shiny gold": 1}},
			expected: 1,
			error:    false,
		}, {
			name: "Example 3",
			input: map[string]map[string]int{
				"light red":    {"bright white": 1, "muted yellow": 2},
				"dark orange":  {"bright white": 3, "muted yellow": 4},
				"bright white": {"shiny gold": 1},
				"muted yellow": {"shiny gold": 2, "faded blue": 9},
				"shiny gold":   {"dark olive": 1, "vibrant plum": 2},
				"dark olive":   {"faded blue": 3, "dotted black": 4},
				"vibrant plum": {"faded blue": 5, "dotted black": 6},
				"faded blue":   {},
				"dotted black": {},
			},
			expected: 4,
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
		input    map[string]map[string]int
		expected int
		error    bool
	}{
		{
			name: "Example 1",
			input: map[string]map[string]int{
				"shiny gold":  {"dark red": 2},
				"dark red":    {"dark orange": 2},
				"dark orange": {"dark yellow": 2},
				"dark yellow": {"dark green": 2},
				"dark green":  {"dark blue": 2},
				"dark blue":   {"dark violet": 2},
				"dark violet": {},
			},
			expected: 126,
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
