package tickettranslation_test

import (
	. "github.com/pjaneiro/advent-of-code-2020/tickettranslation"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		ticket   []int
		tickets  [][]int
		fields   []Field
		expected int
	}{
		{
			name:   "Example 1",
			ticket: []int{7, 1, 14},
			tickets: [][]int{
				{7, 3, 47},
				{40, 4, 50},
				{55, 2, 20},
				{38, 6, 12},
			},
			fields: []Field{
				Field{Name: "class", Min1: 1, Max1: 3, Min2: 5, Max2: 7},
				Field{Name: "row", Min1: 6, Max1: 11, Min2: 33, Max2: 44},
				Field{Name: "seat", Min1: 13, Max1: 40, Min2: 45, Max2: 50},
			},
			expected: 71,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := Challenge1(tc.tickets, tc.ticket, tc.fields)
			if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v, %v) = %d, want %d", tc.tickets, tc.ticket, tc.fields, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		ticket   []int
		tickets  [][]int
		fields   []Field
		expected int
	}{
		{
			name:   "Example 1",
			ticket: []int{11, 12, 13},
			tickets: [][]int{
				{3, 9, 18},
				{15, 1, 5},
				{5, 14, 9},
			},
			fields: []Field{
				Field{Name: "class", Min1: 0, Max1: 1, Min2: 4, Max2: 19},
				Field{Name: "row", Min1: 0, Max1: 5, Min2: 8, Max2: 19},
				Field{Name: "seat", Min1: 0, Max1: 13, Min2: 16, Max2: 19},
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := Challenge2(tc.tickets, tc.ticket, tc.fields)
			if actual != tc.expected {
				t.Errorf("Challenge2(%v, %v, %v) = %d, want %d", tc.tickets, tc.ticket, tc.fields, actual, tc.expected)
			}
		})
	}
}
