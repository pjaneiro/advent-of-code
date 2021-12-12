package smokebasin

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func findLowPoints(data [][]int) []Point {
	var locations []Point
	// All but edges and corners
	var val int
	for i := 1; i < len(data)-1; i++ {
		for j := 1; j < len(data[i])-1; j++ {
			val = data[i][j]
			if data[i-1][j] > val && data[i+1][j] > val && data[i][j-1] > val && data[i][j+1] > val {
				locations = append(locations, Point{X: i, Y: j})
			}
		}
	}

	// Edges
	for i := 1; i < len(data)-1; i++ {
		val = data[i][0]
		if data[i-1][0] > val && data[i+1][0] > val && data[i][1] > val {
			locations = append(locations, Point{X: i, Y: 0})
		}
		val = data[i][len(data[i])-1]
		if data[i-1][len(data[i])-1] > val && data[i+1][len(data[i])-1] > val && data[i][len(data[i])-2] > val {
			locations = append(locations, Point{X: i, Y: len(data[i]) - 1})
		}
	}
	for j := 1; j < len(data[0])-1; j++ {
		val = data[0][j]
		if data[0][j-1] > val && data[0][j+1] > val && data[1][j] > val {
			locations = append(locations, Point{X: 0, Y: j})
		}
		val = data[len(data)-1][j]
		if data[len(data)-1][j-1] > val && data[len(data)-1][j+1] > val && data[len(data)-2][j] > val {
			locations = append(locations, Point{X: len(data) - 1, Y: j})
		}
	}

	// Corners
	val = data[0][0]
	if data[0][1] > val && data[1][0] > val {
		locations = append(locations, Point{X: 0, Y: 0})
	}
	val = data[0][len(data[0])-1]
	if data[0][len(data[0])-2] > val && data[1][len(data[0])-1] > val {
		locations = append(locations, Point{X: 0, Y: len(data[0]) - 1})
	}
	val = data[len(data)-1][0]
	if data[len(data)-1][1] > val && data[1][len(data[0])-2] > val {
		locations = append(locations, Point{X: len(data) - 1, Y: 0})
	}
	val = data[len(data)-1][len(data[0])-1]
	if data[len(data)-1][len(data[0])-2] > val && data[len(data)-2][len(data[0])-1] > val {
		locations = append(locations, Point{X: len(data) - 1, Y: len(data[0]) - 1})
	}
	return locations
}

func basinSize(data [][]int, visited map[Point]bool, x int, y int) int {
	if x < 0 || y < 0 || x >= len(data) || y >= len(data[0]) || data[x][y] == 9 {
		return 0
	} else {
		val := data[x][y]
		var up, down, left, right int
		if x > 0 && data[x-1][y] > val {
			if _, ok := visited[Point{X: x - 1, Y: y}]; !ok {
				up = basinSize(data, visited, x-1, y)
			}
		}
		if x < len(data)-1 && data[x+1][y] > val {
			if _, ok := visited[Point{X: x + 1, Y: y}]; !ok {
				down = basinSize(data, visited, x+1, y)
			}
		}
		if y > 0 && data[x][y-1] > val {
			if _, ok := visited[Point{X: x, Y: y - 1}]; !ok {
				left = basinSize(data, visited, x, y-1)
			}
		}
		if y < len(data[0])-1 && data[x][y+1] > val {
			if _, ok := visited[Point{X: x, Y: y + 1}]; !ok {
				right = basinSize(data, visited, x, y+1)
			}
		}
		visited[Point{X: x, Y: y}] = true
		return 1 + up + down + left + right
	}
}

func Challenge1(data [][]int) (int, error) {
	locations := findLowPoints(data)
	total := 0
	for _, cur := range locations {
		total = total + data[cur.X][cur.Y] + 1
	}
	return total, nil
}

func Challenge2(data [][]int) (int, error) {
	locations := findLowPoints(data)
	var sizes []int
	for _, cur := range locations {
		visited := make(map[Point]bool)
		curSize := basinSize(data, visited, cur.X, cur.Y)
		sizes = append(sizes, curSize)
	}
	sort.Ints(sizes)
	return (sizes[len(sizes)-3] * sizes[len(sizes)-2] * sizes[len(sizes)-1]), nil
}

func Run() {
	fmt.Println("Day 9 - Smoke Basin")
	path := "smokebasin/input.txt"
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
