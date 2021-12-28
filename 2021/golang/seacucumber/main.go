package seacucumber

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(path string) ([][]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]byte = make([][]byte, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, []byte(scanner.Text()))
	}

	return result, nil
}

func printState(data [][]byte) {
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			fmt.Printf("%s", string(data[y][x]))
		}
		fmt.Println()
	}
}

func Challenge1(data [][]byte) (int, error) {
	steps := 0
	var newState [][]byte
	for true {
		newState = make([][]byte, len(data))
		for i := 0; i < len(data); i++ {
			newState[i] = []byte(strings.Repeat(".", len(data[i])))
		}
		moved := false

		for y := 0; y < len(data); y++ {
			for x := 0; x < len(data[y]); x++ {
				if data[y][x] == '>' {
					if data[y][(x+1)%len(data[y])] == '.' {
						newState[y][(x+1)%len(data[y])] = '>'
						moved = true
					} else {
						newState[y][x] = '>'
					}
				} else if data[y][x] == 'v' {
					newState[y][x] = data[y][x]
				}
			}
		}
		data = newState

		newState = make([][]byte, len(data))
		for i := 0; i < len(data); i++ {
			newState[i] = []byte(strings.Repeat(".", len(data[i])))
		}

		for y := 0; y < len(data); y++ {
			for x := 0; x < len(data[y]); x++ {
				if data[y][x] == 'v' {
					if data[(y+1)%len(data)][x] == '.' {
						newState[(y+1)%len(data)][x] = 'v'
						moved = true
					} else {
						newState[y][x] = 'v'
					}
				} else if data[y][x] == '>' {
					newState[y][x] = '>'
				}
			}
		}
		data = newState

		steps++
		if !moved {
			break
		}
	}
	return steps, nil
}

func Run() {
	fmt.Println("Day 25 - Sea Cucumber")
	path := "seacucumber/input.txt"
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
}
