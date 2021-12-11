/*
Package passwordphilosophy implements a solution to the second daily challenge of Advent of Code 2020.

More detailed information can be found in: https://adventofcode.com/2020/day/2.
*/
package passwordphilosophy

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Password holds the info known about a password, including constraints.
type Password struct {
	Content string
	Char    string
	Min     int
	Max     int
}

// readLines reads lines from input file to an array of Passwords.
func readLines(path string) ([]Password, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Password
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r, err := regexp.Compile(`(\d+)-(\d+) (\w): (\w+)`)
		if err != nil {
			return lines, errors.New("couldn't parse regex expression")
		}

		var elem Password
		var data []string = r.FindStringSubmatch(scanner.Text())
		if data == nil {
			return lines, nil
		}

		elem.Min, err = strconv.Atoi(data[1])
		if err != nil {
			return lines, errors.New("couldn't extract all data")
		}

		elem.Max, err = strconv.Atoi(data[2])
		if err != nil {
			return lines, errors.New("couldn't extract all data")
		}

		elem.Char = data[3]
		elem.Content = data[4]

		lines = append(lines, elem)
	}
	return lines, nil
}

/*
Challenge1 runs the first challenge of the day.

From a list of passwords, check how many comply to their defined constraints.
*/
func Challenge1(data []Password) (int, error) {
	var result int = 0
	for _, c := range data {
		count := strings.Count(c.Content, c.Char)
		if count >= c.Min && count <= c.Max {
			result++
		}
	}
	return result, nil
}

/*
Challenge2 runs the second challenge of the day.

From a list of passwords, check how many comply to their defined constraints.
*/
func Challenge2(data []Password) (int, error) {
	var result int = 0
	for _, c := range data {
		var min string = c.Content[c.Min-1 : c.Min]
		var max string = c.Content[c.Max-1 : c.Max]
		if (min == c.Char) != (max == c.Char) {
			result++
		}
	}
	return result, nil
}

// Run runs the daily challenge using the other methods.
func Run() {
	fmt.Println("Day 2 - Password Philosophy")
	path := "passwordphilosophy/input.txt"
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
