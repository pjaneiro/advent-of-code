package supplystacks

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Instruction struct {
	Amt, From, To int
}

func readLines(path string) ([]Instruction, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r, err := regexp.Compile(`^move (\d+) from (\d+) to (\d+)$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}

	var result []Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := r.FindStringSubmatch(line)
		var entry Instruction
		if entry.Amt, err = strconv.Atoi(data[1]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if entry.From, err = strconv.Atoi(data[2]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if entry.To, err = strconv.Atoi(data[3]); err != nil {
			return nil, errors.New("couldn't parse input")
		}

		result = append(result, entry)
	}

	return result, nil
}

func Challenge1(state [][]byte, data []Instruction) (string, error) {
	for _, entry := range data {
		for i := 0; i < entry.Amt; i++ {
			size := len(state[entry.From-1])
			val := state[entry.From-1][size-1]
			state[entry.From-1] = state[entry.From-1][0 : size-1]
			state[entry.To-1] = append(state[entry.To-1], val)
		}
	}
	var result string
	for _, col := range state {
		result += string(rune(col[len(col)-1]))
	}
	return result, nil
}

func Challenge2(state [][]byte, data []Instruction) (string, error) {
	for _, entry := range data {
		size := len(state[entry.From-1])
		val := state[entry.From-1][size-entry.Amt:]
		state[entry.From-1] = state[entry.From-1][0 : size-entry.Amt]
		state[entry.To-1] = append(state[entry.To-1], val...)
	}
	var result string
	for _, col := range state {
		result += string(rune(col[len(col)-1]))
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 5 - Supply Stacks")
	path := "supplystacks/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	state := [][]byte{
		[]byte{'N', 'R', 'G', 'P'},
		[]byte{'J', 'T', 'B', 'L', 'F', 'G', 'D', 'C'},
		[]byte{'M', 'S', 'V'},
		[]byte{'L', 'S', 'R', 'C', 'Z', 'P'},
		[]byte{'P', 'S', 'L', 'V', 'C', 'W', 'D', 'Q'},
		[]byte{'C', 'T', 'N', 'W', 'D', 'M', 'S'},
		[]byte{'H', 'D', 'G', 'W', 'P'},
		[]byte{'Z', 'L', 'P', 'H', 'S', 'C', 'M', 'V'},
		[]byte{'R', 'P', 'F', 'L', 'W', 'G', 'Z'},
	}

	var result string
	result, err = Challenge1(state, data)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %s\n", result)
	}

	state = [][]byte{
		[]byte{'N', 'R', 'G', 'P'},
		[]byte{'J', 'T', 'B', 'L', 'F', 'G', 'D', 'C'},
		[]byte{'M', 'S', 'V'},
		[]byte{'L', 'S', 'R', 'C', 'Z', 'P'},
		[]byte{'P', 'S', 'L', 'V', 'C', 'W', 'D', 'Q'},
		[]byte{'C', 'T', 'N', 'W', 'D', 'M', 'S'},
		[]byte{'H', 'D', 'G', 'W', 'P'},
		[]byte{'Z', 'L', 'P', 'H', 'S', 'C', 'M', 'V'},
		[]byte{'R', 'P', 'F', 'L', 'W', 'G', 'Z'},
	}

	result, err = Challenge2(state, data)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %s\n", result)
	}
}
