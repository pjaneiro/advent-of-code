package bridgerepair

import (
	"bufio"
	"fmt"
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

func parseInput(data []string) ([]int, [][]int, error) {
	var tests []int = make([]int, 0, len(data))
	var values [][]int = make([][]int, 0, len(data))

	for _, row := range data {
		parts := strings.Split(row, ":")
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return tests, values, err
		}
		tests = append(tests, left)
		operandStrings := strings.Split(strings.Trim(parts[1], " "), " ")
		var operands []int = make([]int, 0, len(operandStrings))
		for _, op := range operandStrings {
			val, err := strconv.Atoi(op)
			if err != nil {
				return tests, values, err
			}
			operands = append(operands, val)
		}
		values = append(values, operands)
	}

	return tests, values, nil
}

func testOperator1(test int, operands []int, carry int, index int) bool {
	if index == len(operands)-1 {
		return carry == test
	}
	return testOperator1(test, operands, carry+operands[index+1], index+1) || testOperator1(test, operands, carry*operands[index+1], index+1)
}

func Challenge1(data []string) (int, error) {
	tests, values, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	for idx, test := range tests {
		if testOperator1(test, values[idx], values[idx][0], 0) {
			result += test
		}
	}
	return result, nil
}

func concatenate(a, b int) (int, error) {
	sa, sb := strconv.Itoa(a), strconv.Itoa(b)
	sc := sa + sb
	return strconv.Atoi(sc)
}

func testOperator2(test int, operands []int, carry int, index int) bool {
	if index == len(operands)-1 {
		return carry == test
	}
	conc, err := concatenate(carry, operands[index+1])
	if err != nil {
		return false
	}
	return testOperator2(test, operands, carry+operands[index+1], index+1) || testOperator2(test, operands, carry*operands[index+1], index+1) || testOperator2(test, operands, conc, index+1)
}

func Challenge2(data []string) (int, error) {
	tests, values, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	for idx, test := range tests {
		if testOperator2(test, values[idx], values[idx][0], 0) {
			result += test
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 7 - Bridge Repair")
	path := "bridgerepair/input.txt"
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
