package rucksackreorganization

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

func Challenge1(data []string) (int, error) {
	result := 0
	for _, entry := range data {
		left, right := entry[:len(entry)/2], entry[len(entry)/2:]
	outerloop:
		for _, curLeft := range left {
			for _, curRight := range right {
				if curLeft == curRight {
					if curLeft >= 97 {
						result += int(curLeft) - 97 + 1
					} else {
						result += int(curLeft) - 65 + 27
					}
					break outerloop
				}
			}
		}
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	result := 0
	for i := 0; i < len(data); i += 3 {
		first, second, third := data[i], data[i+1], data[i+2]
	outerloop:
		for _, curFirst := range first {
			for _, curSecond := range second {
				for _, curThird := range third {
					if curFirst == curSecond && curFirst == curThird {
						if curFirst >= 97 {
							result += int(curFirst) - 97 + 1
						} else {
							result += int(curFirst) - 65 + 27
						}
						break outerloop
					}
				}
			}
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 3 - Rucksack Reorganization")
	path := "rucksackreorganization/input.txt"
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
