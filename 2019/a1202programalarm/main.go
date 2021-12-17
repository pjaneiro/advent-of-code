package a1202programalarm

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		line := scanner.Text()
		for _, v := range strings.Split(line, ",") {
			n, err := strconv.Atoi(v)
			if err != nil {
				return lines, errors.New("something went wrong scanning int")
			}
			lines = append(lines, n)
		}
	}

	return lines, nil
}

func Challenge1(input []int, mode1201 bool) (int, error) {
	data := make([]int, len(input))
	copy(data, input)
	if mode1201 {
		data[1] = 12
		data[2] = 2
	}
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case 1:
			data[data[i+3]] = data[data[i+1]] + data[data[i+2]]
			i = i + 3
		case 2:
			data[data[i+3]] = data[data[i+1]] * data[data[i+2]]
			i = i + 3
		case 99:
			return data[0], nil
		default:
			return 0, errors.New("invalid opcode")
		}
	}
	return 0, errors.New("something went wrong")
}

func challenge2loop(data []int, n int, v int) (int, error) {
	data[1] = n
	data[2] = v
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case 1:
			if data[i+3] >= len(data) || data[i+2] >= len(data) || data[i+1] >= len(data) {
				return 0, errors.New("out of bounds")
			}
			data[data[i+3]] = data[data[i+1]] + data[data[i+2]]
			i = i + 3
		case 2:
			if data[i+3] >= len(data) || data[i+2] >= len(data) || data[i+1] >= len(data) {
				return 0, errors.New("out of bounds")
			}
			data[data[i+3]] = data[data[i+1]] * data[data[i+2]]
			i = i + 3
		case 99:
			return data[0], nil
		default:
			return 0, errors.New("invalid opcode")
		}
	}
	return 0, errors.New("something went wrong")
}

func Challenge2(data []int) (int, error) {
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			temp := make([]int, len(data))
			copy(temp, data)
			res, err := challenge2loop(temp, n, v)
			if err == nil && res == 19690720 {
				return 100*n + v, nil
			}
		}
	}
	return 0, errors.New("no solution found")
}

func Run() {
	fmt.Println("Day 2 - 1202 Program Alarm")
	path := "a1202programalarm/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(data, true)
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
