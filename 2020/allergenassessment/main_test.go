package allergenassessment_test

import (
	. "github.com/pjaneiro/advent-of-code/2020/allergenassessment"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Recipe
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Recipe{Recipe{Ingredients: []string{"mxmxvkd", "kfcds", "sqjhc", "nhms"}, Allergens: []string{"dairy", "fish"}}, Recipe{Ingredients: []string{"trh", "fvjkl", "sbzzf", "mxmxvkd"}, Allergens: []string{"dairy"}}, Recipe{Ingredients: []string{"sqjhc", "fvjkl"}, Allergens: []string{"soy"}}, Recipe{Ingredients: []string{"sqjhc", "mxmxvkd", "sbzzf"}, Allergens: []string{"fish"}}},
			expected: 5,
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
		input    []Recipe
		expected string
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Recipe{Recipe{Ingredients: []string{"mxmxvkd", "kfcds", "sqjhc", "nhms"}, Allergens: []string{"dairy", "fish"}}, Recipe{Ingredients: []string{"trh", "fvjkl", "sbzzf", "mxmxvkd"}, Allergens: []string{"dairy"}}, Recipe{Ingredients: []string{"sqjhc", "fvjkl"}, Allergens: []string{"soy"}}, Recipe{Ingredients: []string{"sqjhc", "mxmxvkd", "sbzzf"}, Allergens: []string{"fish"}}},
			expected: "mxmxvkd,sqjhc,fvjkl",
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v) threw '%v', want %s", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v) = %s, want to throw", tc.input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v) = %s, want %s", tc.input, actual, tc.expected)
			}
		})
	}
}
