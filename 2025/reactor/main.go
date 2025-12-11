package reactor

import (
	"bufio"
	"fmt"
	"os"
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

func parseInput(data []string) map[string][]string {
	var result map[string][]string = make(map[string][]string)
	for _, line := range data {
		blob := strings.Split(line, ": ")
		result[blob[0]] = strings.Split(blob[1], " ")
	}
	return result
}

func traverse(mapping map[string][]string, path string, current string, goal string, memory map[string]int) int {
	if current == goal {
		memory[current] = 1
		return 1
	}
	if val, ok := memory[current]; ok {
		return val
	}

	path += " " + current
	var result int = 0
	for _, output := range mapping[current] {
		if !strings.Contains(path, output) {
			result += traverse(mapping, path, output, goal, memory)
		}
	}
	memory[current] = result
	return result
}

func Challenge1(data []string) (int, error) {
	var mapping map[string][]string = parseInput(data)

	return traverse(mapping, "", "you", "out", make(map[string]int)), nil
}

func Challenge2(data []string) (int, error) {
	var mapping map[string][]string = parseInput(data)

	svr_dac := traverse(mapping, "", "svr", "dac", make(map[string]int))
	dac_fft := traverse(mapping, "", "dac", "fft", make(map[string]int))
	fft_out := traverse(mapping, "", "fft", "out", make(map[string]int))

	svr_fft := traverse(mapping, "", "svr", "fft", make(map[string]int))
	fft_dac := traverse(mapping, "", "fft", "dac", make(map[string]int))
	dac_out := traverse(mapping, "", "dac", "out", make(map[string]int))

	return (svr_dac * dac_fft * fft_out) + (svr_fft * fft_dac * dac_out), nil
}

func Run() {
	fmt.Println("Day 11 - Reactor")
	path := "reactor/input.txt"
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
