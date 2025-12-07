package guardgallivant

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type vector struct {
	dy, dx int
}

type point struct {
	y, x int
}

type status struct {
	pos point
	dir vector
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

func getStart(data []string) (point, error) {
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == '^' {
				return point{y, x}, nil
			}
		}
	}
	return point{}, errors.New("something went wrong")
}

func getVisited(data []string) (map[point]struct{}, error) {
	var directions [4]vector = [4]vector{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var visited map[point]struct{} = make(map[point]struct{}, 0)
	var dirIndex int = 0
	curPosition, err := getStart(data)
	if err != nil {
		return nil, err
	}
	visited[curPosition] = struct{}{}
	for {
		next := point{curPosition.y + directions[dirIndex].dy, curPosition.x + directions[dirIndex].dx}
		if next.y < 0 || next.x < 0 || next.y >= len(data) || next.x >= len(data[0]) {
			return visited, nil
		}
		if data[next.y][next.x] != '#' {
			curPosition = next
			visited[curPosition] = struct{}{}
			continue
		}
		for i := 0; i < 3 && data[next.y][next.x] == '#'; i++ {
			dirIndex = (dirIndex + 1) % 4
			next = point{curPosition.y + directions[dirIndex].dy, curPosition.x + directions[dirIndex].dx}
			if next.y < 0 || next.x < 0 || next.y >= len(data) || next.x >= len(data[0]) {
				return visited, nil
			}
		}
		if data[next.y][next.x] == '#' {
			return nil, errors.New("impossible situation")
		}
		curPosition = next
		visited[curPosition] = struct{}{}
	}
}

func Challenge1(data []string) (int, error) {
	visited, err := getVisited(data)
	if err != nil {
		return 0, err
	}
	return len(visited), nil
}

func Challenge2(data []string) (int, error) {
	var result int = 0
	var directions [4]vector = [4]vector{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var dirIndex, startIndex int = 0, 0
	curPosition, err := getStart(data)
	if err != nil {
		return 0, err
	}
	startPosition := curPosition
	routes, err := getVisited(data)
	if err != nil {
		return 0, err
	}
	for obstacle := range routes {
		if obstacle.y == startPosition.y && obstacle.x == startPosition.x {
			continue
		}
		var visited map[status]struct{} = make(map[status]struct{})
		curPosition = startPosition
		dirIndex = startIndex
	outer:
		for {
			next := point{curPosition.y + directions[dirIndex].dy, curPosition.x + directions[dirIndex].dx}
			for next.y >= 0 && next.x >= 0 && next.y < len(data) && next.x < len(data[0]) && data[next.y][next.x] != '#' && (next.y != obstacle.y || next.x != obstacle.x) {
				curPosition = next
				next = point{curPosition.y + directions[dirIndex].dy, curPosition.x + directions[dirIndex].dx}
			}
			if next.y < 0 || next.x < 0 || next.y >= len(data) || next.x >= len(data[0]) {
				break outer
			}
			for i := 0; i < 3 && (data[next.y][next.x] == '#' || next.y == obstacle.y && next.x == obstacle.x); i++ {
				dirIndex = (dirIndex + 1) % 4
				next = point{curPosition.y + directions[dirIndex].dy, curPosition.x + directions[dirIndex].dx}
				if next.y < 0 || next.x < 0 || next.y >= len(data) || next.x >= len(data[0]) {
					break outer
				}
			}
			if data[next.y][next.x] == '#' || next.y == obstacle.y && next.x == obstacle.x {
				return 0, errors.New("impossible situation")
			}
			curPosition = next
			if _, ok := visited[status{curPosition, directions[dirIndex]}]; ok {
				result++
				break
			}
			visited[status{curPosition, directions[dirIndex]}] = struct{}{}
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 6 - Guard Gallivant")
	path := "guardgallivant/input.txt"
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
