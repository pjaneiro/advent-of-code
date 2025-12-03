package giftshop

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	First, Last int
}

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

func parseInput(data []string) ([]Range, error) {
	var result []Range = make([]Range, 0)
	for _, row := range data {
		pairs := strings.Split(row, ",")
		for _, pair := range pairs {
			split := strings.Split(pair, "-")
			first, err := strconv.Atoi(split[0])
			if err != nil {
				return nil, err
			}
			last, err := strconv.Atoi(split[1])
			if err != nil {
				return nil, err
			}
			result = append(result, Range{first, last})
		}
	}
	return result, nil
}

func Challenge1(data []string) (int, error) {
	ranges, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, entry := range ranges {
		for i := entry.First; i <= entry.Last; i++ {
			stringed := strconv.Itoa(i)
			length := len(stringed)
			if length%2 != 0 {
				continue
			}
			if stringed[0:length/2] == stringed[length/2:length] {
				result += i
			}
		}
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	ranges, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, entry := range ranges {
		for val := entry.First; val <= entry.Last; val++ {
			stringed := strconv.Itoa(val)
			for length := 1; length <= len(stringed)/2; length++ {
				if len(stringed)%length != 0 {
					continue
				}
				compare := stringed[0:length]
				found := true
				for j := length; j <= len(stringed)-length; j += length {
					if stringed[j:j+length] != compare {
						found = false
						break
					}
				}
				if found {
					result += val
					break
				}
			}
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 2 - Gift Shop")
	path := "giftshop/input.txt"
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
