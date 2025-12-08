package plutonianpebbles

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func parseInput(data []string) ([]int, error) {
	var result []int = make([]int, 0)
	values := strings.Split(data[0], " ")
	for _, v := range values {
		x, err := strconv.Atoi(v)
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, nil
}

func iterate(stones []int, generations int) int {
	var combos map[int]int = make(map[int]int)
	for _, s := range stones {
		if _, ok := combos[s]; !ok {
			combos[s] = 0
		}
		combos[s]++
	}

	for range generations {
		old := combos
		combos = make(map[int]int)
		for k, v := range old {
			if k == 0 {
				if _, ok := combos[1]; !ok {
					combos[1] = 0
				}
				combos[1] += v
			} else if ndigs := len(strconv.Itoa(k)); ndigs%2 == 0 {
				threshold := int(math.Pow(10, float64(ndigs/2)))
				left, right := k/threshold, k%threshold
				if _, ok := combos[left]; !ok {
					combos[left] = 0
				}
				if _, ok := combos[right]; !ok {
					combos[right] = 0
				}
				combos[left] += v
				combos[right] += v
			} else {
				if _, ok := combos[k*2024]; !ok {
					combos[k*2024] = 0
				}
				combos[k*2024] += v
			}
		}
	}

	result := 0
	for _, v := range combos {
		result += v
	}
	return result
}

func Challenge1(data []string) (int, error) {
	stones, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	return iterate(stones, 25), nil
}

func Challenge2(data []string) (int, error) {
	stones, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	return iterate(stones, 75), nil
}

func Run() {
	fmt.Println("Day 11 - Plutonian Pebbles")
	path := "plutonianpebbles/input.txt"
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
