package playground

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type box struct {
	x, y, z int
}

type pair struct {
	a, b box
	dist float64
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

func parseInput(data []string) ([]box, error) {
	var result []box = make([]box, 0, len(data))
	for _, row := range data {
		params := strings.Split(row, ",")
		x, err := strconv.Atoi(params[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(params[1])
		if err != nil {
			return nil, err
		}
		z, err := strconv.Atoi(params[2])
		if err != nil {
			return nil, err
		}
		result = append(result, box{x, y, z})
	}
	return result, nil
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

func distance(a, b box) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	dz := float64(a.z - b.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func merge(a, b []box) []box {
	var result []box = make([]box, 0, len(a)+len(b))
	var seen map[box]struct{} = make(map[box]struct{})

	for _, bx := range a {
		if _, ok := seen[bx]; !ok {
			seen[bx] = struct{}{}
			result = append(result, bx)
		}
	}
	for _, bx := range b {
		if _, ok := seen[bx]; !ok {
			seen[bx] = struct{}{}
			result = append(result, bx)
		}
	}
	return result
}

func Challenge1(data []string, steps int) (int, error) {
	boxes, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	var byDistance []pair = make([]pair, 0, len(boxes)*len(boxes))
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			byDistance = append(byDistance, pair{boxes[i], boxes[j], distance(boxes[i], boxes[j])})
		}
	}
	slices.SortFunc(byDistance, func(a, b pair) int {
		if a.dist < b.dist {
			return -1
		}
		return 1
	})
	var sets [][]box = make([][]box, 0)
	for _, b := range boxes {
		sets = append(sets, []box{b})
	}
	for i := 0; i < steps; i++ {
		pair := byDistance[i]
		var sa, sb int
		for idx, set := range sets {
			if slices.Contains(set, pair.a) {
				sa = idx
			}
			if slices.Contains(set, pair.b) {
				sb = idx
			}
		}
		if sa == sb {
			continue
		}
		s1, s2 := min(sa, sb), max(sa, sb)
		newSet := merge(sets[s1], sets[s2])
		newSets := append(sets[:s1], sets[s1+1:s2]...)
		newSets = append(newSets, sets[s2+1:]...)
		newSets = append(newSets, newSet)
		sets = newSets
	}
	slices.SortFunc(sets, func(a, b []box) int {
		return len(b) - len(a)
	})
	result := 1
	for i := 0; i < 3; i++ {
		result *= len(sets[i])
	}

	return result, nil
}

func Challenge2(data []string) (int, error) {
	boxes, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	var byDistance []pair = make([]pair, 0, len(boxes)*len(boxes))
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			byDistance = append(byDistance, pair{boxes[i], boxes[j], distance(boxes[i], boxes[j])})
		}
	}
	slices.SortFunc(byDistance, func(a, b pair) int {
		if a.dist < b.dist {
			return -1
		}
		return 1
	})
	var sets [][]box = make([][]box, 0)
	for _, b := range boxes {
		sets = append(sets, []box{b})
	}
	var i int
	for i = 0; len(sets) >= 2; i++ {
		pair := byDistance[i]
		var sa, sb int
		for idx, set := range sets {
			if slices.Contains(set, pair.a) {
				sa = idx
			}
			if slices.Contains(set, pair.b) {
				sb = idx
			}
		}
		if sa == sb {
			continue
		}
		s1, s2 := min(sa, sb), max(sa, sb)
		newSet := merge(sets[s1], sets[s2])
		newSets := append(sets[:s1], sets[s1+1:s2]...)
		newSets = append(newSets, sets[s2+1:]...)
		newSets = append(newSets, newSet)
		sets = newSets
	}
	finisher := byDistance[i-1]
	return finisher.a.x * finisher.b.x, nil
}

func Run() {
	fmt.Println("Day 8 - Playground")
	path := "playground/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(data, 1000)
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
