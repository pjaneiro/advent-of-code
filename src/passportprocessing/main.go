package passportprocessing

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Passport struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

func readLines(path string) ([]Passport, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Passport = []Passport{{}}
	var index int = 0
	var current *Passport
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current = &lines[index]
		line := scanner.Text()
		words := strings.Fields(line)
		if len(words) == 0 {
			index++
			lines = append(lines, Passport{})
			current = &lines[index]
		}

		for _, word := range words {
			pair := strings.Split(word, ":")

			switch pair[0] {
			case "byr":
				current.Byr = pair[1]
			case "iyr":
				current.Iyr = pair[1]
			case "eyr":
				current.Eyr = pair[1]
			case "hgt":
				current.Hgt = pair[1]
			case "hcl":
				current.Hcl = pair[1]
			case "ecl":
				current.Ecl = pair[1]
			case "pid":
				current.Pid = pair[1]
			case "cid":
				current.Cid = pair[1]
			}
		}
	}
	return lines, scanner.Err()
}

func Challenge1(data []Passport) (int, error) {
	var result int = 0
	for _, pass := range data {
		if pass.Byr != "" && pass.Iyr != "" && pass.Eyr != "" && pass.Hgt != "" && pass.Hcl != "" && pass.Ecl != "" && pass.Pid != "" {
			result++
		}
	}
	return result, nil
}

func Challenge2(data []Passport) (int, error) {
	var result int = 0
	for _, pass := range data {
		if pass.Byr == "" || pass.Iyr == "" || pass.Eyr == "" || pass.Hgt == "" || pass.Hcl == "" || pass.Ecl == "" || pass.Pid == "" {
			continue
		}
		var r *regexp.Regexp
		var err error

		// byr
		r, err = regexp.Compile(`^(19[2-9][0-9]|200[0-2])$`)
		if err != nil {
			return 0, errors.New("couldn't parse regex expression")
		}
		if !r.MatchString(pass.Byr) {
			continue
		}

		// iyr
		r, err = regexp.Compile(`^20(1\d|20)$`)
		if err != nil {
			return 0, errors.New("couldn't parse regex expression")
		}
		if !r.MatchString(pass.Iyr) {
			continue
		}

		// eyr
		r, err = regexp.Compile(`^20(2\d|30)$`)
		if err != nil {
			return 0, errors.New("couldn't parse regex expression")
		}
		if !r.MatchString(pass.Eyr) {
			continue
		}

		// hgt
		r, err = regexp.Compile(`^((1[5-8][0-9]|19[0-3])cm)|((59|6[0-9]|7[0-6])in)$`)
		if err != nil {
			return 0, errors.New("couldn't parse regex expression")
		}
		if !r.MatchString(pass.Hgt) {
			continue
		}

		// hcl
		r, err = regexp.Compile(`^#[0-9a-f]{6}$`)
		if err != nil {
			return 0, errors.New("couldn't parse regex expression")
		}
		if !r.MatchString(pass.Hcl) {
			continue
		}

		// ecl
		r, err = regexp.Compile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
		if err != nil {
			return 0, errors.New("couldn't parse regex expression")
		}
		if !r.MatchString(pass.Ecl) {
			continue
		}

		// pid
		r, err = regexp.Compile(`^[0-9]{9}$`)
		if err != nil {
			return 0, errors.New("couldn't parse regex expression")
		}
		if !r.MatchString(pass.Pid) {
			continue
		}

		result++
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 3")
	path := "src/passportprocessing/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(data)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(data)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
