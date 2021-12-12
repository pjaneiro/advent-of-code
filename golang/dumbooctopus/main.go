package dumbooctopus

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func readLines(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []int
		line := scanner.Text()
		for _, chr := range line {
			val, err := strconv.Atoi(string(chr))
			if err != nil {
				return nil, err
			}

			row = append(row, val)
		}
		matrix = append(matrix, row)
	}

	return matrix, nil
}

func flash(data [][]int, flashes map[Point]bool, x int, y int) {
	if flashes[Point{X: x, Y: y}] == true {
		return
	} else {
		flashes[Point{X: x, Y: y}] = true
	}
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			if dx == 0 && dy == 0 {
				continue
			}
			newX, newY := x+dx, y+dy
			if newX >= 0 && newX < len(data) && newY >= 0 && newY < len(data) {
				data[newX][newY]++
				if data[newX][newY] == 10 {
					flash(data, flashes, newX, newY)
				}
			}
		}
	}
}

func Challenge1(data [][]int) (int, error) {
	count := 0
	for step := 0; step < 100; step++ {
		flashes := make(map[Point]bool)
		for i, _ := range data {
			for j, _ := range data[i] {
				data[i][j]++
				if data[i][j] > 9 {
					flash(data, flashes, i, j)
				}
			}
		}
		for i, _ := range data {
			for j, _ := range data[i] {
				if data[i][j] > 9 {
					count++
					data[i][j] = 0
				}
			}
		}
	}
	return count, nil
}

func Challenge2(data [][]int) (int, error) {
	for step := 0; true; step++ {
		flashes := make(map[Point]bool)
		for i, _ := range data {
			for j, _ := range data[i] {
				data[i][j]++
				if data[i][j] > 9 {
					flash(data, flashes, i, j)
				}
			}
		}
		for i, _ := range data {
			for j, _ := range data[i] {
				if data[i][j] > 9 {
					data[i][j] = 0
				}
			}
		}
		if len(flashes) == len(data)*len(data[0]) {
			return step + 1, nil
		}
	}
	return 0, errors.New("something went wrong")
}

func Run() {
	fmt.Println("Day 11 - Dumbo Octopus")
	path := "dumbooctopus/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var cpy [][]int
	var result int

	cpy = make([][]int, len(data))
	for i, _ := range data {
		cpy[i] = make([]int, len(data[i]))
		copy(cpy[i], data[i])
	}
	result, err = Challenge1(cpy)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	cpy = make([][]int, len(data))
	for i, _ := range data {
		cpy[i] = make([]int, len(data[i]))
		copy(cpy[i], data[i])
	}
	result, err = Challenge2(cpy)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
