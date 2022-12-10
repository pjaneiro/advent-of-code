package ropebridge

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	X, Y int
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func solve(data []string, nodes int) (int, error) {
	var visited map[Point]struct{} = make(map[Point]struct{})

	var knots []Point
	for i := 0; i < nodes; i++ {
		knots = append(knots, Point{X: 0, Y: 0})
	}
	visited[knots[nodes-1]] = struct{}{}

	for _, line := range data {
		direction := line[0]
		value, err := strconv.Atoi(line[2:])
		if err != nil {
			return 0, errors.New("something went wrong")
		}

		for i := 0; i < value; i++ {
			switch direction {
			case 'U':
				knots[0].Y++
			case 'D':
				knots[0].Y--
			case 'L':
				knots[0].X--
			case 'R':
				knots[0].X++
			default:
				return 0, errors.New("invalid direction")
			}

			for j := 0; j < nodes-1; j++ {
				if knots[j+1].X == knots[j].X {
					if knots[j+1].Y-knots[j].Y > 1 {
						knots[j+1].Y--
					} else if knots[j+1].Y-knots[j].Y < -1 {
						knots[j+1].Y++
					}
				} else if knots[j+1].Y == knots[j].Y {
					if knots[j+1].X-knots[j].X > 1 {
						knots[j+1].X--
					} else if knots[j+1].X-knots[j].X < -1 {
						knots[j+1].X++
					}
				} else {
					if abs(knots[j+1].X-knots[j].X)+abs(knots[j+1].Y-knots[j].Y) <= 2 {
						continue
					}
					if knots[j+1].X > knots[j].X {
						knots[j+1].X--
					} else {
						knots[j+1].X++
					}
					if knots[j+1].Y > knots[j].Y {
						knots[j+1].Y--
					} else {
						knots[j+1].Y++
					}
				}
			}

			visited[knots[nodes-1]] = struct{}{}
		}
	}

	return len(visited), nil
}

func Challenge1(data []string) (int, error) {
	return solve(data, 2)
}

func Challenge2(data []string) (int, error) {
	return solve(data, 10)
}

func Run() {
	fmt.Println("Day 9 - Rope Bridge")
	path := "ropebridge/input.txt"
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
