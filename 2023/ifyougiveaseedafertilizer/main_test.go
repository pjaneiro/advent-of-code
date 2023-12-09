package ifyougiveaseedafertilizer_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2023/ifyougiveaseedafertilizer"
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
			input:    []string{"seeds: 79 14 55 13", "", "seed-to-soil map:", "50 98 2", "52 50 48", "", "soil-to-fertilizer map:", "0 15 37", "37 52 2", "39 0 15", "", "fertilizer-to-water map:", "49 53 8", "0 11 42", "42 0 7", "57 7 4", "", "water-to-light map:", "88 18 7", "18 25 70", "", "light-to-temperature map:", "45 77 23", "81 45 19", "68 64 13", "", "temperature-to-humidity map:", "0 69 1", "1 0 69", "", "humidity-to-location map:", "60 56 37", "56 93 4"},
			expected: 35,
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
		input    []string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []string{"seeds: 79 14 55 13", "", "seed-to-soil map:", "50 98 2", "52 50 48", "", "soil-to-fertilizer map:", "0 15 37", "37 52 2", "39 0 15", "", "fertilizer-to-water map:", "49 53 8", "0 11 42", "42 0 7", "57 7 4", "", "water-to-light map:", "88 18 7", "18 25 70", "", "light-to-temperature map:", "45 77 23", "81 45 19", "68 64 13", "", "temperature-to-humidity map:", "0 69 1", "1 0 69", "", "humidity-to-location map:", "60 56 37", "56 93 4"},
			expected: 46,
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
