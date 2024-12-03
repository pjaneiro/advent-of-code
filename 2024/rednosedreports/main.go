package rednosedreports

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	ASC = iota
	DESC
	UND
)

func readLines(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var list [][]int = make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var row []int = make([]int, 0)
		for _, v := range line {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			row = append(row, n)
		}
		list = append(list, row)
	}
	return list, nil
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func evaluateRow(row []int) bool {
	if row[1] == row[0] || abs(row[1]-row[0]) < 1 || abs(row[1]-row[0]) > 3 {
		return false
	}
	direction := ASC
	if row[1] < row[0] {
		direction = DESC
	}
	valid := true
	for i := 2; i < len(row); i++ {
		if direction == ASC && row[i] <= row[i-1] || direction == DESC && row[i] >= row[i-1] || abs(row[i]-row[i-1]) < 1 || abs(row[i]-row[i-1]) > 3 {
			valid = false
			break
		}
	}
	return valid
}

func Challenge1(input [][]int) (int, error) {
	result := 0
	for _, row := range input {
		if evaluateRow(row) {
			result++
		}
	}
	return result, nil
}

func Challenge2(input [][]int) (int, error) {
	result := 0
	for _, row := range input {
		if evaluateRow(row) {
			result++
			continue
		}
		for i := 0; i < len(row); i++ {
			if evaluateRow(append(slices.Clone(row[:i]), row[i+1:]...)) {
				result++
				break
			}
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 2 - Red-Nosed Reports")
	path := "rednosedreports/input.txt"
	input, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(input)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(input)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
