package extendedpolymerization

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
)

func readLines(path string) (string, map[string]map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	var starter string
	var rules map[string]map[string]string = make(map[string]map[string]string)

	r, err := regexp.Compile(`^(\w)(\w) -> (\w)$`)
	if err != nil {
		return "", nil, errors.New("couldn't parse regex expression")
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	starter = scanner.Text()
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		data := r.FindStringSubmatch(line)
		if _, ok := rules[data[1]]; !ok {
			rules[data[1]] = make(map[string]string)
		}
		rules[data[1]][data[2]] = data[3]
	}

	return starter, rules, nil
}

func Challenge1(starter string, rules map[string]map[string]string) (int64, error) {
	for i := 0; i < 10; i++ {
		var newString string = starter[0:1]
		for j := 1; j < len(starter); j++ {
			if val, ok := rules[starter[j-1:j]][starter[j:j+1]]; ok {
				newString = newString + val + starter[j:j+1]
			}
		}
		starter = newString
	}
	var counts map[string]int64 = make(map[string]int64)
	for _, cur := range starter {
		counts[string(cur)]++
	}
	maxCount, minCount := int64(0), int64(math.MaxInt64)
	for _, v := range counts {
		if v > maxCount {
			maxCount = v
		}
		if v < minCount {
			minCount = v
		}
	}
	return maxCount - minCount, nil
}

func Challenge2(starter string, rules map[string]map[string]string) (int64, error) {
	var counts map[string]int64 = make(map[string]int64)
	for i := 0; i < len(starter)-1; i++ {
		counts[starter[i:i+2]]++
	}
	for i := 0; i < 40; i++ {
		var newCounts map[string]int64 = make(map[string]int64)
		for k, v := range counts {
			left, right := k[0:1], k[1:2]
			mid := rules[left][right]
			newCounts[left+mid] = newCounts[left+mid] + v
			newCounts[mid+right] = newCounts[mid+right] + v
		}
		counts = newCounts
	}
	var totals map[string]int64 = make(map[string]int64)
	for k, v := range counts {
		totals[k[0:1]] = totals[k[0:1]] + v
	}
	totals[starter[len(starter)-1:len(starter)]]++
	maxCount, minCount := int64(0), int64(math.MaxInt64)
	for _, v := range totals {
		if v > maxCount {
			maxCount = v
		}
		if v < minCount {
			minCount = v
		}
	}
	return maxCount - minCount, nil
}

func Run() {
	fmt.Println("Day 14 - Extended Polymerization")
	path := "extendedpolymerization/input.txt"
	points, folds, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int64
	result, err = Challenge1(points, folds)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(points, folds)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
