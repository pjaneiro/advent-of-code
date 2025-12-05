package printingdepartment

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct {
	y, x int
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

func countReachable(floor []string) []point {
	var reachable []point = make([]point, 0)
	for y := 0; y < len(floor); y++ {
		for x := 0; x < len(floor[y]); x++ {
			if floor[y][x] != '@' {
				continue
			}
			count := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dy == 0 && dx == 0 {
						continue
					}
					if y+dy < 0 || y+dy >= len(floor) || x+dx < 0 || x+dx >= len(floor[y]) {
						continue
					}
					if floor[y+dy][x+dx] == '@' {
						count++
					}
				}
			}
			if count < 4 {
				reachable = append(reachable, point{y, x})
			}
		}
	}
	return reachable
}

func Challenge1(data []string) (int, error) {
	return len(countReachable(data)), nil
}

func Challenge2(data []string) (int, error) {
	result, missing := 0, math.MaxInt
	for missing != 0 {
		reachable := countReachable(data)
		missing = len(reachable)
		result += missing
		for _, cur := range reachable {
			data[cur.y] = data[cur.y][:cur.x] + "." + data[cur.y][cur.x+1:]
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 4 - Printing Department")
	path := "printingdepartment/input.txt"
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
