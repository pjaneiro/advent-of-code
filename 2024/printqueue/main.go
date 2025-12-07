package printqueue

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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

func parseInput(data []string) (map[int][]int, [][]int, error) {
	var rules map[int][]int = make(map[int][]int, 0)
	var updates [][]int = make([][]int, 0)
	parsingUpdates := false

	for _, row := range data {
		if len(row) == 0 {
			parsingUpdates = true
			continue
		} else if parsingUpdates {
			var rowUpdates []int = make([]int, 0)
			values := strings.Split(row, ",")
			for _, v := range values {
				u, err := strconv.Atoi(v)
				if err != nil {
					return nil, nil, err
				}
				rowUpdates = append(rowUpdates, u)
			}
			updates = append(updates, rowUpdates)
		} else {
			values := strings.Split(row, "|")
			a, err := strconv.Atoi(values[0])
			if err != nil {
				return nil, nil, err
			}
			b, err := strconv.Atoi(values[1])
			if err != nil {
				return nil, nil, err
			}
			if _, ok := rules[a]; !ok {
				rules[a] = make([]int, 0)
			}
			rules[a] = append(rules[a], b)
		}
	}

	return rules, updates, nil
}

func Challenge1(data []string) (int, error) {
	result := 0
	rules, updates, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	for _, row := range updates {
		var found map[int]struct{} = make(map[int]struct{}, 0)
		breaking := false
		for _, val := range row {
			for _, after := range rules[val] {
				if _, ok := found[after]; ok {
					breaking = true
					break
				}
			}
			found[val] = struct{}{}
		}
		if !breaking {
			result += row[len(row)/2]
		}
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	result := 0
	rules, updates, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	for _, row := range updates {
		var found map[int]struct{} = make(map[int]struct{}, 0)
		breaking := false
		for _, val := range row {
			for _, after := range rules[val] {
				if _, ok := found[after]; ok {
					breaking = true
					break
				}
			}
			found[val] = struct{}{}
		}
		if breaking {
			sort.Slice(row, func(i, j int) bool {
				if slices.Index(rules[row[i]], row[j]) > 0 {
					return true
				}
				return slices.Index(rules[row[j]], row[i]) < 0
			})
			result += row[len(row)/2]
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 5 - Print Queue")
	path := "printqueue/input.txt"
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
