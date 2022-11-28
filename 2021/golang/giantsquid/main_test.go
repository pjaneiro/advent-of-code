package giantsquid_test

import (
	. "github.com/pjaneiro/advent-of-code/2021/golang/giantsquid"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		order    []int
		boards   [][][]House
		expected int
		error    bool
	}{
		{
			name:  "Example 1",
			order: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
			boards: [][][]House{
				{
					{House{22, false}, House{13, false}, House{17, false}, House{11, false}, House{0, false}},
					{House{8, false}, House{2, false}, House{23, false}, House{4, false}, House{24, false}},
					{House{21, false}, House{9, false}, House{14, false}, House{16, false}, House{7, false}},
					{House{6, false}, House{10, false}, House{3, false}, House{18, false}, House{5, false}},
					{House{1, false}, House{12, false}, House{20, false}, House{15, false}, House{19, false}},
				},
				{
					{House{3, false}, House{15, false}, House{0, false}, House{2, false}, House{22, false}},
					{House{9, false}, House{18, false}, House{13, false}, House{17, false}, House{5, false}},
					{House{19, false}, House{8, false}, House{7, false}, House{25, false}, House{23, false}},
					{House{20, false}, House{11, false}, House{10, false}, House{24, false}, House{4, false}},
					{House{14, false}, House{21, false}, House{16, false}, House{12, false}, House{6, false}},
				},
				{
					{House{14, false}, House{21, false}, House{17, false}, House{24, false}, House{4, false}},
					{House{10, false}, House{16, false}, House{15, false}, House{9, false}, House{19, false}},
					{House{18, false}, House{8, false}, House{23, false}, House{26, false}, House{20, false}},
					{House{22, false}, House{11, false}, House{13, false}, House{6, false}, House{5, false}},
					{House{2, false}, House{0, false}, House{12, false}, House{3, false}, House{7, false}},
				},
			},
			expected: 4512,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.order, tc.boards)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %v) threw '%v', want %d", tc.order, tc.boards, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %v) = %d, want to throw", tc.order, tc.boards, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v) = %d, want %d", tc.order, tc.boards, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		order    []int
		boards   [][][]House
		expected int
		error    bool
	}{
		{
			name:  "Example 1",
			order: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
			boards: [][][]House{
				{
					{House{22, false}, House{13, false}, House{17, false}, House{11, false}, House{0, false}},
					{House{8, false}, House{2, false}, House{23, false}, House{4, false}, House{24, false}},
					{House{21, false}, House{9, false}, House{14, false}, House{16, false}, House{7, false}},
					{House{6, false}, House{10, false}, House{3, false}, House{18, false}, House{5, false}},
					{House{1, false}, House{12, false}, House{20, false}, House{15, false}, House{19, false}},
				},
				{
					{House{3, false}, House{15, false}, House{0, false}, House{2, false}, House{22, false}},
					{House{9, false}, House{18, false}, House{13, false}, House{17, false}, House{5, false}},
					{House{19, false}, House{8, false}, House{7, false}, House{25, false}, House{23, false}},
					{House{20, false}, House{11, false}, House{10, false}, House{24, false}, House{4, false}},
					{House{14, false}, House{21, false}, House{16, false}, House{12, false}, House{6, false}},
				},
				{
					{House{14, false}, House{21, false}, House{17, false}, House{24, false}, House{4, false}},
					{House{10, false}, House{16, false}, House{15, false}, House{9, false}, House{19, false}},
					{House{18, false}, House{8, false}, House{23, false}, House{26, false}, House{20, false}},
					{House{22, false}, House{11, false}, House{13, false}, House{6, false}, House{5, false}},
					{House{2, false}, House{0, false}, House{12, false}, House{3, false}, House{7, false}},
				},
			},
			expected: 1924,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.order, tc.boards)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v, %v) threw '%v', want %d", tc.order, tc.boards, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v, %v) = %d, want to throw", tc.order, tc.boards, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v, %v) = %d, want %d", tc.order, tc.boards, actual, tc.expected)
			}
		})
	}
}
