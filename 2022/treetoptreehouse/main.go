package treetoptreehouse

import (
	"bufio"
	"fmt"
	"os"
)

type Tree struct {
	Height  int
	Visible bool
}

func readLines(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result [][]int
	for scanner.Scan() {
		data := scanner.Text()
		var line []int
		for _, val := range data {
			line = append(line, int(val)-48)
		}
		result = append(result, line)
	}

	return result, nil
}

func intsToTrees(data [][]int) [][]Tree {
	var result [][]Tree

	for _, line := range data {
		var treeLine []Tree
		for _, tree := range line {
			treeLine = append(treeLine, Tree{Height: tree, Visible: false})
		}
		result = append(result, treeLine)
	}

	return result
}

func Challenge1(data [][]int) (int, error) {
	trees := intsToTrees(data)

	nLines, nCols := len(data), len(data[0])

	for y := 0; y < nLines; y++ {
		trees[y][0].Visible = true
		tallest := trees[y][0].Height
		for x := 1; x < nCols; x++ {
			if trees[y][x].Height > tallest {
				trees[y][x].Visible = true
				tallest = trees[y][x].Height
			}
		}

		trees[y][nCols-1].Visible = true
		tallest = trees[y][nCols-1].Height
		for x := nCols - 2; x >= 0; x-- {
			if trees[y][x].Height > tallest {
				trees[y][x].Visible = true
				tallest = trees[y][x].Height
			}
		}
	}

	for x := 0; x < nCols; x++ {
		trees[0][x].Visible = true
		tallest := trees[0][x].Height
		for y := 1; y < nLines; y++ {
			if trees[y][x].Height > tallest {
				trees[y][x].Visible = true
				tallest = trees[y][x].Height
			}
		}

		trees[nLines-1][x].Visible = true
		tallest = trees[nLines-1][x].Height
		for y := nLines - 2; y >= 0; y-- {
			if trees[y][x].Height > tallest {
				trees[y][x].Visible = true
				tallest = trees[y][x].Height
			}
		}
	}

	result := 0
	for _, line := range trees {
		for _, tree := range line {
			if tree.Visible {
				result++
			}
		}
	}

	return result, nil
}

func Challenge2(data [][]int) (int, error) {
	trees := intsToTrees(data)

	nLines, nCols := len(data), len(data[0])

	result := 0

	for y := 0; y < nLines; y++ {
		for x := 0; x < nCols; x++ {
			curTree := trees[y][x]
			height := curTree.Height
			up, down, left, right := 0, 0, 0, 0

			for x2 := x - 1; x2 >= 0; x2-- {
				left++
				if trees[y][x2].Height >= height {
					break
				}
			}
			for x2 := x + 1; x2 < nCols; x2++ {
				right++
				if trees[y][x2].Height >= height {
					break
				}
			}
			for y2 := y - 1; y2 >= 0; y2-- {
				up++
				if trees[y2][x].Height >= height {
					break
				}
			}
			for y2 := y + 1; y2 < nLines; y2++ {
				down++
				if trees[y2][x].Height >= height {
					break
				}
			}

			if up*down*left*right > result {
				result = up * down * left * right
			}
		}
	}

	return result, nil
}

func Run() {
	fmt.Println("Day 8 - Treetop Tree House")
	path := "treetoptreehouse/input.txt"
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
