/*
Package binaryboarding implements a solution to the fifth daily challenge of Advent of Code 2020.

More detailed information can be found in: https://adventofcode.com/2020/day/5.
*/
package binaryboarding

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// readLines reads lines from input file to an array of strings.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

/*
Challenge1 runs the first challenge of the day.

From a list of boarding passes, find the highest boarding pass ID.
*/
func Challenge1(data []string) (int, error) {
	result := 0
	for _, seat := range data {
		rowdata, coldata := seat[0:7], seat[7:10]
		row, col := 0, 0
		for i, char := range rowdata {
			if char == 'B' {
				row = row | 1<<(6-i)
			}
		}
		for i, char := range coldata {
			if char == 'R' {
				col = col | 1<<(2-i)
			}
		}
		id := (row * 8) + col
		if id < result {
			continue
		}
		result = id
	}
	return result, nil
}

/*
Challenge2 runs the first challenge of the day.

From a list of boarding passes, guess your boarding pass ID.
*/
func Challenge2(data []string) (int, error) {
	ids := make(map[int]struct{})
	for _, seat := range data {
		rowdata, coldata := seat[0:7], seat[7:10]
		row, col := 0, 0
		for i, char := range rowdata {
			if char == 'B' {
				row = row | 1<<(6-i)
			}
		}
		for i, char := range coldata {
			if char == 'R' {
				col = col | 1<<(2-i)
			}
		}
		ids[(row*8)+col] = struct{}{}
	}
	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			id := (i * 8) + j
			if _, ok := ids[id]; ok {
				continue
			}
			if _, ok := ids[id+1]; !ok {
				continue
			}

			if _, ok := ids[id-1]; !ok {
				continue
			}

			return id, nil
		}
	}
	return 0, errors.New("couldn't find a satisfying solution")
}

// Run runs the daily challenge using the other methods.
func Run() {
	fmt.Println("Day 5 - Binary Boarding")
	path := "binaryboarding/input.txt"
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
