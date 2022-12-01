package chronalcoordinates

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Coordinate struct {
	X, Y int
}

func readLines(path string) ([]Coordinate, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r, err := regexp.Compile(`^(\d+), (\d+)$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}

	var lines []Coordinate
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := r.FindStringSubmatch(line)
		var coord Coordinate
		if coord.X, err = strconv.Atoi(data[1]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if coord.Y, err = strconv.Atoi(data[2]); err != nil {
			return nil, errors.New("couldn't parse input")
		}

		lines = append(lines, coord)
	}
	return lines, nil
}

func boundaries(data []Coordinate) (int, int, int, int) {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, 0, 0

	for _, v := range data {
		if v.X < minX {
			minX = v.X
		}
		if v.X > maxX {
			maxX = v.X
		}
		if v.Y < minY {
			minY = v.Y
		}
		if v.Y > maxY {
			maxY = v.Y
		}
	}

	return minX, minY, maxX, maxY
}

func Challenge1(data []Coordinate) (int, error) {
	minX, minY, maxX, maxY := boundaries(data)

	var infinite map[Coordinate]struct{} = make(map[Coordinate]struct{})
	var distances map[Coordinate]Coordinate = make(map[Coordinate]Coordinate)

	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			min, count := math.MaxInt, 0
			coords := []Coordinate{}
			for _, coord := range data {
				dist := int(math.Abs(float64(coord.X)-float64(x))) + int(math.Abs(float64(coord.Y)-float64(y)))
				if dist < min {
					min = dist
					count = 1
					coords = []Coordinate{coord}
				} else if dist == min {
					count++
					coords = append(coords, coord)
				}
			}
			if count == 1 {
				distances[Coordinate{X: x, Y: y}] = coords[0]
				if x == minX || x == maxX || y == minY || y == maxY {
					infinite[coords[0]] = struct{}{}
				}
			}
		}
	}

	counts := make(map[Coordinate]int)
	max := 0
	for _, value := range distances {
		if _, ok := infinite[value]; ok {
			continue
		}
		counts[value]++
		if counts[value] > max {
			max = counts[value]
		}
	}

	return max, nil
}

func Challenge2(data []Coordinate, reach int) (int, error) {
	minX, minY, maxX, maxY := boundaries(data)

	var distances map[Coordinate]struct{} = make(map[Coordinate]struct{})

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			total := 0
			for _, coord := range data {
				total += int(math.Abs(float64(coord.X)-float64(x))) + int(math.Abs(float64(coord.Y)-float64(y)))
			}
			if total < reach {
				distances[Coordinate{X: x, Y: y}] = struct{}{}
			}
		}
	}

	return len(distances), nil
}

func Run() {
	fmt.Println("Day 6 - Chronal Coordinates")
	path := "chronalcoordinates/input.txt"
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

	result, err = Challenge2(data, 10000)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
