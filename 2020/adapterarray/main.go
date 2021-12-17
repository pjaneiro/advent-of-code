package adapterarray

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return lines, errors.New("something went wrong scanning int")
		}
		lines = append(lines, line)
	}

	return lines, nil
}

func Challenge1(data []int) (int, error) {
	sorted := make([]int, len(data))
	copy(sorted, data)
	sorted = append(sorted, 0)
	sort.Ints(sorted)
	var diffs map[int]int = map[int]int{1: 0, 2: 0, 3: 0}
	for i := 0; i < len(sorted)-1; i++ {
		diffs[sorted[i+1]-sorted[i]]++
	}
	diffs[3]++
	return diffs[1] * diffs[3], nil
}

func challenge2loop(data []int, prev map[int]int, i int) int {
	if v, ok := prev[i]; ok {
		return v
	}
	if i >= len(data)-1 {
		return 1
	}
	result := 0
	for j := i + 1; j < len(data); j++ {
		if data[j]-data[i] > 3 {
			break
		}
		result = result + challenge2loop(data, prev, j)
	}
	prev[i] = result
	return result
}

func Challenge2(data []int) (int, error) {
	sorted := make([]int, len(data))
	copy(sorted, data)
	sorted = append(sorted, 0)
	sort.Ints(sorted)
	prev := make(map[int]int)
	return challenge2loop(sorted, prev, 0), nil
}

func Run() {
	fmt.Println("Day 10 - Adapter Array")
	path := "adapterarray/input.txt"
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
