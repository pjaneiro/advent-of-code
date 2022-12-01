package trenchmap

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func readLines(path string) (string, []string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	algo := scanner.Text()

	var data []string = make([]string, 0)
	scanner.Scan()
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return algo, data, nil
}

func parseMap(data []string) map[Coord]bool {
	var result map[Coord]bool = make(map[Coord]bool)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == '#' {
				result[Coord{x: j, y: i}] = true
			} else {
				result[Coord{x: j, y: i}] = false
			}
		}
	}
	return result
}

func calculateVal(algo string, key []string) bool {
	parsedKey := strings.Join(key, "")
	var actualKey int
	if i, err := strconv.ParseInt(parsedKey, 2, 32); err != nil {
		return false
	} else {
		actualKey = int(i)
	}
	if algo[actualKey] == '#' {
		return true
	} else {
		return false
	}
}

func getBoundaries(data map[Coord]bool) (int, int, int, int) {
	minX, maxX, minY, maxY := math.MaxInt32, math.MinInt32, math.MaxInt32, math.MinInt32
	for k, _ := range data {
		if k.x < minX {
			minX = k.x
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.y < minY {
			minY = k.y
		}
		if k.y > maxY {
			maxY = k.y
		}
	}
	return minX, maxX, minY, maxY
}

func solve(algo string, data []string, rounds int) map[Coord]bool {
	parsedMap := parseMap(data)

	for round := 0; round < rounds; round++ {
		minX, maxX, minY, maxY := getBoundaries(parsedMap)
		var newMap map[Coord]bool = make(map[Coord]bool)

		for y := minY - 1; y <= maxY+1; y++ {
			for x := minX - 1; x <= maxX+1; x++ {
				curCoord := Coord{x: x, y: y}
				var key []string = make([]string, 9)
				counter := 0
				for dy := -1; dy < 2; dy++ {
					for dx := -1; dx < 2; dx++ {
						if v, ok := parsedMap[Coord{x: x + dx, y: y + dy}]; ok {
							if v {
								key[counter] = "1"
							} else {
								key[counter] = "0"
							}
						} else {
							if round%2 == 1 {
								key[counter] = "1"
							} else {
								key[counter] = "0"
							}
						}
						counter++
					}
				}
				if calculateVal(algo, key) {
					newMap[curCoord] = true
				} else {
					newMap[curCoord] = false
				}
			}
		}

		parsedMap = newMap
	}
	return parsedMap
}

func Challenge1(algo string, data []string) (int, error) {
	solved := solve(algo, data, 2)
	result := 0
	for _, v := range solved {
		if v {
			result++
		}
	}
	return result, nil
}

func Challenge2(algo string, data []string) (int, error) {
	solved := solve(algo, data, 50)
	result := 0
	for _, v := range solved {
		if v {
			result++
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 20 - Trench Map")
	path := "trenchmap/input.txt"
	algo, data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(algo, data)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(algo, data)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
