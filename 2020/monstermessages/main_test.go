package monstermessages_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2020/monstermessages"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		rules    map[string]string
		messages []string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			rules:    map[string]string{"0": "4 1 5", "1": "2 3 | 3 2", "2": "4 4 | 5 5", "3": "4 5 | 5 4", "4": "\"a\"", "5": "\"b\""},
			messages: []string{"ababbb", "bababa", "abbbab", "aaabbb", "aaaabbb"},
			expected: 2,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.rules, tc.messages)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v, %v) threw '%v', want %d", tc.rules, tc.messages, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v, %v) = %d, want to throw", tc.rules, tc.messages, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v, %v) = %d, want %d", tc.rules, tc.messages, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		rules    map[string]string
		messages []string
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			rules:    map[string]string{"0": "4 1 5", "1": "2 3 | 3 2", "2": "4 4 | 5 5", "3": "4 5 | 5 4", "4": "\"a\"", "5": "\"b\""},
			messages: []string{"ababbb", "bababa", "abbbab", "aaabbb", "aaaabbb"},
			expected: 2,
			error:    false,
		}, {
			name:     "Example 2",
			rules:    map[string]string{"42": "9 14 | 10 1", "9": "14 27 | 1 26", "10": "23 14 | 28 1", "1": "\"a\"", "11": "42 31", "5": "1 14 | 15 1", "19": "14 1 | 14 14", "12": "24 14 | 19 1", "16": "15 1 | 14 14", "31": "14 17 | 1 13", "6": "14 14 | 1 14", "2": "1 24 | 14 4", "0": "8 11", "13": "14 3 | 1 12", "15": "1 | 14", "17": "14 2 | 1 7", "23": "25 1 | 22 14", "28": "16 1", "4": "1 1", "20": "14 14 | 1 15", "3": "5 14 | 16 1", "27": "1 6 | 14 18", "14": "\"b\"", "21": "14 1 | 1 14", "25": "1 1 | 1 14", "22": "14 14", "8": "42", "26": "14 22 | 1 20", "18": "15 15", "7": "14 5 | 1 21", "24": "14 1"},
			messages: []string{"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa", "bbabbbbaabaabba", "babbbbaabbbbbabbbbbbaabaaabaaa", "aaabbbbbbaaaabaababaabababbabaaabbababababaaa", "bbbbbbbaaaabbbbaaabbabaaa", "bbbababbbbaaaaaaaabbababaaababaabab", "ababaaaaaabaaab", "ababaaaaabbbaba", "baabbaaaabbaaaababbaababb", "abbbbabbbbaaaababbbbbbaaaababb", "aaaaabbaabaaaaababaa", "aaaabbaaaabbaaa", "aaaabbaabbaaaaaaabbbabbbaaabbaabaaa", "babaaabbbaaabaababbaabababaaab", "aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba"},
			expected: 12,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.rules, tc.messages)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v, %v) threw '%v', want %d", tc.rules, tc.messages, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v, %v) = %d, want to throw", tc.rules, tc.messages, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v, %v) = %d, want %d", tc.rules, tc.messages, actual, tc.expected)
			}
		})
	}
}
