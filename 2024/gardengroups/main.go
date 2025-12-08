package gardengroups

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
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

func getRegionData(data []string, visited *map[coord]struct{}, point coord) (perimeter, area, corners int) {
	(*visited)[point] = struct{}{}
	area++
	value := data[point.y][point.x]
	top, right, bottom, left := false, false, false, false
	// top
	if point.y == 0 || data[point.y-1][point.x] != value {
		top = true
		perimeter++
	} else if _, ok := (*visited)[coord{point.y - 1, point.x}]; !ok {
		tp, ta, tc := getRegionData(data, visited, coord{point.y - 1, point.x})
		perimeter, area, corners = perimeter+tp, area+ta, corners+tc
	}
	// right
	if point.x == len(data[0])-1 || data[point.y][point.x+1] != value {
		right = true
		perimeter++
	} else if _, ok := (*visited)[coord{point.y, point.x + 1}]; !ok {
		tp, ta, tc := getRegionData(data, visited, coord{point.y, point.x + 1})
		perimeter, area, corners = perimeter+tp, area+ta, corners+tc
	}
	// bottom
	if point.y == len(data)-1 || data[point.y+1][point.x] != value {
		bottom = true
		perimeter++
	} else if _, ok := (*visited)[coord{point.y + 1, point.x}]; !ok {
		tp, ta, tc := getRegionData(data, visited, coord{point.y + 1, point.x})
		perimeter, area, corners = perimeter+tp, area+ta, corners+tc
	}
	// left
	if point.x == 0 || data[point.y][point.x-1] != value {
		left = true
		perimeter++
	} else if _, ok := (*visited)[coord{point.y, point.x - 1}]; !ok {
		tp, ta, tc := getRegionData(data, visited, coord{point.y, point.x - 1})
		perimeter, area, corners = perimeter+tp, area+ta, corners+tc
	}

	if (top && right) || (!top && !right && data[point.y-1][point.x+1] != value) {
		corners++
	}
	if (right && bottom) || (!right && !bottom && data[point.y+1][point.x+1] != value) {
		corners++
	}
	if (bottom && left) || (!bottom && !left && data[point.y+1][point.x-1] != value) {
		corners++
	}
	if (left && top) || (!left && !top && data[point.y-1][point.x-1] != value) {
		corners++
	}

	return perimeter, area, corners
}

func Challenge1(data []string) (int, error) {
	var visited map[coord]struct{} = make(map[coord]struct{})
	result := 0
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if _, ok := visited[coord{y, x}]; !ok {
				p, a, _ := getRegionData(data, &visited, coord{y, x})
				result += p * a
			}
		}
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	var visited map[coord]struct{} = make(map[coord]struct{})
	result := 0
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if _, ok := visited[coord{y, x}]; !ok {
				_, a, c := getRegionData(data, &visited, coord{y, x})
				result += c * a
			}
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 12 - Garden Groups")
	path := "gardengroups/input.txt"
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
