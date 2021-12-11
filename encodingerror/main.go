package encodingerror

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
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

func Challenge1(data []int, preamble int) (int, error) {
OUTER:
	for i, n := range data {
		if i < preamble {
			continue
		}
		for j := i - preamble; j < i; j++ {
			for k := j; k < i; k++ {
				if j == k {
					continue
				}
				if data[j]+data[k] == n {
					continue OUTER
				}
			}
		}
		return n, nil
	}
	return 0, errors.New("not implemented yet")
}

func Challenge2(data []int, preamble int) (int, error) {
	target, _ := Challenge1(data, preamble)
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			try := data[i:j]
			result := 0
			min := math.MaxInt32
			max := 0
			for _, k := range try {
				result = result + k
				if k < min {
					min = k
				}
				if k > max {
					max = k
				}
			}
			if result == target {
				return min + max, nil
			}
		}
	}
	return 0, errors.New("not implemented yet")
}

func Run() {
	fmt.Println("Day 9")
	path := "encodingerror/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(data, 25)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(data, 25)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
