package rainrisk

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Instruction struct {
	Dir rune
	Val int
}

func readLines(path string) ([]Instruction, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var cur Instruction = Instruction{}
		cur.Dir = []rune(line)[0]
		cur.Val, err = strconv.Atoi(line[1:len(line)])
		if err != nil {
			return nil, err
		}
		lines = append(lines, cur)
	}
	return lines, scanner.Err()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Challenge1(data []Instruction) (int, error) {
	posX, posY, ang := 0, 0, 0
	for _, inst := range data {
		switch inst.Dir {
		case 'N':
			posY += inst.Val
		case 'S':
			posY -= inst.Val
		case 'E':
			posX += inst.Val
		case 'W':
			posX -= inst.Val
		case 'L':
			ang += inst.Val
		case 'R':
			ang -= inst.Val
		case 'F':
			for ang < 0 {
				ang += 360
			}
			switch ang % 360 {
			case 0:
				posX += inst.Val
			case 90:
				posY += inst.Val
			case 180:
				posX -= inst.Val
			case 270:
				posY -= inst.Val
			default:
				fmt.Println(ang)
				return 0, errors.New("invalid angle")
			}
		default:
			return 0, errors.New("invalid instruction")
		}
	}
	return abs(posX) + abs(posY), nil
}

func Challenge2(data []Instruction) (int, error) {
	posX, posY := 0, 0
	wayX, wayY := 10, 1
	for _, inst := range data {
		switch inst.Dir {
		case 'N':
			wayY += inst.Val
		case 'S':
			wayY -= inst.Val
		case 'E':
			wayX += inst.Val
		case 'W':
			wayX -= inst.Val
		case 'L':
			ang := inst.Val
			for ang < 0 {
				ang += 360
			}
			ang %= 360
			distX, distY := wayX-posX, wayY-posY
			switch ang {
			case 90:
				wayX = posX - distY
				wayY = posY + distX
			case 180:
				wayX = posX - distX
				wayY = posY - distY
			case 270:
				wayX = posX + distY
				wayY = posY - distX
			}
		case 'R':
			ang := inst.Val
			for ang < 0 {
				ang += 360
			}
			ang %= 360
			distX, distY := wayX-posX, wayY-posY
			switch ang {
			case 90:
				wayX = posX + distY
				wayY = posY - distX
			case 180:
				wayX = posX - distX
				wayY = posY - distY
			case 270:
				wayX = posX - distY
				wayY = posY + distX
			}
		case 'F':
			distX, distY := wayX-posX, wayY-posY
			posX = posX + (inst.Val * distX)
			posY = posY + (inst.Val * distY)
			wayX = posX + distX
			wayY = posY + distY
		default:
			return 0, errors.New("invalid instruction")
		}
	}
	return abs(posX) + abs(posY), nil
}

func Run() {
	fmt.Println("Day 12")
	path := "rainrisk/input.txt"
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
