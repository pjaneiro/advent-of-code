package campcleanup

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Range struct {
	Min1, Min2, Max1, Max2 int
}

func readLines(path string) ([]Range, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r, err := regexp.Compile(`^(\d+)\-(\d+),(\d+)-(\d+)$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}

	var result []Range
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := r.FindStringSubmatch(line)
		var entry Range
		if entry.Min1, err = strconv.Atoi(data[1]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if entry.Max1, err = strconv.Atoi(data[2]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if entry.Min2, err = strconv.Atoi(data[3]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if entry.Max2, err = strconv.Atoi(data[4]); err != nil {
			return nil, errors.New("couldn't parse input")
		}

		result = append(result, entry)
	}

	return result, nil
}

func Challenge1(data []Range) (int, error) {
	result := 0
	for _, entry := range data {
		if entry.Min1 <= entry.Min2 && entry.Max1 >= entry.Max2 ||
			entry.Min1 >= entry.Min2 && entry.Max1 <= entry.Max2 {
			result++
		}
	}
	return result, nil
}

func Challenge2(data []Range) (int, error) {
	result := 0
	for _, entry := range data {
		if entry.Min1 <= entry.Min2 && entry.Max1 >= entry.Min2 ||
			entry.Min1 <= entry.Max2 && entry.Max1 >= entry.Max2 ||
			entry.Min2 <= entry.Min1 && entry.Max2 >= entry.Min1 ||
			entry.Min2 <= entry.Max1 && entry.Max2 >= entry.Max1 {
			result++
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 4 - Camp Cleanup")
	path := "campcleanup/input.txt"
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
