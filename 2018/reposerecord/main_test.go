package reposerecord_test

import (
	"testing"

	. "github.com/pjaneiro/advent-of-code/2018/reposerecord"
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
			input:    []Entry{Entry{Year: 1518, Month: 11, Day: 1, Hour: 0, Minute: 0, Command: "Guard #10 begins shift"}, Entry{Year: 1518, Month: 11, Day: 1, Hour: 0, Minute: 5, Command: "falls asleep"}, Entry{Year: 1518, Month: 11, Day: 1, Hour: 0, Minute: 25, Command: "wakes up"}, Entry{Year: 1518, Month: 11, Day: 1, Hour: 0, Minute: 30, Command: "falls asleep"}, Entry{Year: 1518, Month: 11, Day: 1, Hour: 0, Minute: 55, Command: "wakes up"}, Entry{Year: 1518, Month: 11, Day: 1, Hour: 23, Minute: 58, Command: "Guard #99 begins shift"}, Entry{Year: 1518, Month: 11, Day: 2, Hour: 0, Minute: 40, Command: "falls asleep"}, Entry{Year: 1518, Month: 11, Day: 2, Hour: 0, Minute: 50, Command: "wakes up"}, Entry{Year: 1518, Month: 11, Day: 3, Hour: 0, Minute: 5, Command: "Guard #10 begins shift"}, Entry{Year: 1518, Month: 11, Day: 3, Hour: 0, Minute: 24, Command: "falls asleep"}, Entry{Year: 1518, Month: 11, Day: 3, Hour: 0, Minute: 29, Command: "wakes up"}, Entry{Year: 1518, Month: 11, Day: 4, Hour: 0, Minute: 2, Command: "Guard #99 begins shift"}, Entry{Year: 1518, Month: 11, Day: 4, Hour: 0, Minute: 36, Command: "falls asleep"}, Entry{Year: 1518, Month: 11, Day: 4, Hour: 0, Minute: 46, Command: "wakes up"}, Entry{Year: 1518, Month: 11, Day: 5, Hour: 0, Minute: 3, Command: "Guard #99 begins shift"}, Entry{Year: 1518, Month: 11, Day: 5, Hour: 0, Minute: 45, Command: "falls asleep"}, Entry{Year: 1518, Month: 11, Day: 5, Hour: 0, Minute: 55, Command: "wakes up"}},
			expected: 240,
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
			input:    []Entry{Entry{Year: 1518, Month: 11, Day: 1, Hour: 0, Minute: 0, Command: "Guard #10 begins shift"}, Entry{Year: 1518, Month: 11, Day: 1, Hour: 0, Minute: 5, Command: "falls asleep"}, Entry{Year: 1518, Month: 11, Day: 1, Hour: 0, Minute: 25, Command: "wakes up"}, Entry{Year: 1518, Month: 11, Day: 1, Hour: 0, Minute: 30, Command: "falls asleep"}, Entry{Year: 1518, Month: 11, Day: 1, Hour: 0, Minute: 55, Command: "wakes up"}, Entry{Year: 1518, Month: 11, Day: 1, Hour: 23, Minute: 58, Command: "Guard #99 begins shift"}, Entry{Year: 1518, Month: 11, Day: 2, Hour: 0, Minute: 40, Command: "falls asleep"}, Entry{Year: 1518, Month: 11, Day: 2, Hour: 0, Minute: 50, Command: "wakes up"}, Entry{Year: 1518, Month: 11, Day: 3, Hour: 0, Minute: 5, Command: "Guard #10 begins shift"}, Entry{Year: 1518, Month: 11, Day: 3, Hour: 0, Minute: 24, Command: "falls asleep"}, Entry{Year: 1518, Month: 11, Day: 3, Hour: 0, Minute: 29, Command: "wakes up"}, Entry{Year: 1518, Month: 11, Day: 4, Hour: 0, Minute: 2, Command: "Guard #99 begins shift"}, Entry{Year: 1518, Month: 11, Day: 4, Hour: 0, Minute: 36, Command: "falls asleep"}, Entry{Year: 1518, Month: 11, Day: 4, Hour: 0, Minute: 46, Command: "wakes up"}, Entry{Year: 1518, Month: 11, Day: 5, Hour: 0, Minute: 3, Command: "Guard #99 begins shift"}, Entry{Year: 1518, Month: 11, Day: 5, Hour: 0, Minute: 45, Command: "falls asleep"}, Entry{Year: 1518, Month: 11, Day: 5, Hour: 0, Minute: 55, Command: "wakes up"}},
			expected: 4455,
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
