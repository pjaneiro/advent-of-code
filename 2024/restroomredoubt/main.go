package restroomredoubt

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type robot struct {
	px, py, vx, vy int
}

type coord struct {
	x, y int
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

func parseInput(data []string) ([]robot, error) {
	var result []robot = make([]robot, 0)
	r, err := regexp.Compile(`^p=(\-*\d+),(\-*\d+) v=(\-*\d+),(\-*\d+)$`)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(data); i++ {
		var cur robot
		var err error
		found := r.FindStringSubmatch(data[i])
		if cur.px, err = strconv.Atoi(found[1]); err != nil {
			return nil, err
		}
		if cur.py, err = strconv.Atoi(found[2]); err != nil {
			return nil, err
		}
		if cur.vx, err = strconv.Atoi(found[3]); err != nil {
			return nil, err
		}
		if cur.vy, err = strconv.Atoi(found[4]); err != nil {
			return nil, err
		}
		result = append(result, cur)
	}
	return result, nil
}

func Challenge1(data []string, width, height int) (int, error) {
	robots, err := parseInput(data)
	if err != nil {
		return 0, nil
	}
	nw, ne, sw, se := 0, 0, 0, 0
	tx, ty := width/2, height/2
	for _, r := range robots {
		nx, ny := ((r.px+100*r.vx)%width+width)%width, ((r.py+100*r.vy)%height+height)%height
		if nx < tx && ny < ty {
			nw++
		}
		if nx > tx && ny < ty {
			ne++
		}
		if nx < tx && ny > ty {
			sw++
		}
		if nx > tx && ny > ty {
			se++
		}
	}
	return nw * ne * sw * se, nil
}

func entropy(data []int, width, height int) float64 {
	bins, occupied := 256, 0
	histogram := make([]float64, bins)
	for _, value := range data {
		histogram[value]++
		occupied++
	}
	histogram[0] = float64(width*height - occupied)

	totalPixels := float64(width * height)
	for i := range histogram {
		histogram[i] /= totalPixels
	}

	entropy := 0.0
	for _, p := range histogram {
		if p > 0 {
			entropy -= p * math.Log2(p)
		}
	}

	return entropy
}

func Challenge2(data []string, width, height int) (int, error) {
	robots, err := parseInput(data)
	if err != nil {
		return 0, nil
	}

	result, minEntropy := -1, math.MaxFloat64

	for i := 0; i < width*height; i++ {
		var flatMap []int = make([]int, width*height)
		for _, r := range robots {
			flatMap[r.py*width+r.px]++
		}
		e := entropy(flatMap, width, height)
		if e < minEntropy {
			result, minEntropy = i, e
		}
		for j := 0; j < len(robots); j++ {
			r := robots[j]
			robots[j].px, robots[j].py = ((r.px+r.vx)%width+width)%width, ((r.py+r.vy)%height+height)%height
		}
	}

	return result, nil
}

func Run() {
	fmt.Println("Day 14 - Restroom Redoubt")
	path := "restroomredoubt/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(data, 101, 103)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(data, 101, 103)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
