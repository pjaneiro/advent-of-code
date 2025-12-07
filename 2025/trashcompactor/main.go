package trashcompactor

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation int

const (
	Add  Operation = iota
	Mult Operation = iota
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

func parseInput(data []string) ([][]int, []Operation, error) {
	var matrix [][]int = make([][]int, 0)
	var operations []Operation = make([]Operation, 0)
	var i int
	for i = 0; i < len(data)-1; i++ {
		values := strings.Split(data[i], " ")
		var newRow []int = make([]int, 0)
		for _, val := range values {
			val = strings.Trim(val, " ")
			if len(val) == 0 {
				continue
			}
			x, err := strconv.Atoi(val)
			if err != nil {
				return nil, nil, err
			}
			newRow = append(newRow, x)
		}
		matrix = append(matrix, newRow)
	}

	ops := strings.Split(data[len(data)-1], " ")
	for _, op := range ops {
		op = strings.Trim(op, " ")
		if len(op) == 0 {
			continue
		}
		switch op {
		case "+":
			operations = append(operations, Add)
		case "*":
			operations = append(operations, Mult)
		default:
			return nil, nil, errors.New("invalid operation")
		}
	}
	return matrix, operations, nil
}

func Challenge1(data []string) (int, error) {
	matrix, operations, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	var inters []int = make([]int, len(matrix[0]))
	copy(inters, matrix[0])
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(inters); j++ {
			switch operations[j] {
			case Add:
				inters[j] += matrix[i][j]
			case Mult:
				inters[j] *= matrix[i][j]
			default:
				return 0, errors.New("invalid operation")
			}
		}
	}
	for _, val := range inters {
		result += val
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	result := 0
	var curOp Operation = -1
	var curValues []int = make([]int, 0)

	for i := len(data[0]) - 1; i >= 0; i-- {
		var curCol int = 0
		empty := true
		for j := 0; j < len(data); j++ {
			if data[j][i] >= 48 && data[j][i] <= 57 {
				empty = false
				curCol = (curCol * 10) + int(data[j][i]) - 48
			} else if data[j][i] == 43 {
				empty = false
				curOp = Add
			} else if data[j][i] == 42 {
				empty = false
				curOp = Mult
			}
		}
		if i == 0 {
			empty = true
			curValues = append(curValues, curCol)
		}
		if empty {
			var curRes int
			switch curOp {
			case Add:
				curRes = 0
				for _, cur := range curValues {
					curRes += cur
				}
			case Mult:
				curRes = 1
				for _, cur := range curValues {
					curRes *= cur
				}
			default:
				return 0, errors.New("you fucked up")
			}
			result += curRes
			curValues = make([]int, 0)
		} else {
			curValues = append(curValues, curCol)
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 6 - Trash Compactor")
	path := "trashcompactor/input.txt"
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
