package dockingdata

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Operation struct {
	Mask    string
	Address int64
	Value   int64
}

func readLines(path string) ([]Operation, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Operation
	var currentmask string
	rmask, err := regexp.Compile(`^mask = ([01X]{36})$`)
	if err != nil {
		return lines, errors.New("couldn't parse regex expression")
	}
	rop, err := regexp.Compile(`^mem\[(\d+)\] = (\d+)$`)
	if err != nil {
		return lines, errors.New("couldn't parse regex expression")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if rmask.MatchString(line) {
			data := rmask.FindStringSubmatch(line)
			currentmask = data[1]
		} else if rop.MatchString(line) {
			data := rop.FindStringSubmatch(line)
			add, err := strconv.ParseInt(data[1], 10, 64)
			if err != nil {
				return lines, errors.New("couldn't extract all data")
			}
			val, err := strconv.ParseInt(data[2], 10, 64)
			if err != nil {
				return lines, errors.New("couldn't extract all data")
			}
			op := Operation{Mask: currentmask, Address: add, Value: val}
			lines = append(lines, op)
		} else {
			fmt.Println(line)
			return lines, errors.New("weird line found")
		}
	}
	return lines, nil
}

func Challenge1(data []Operation) (int64, error) {
	var mem map[int64]int64 = make(map[int64]int64)
	for _, op := range data {
		value := op.Value
		for i, j := range op.Mask {
			if j == 'X' {
				continue
			} else if j == '1' {
				value = value | (1 << (35 - i))
			} else {
				value = value & (^(1 << (35 - i)) & 0b0000000000000000000000000000111111111111111111111111111111111111)
			}
		}
		mem[op.Address] = value
	}
	result := int64(0)
	for _, x := range mem {
		result += x
	}
	return result, nil
}

func allIndexes(data string, char rune) []int {
	result := []int{}
	for i, j := range data {
		if j == char {
			result = append(result, i)
		}
	}
	return result
}

func allVariants(data string, indexes []int) []string {
	chars := []rune(data)
	for _, j := range indexes {
		chars[j] = 'X'
	}
	data = string(chars)
	result := []string{data}
	for true {
		done := true
		for _, x := range result {
			if strings.Index(x, "X") != -1 {
				done = false
				break
			}
		}
		if done {
			break
		}
		for i, j := range result {
			index := strings.Index(j, "X")
			if index == -1 {
				continue
			}
			op0 := strings.Replace(j, "X", "0", 1)
			op1 := strings.Replace(j, "X", "1", 1)
			result = append(result[:i], append([]string{op0, op1}, result[i+1:]...)...)
			break
		}

	}
	return result
}

func Challenge2(data []Operation) (int64, error) {
	var mem map[int64]int64 = make(map[int64]int64)
	for _, op := range data {
		address := op.Address
		for i, j := range op.Mask {
			if j == '0' || j == 'X' {
				continue
			}
			address = address | (1 << (35 - i))
		}
		addrstring := fmt.Sprintf("%036s", strconv.FormatInt(address, 2))
		indexes := allIndexes(op.Mask, 'X')
		for _, j := range allVariants(addrstring, indexes) {
			address, _ = strconv.ParseInt(j, 2, 64)
			mem[address] = op.Value
		}

	}
	result := int64(0)
	for _, x := range mem {
		result += x
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 14")
	path := "dockingdata/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int64
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
