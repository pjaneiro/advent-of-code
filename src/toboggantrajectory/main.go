package toboggantrajectory

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
	return lines, scanner.Err()
}

func Challenge1(data []string) (int, error) {
	var result, i, j int = 0, 0, 0
	var height, width int = len(data), len(data[0])
	for true {
		if i == height {
			return result, nil
		}

		if data[i][j:j+1] == "#" {
			result++
		}

		i++
		j = (j + 3) % width
	}
	return 0, errors.New("something went wrong")
}

func Challenge2(data []string) (int, error) {
	var result int = 1
	var height, width int = len(data), len(data[0])
	var slopes [][]int = [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, s := range slopes {
		var sresult, i, j int = 0, 0, 0
		for true {
			if i >= height {
				result = result * sresult
				break
			}

			if data[i][j:j+1] == "#" {
				sresult++
			}

			i = i + s[1]
			j = (j + s[0]) % width
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 3")
	path := "src/toboggantrajectory/input.txt"
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
