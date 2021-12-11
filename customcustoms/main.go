/*
Package customcustoms implements a solution to the sixth daily challenge of Advent of Code 2020.

More detailed information can be found in: https://adventofcode.com/2020/day/6.
*/
package customcustoms

import (
	"bufio"
	"fmt"
	"os"
)

// Form holds customs information for a specific group of people.
type Form struct {
	Npeople int
	Answers map[rune]int
}

func readLines(path string) ([]Form, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Form = []Form{Form{}}
	var index int = 0
	var current *Form
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current = &lines[index]
		if current.Answers == nil {
			current.Answers = make(map[rune]int)
		}
		line := scanner.Text()
		if len(line) == 0 {
			index++
			lines = append(lines, Form{})
			current = &lines[index]
		} else {
			current.Npeople++
		}

		for _, char := range line {
			current.Answers[char]++
		}
	}
	return lines, scanner.Err()
}

/*
Challenge1 runs the first challenge of the day.

From a list of filled forms, check how many items are checked.
*/
func Challenge1(data []Form) (int, error) {
	result := 0
	for _, form := range data {
		result = result + len(form.Answers)
	}
	return result, nil
}

/*
Challenge2 runs the second challenge of the day.

From a list of filled forms, check how many items are checked by all people in the group.
*/
func Challenge2(data []Form) (int, error) {
	result := 0
	for _, form := range data {
		for _, c := range form.Answers {
			if c != form.Npeople {
				continue
			}
			result++
		}
	}
	return result, nil
}

// Run runs the daily challenge using the other methods.
func Run() {
	fmt.Println("Day 6")
	path := "customcustoms/input.txt"
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
