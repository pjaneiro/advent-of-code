package spacepolice

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

const OUTPUT = false
const DEBUG = false

const PARAMMODE_POSITION = 0
const PARAMMODE_IMMEDIATE = 1
const PARAMMODE_RELATIVE = 2

const OPCODE_ADD int64 = 1
const OPCODE_MULT int64 = 2
const OPCODE_WRITE int64 = 3
const OPCODE_PRINT int64 = 4
const OPCODE_JUMP_IF_TRUE int64 = 5
const OPCODE_JUMP_IF_FALSE int64 = 6
const OPCODE_LESS_THAN int64 = 7
const OPCODE_EQUALS int64 = 8
const OPCODE_ADJUST_RB int64 = 9
const OPCODE_HALT int64 = 99

const NORTH int = 0
const WEST int = 1
const SOUTH int = 2
const EAST int = 3

func readLines(path string) (map[int64]int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines map[int64]int64 = make(map[int64]int64)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i, v := range strings.Split(line, ",") {
			n, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return lines, errors.New("something went wrong scanning int")
			}
			lines[int64(i)] = n
		}
	}

	return lines, nil
}

func readParam(data map[int64]int64, parammode int64, ip int64, pos int64, rb int64) int64 {
	if parammode == PARAMMODE_POSITION {
		return data[ip+pos]
	} else if parammode == PARAMMODE_IMMEDIATE {
		return ip + pos
	} else if parammode == PARAMMODE_RELATIVE {
		return rb + data[ip+pos]
	}
	return 0
}

func intCode(data map[int64]int64, input []int64, ip *int64, ii *int64, out *int64, rb *int64) (output int64, done bool, err error) {
	for true {
		opcode, instruction := data[*ip]%100, data[*ip]/100
		parammodes := []int64{}
		for instruction > 0 {
			parammodes = append(parammodes, instruction%10)
			instruction /= 10
		}
		for len(parammodes) < 3 {
			parammodes = append(parammodes, 0)
		}
		switch opcode {
		case OPCODE_ADD:
			var val1, val2, val3 int64 = readParam(data, parammodes[0], *ip, 1, *rb), readParam(data, parammodes[1], *ip, 2, *rb), readParam(data, parammodes[2], *ip, 3, *rb)
			data[val3] = data[val1] + data[val2]
			*ip += 4
		case OPCODE_MULT:
			var val1, val2, val3 int64 = readParam(data, parammodes[0], *ip, 1, *rb), readParam(data, parammodes[1], *ip, 2, *rb), readParam(data, parammodes[2], *ip, 3, *rb)
			data[val3] = data[val1] * data[val2]
			*ip += 4
		case OPCODE_WRITE:
			var val1 int64 = readParam(data, parammodes[0], *ip, 1, *rb)
			data[val1] = input[*ii]
			*ip += 2
			*ii++
		case OPCODE_PRINT:
			*out = data[readParam(data, parammodes[0], *ip, 1, *rb)]
			if OUTPUT {
				fmt.Println(*out)
			}
			*ip += 2
			return *out, false, nil
		case OPCODE_JUMP_IF_TRUE:
			var val1, val2 int64 = readParam(data, parammodes[0], *ip, 1, *rb), readParam(data, parammodes[1], *ip, 2, *rb)
			if data[val1] != 0 {
				*ip = data[val2]
			} else {
				*ip += 3
			}
		case OPCODE_JUMP_IF_FALSE:
			var val1, val2 int64 = readParam(data, parammodes[0], *ip, 1, *rb), readParam(data, parammodes[1], *ip, 2, *rb)
			if data[val1] == 0 {
				*ip = data[val2]
			} else {
				*ip += 3
			}
		case OPCODE_LESS_THAN:
			var val1, val2, val3 int64 = readParam(data, parammodes[0], *ip, 1, *rb), readParam(data, parammodes[1], *ip, 2, *rb), readParam(data, parammodes[2], *ip, 3, *rb)
			if data[val1] < data[val2] {
				data[val3] = 1
			} else {
				data[val3] = 0
			}
			*ip += 4
		case OPCODE_EQUALS:
			var val1, val2, val3 int64 = readParam(data, parammodes[0], *ip, 1, *rb), readParam(data, parammodes[1], *ip, 2, *rb), readParam(data, parammodes[2], *ip, 3, *rb)
			if data[val1] == data[val2] {
				data[val3] = 1
			} else {
				data[val3] = 0
			}
			*ip += 4
		case OPCODE_ADJUST_RB:
			var offset int64 = readParam(data, parammodes[0], *ip, 1, *rb)
			*rb = *rb + data[offset]
			*ip += 2
		case OPCODE_HALT:
			return *out, true, nil
		default:
			if DEBUG {
				fmt.Printf("Tried to use opcode %d\n", opcode)
			}
			return 0, false, errors.New("invalid opcode")
		}
	}
	return 0, false, errors.New("something went wrong")
}

func HullPaintingRobot(data map[int64]int64, input []int64) (int64, error) {
	var ip, ii, out, rb int64 = 0, 0, 0, 0
	var state map[Point]int64 = make(map[Point]int64)
	curPos := Point{X: 0, Y: 0}
	curDir := NORTH

	minX, maxX, minY, maxY := math.MaxInt32, math.MinInt32, math.MaxInt32, math.MinInt32

	for true {
		newColor, done, _ := intCode(data, input, &ip, &ii, &out, &rb)
		if done {
			break
		}
		nextMove, done, _ := intCode(data, input, &ip, &ii, &out, &rb)
		if done {
			break
		}
		state[curPos] = newColor
		if nextMove == 0 {
			curDir = (curDir + 1) % 4
		} else if nextMove == 1 {
			curDir = (curDir + 3) % 4
		}
		switch curDir {
		case NORTH:
			curPos = Point{X: curPos.X, Y: curPos.Y - 1}
		case WEST:
			curPos = Point{X: curPos.X - 1, Y: curPos.Y}
		case SOUTH:
			curPos = Point{X: curPos.X, Y: curPos.Y + 1}
		case EAST:
			curPos = Point{X: curPos.X + 1, Y: curPos.Y}
		}
		input = append(input, state[curPos])
		if curPos.X < minX {
			minX = curPos.X
		}
		if curPos.X > maxX {
			maxX = curPos.X
		}
		if curPos.Y < minY {
			minY = curPos.Y
		}
		if curPos.Y > maxY {
			maxY = curPos.Y
		}
	}

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			if state[Point{X: j, Y: i}] == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}

	return int64(len(state)), nil
}

func Challenge1(data map[int64]int64) (int64, error) {
	return HullPaintingRobot(data, []int64{0})
}

func Challenge2(data map[int64]int64) (int64, error) {
	return HullPaintingRobot(data, []int64{1})
}

func Run() {
	fmt.Println("Day 11 - Space Police")
	path := "spacepolice/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int64

	var cpy1 map[int64]int64 = make(map[int64]int64)
	for k, v := range data {
		cpy1[k] = v
	}
	result, err = Challenge1(cpy1)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	var cpy2 map[int64]int64 = make(map[int64]int64)
	for k, v := range data {
		cpy2[k] = v
	}
	result, err = Challenge2(cpy2)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
