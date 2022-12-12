package hillclimbingalgorithm

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Coord struct {
	X, Y int
}

type Item struct {
	coord Coord
	cost  int
}

func readLines(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, []rune(scanner.Text()))
	}
	return result, nil
}

func copyData(data [][]rune) [][]rune {
	result := make([][]rune, len(data))
	for i, line := range data {
		result[i] = make([]rune, len(line))
		copy(result[i], line)
	}
	return result
}

func travel(data [][]rune, start Coord, goal Coord, climbing bool) int {
	var queue []Item
	if climbing {
		queue = append(queue, Item{coord: start, cost: 0})
	} else {
		queue = append(queue, Item{coord: goal, cost: 0})
	}

	var visited map[Coord]struct{} = make(map[Coord]struct{}, len(data)*len(data[0]))

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		coord, cost := current.coord, current.cost
		x, y := coord.X, coord.Y
		if climbing && coord == goal || !climbing && data[y][x] == 'a' {
			return cost
		}
		if _, ok := visited[coord]; ok {
			continue
		}
		visited[coord] = struct{}{}
		neighbours := []Coord{Coord{X: x, Y: y - 1}, Coord{X: x, Y: y + 1}, Coord{X: x - 1, Y: y}, Coord{X: x + 1, Y: y}}
		for _, neighbour := range neighbours {
			if neighbour.X < 0 || neighbour.Y < 0 || neighbour.X >= len(data[0]) || neighbour.Y >= len(data) {
				continue
			}
			if climbing && data[neighbour.Y][neighbour.X]-data[y][x] <= 1 || !climbing && data[y][x]-data[neighbour.Y][neighbour.X] <= 1 {
				queue = append(queue, Item{coord: neighbour, cost: cost + 1})
			}
		}
	}
	return math.MaxInt
}

func Challenge1(data [][]rune) (int, error) {
	mapping := copyData(data)
	var start, goal Coord
	for y := 0; y < len(mapping); y++ {
		for x := 0; x < len(mapping[y]); x++ {
			if mapping[y][x] == 'S' {
				start = Coord{X: x, Y: y}
				mapping[y][x] = 'a'
			}
			if mapping[y][x] == 'E' {
				goal = Coord{X: x, Y: y}
				mapping[y][x] = 'z'
			}
		}
	}

	return travel(mapping, start, goal, true), nil
}

func Challenge2(data [][]rune) (int, error) {
	mapping := copyData(data)
	var start, goal Coord
	for y := 0; y < len(mapping); y++ {
		for x := 0; x < len(mapping[y]); x++ {
			if mapping[y][x] == 'S' {
				start = Coord{X: x, Y: y}
				mapping[y][x] = 'a'
			}
			if mapping[y][x] == 'E' {
				goal = Coord{X: x, Y: y}
				mapping[y][x] = 'z'
			}
		}
	}

	return travel(mapping, start, goal, false), nil
}

func Run() {
	fmt.Println("Day 12 - Hill Climbing Algorithm")
	path := "hillclimbingalgorithm/input.txt"

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
