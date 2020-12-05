package binaryboarding

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

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

func Challenge1(data []string) (int, error) {
	var result int = 0
	for _, seat := range data {
		var rowdata string = seat[0:7]
		var coldata string = seat[7:10]
		var i, row, col int = 0, 0, 0
		for i = 0; i < len(rowdata); i++ {
			if rowdata[i:i+1] == "B" {
				row = row | 1<<(6-i)
			}
		}
		for i = 0; i < len(coldata); i++ {
			if coldata[i:i+1] == "R" {
				col = col | 1<<(2-i)
			}
		}
		if (row*8)+col > result {
			result = (row * 8) + col
		}
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	var plane [128][8]bool
	var ids map[int]bool = make(map[int]bool)
	for _, seat := range data {
		var rowdata string = seat[0:7]
		var coldata string = seat[7:10]
		var i, row, col int = 0, 0, 0
		for i = 0; i < len(rowdata); i++ {
			if rowdata[i:i+1] == "B" {
				row = row | 1<<(6-i)
			}
		}
		for i = 0; i < len(coldata); i++ {
			if coldata[i:i+1] == "R" {
				col = col | 1<<(2-i)
			}
		}
		plane[row][col] = true
		ids[(row*8)+col] = true
	}
	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			if plane[i][j] == false {
				id := (i * 8) + j
				if ids[id+1] == true && ids[id-1] == true {
					return id, nil
				}
			}
		}
	}
	return 0, errors.New("couldn't find a satisfying solution")
}

func Run() {
	fmt.Println("Day 5")
	path := "src/binaryboarding/input.txt"
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
