package passportprocessing_test

import (
	"fmt"
	. "passportprocessing"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Passport
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Passport{},
			expected: 0,
			error:    false,
		}, {
			name:     "Example 2",
			input:    []Passport{{Ecl: "gry", Pid: "860033327", Eyr: "2020", Hcl: "#fffffd", Byr: "1937", Iyr: "2017", Cid: "147", Hgt: "183cm"}},
			expected: 1,
			error:    false,
		}, {
			name:     "Example 3",
			input:    []Passport{{Ecl: "gry", Pid: "860033327", Eyr: "2020", Hcl: "#fffffd", Byr: "1937", Iyr: "2017", Cid: "147", Hgt: "183cm"}, {Iyr: "2013", Ecl: "amb", Cid: "350", Eyr: "2023", Pid: "028048884", Hcl: "#cfa07d", Byr: "1929"}},
			expected: 1,
			error:    false,
		}, {
			name:     "Example 4",
			input:    []Passport{{Ecl: "gry", Pid: "860033327", Eyr: "2020", Hcl: "#fffffd", Byr: "1937", Iyr: "2017", Cid: "147", Hgt: "183cm"}, {Iyr: "2013", Ecl: "amb", Cid: "350", Eyr: "2023", Pid: "028048884", Hcl: "#cfa07d", Byr: "1929"}, {Hcl: "#ae17e1", Iyr: "2013", Eyr: "2024", Ecl: "brn", Pid: "760753108", Byr: "1931", Hgt: "179cm"}},
			expected: 2,
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

func ExampleChallenge1() {
	var input []Passport = []Passport{
		{Ecl: "gry", Pid: "860033327", Eyr: "2020", Hcl: "#fffffd", Byr: "1937", Iyr: "2017", Cid: "147", Hgt: "183cm"},
		{Iyr: "2013", Ecl: "amb", Cid: "350", Eyr: "2023", Pid: "028048884", Hcl: "#cfa07d", Byr: "1929"},
		{Hcl: "#ae17e1", Iyr: "2013", Eyr: "2024", Ecl: "brn", Pid: "760753108", Byr: "1931", Hgt: "179cm"},
	}
	result, err := Challenge1(input)
	fmt.Printf("Result is %d, error is %v\n", result, err)
	// Output: Result is 2, error is <nil>
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Passport
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Passport{},
			expected: 0,
			error:    false,
		}, {
			name:     "Example 2",
			input:    []Passport{{Eyr: "1972", Cid: "100", Hcl: "#18171d", Ecl: "amb", Hgt: "170", Pid: "186cm", Iyr: "2018", Byr: "1926"}, {Iyr: "2019", Hcl: "#602927", Eyr: "1967", Hgt: "170cm", Ecl: "grn", Pid: "012533040", Byr: "1946"}, {Hcl: "dab227", Iyr: "2012", Ecl: "brn", Hgt: "182cm", Pid: "021572410", Eyr: "2020", Byr: "1992", Cid: "277"}, {Hgt: "59cm", Ecl: "zzz", Eyr: "2038", Hcl: "74454a", Iyr: "2023", Pid: "3556412378", Byr: "2007"}},
			expected: 0,
			error:    false,
		}, {
			name:     "Example 3",
			input:    []Passport{{Pid: "087499704", Hgt: "74in", Ecl: "grn", Iyr: "2012", Eyr: "2030", Byr: "1980", Hcl: "#623a2f"}, {Eyr: "2029", Ecl: "blu", Cid: "129", Byr: "1989", Iyr: "2014", Pid: "896056539", Hcl: "#a97842", Hgt: "165cm"}, {Hcl: "#888785", Hgt: "164cm", Byr: "2001", Iyr: "2015", Cid: "88", Pid: "545766238", Ecl: "hzl", Eyr: "2022"}, {Iyr: "2010", Hgt: "158cm", Hcl: "#b6652a", Ecl: "blu", Byr: "1944", Eyr: "2021", Pid: "093154719"}},
			expected: 4,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.input)
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

func ExampleChallenge2() {
	var input []Passport = []Passport{
		{Pid: "087499704", Hgt: "74in", Ecl: "grn", Iyr: "2012", Eyr: "2030", Byr: "1980", Hcl: "#623a2f"},
		{Eyr: "2029", Ecl: "blu", Cid: "129", Byr: "1989", Iyr: "2014", Pid: "896056539", Hcl: "#a97842", Hgt: "165cm"},
		{Hcl: "#888785", Hgt: "164cm", Byr: "2001", Iyr: "2015", Cid: "88", Pid: "545766238", Ecl: "hzl", Eyr: "2022"},
		{Iyr: "2010", Hgt: "158cm", Hcl: "#b6652a", Ecl: "blu", Byr: "1944", Eyr: "2021", Pid: "093154719"},
	}
	result, err := Challenge2(input)
	fmt.Printf("Result is %d, error is %v\n", result, err)
	// Output: Result is 4, error is <nil>
}
