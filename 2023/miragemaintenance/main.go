package miragemaintenance

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	return lines, nil
}

func parseInput(data []string) ([][]int, error) {
	var result [][]int = make([][]int, 0)
	for _, line := range data {
		var values []int = make([]int, 0)
		fields := strings.Fields(line)
		for _, field := range fields {
			t, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			values = append(values, t)
		}
		result = append(result, values)
	}
	return result, nil
}

func isZeroed(input []int) bool {
	for _, c := range input {
		if c != 0 {
			return false
		}
	}
	return true
}

func value1(input []int) int {
	if isZeroed(input) {
		return 0
	}
	var diffs []int = []int{input[1] - input[0]}
	for i := 2; i < len(input); i++ {
		diffs = append(diffs, input[i]-input[i-1])
	}
	return input[len(input)-1] + value1(diffs)
}

func value2(input []int) int {
	if isZeroed(input) {
		return 0
	}
	var diffs []int = []int{input[1] - input[0]}
	for i := 2; i < len(input); i++ {
		diffs = append(diffs, input[i]-input[i-1])
	}
	return input[0] - value2(diffs)
}

func Challenge1(data []string) (int, error) {
	parsedData, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, c := range parsedData {
		result = result + value1(c)
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	parsedData, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, c := range parsedData {
		result = result + value2(c)
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 9 - Mirage Maintenance")
	path := "miragemaintenance/input.txt"
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
