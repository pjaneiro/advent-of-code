package hydrothermalventure

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	X int
	Y int
}

type Vector struct {
	From Point
	To   Point
}

func readLines(path string) ([]Vector, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Vector
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r, err := regexp.Compile(`^(\d+),(\d+) -> (\d+),(\d+)$`)
		if err != nil {
			return lines, errors.New("couldn't parse regex expression")
		}

		var elem Vector
		var data []string = r.FindStringSubmatch(scanner.Text())
		if data == nil {
			return lines, nil
		}

		elem.From.X, err = strconv.Atoi(data[1])
		elem.From.Y, err = strconv.Atoi(data[2])
		elem.To.X, err = strconv.Atoi(data[3])
		elem.To.Y, err = strconv.Atoi(data[4])

		lines = append(lines, elem)
	}

	return lines, nil
}

func createBoard(x int, y int) [][]int {
	board := make([][]int, x)
	for i, _ := range board {
		board[i] = make([]int, y)
	}
	return board
}

func maxes(data []Vector) (int, int) {
	maxX, maxY := 0, 0
	for _, cur := range data {
		if cur.From.X > maxX {
			maxX = cur.From.X
		}
		if cur.To.X > maxX {
			maxX = cur.To.X
		}
		if cur.From.Y > maxY {
			maxY = cur.From.Y
		}
		if cur.To.Y > maxY {
			maxY = cur.To.Y
		}
	}
	return maxX, maxY
}

func markPoints(board [][]int, fromX int, fromY int, toX int, toY int) {
	if fromX == toX {
		if fromY < toY {
			for j := fromY; j <= toY; j++ {
				board[fromX][j]++
			}
		} else {
			for j := fromY; j >= toY; j-- {
				board[fromX][j]++
			}
		}
	} else if fromY == toY {
		if fromX < toX {
			for i := fromX; i <= toX; i++ {
				board[i][fromY]++
			}
		} else {
			for i := fromX; i >= toX; i-- {
				board[i][fromY]++
			}
		}
	} else if fromX < toX {
		steps := toX - fromX
		if fromY < toY {
			for i := 0; i <= steps; i++ {
				board[fromX+i][fromY+i]++
			}
		} else {
			for i := 0; i <= steps; i++ {
				board[fromX+i][fromY-i]++
			}
		}
	} else {
		steps := fromX - toX
		if fromY < toY {
			for i := 0; i <= steps; i++ {
				board[fromX-i][fromY+i]++
			}
		} else {
			for i := 0; i <= steps; i++ {
				board[fromX-i][fromY-i]++
			}
		}
	}
}

func countPoints(board [][]int) int {
	result := 0
	for i, _ := range board {
		for j, _ := range board[i] {
			if board[i][j] >= 2 {
				result++
			}
		}
	}
	return result
}

func Challenge1(data []Vector) (int, error) {
	maxX, maxY := maxes(data)

	board := createBoard(maxX+1, maxY+1)

	for _, cur := range data {
		if cur.From.X == cur.To.X || cur.From.Y == cur.To.Y {
			markPoints(board, cur.From.X, cur.From.Y, cur.To.X, cur.To.Y)
		}
	}

	return countPoints(board), nil
}

func Challenge2(data []Vector) (int, error) {
	maxX, maxY := maxes(data)

	board := createBoard(maxX+1, maxY+1)

	for _, cur := range data {
		markPoints(board, cur.From.X, cur.From.Y, cur.To.X, cur.To.Y)
	}

	return countPoints(board), nil
}

func Run() {
	fmt.Println("Day 5 - Hydrothermal Venture")
	path := "hydrothermalventure/input.txt"
	inputData, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(inputData)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(inputData)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
