package passagepathing

import (
	"bufio"
	"fmt"
	"os"
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

func makeMapping(data []string) map[string][]string {
	var mapping map[string][]string = make(map[string][]string)

	for _, cur := range data {
		parts := strings.Split(cur, "-")
		if cur, present := mapping[parts[0]]; present {
			mapping[parts[0]] = append(cur, parts[1])
		} else {
			mapping[parts[0]] = []string{parts[1]}
		}
		if cur, present := mapping[parts[1]]; present {
			mapping[parts[1]] = append(cur, parts[0])
		} else {
			mapping[parts[1]] = []string{parts[0]}
		}
	}

	return mapping
}

func isLowercase(input string) bool {
	if input == "start" || input == "end" {
		return false
	}
	for _, cur := range input {
		if cur < 'a' || cur > 'z' {
			return false
		}
	}
	return true
}

func traversePart1(path []string, score int, mapping map[string][]string) int {
	lastLocation := path[len(path)-1]
	if lastLocation == "end" {
		return 1
	} else if lastLocation == "start" && len(path) != 1 {
		return 0
	} else if isLowercase(lastLocation) {
		for _, cur := range path[1 : len(path)-1] {
			if cur == lastLocation {
				return 0
			}
		}
	}
	result := 0
	for _, cur := range mapping[lastLocation] {
		result = result + traversePart1(append(path, cur), score, mapping)
	}
	return result
}

func traversePart2(path []string, score int, mapping map[string][]string, rule bool) int {
	lastLocation := path[len(path)-1]
	if lastLocation == "end" {
		return 1
	} else if lastLocation == "start" && len(path) != 1 {
		return 0
	} else if isLowercase(lastLocation) {
		for _, cur := range path[1 : len(path)-1] {
			if cur == lastLocation {
				if rule {
					return 0
				} else {
					rule = true
				}
			}
		}
	}
	result := 0
	for _, cur := range mapping[lastLocation] {
		result = result + traversePart2(append(path, cur), score, mapping, rule)
	}
	return result
}

func Challenge1(data []string) (int, error) {
	mapping := makeMapping(data)

	score := traversePart1([]string{"start"}, 0, mapping)

	return score, nil
}

func Challenge2(data []string) (int, error) {
	mapping := makeMapping(data)

	score := traversePart2([]string{"start"}, 0, mapping, false)

	return score, nil
}

func Run() {
	fmt.Println("Day 12 - Passage Pathing")
	path := "passagepathing/input.txt"
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
