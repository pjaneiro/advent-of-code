package crossedwires

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Dir rune
	Val int
}

type pos struct {
	visited1 bool
	visited2 bool
	steps1   int
	steps2   int
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func readLines(path string) ([][]Instruction, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var block []Instruction
		for _, v := range strings.Split(line, ",") {
			var ins Instruction
			for _, r := range v {
				ins.Dir = r
				break
			}
			n, err := strconv.Atoi(v[1:len(v)])
			if err != nil {
				return lines, errors.New("something went wrong scanning int")
			}
			ins.Val = n
			block = append(block, ins)
		}
		lines = append(lines, block)
	}

	return lines, nil
}

func Challenge1(data [][]Instruction) (int, error) {
	var panel map[int]map[int]int = make(map[int]map[int]int)
	panel[0] = make(map[int]int)
	i, j := 0, 0
	for _, v := range data[0] {
		switch v.Dir {
		case 'R':
			for k := 1; k <= v.Val; k++ {
				j++
				panel[i][j] = 1
			}
			break
		case 'L':
			for k := 1; k <= v.Val; k++ {
				j--
				panel[i][j] = 1
			}
			break
		case 'U':
			for k := 1; k <= v.Val; k++ {
				i--
				if _, ok := panel[i]; !ok {
					panel[i] = make(map[int]int)
				}
				panel[i][j] = 1
			}
			break
		case 'D':
			for k := 1; k <= v.Val; k++ {
				i++
				if _, ok := panel[i]; !ok {
					panel[i] = make(map[int]int)
				}
				panel[i][j] = 1
			}
			break
		default:
			break
		}
	}
	min := math.MaxInt32
	i, j = 0, 0
	for _, v := range data[1] {
		switch v.Dir {
		case 'R':
			for k := 1; k <= v.Val; k++ {
				j++
				if panel[i][j] == 1 {
					dist := abs(i) + abs(j)
					if dist < min {
						min = dist
					}
				}
			}
			break
		case 'L':
			for k := 1; k <= v.Val; k++ {
				j--
				if panel[i][j] == 1 {
					dist := abs(i) + abs(j)
					if dist < min {
						min = dist
					}
				}
			}
			break
		case 'U':
			for k := 1; k <= v.Val; k++ {
				i--
				if _, ok := panel[i]; !ok {
					panel[i] = make(map[int]int)
				}
				if panel[i][j] == 1 {
					dist := abs(i) + abs(j)
					if dist < min {
						min = dist
					}
				}
			}
			break
		case 'D':
			for k := 1; k <= v.Val; k++ {
				i++
				if _, ok := panel[i]; !ok {
					panel[i] = make(map[int]int)
				}
				if panel[i][j] == 1 {
					dist := abs(i) + abs(j)
					if dist < min {
						min = dist
					}
				}
			}
			break
		default:
			break
		}
	}
	return min, nil
}

func Challenge2(data [][]Instruction) (int, error) {
	var panel map[int]map[int]*pos = make(map[int]map[int]*pos)
	panel[0] = make(map[int]*pos)
	i, j, steps := 0, 0, 0
	for _, v := range data[0] {
		switch v.Dir {
		case 'R':
			for k := 1; k <= v.Val; k++ {
				j++
				steps++
				if _, ok := panel[i][j]; !ok {
					panel[i][j] = &pos{}
				}
				panel[i][j].visited1 = true
				if panel[i][j].steps1 == 0 {
					panel[i][j].steps1 = steps
				}
			}
			break
		case 'L':
			for k := 1; k <= v.Val; k++ {
				j--
				steps++
				if _, ok := panel[i][j]; !ok {
					panel[i][j] = &pos{}
				}
				panel[i][j].visited1 = true
				if panel[i][j].steps1 == 0 {
					panel[i][j].steps1 = steps
				}
			}
			break
		case 'U':
			for k := 1; k <= v.Val; k++ {
				i--
				if _, ok := panel[i]; !ok {
					panel[i] = make(map[int]*pos)
				}
				steps++
				if _, ok := panel[i][j]; !ok {
					panel[i][j] = &pos{}
				}
				panel[i][j].visited1 = true
				if panel[i][j].steps1 == 0 {
					panel[i][j].steps1 = steps
				}
			}
			break
		case 'D':
			for k := 1; k <= v.Val; k++ {
				i++
				if _, ok := panel[i]; !ok {
					panel[i] = make(map[int]*pos)
				}
				steps++
				if _, ok := panel[i][j]; !ok {
					panel[i][j] = &pos{}
				}
				panel[i][j].visited1 = true
				if panel[i][j].steps1 == 0 {
					panel[i][j].steps1 = steps
				}
			}
			break
		default:
			break
		}
	}
	min := math.MaxInt32
	i, j, steps = 0, 0, 0
	for _, v := range data[1] {
		switch v.Dir {
		case 'R':
			for k := 1; k <= v.Val; k++ {
				j++
				steps++
				if _, ok := panel[i][j]; !ok {
					panel[i][j] = &pos{}
				}
				panel[i][j].visited2 = true
				if panel[i][j].steps2 == 0 {
					panel[i][j].steps2 = steps
				}
				if panel[i][j].visited1 {
					dist := panel[i][j].steps1 + panel[i][j].steps2
					if dist < min {
						min = dist
					}
				}
			}
			break
		case 'L':
			for k := 1; k <= v.Val; k++ {
				j--
				steps++
				if _, ok := panel[i][j]; !ok {
					panel[i][j] = &pos{}
				}
				panel[i][j].visited2 = true
				if panel[i][j].steps2 == 0 {
					panel[i][j].steps2 = steps
				}
				if panel[i][j].visited1 {
					dist := panel[i][j].steps1 + panel[i][j].steps2
					if dist < min {
						min = dist
					}
				}
			}
			break
		case 'U':
			for k := 1; k <= v.Val; k++ {
				i--
				steps++
				if _, ok := panel[i]; !ok {
					panel[i] = make(map[int]*pos)
				}
				if _, ok := panel[i][j]; !ok {
					panel[i][j] = &pos{}
				}
				panel[i][j].visited2 = true
				if panel[i][j].steps2 == 0 {
					panel[i][j].steps2 = steps
				}
				if _, ok := panel[i]; !ok {
					panel[i] = make(map[int]*pos)
				}
				if panel[i][j].visited1 {
					dist := panel[i][j].steps1 + panel[i][j].steps2
					if dist < min {
						min = dist
					}
				}
			}
			break
		case 'D':
			for k := 1; k <= v.Val; k++ {
				i++
				steps++
				if _, ok := panel[i]; !ok {
					panel[i] = make(map[int]*pos)
				}
				if _, ok := panel[i][j]; !ok {
					panel[i][j] = &pos{}
				}
				panel[i][j].visited2 = true
				if panel[i][j].steps2 == 0 {
					panel[i][j].steps2 = steps
				}
				if _, ok := panel[i]; !ok {
					panel[i] = make(map[int]*pos)
				}
				if panel[i][j].visited1 {
					dist := panel[i][j].steps1 + panel[i][j].steps2
					if dist < min {
						min = dist
					}
				}
			}
			break
		default:
			break
		}
	}
	return min, nil
}

func Run() {
	fmt.Println("Day 3 - Crossed Wires")
	path := "crossedwires/input.txt"
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
