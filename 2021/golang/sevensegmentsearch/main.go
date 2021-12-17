package sevensegmentsearch

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Entry struct {
	Input  []string
	Output []string
}

func readLines(path string) ([]Entry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Entry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")

		var elem Entry
		elem.Input = strings.Split(strings.TrimSpace(line[0]), " ")
		elem.Output = strings.Split(strings.TrimSpace(line[1]), " ")

		lines = append(lines, elem)
	}
	return lines, nil
}

func Challenge1(data []Entry) (int, error) {
	result := 0
	for _, entry := range data {
		for _, elem := range entry.Output {
			length := len(elem)
			if length == 7 || length < 5 {
				result++
			}
		}
	}
	return result, nil
}

func sortString(input string) string {
	result := strings.Split(input, "")
	sort.Strings(result)
	return strings.Join(result, "")
}

func contains(list string, elem byte) bool {
	for _, chr := range list {
		if byte(chr) == elem {
			return true
		}
	}
	return false
}

func difference(bigger string, smaller string) string {
	var result string
	for _, chr := range bigger {
		if !contains(smaller, byte(chr)) {
			result = result + string(chr)
		}
	}
	return result
}

func Challenge2(data []Entry) (int, error) {
	total := 0
	powers := []int{1000, 100, 10, 1}
	for _, entry := range data {
		sum := 0
		var one, four, fourminusone string
		for _, cur := range entry.Input {
			if len(cur) == 2 {
				one = cur
			} else if len(cur) == 4 {
				four = cur
			}
		}
		fourminusone = string(difference(four, one))
		for k, v := range entry.Output {
			if len(v) == 2 {
				sum = sum + powers[k]
			} else if len(v) == 3 {
				sum = sum + (7 * powers[k])
			} else if len(v) == 4 {
				sum = sum + (4 * powers[k])
			} else if len(v) == 7 {
				sum = sum + (8 * powers[k])
			} else if len(v) == 5 {
				if contains(v, one[0]) && contains(v, one[1]) {
					sum = sum + (3 * powers[k])
				} else if contains(v, fourminusone[0]) && contains(v, fourminusone[1]) {
					sum = sum + (5 * powers[k])
				} else {
					sum = sum + (2 * powers[k])
				}
			} else {
				if contains(v, one[0]) && contains(v, one[1]) && contains(v, fourminusone[0]) && contains(v, fourminusone[1]) {
					sum = sum + (9 * powers[k])
				} else if contains(v, fourminusone[0]) && contains(v, fourminusone[1]) {
					sum = sum + (6 * powers[k])
				}
			}
		}
		total = total + sum
	}
	return total, nil
}

func Run() {
	fmt.Println("Day 8 - Seven Segment Search")
	path := "sevensegmentsearch/input.txt"
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
