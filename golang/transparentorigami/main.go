package transparentorigami

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Fold struct {
	Axis string
	Pos  int
}

func readLines(path string) ([]Point, []Fold, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var points []Point
	var folds []Fold

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		data := strings.Split(line, ",")
		x, err := strconv.Atoi(data[0])
		if err != nil {
			return points, folds, err
		}

		y, err := strconv.Atoi(data[1])
		if err != nil {
			return points, folds, err
		}

		points = append(points, Point{X: x, Y: y})
	}

	r, err := regexp.Compile(`^fold along ([xy])=(\d+)$`)
	if err != nil {
		return points, folds, err
	}

	for scanner.Scan() {
		line := scanner.Text()
		data := r.FindStringSubmatch(line)
		point, err := strconv.Atoi(data[2])
		if err != nil {
			return points, folds, err
		}
		folds = append(folds, Fold{Axis: data[1], Pos: point})
	}

	return points, folds, nil
}

func Challenge1(points []Point, folds []Fold) (int, error) {
	var matrix map[Point]int = make(map[Point]int)
	for _, cur := range points {
		matrix[cur]++
	}
	for _, cur := range folds {
		if cur.Axis == "x" {
			for k, _ := range matrix {
				if k.X > cur.Pos {
					diff := k.X - cur.Pos
					matrix[Point{X: cur.Pos - diff, Y: k.Y}]++
					delete(matrix, k)
				}
			}
		} else {
			for k, _ := range matrix {
				if k.Y > cur.Pos {
					diff := k.Y - cur.Pos
					matrix[Point{X: k.X, Y: cur.Pos - diff}]++
					delete(matrix, k)
				}
			}
		}
		break
	}
	return len(matrix), nil
}

func Challenge2(points []Point, folds []Fold) {
	var matrix map[Point]int = make(map[Point]int)
	for _, cur := range points {
		matrix[cur]++
	}
	for _, cur := range folds {
		if cur.Axis == "x" {
			for k, _ := range matrix {
				if k.X > cur.Pos {
					diff := k.X - cur.Pos
					matrix[Point{X: cur.Pos - diff, Y: k.Y}]++
					delete(matrix, k)
				}
			}
		} else {
			for k, _ := range matrix {
				if k.Y > cur.Pos {
					diff := k.Y - cur.Pos
					matrix[Point{X: k.X, Y: cur.Pos - diff}]++
					delete(matrix, k)
				}
			}
		}
	}
	lastX, lastY := math.MaxInt32, math.MaxInt32
	for _, cur := range folds {
		if cur.Axis == "x" {
			if cur.Pos < lastX {
				lastX = cur.Pos
			}
		} else {
			if cur.Pos < lastY {
				lastY = cur.Pos
			}
		}
	}
	for i := 0; i < lastY; i++ {
		for j := 0; j < lastX; j++ {
			if _, ok := matrix[Point{X: j, Y: i}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func Run() {
	fmt.Println("Day 13 - Transparent Origami")
	path := "transparentorigami/input.txt"
	points, folds, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(points, folds)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	fmt.Println("Challenge 2:")
	Challenge2(points, folds)
}
