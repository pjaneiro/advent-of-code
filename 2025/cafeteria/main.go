package cafeteria

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	min, max int
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

func parseInput(data []string) ([]Range, []int, error) {
	var ranges []Range = make([]Range, 0)
	var ids []int = make([]int, 0)
	var i int
	for i = 0; i < len(data); i++ {
		if len(data[i]) == 0 {
			i++
			break
		}
		parts := strings.Split(data[i], "-")
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return ranges, ids, err
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return ranges, ids, err
		}
		ranges = append(ranges, Range{left, right})
	}
	for ; i < len(data); i++ {
		id, err := strconv.Atoi(data[i])
		if err != nil {
			return ranges, ids, err
		}
		ids = append(ids, id)
	}
	return ranges, ids, nil
}

func Challenge1(data []string) (int, error) {
	ranges, ids, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.min && id <= r.max {
				result++
				break
			}
		}
	}
	return result, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Challenge2(data []string) (int, error) {
	ranges, _, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].min == ranges[j].min {
			return ranges[i].max < ranges[j].max
		}
		return ranges[i].min < ranges[j].min
	})
	var newRanges []Range = make([]Range, 0)
	curMin, curMax := ranges[0].min, ranges[0].max
	for i := 1; i < len(ranges); i++ {
		if ranges[i].min > curMax {
			newRanges = append(newRanges, Range{curMin, curMax})
			curMin, curMax = ranges[i].min, ranges[i].max
		} else {
			curMax = max(curMax, ranges[i].max)
		}
	}
	newRanges = append(newRanges, Range{curMin, curMax})
	for _, r := range newRanges {
		result += r.max - r.min + 1
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 5 - Cafeteria")
	path := "cafeteria/input.txt"
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
