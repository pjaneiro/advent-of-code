package arithmeticlogicunit

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Command string
	Op1     string
	Op2     string
}

func max(a int, b int) byte {
	if a > b {
		return byte(a)
	}
	return byte(b)
}

func min(a int, b int) byte {
	if a < b {
		return byte(a)
	}
	return byte(b)
}

func readLines(path string) ([]Instruction, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []Instruction = make([]Instruction, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Fields(line)
		var cur Instruction
		cur.Command = data[0]
		cur.Op1 = data[1]
		if len(data) > 2 {
			cur.Op2 = data[2]
		}
		result = append(result, cur)
	}

	return result, nil
}

func solve(data []Instruction, highest bool) string {
	var subsets [][]Instruction = make([][]Instruction, 0)
	var stack [][]int = make([][]int, 0)
	var carry map[int]int = make(map[int]int)
	for step := 0; step < 14; step++ {
		subsets = append(subsets, data[(18*step)+0:(18*step)+18])
	}

	var result []byte = make([]byte, 14)
	for i := 0; i < 14; i++ {
		a, _ := strconv.Atoi(subsets[i][5].Op2)
		b, _ := strconv.Atoi(subsets[i][15].Op2)
		if a > 0 {
			stack = append(stack, []int{i, b})
		} else {
			xval := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			carry[i] = xval[1] + a
			carry[xval[0]] = -(xval[1] + a)
		}
	}
	if highest {
		for key, val := range carry {
			result[key] = 48 + min(9, 9+val)
		}
	} else {
		for key, val := range carry {
			result[key] = 48 + max(1, 1+val)
		}
	}
	return string(result)
}

func Challenge1(data []Instruction) (string, error) {
	return solve(data, true), nil
}

func Challenge2(data []Instruction) (string, error) {
	return solve(data, false), nil
}

func Run() {
	fmt.Println("Day 24 - Arithmetic Logic Unit")
	path := "arithmeticlogicunit/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result string

	result, err = Challenge1(data)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %s\n", result)
	}

	result, err = Challenge2(data)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %s\n", result)
	}
}
