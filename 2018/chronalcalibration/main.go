package chronalcalibration

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Operation struct {
	Op  byte
	Val int
}

func readLines(path string) ([]Operation, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Operation
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var op Operation
		op.Op = line[0]
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			return lines, errors.New("something went wrong scanning input")
		}
		op.Val = val
		lines = append(lines, op)
	}

	return lines, nil
}

func Challenge1(data []Operation) (int, error) {
	result := 0
	for _, op := range data {
		switch op.Op {
		case '+':
			result += op.Val
		case '-':
			result -= op.Val
		default:
			return 0, errors.New("invalid operation")
		}
	}
	return result, nil
}

func Challenge2(data []Operation) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("invalid input")
	}
	result, index := 0, 0
	var hits map[int]struct{} = make(map[int]struct{})
	hits[0] = struct{}{}
	for true {
		op := data[index]
		switch op.Op {
		case '+':
			result += op.Val
		case '-':
			result -= op.Val
		default:
			return 0, errors.New("invalid operation")
		}
		if _, ok := hits[result]; ok {
			return result, nil
		}
		hits[result] = struct{}{}
		index = (index + 1) % len(data)
	}
	return 0, errors.New("something went wrong")
}

func Run() {
	fmt.Println("Day 1 - Chronal Calibration")
	path := "chronalcalibration/input.txt"
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
