package secretentrance

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Rotation struct {
	Direction int
	Distance  int
}

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

func parseInput(data []string) ([]Rotation, error) {
	var pairs []Rotation = make([]Rotation, 0)
	r, err := regexp.Compile(`^([LR])(\d+)$`)
	if err != nil {
		return pairs, err
	}
	for _, row := range data {
		var rotation Rotation
		found := r.FindStringSubmatch(row)
		rotation.Direction = -1
		if found[1] == "R" {
			rotation.Direction = 1
		}
		distance, err := strconv.Atoi(found[2])
		if err != nil {
			return pairs, err
		}
		rotation.Distance = distance
		pairs = append(pairs, rotation)
	}
	return pairs, nil
}

func Challenge1(data []string) (int, error) {
	pairs, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result, location := 0, 50
	for _, step := range pairs {
		location = (location + (step.Direction * step.Distance) + 100) % 100
		if location == 0 {
			result++
		}
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	pairs, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result, location := 0, 50
	for _, step := range pairs {
		result += step.Distance / 100
		step.Distance %= 100
		if (location > 0 && step.Direction == -1 && location-step.Distance <= 0) || (step.Direction == 1 && location+step.Distance > 99) {
			result++
		}
		location = (location + (step.Direction * step.Distance) + 100) % 100
		fmt.Println(location, result)
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 1 - Secret Entrance")
	path := "secretentrance/input.txt"
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
