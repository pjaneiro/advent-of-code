package resonantcollinearity

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

func Challenge1(data []string) (int, error) {
	var antennas map[byte][]point = make(map[byte][]point)
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == '.' {
				continue
			}
			if _, ok := antennas[data[y][x]]; !ok {
				antennas[data[y][x]] = make([]point, 0)
			}
			antennas[data[y][x]] = append(antennas[data[y][x]], point{y, x})
		}
	}
	var antinodes map[point]struct{} = make(map[point]struct{})
	for _, list := range antennas {
		for i := 0; i < len(list); i++ {
			for j := i + 1; j < len(list); j++ {
				dy, dx := list[j].y-list[i].y, list[j].x-list[i].x
				antinodeA := point{list[i].y - dy, list[i].x - dx}
				if !(antinodeA.y < 0 || antinodeA.x < 0 || antinodeA.y >= len(data) || antinodeA.x >= len(data[0])) {
					antinodes[antinodeA] = struct{}{}
				}
				antinodeB := point{list[j].y + dy, list[j].x + dx}
				if !(antinodeB.y < 0 || antinodeB.x < 0 || antinodeB.y >= len(data) || antinodeB.x >= len(data[0])) {
					antinodes[antinodeB] = struct{}{}
				}
			}
		}
	}
	return len(antinodes), nil
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func Challenge2(data []string) (int, error) {
	var antennas map[byte][]point = make(map[byte][]point)
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == '.' {
				continue
			}
			if _, ok := antennas[data[y][x]]; !ok {
				antennas[data[y][x]] = make([]point, 0)
			}
			antennas[data[y][x]] = append(antennas[data[y][x]], point{y, x})
		}
	}
	var antinodes map[point]struct{} = make(map[point]struct{})
	for _, list := range antennas {
		if len(list) == 1 {
			continue
		}
		for i := 0; i < len(list); i++ {
			antinodes[list[i]] = struct{}{}
			for j := i + 1; j < len(list); j++ {
				dy, dx := list[j].y-list[i].y, list[j].x-list[i].x
				factor := gcd(dy, dx)
				dy, dx = dy/factor, dx/factor
				var antinode point = list[i]
				for {
					antinode = point{antinode.y - dy, antinode.x - dx}
					if antinode.y < 0 || antinode.x < 0 || antinode.y >= len(data) || antinode.x >= len(data[0]) {
						break
					}
					antinodes[antinode] = struct{}{}
				}
				for {
					antinode = point{antinode.y + dy, antinode.x + dx}
					if antinode.y < 0 || antinode.x < 0 || antinode.y >= len(data) || antinode.x >= len(data[0]) {
						break
					}
					antinodes[antinode] = struct{}{}
				}
			}
		}
	}
	return len(antinodes), nil
}

func Run() {
	fmt.Println("Day 8 - Resonant Collinearity")
	path := "resonantcollinearity/input.txt"
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
