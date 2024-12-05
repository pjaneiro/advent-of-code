package ceressearch

import (
	"bufio"
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
	return lines, nil
}

func Challenge1(data []string) (int, error) {
	result := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] != 'X' {
				continue
			}
			if i < len(data)-3 && data[i+1][j] == 'M' && data[i+2][j] == 'A' && data[i+3][j] == 'S' {
				result++
			}
			if i > 2 && data[i-1][j] == 'M' && data[i-2][j] == 'A' && data[i-3][j] == 'S' {
				result++
			}
			if j < len(data[i])-3 && data[i][j+1] == 'M' && data[i][j+2] == 'A' && data[i][j+3] == 'S' {
				result++
			}
			if j > 2 && data[i][j-1] == 'M' && data[i][j-2] == 'A' && data[i][j-3] == 'S' {
				result++
			}
			if i < len(data)-3 && j < len(data[i])-3 && data[i+1][j+1] == 'M' && data[i+2][j+2] == 'A' && data[i+3][j+3] == 'S' {
				result++
			}
			if i > 2 && j < len(data[i])-3 && data[i-1][j+1] == 'M' && data[i-2][j+2] == 'A' && data[i-3][j+3] == 'S' {
				result++
			}
			if i < len(data)-3 && j > 2 && data[i+1][j-1] == 'M' && data[i+2][j-2] == 'A' && data[i+3][j-3] == 'S' {
				result++
			}
			if i > 2 && j > 2 && data[i-1][j-1] == 'M' && data[i-2][j-2] == 'A' && data[i-3][j-3] == 'S' {
				result++
			}
		}
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	result := 0
	for i := 0; i < len(data)-2; i++ {
		for j := 0; j < len(data[i])-2; j++ {
			if data[i][j] != 'M' && data[i][j] != 'S' {
				continue
			}
			if data[i][j] == 'M' && data[i+1][j+1] == 'A' && data[i+2][j+2] == 'S' || data[i][j] == 'S' && data[i+1][j+1] == 'A' && data[i+2][j+2] == 'M' {
				if data[i+2][j] == 'M' && data[i+1][j+1] == 'A' && data[i][j+2] == 'S' || data[i+2][j] == 'S' && data[i+1][j+1] == 'A' && data[i][j+2] == 'M' {
					result++
				}
			}
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 4 - Ceres Search")
	path := "ceressearch/input.txt"
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
