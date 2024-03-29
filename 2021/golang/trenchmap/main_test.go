package trenchmap_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2021/golang/trenchmap"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		algo     string
		input    []string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			algo:     "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#",
			input:    []string{"#..#.", "#....", "##..#", "..#..", "..###"},
			expected: 35,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.algo, tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %v) threw '%v', want %d", tc.algo, tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %v) = %d, want to throw", tc.algo, tc.input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v) = %d, want %d", tc.algo, tc.input, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		algo     string
		input    []string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			algo:     "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#",
			input:    []string{"#..#.", "#....", "##..#", "..#..", "..###"},
			expected: 3351,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.algo, tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v, %v) threw '%v', want %d", tc.algo, tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v, %v) = %d, want to throw", tc.algo, tc.input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v, %v) = %d, want %d", tc.algo, tc.input, actual, tc.expected)
			}
		})
	}
}
