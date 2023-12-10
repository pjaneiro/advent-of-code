package hauntedwasteland

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Mapping map[string]string

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

func parseInput(data []string) (string, []string, Mapping, Mapping, error) {
	var instructions string = data[0]
	var keys []string = make([]string, 0)
	var left, right Mapping = make(Mapping), make(Mapping)
	r, err := regexp.Compile(`^(\w{3}) = \((\w{3}), (\w{3})\)$`)
	if err != nil {
		return instructions, nil, nil, nil, err
	}
	for i := 2; i < len(data); i++ {
		line := data[i]
		fields := r.FindStringSubmatch(line)
		keys = append(keys, fields[1])
		left[fields[1]] = fields[2]
		right[fields[1]] = fields[3]
	}
	return instructions, keys, left, right, nil
}

func Challenge1(data []string) (int, error) {
	instructions, _, left, right, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	cur := "AAA"
	for i := 0; ; i = (i + 1) % len(instructions) {
		result++
		switch instructions[i] {
		case 'L':
			cur = left[cur]
		case 'R':
			cur = right[cur]
		}
		if cur == "ZZZ" {
			break
		}
	}
	return result, nil
}

func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func Challenge2(data []string) (int, error) {
	instructions, keys, left, right, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	var distances []int = make([]int, 0)
	for _, start := range keys {
		if start[2] != 'A' {
			continue
		}
		steps := 0
		cur := start
		for i := 0; ; i = (i + 1) % len(instructions) {
			steps++
			switch instructions[i] {
			case 'L':
				cur = left[cur]
			case 'R':
				cur = right[cur]
			}
			if cur[2] == 'Z' {
				distances = append(distances, steps)
				break
			}
		}
	}
	result := lcm(distances[0], distances[1])
	for i := 2; i < len(distances); i++ {
		result = lcm(result, distances[i])
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 8 - Haunted Wastelands")
	path := "hauntedwasteland/input.txt"
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
