package thestarsalign

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Star struct {
	X, Y, Vx, Vy int
}

func readLines(path string) ([]Star, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r, err := regexp.Compile(`^position=<\s*(-?\d+), \s*(-?\d+)> velocity=<\s*(-?\d+), \s*(-?\d+)>$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}

	var result []Star = make([]Star, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var star Star
		line := scanner.Text()
		data := r.FindStringSubmatch(line)
		if star.X, err = strconv.Atoi(data[1]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if star.Y, err = strconv.Atoi(data[2]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if star.Vx, err = strconv.Atoi(data[3]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if star.Vy, err = strconv.Atoi(data[4]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		result = append(result, star)
	}
	return result, nil
}

func run(data []Star) ([]string, int) {
	var minX, minY, maxX, maxY int
	var round int
	for round = 1; ; round++ {
		minX, minY, maxX, maxY = math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
		for i := 0; i < len(data); i++ {
			data[i].X += data[i].Vx
			data[i].Y += data[i].Vy
			if data[i].X > maxX {
				maxX = data[i].X
			}
			if data[i].X < minX {
				minX = data[i].X
			}
			if data[i].Y > maxY {
				maxY = data[i].Y
			}
			if data[i].Y < minY {
				minY = data[i].Y
			}
		}

		if round >= 3 && maxY-minY < 10 {
			break
		}
	}

	var mapping [][]rune = make([][]rune, maxY-minY+1)
	for i := 0; i < maxY-minY+1; i++ {
		mapping[i] = make([]rune, maxX-minX+1)
		for j := 0; j < maxX-minX+1; j++ {
			mapping[i][j] = '.'
		}
	}
	for _, star := range data {
		mapping[star.Y-minY][star.X-minX] = '#'
	}

	var result []string = make([]string, maxY-minY+1)
	for i := 0; i < maxY-minY+1; i++ {
		result[i] = string(mapping[i])
	}

	return result, round
}

func Challenge1(data []Star) ([]string, error) {
	result, _ := run(data)
	return result, nil
}

func Challenge2(data []Star) (int, error) {
	_, result := run(data)
	return result, nil
}

func Run() {
	fmt.Println("Day 10 - The Stars Align")
	path := "thestarsalign/input.txt"

	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result1 []string
	result1, err = Challenge1(data)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Println("Challenge 1:")
		for _, line := range result1 {
			fmt.Println(line)
		}
	}

	data, err = readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result2 int
	result2, err = Challenge2(data)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result2)
	}
}
