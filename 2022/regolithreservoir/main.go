package regolithreservoir

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, nil
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func inputToMap(data []string) (map[Coord]struct{}, int, error) {
	maxY := 0
	var blocks map[Coord]struct{} = make(map[Coord]struct{})

	for _, line := range data {
		points := strings.Split(line, " -> ")
		start := strings.Split(points[0], ",")
		startX, err := strconv.Atoi(start[0])
		if err != nil {
			return nil, 0, errors.New("error parsing")
		}
		startY, err := strconv.Atoi(start[1])
		if err != nil {
			return nil, 0, errors.New("error parsing")
		}
		coord := Coord{X: startX, Y: startY}
		blocks[coord] = struct{}{}
		for i := 1; i < len(points); i++ {
			target := strings.Split(points[i], ",")
			targetX, err := strconv.Atoi(target[0])
			if err != nil {
				return nil, 0, errors.New("error parsing")
			}
			targetY, err := strconv.Atoi(target[1])
			if err != nil {
				return nil, 0, errors.New("error parsing")
			}

			for x := min(startX, targetX); x <= max(startX, targetX); x++ {
				for y := min(startY, targetY); y <= max(startY, targetY); y++ {
					if y > maxY {
						maxY = y
					}
					coord = Coord{X: x, Y: y}
					blocks[coord] = struct{}{}
				}
			}
			startX, startY = targetX, targetY
		}
	}
	return blocks, maxY, nil
}

func Challenge1(data []string) (int, error) {
	result := 0
	blocks, maxY, err := inputToMap(data)
	if err != nil {
		return 0, err
	}

	for {
		sand := Coord{X: 500, Y: 0}
		if _, ok := blocks[sand]; ok {
			break
		}
		for {
			if sand.Y >= maxY {
				return result, nil
			}
			down := Coord{X: sand.X, Y: sand.Y + 1}
			if _, ok := blocks[down]; !ok {
				sand = down
				continue
			}
			downLeft := Coord{X: sand.X - 1, Y: sand.Y + 1}
			if _, ok := blocks[downLeft]; !ok {
				sand = downLeft
				continue
			}
			downRight := Coord{X: sand.X + 1, Y: sand.Y + 1}
			if _, ok := blocks[downRight]; !ok {
				sand = downRight
				continue
			}
			blocks[sand] = struct{}{}
			result++
			break
		}
	}

	return 0, errors.New("something went wrong")
}

func Challenge2(data []string) (int, error) {
	result := 0
	blocks, maxY, err := inputToMap(data)
	if err != nil {
		return 0, err
	}

	for {
		sand := Coord{X: 500, Y: 0}
		if _, ok := blocks[sand]; ok {
			return result, nil
		}
		for {
			if sand.Y > maxY {
				blocks[sand] = struct{}{}
				result++
				break
			}
			down := Coord{X: sand.X, Y: sand.Y + 1}
			if _, ok := blocks[down]; !ok {
				sand = down
				continue
			}
			downLeft := Coord{X: sand.X - 1, Y: sand.Y + 1}
			if _, ok := blocks[downLeft]; !ok {
				sand = downLeft
				continue
			}
			downRight := Coord{X: sand.X + 1, Y: sand.Y + 1}
			if _, ok := blocks[downRight]; !ok {
				sand = downRight
				continue
			}
			blocks[sand] = struct{}{}
			result++
			break
		}
	}

	return result, nil
}

func Run() {
	fmt.Println("Day 14 - Regolith Reservoir")
	path := "regolithreservoir/input.txt"

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
