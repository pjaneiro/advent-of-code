/*
Package reportrepair implements a solution to the first daily challenge of Advent of Code 2020.

More detailed information can be found in: https://adventofcode.com/2020/day/1.
*/
package reportrepair

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// readLines reads lines from input file to an array of ints.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, num)
	}
	return lines, scanner.Err()
}

/*
Challenge1 runs the first challenge of the day.

Having an array of integers, find two that add up to 2020, and return their product.
*/
func Challenge1(data []int) (int, error) {
	for _, a := range data {
		for _, b := range data {
			if a+b == 2020 {
				return a * b, nil
			}
		}
	}
	return 0, errors.New("couldn't find a satisfying solution")
}

/*
Challenge2 runs the second challenge of the day.

Having an array of integers, find three that add up to 2020, and return their product.
*/
func Challenge2(data []int) (int, error) {
	for _, a := range data {
		for _, b := range data {
			for _, c := range data {
				if a+b+c == 2020 {
					return a * b * c, nil
				}
			}
		}
	}
	return 0, errors.New("couldn't find a satisfying solution")
}

// Run runs the daily challenge using the other methods
func Run() {
	fmt.Println("Day 1")
	path := "src/reportrepair/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Println(err)
	}

	var result int
	result, err = Challenge1(data)
	if err != nil {
		fmt.Printf("Error running challenge: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(data)
	if err != nil {
		fmt.Printf("Error running challenge: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
