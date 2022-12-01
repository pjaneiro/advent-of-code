package sevensegmentsearch_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2021/golang/sevensegmentsearch"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Entry
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Entry{{Input: []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"}, Output: []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}}, {Input: []string{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"}, Output: []string{"fcgedb", "cgb", "dgebacf", "gc"}}, {Input: []string{"fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef"}, Output: []string{"cg", "cg", "fdcagb", "cbg"}}, {Input: []string{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega"}, Output: []string{"efabcd", "cedba", "gadfec", "cb"}}, {Input: []string{"aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga"}, Output: []string{"gecf", "egdcabf", "bgf", "bfgea"}}, {Input: []string{"fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf"}, Output: []string{"gebdcfa", "ecba", "ca", "fadegcb"}}, {Input: []string{"dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf"}, Output: []string{"cefg", "dcbef", "fcge", "gbcadfe"}}, {Input: []string{"bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd"}, Output: []string{"ed", "bcgafe", "cdgba", "cbgef"}}, {Input: []string{"egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg"}, Output: []string{"gbdfcae", "bgc", "cg", "cgb"}}, {Input: []string{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc"}, Output: []string{"fgae", "cfgab", "fg", "bagce"}}},
			expected: 26,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []Entry{},
			expected: 0,
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
		input    []Entry
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Entry{{Input: []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"}, Output: []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}}, {Input: []string{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"}, Output: []string{"fcgedb", "cgb", "dgebacf", "gc"}}, {Input: []string{"fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef"}, Output: []string{"cg", "cg", "fdcagb", "cbg"}}, {Input: []string{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega"}, Output: []string{"efabcd", "cedba", "gadfec", "cb"}}, {Input: []string{"aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga"}, Output: []string{"gecf", "egdcabf", "bgf", "bfgea"}}, {Input: []string{"fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf"}, Output: []string{"gebdcfa", "ecba", "ca", "fadegcb"}}, {Input: []string{"dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf"}, Output: []string{"cefg", "dcbef", "fcge", "gbcadfe"}}, {Input: []string{"bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd"}, Output: []string{"ed", "bcgafe", "cdgba", "cbgef"}}, {Input: []string{"egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg"}, Output: []string{"gbdfcae", "bgc", "cg", "cgb"}}, {Input: []string{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc"}, Output: []string{"fgae", "cfgab", "fg", "bagce"}}},
			expected: 61229,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []Entry{},
			expected: 0,
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
