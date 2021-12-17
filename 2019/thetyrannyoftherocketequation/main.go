package thetyrannyoftherocketequation

import (
	"bufio"
	"errors"
	"fmt"
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

func Challenge1(data []int) (int, error) {
	result := 0
	for _, v := range data {
		result = result + v/3 - 2
	}
	return result, nil
}

func Challenge2(data []int) (int, error) {
	result := 0
	for _, v := range data {
		for true {
			if v = v/3 - 2; v <= 0 {
				break
			}
			result = result + v
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 1 - The Tyranny of the Rocket Equation")
	path := "thetyrannyoftherocketequation/input.txt"
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
