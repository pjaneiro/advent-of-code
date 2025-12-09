package movietheater

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type tile struct {
	y, x int
}

type line struct {
	t1, t2 tile
}

type rectangle struct {
	c1, c2 tile
	area   int
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

func parseInput(data []string) ([]tile, error) {
	var result []tile = make([]tile, 0)

	for _, row := range data {
		parts := strings.Split(row, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return result, err
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return result, err
		}
		result = append(result, tile{y, x})
	}

	return result, nil
}

func abs(val int) int {
	if val >= 0 {
		return val
	}
	return -val
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Challenge1(data []string) (int, error) {
	tiles, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			area := (abs(tiles[i].y-tiles[j].y) + 1) * (abs(tiles[i].x-tiles[j].x) + 1)
			if area > result {
				result = area
			}
		}
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	tiles, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	var lines []line = make([]line, 0, len(tiles))
	for i := 0; i < len(tiles); i++ {
		t1, t2 := tiles[i], tiles[(i+1)%len(tiles)]
		lines = append(lines, line{t1: t1, t2: t2})
	}
	var rectangles []rectangle = make([]rectangle, 0, len(tiles)*len(tiles))
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles); j++ {
			if i == j {
				continue
			}
			c1, c2 := tiles[i], tiles[j]
			area := (abs(c1.y-c2.y) + 1) * (abs(c1.x-c2.x) + 1)
			rectangles = append(rectangles, rectangle{c1: c1, c2: c2, area: area})
		}
	}
	slices.SortFunc(rectangles, func(a, b rectangle) int {
		return b.area - a.area
	})
outer:
	for _, r := range rectangles {
		for _, l := range lines {
			if l.t1.x == l.t2.x {
				x := l.t1.x
				// vertical
				if x > min(r.c1.x, r.c2.x) && x < max(r.c1.x, r.c2.x) && min(l.t1.y, l.t2.y) < max(r.c1.y, r.c2.y) && max(l.t1.y, l.t2.y) > min(r.c1.y, r.c2.y) {
					continue outer
				}
			} else {
				y := l.t1.y
				// horizontal
				if y > min(r.c1.y, r.c2.y) && y < max(r.c1.y, r.c2.y) && min(l.t1.x, l.t2.x) < max(r.c1.x, r.c2.x) && max(l.t1.x, l.t2.x) > min(r.c1.x, r.c2.x) {
					continue outer
				}
			}
		}
		return r.area, nil
	}
	return 0, errors.New("something went wrong")
}

func Run() {
	fmt.Println("Day 9 - Movie Theater")
	path := "movietheater/input.txt"
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
