package laboratories

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	row, col int
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

func Challenge1(data []string) (int, error) {
	var result int = 0
	var beams []bool = make([]bool, len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		if data[0][i] == 'S' {
			beams[i] = true
		}
	}
	for i := 1; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == '^' && beams[j] {
				result++
				beams[j] = false
				if j > 0 {
					beams[j-1] = true
				}
				if j < len(data[i])-1 {
					beams[j+1] = true
				}
			}
		}
	}
	return result, nil
}

func testTimeline(data []string, row int, col int, memory map[Point]int) int {
	coord := Point{row, col}
	if row == len(data)-1 {
		memory[coord] = 1
		return 1
	}
	if data[row][col] == '.' {
		if count, ok := memory[coord]; ok {
			return count
		}
		memory[coord] = testTimeline(data, row+1, col, memory)
		return memory[coord]
	} else {
		if count, ok := memory[coord]; ok {
			return count
		}
		memory[coord] = testTimeline(data, row+1, col-1, memory) + testTimeline(data, row+1, col+1, memory)
		return memory[coord]
	}
}

func Challenge2(data []string) (int, error) {
	var start int
	for i := 0; i < len(data[0]); i++ {
		if data[0][i] == 'S' {
			start = i
		}
	}
	var memory map[Point]int = make(map[Point]int, 0)
	return testTimeline(data, 1, start, memory), nil
}

func Run() {
	fmt.Println("Day 7 - Laboratories")
	path := "laboratories/input.txt"
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
