package mullitover

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func parseInput(data []string) ([][]int, error) {
	var pairs [][]int = make([][]int, 0)
	r, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		return pairs, err
	}
	for _, row := range data {
		found := r.FindAllStringSubmatch(row, -1)
		for _, v := range found {
			v1, err := strconv.Atoi(v[1])
			if err != nil {
				return pairs, err
			}
			v2, err := strconv.Atoi(v[2])
			if err != nil {
				return pairs, err
			}
			pairs = append(pairs, []int{v1, v2})
		}
	}
	return pairs, nil
}

func Challenge1(data []string) (int, error) {
	pairs, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, v := range pairs {
		result += v[0] * v[1]
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	result := 0
	enabled := true
	for _, row := range data {
		indexes := []int{0}
		for i := 0; i < len(row); i++ {
			if i+4 <= len(row) && row[i:i+4] == "do()" || i+7 <= len(row) && row[i:i+7] == "don't()" {
				indexes = append(indexes, i)
			}
		}
		indexes = append(indexes, len(row))
		sections := make([]string, 0)
		for i := 1; i < len(indexes); i++ {
			if indexes[i-1]+7 <= len(row) && row[indexes[i-1]:indexes[i-1]+7] != "don't()" {
				if indexes[i-1] == 0 && !enabled {
					continue
				}
				sections = append(sections, row[indexes[i-1]:indexes[i]])
				enabled = true
				continue
			}
			enabled = false
		}
		pairs, err := parseInput(sections)
		if err != nil {
			return 0, err
		}
		for _, v := range pairs {
			result += v[0] * v[1]
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 3 - Mull It Over")
	path := "mullitover/input.txt"
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
