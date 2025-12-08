package hoofit

import (
	"bufio"
	"fmt"
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

func findTrailheads(data []string) []point {
	var result []point = make([]point, 0)
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == '0' {
				result = append(result, point{y, x})
			}
		}
	}
	return result
}

func findTops(data []string, cache *map[point][]point, pos point) []point {
	var result []point = make([]point, 0)
	if data[pos.y][pos.x] == '9' {
		result = append(result, pos)
		(*cache)[pos] = []point{pos}
		return result
	}
	if pos.y > 0 {
		var tmp []point
		coord := point{pos.y - 1, pos.x}
		if data[coord.y][coord.x] == data[pos.y][pos.x]+1 {
			if val, ok := (*cache)[coord]; ok {
				tmp = val
			} else {
				tmp = findTops(data, cache, coord)
			}
			for _, k := range tmp {
				exists := false
				for _, v := range result {
					if k.y == v.y && k.x == v.x {
						exists = true
						break
					}
				}
				if !exists {
					result = append(result, k)
				}
			}
		}
	}
	if pos.y < len(data)-1 {
		var tmp []point
		coord := point{pos.y + 1, pos.x}
		if data[coord.y][coord.x] == data[pos.y][pos.x]+1 {
			if val, ok := (*cache)[coord]; ok {
				tmp = val
			} else {
				tmp = findTops(data, cache, coord)
			}
			for _, k := range tmp {
				exists := false
				for _, v := range result {
					if k.y == v.y && k.x == v.x {
						exists = true
						break
					}
				}
				if !exists {
					result = append(result, k)
				}
			}
		}
	}
	if pos.x > 0 {
		var tmp []point
		coord := point{pos.y, pos.x - 1}
		if data[coord.y][coord.x] == data[pos.y][pos.x]+1 {
			if val, ok := (*cache)[coord]; ok {
				tmp = val
			} else {
				tmp = findTops(data, cache, coord)
			}
			for _, k := range tmp {
				exists := false
				for _, v := range result {
					if k.y == v.y && k.x == v.x {
						exists = true
						break
					}
				}
				if !exists {
					result = append(result, k)
				}
			}
		}
	}
	if pos.x < len(data[0])-1 {
		var tmp []point
		coord := point{pos.y, pos.x + 1}
		if data[coord.y][coord.x] == data[pos.y][pos.x]+1 {
			if val, ok := (*cache)[coord]; ok {
				tmp = val
			} else {
				tmp = findTops(data, cache, coord)
			}
			for _, k := range tmp {
				exists := false
				for _, v := range result {
					if k.y == v.y && k.x == v.x {
						exists = true
						break
					}
				}
				if !exists {
					result = append(result, k)
				}
			}
		}
	}
	(*cache)[pos] = result
	return result
}

func Challenge1(data []string) (int, error) {
	trailHeads := findTrailheads(data)
	var cache map[point][]point = make(map[point][]point)
	result := 0
	for _, th := range trailHeads {
		paths := findTops(data, &cache, th)
		result += len(paths)
	}
	return result, nil
}

func findPaths(data []string, pos point) int {
	var result int
	if data[pos.y][pos.x] == '9' {
		return 1
	}

	if pos.y > 0 {
		coord := point{pos.y - 1, pos.x}
		if data[coord.y][coord.x] == data[pos.y][pos.x]+1 {
			result += findPaths(data, coord)
		}
	}
	if pos.y < len(data)-1 {
		coord := point{pos.y + 1, pos.x}
		if data[coord.y][coord.x] == data[pos.y][pos.x]+1 {
			result += findPaths(data, coord)
		}
	}
	if pos.x > 0 {
		coord := point{pos.y, pos.x - 1}
		if data[coord.y][coord.x] == data[pos.y][pos.x]+1 {
			result += findPaths(data, coord)
		}
	}
	if pos.x < len(data[0])-1 {
		coord := point{pos.y, pos.x + 1}
		if data[coord.y][coord.x] == data[pos.y][pos.x]+1 {
			result += findPaths(data, coord)
		}
	}
	return result
}

func Challenge2(data []string) (int, error) {
	trailHeads := findTrailheads(data)
	result := 0
	for _, th := range trailHeads {
		result += findPaths(data, th)
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 10 - Hoof It")
	path := "hoofit/input.txt"
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
