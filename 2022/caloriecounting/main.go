package caloriecounting

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func Challenge1(data []string) (int, error) {
	max, count := 0, 0
	for _, line := range data {
		if line == "" {
			if count > max {
				max = count
			}
			count = 0
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				return 0, errors.New("something went wrong")
			}
			count += val
		}
	}
	if count > max {
		max = count
	}
	return max, nil
}

func Challenge2(data []string) (int, error) {
	count := 0
	var counts []int
	for _, line := range data {
		if line == "" {
			counts = append(counts, count)
			count = 0
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				return 0, errors.New("something went wrong")
			}
			count += val
		}
	}
	counts = append(counts, count)
	sort.Ints(counts)
	return counts[len(counts)-1] + counts[len(counts)-2] + counts[len(counts)-3], nil
}

func Run() {
	fmt.Println("Day 1 - Calorie Counting")
	path := "caloriecounting/input.txt"
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
