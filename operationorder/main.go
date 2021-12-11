package operationorder

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
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

func solveOperation(operation string) int {
	members := strings.Split(operation, " ")
	result, _ := strconv.Atoi(members[0])
	for i := 1; i < len(members); i += 2 {
		nextNum, _ := strconv.Atoi(members[i+1])
		if members[i] == "*" {
			result *= nextNum
		} else {
			result += nextNum
		}
	}
	return result
}

func Challenge1(data []string) (int, error) {
	result := 0
	r, err := regexp.Compile(`\([\d\s\+\*]+\)`)
	if err != nil {
		return 0, errors.New("couldn't compile regex expression")
	}
	for _, operation := range data {
		for strings.Contains(operation, "(") {
			indexes := r.FindStringSubmatchIndex(operation)
			subresult := solveOperation(operation[indexes[0]+1 : indexes[1]-1])
			operation = operation[:indexes[0]] + strconv.Itoa(subresult) + operation[indexes[1]:]
		}
		result += solveOperation(operation)
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	result := 0
	r1, err := regexp.Compile(`\([\d\s\+\*]+\)`)
	if err != nil {
		return 0, errors.New("couldn't compile regex expression")
	}
	r2, err := regexp.Compile(`\d+\s\+\s\d+`)
	if err != nil {
		return 0, errors.New("couldn't compile regex expression")
	}
	for _, operation := range data {
		for strings.Contains(operation, "(") {
			indexes := r1.FindStringSubmatchIndex(operation)
			suboperation := operation[indexes[0]+1 : indexes[1]-1]
			for strings.Contains(suboperation, "+") {
				subindexes := r2.FindStringSubmatchIndex(suboperation)
				subsuboperation := suboperation[subindexes[0]:subindexes[1]]
				subsubresult := solveOperation(subsuboperation)
				suboperation = suboperation[:subindexes[0]] + strconv.Itoa(subsubresult) + suboperation[subindexes[1]:]
			}
			subresult := solveOperation(suboperation)
			operation = operation[:indexes[0]] + strconv.Itoa(subresult) + operation[indexes[1]:]
		}
		for strings.Contains(operation, "+") {
			indexes := r2.FindStringSubmatchIndex(operation)
			suboperation := operation[indexes[0]:indexes[1]]
			subresult := solveOperation(suboperation)
			operation = operation[:indexes[0]] + strconv.Itoa(subresult) + operation[indexes[1]:]
		}
		result += solveOperation(operation)
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 18 - Operation Order")
	path := "operationorder/input.txt"
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
