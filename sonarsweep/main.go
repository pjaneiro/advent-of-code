package sonarsweep

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
	count := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			count++
		}
	}
	return count, nil
}

func Challenge2(data []int) (int, error) {
	var newData []int
	count := 0
	for i := 0; i < len(data)-2; i++ {
		newData = append(newData, data[i]+data[i+1]+data[i+2])
	}
	for i := 1; i < len(newData); i++ {
		if newData[i] > newData[i-1] {
			count++
		}
	}
	return count, nil
}

func Run() {
	fmt.Println("Day 1 - Sonar Sweep")
	path := "sonarsweep/input.txt"
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
