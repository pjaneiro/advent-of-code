package transparentorigami_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2021/golang/transparentorigami"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		points   []Point
		folds    []Fold
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			points:   []Point{Point{X: 6, Y: 10}, Point{X: 0, Y: 14}, Point{X: 9, Y: 10}, Point{X: 0, Y: 3}, Point{X: 10, Y: 4}, Point{X: 4, Y: 11}, Point{X: 6, Y: 0}, Point{X: 6, Y: 12}, Point{X: 4, Y: 1}, Point{X: 0, Y: 13}, Point{X: 10, Y: 12}, Point{X: 3, Y: 4}, Point{X: 3, Y: 0}, Point{X: 8, Y: 4}, Point{X: 1, Y: 10}, Point{X: 2, Y: 14}, Point{X: 8, Y: 10}, Point{X: 9, Y: 0}},
			folds:    []Fold{Fold{Axis: "y", Pos: 7}, Fold{Axis: "x", Pos: 5}},
			expected: 17,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.points, tc.folds)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %v) threw '%v', want %d", tc.points, tc.folds, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %v) = %d, want to throw", tc.points, tc.folds, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v) = %d, want %d", tc.points, tc.folds, actual, tc.expected)
			}
		})
	}
}
