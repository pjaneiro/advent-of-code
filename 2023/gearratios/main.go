package gearratios

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	Y, X int
}

type Part struct {
	Number int
	Length int
	Used   bool
}

type Gear struct {
	Parts map[Coord]struct{}
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

func locateNumbers(data []string) map[Coord]Part {
	var numbers map[Coord]Part = make(map[Coord]Part)
	currentNumber, currentLength, y, x := 0, 0, 0, 0
	inNumber := false
	for y = 0; y < len(data); y++ {
		for x = 0; x < len(data[y]); x++ {
			val := data[y][x]
			if val >= 48 && val <= 57 {
				currentNumber = (currentNumber * 10) + int(val-48)
				currentLength++
				inNumber = true
			} else if inNumber {
				numbers[Coord{Y: y, X: x - currentLength}] = Part{Number: currentNumber, Length: currentLength, Used: false}
				inNumber = false
				currentNumber, currentLength = 0, 0
			}
		}
		if inNumber {
			numbers[Coord{Y: y, X: x - currentLength}] = Part{Number: currentNumber, Length: currentLength, Used: false}
			inNumber = false
			currentNumber, currentLength = 0, 0
		}
	}
	if inNumber {
		numbers[Coord{Y: y, X: x - currentLength}] = Part{Number: currentNumber, Length: currentLength, Used: false}
		inNumber = false
		currentNumber, currentLength = 0, 0
	}
	return numbers
}

func calculatePart1(data []string, parts map[Coord]Part) (int, error) {
	result := 0
OUTER:
	for coord, part := range parts {
		iY, iX := coord.Y, coord.X
		for houses := 0; houses < part.Length; houses++ {
			for dY := -1; dY <= 1; dY++ {
				for dX := -1; dX <= 1; dX++ {
					coord := Coord{Y: iY + dY, X: iX + houses + dX}
					if coord.Y >= 0 && coord.Y < len(data) && coord.X >= 0 && coord.X < len(data[coord.Y]) {
						if data[coord.Y][coord.X] != '.' && (data[coord.Y][coord.X] < 48 || data[coord.Y][coord.X] > 57) {
							result += part.Number
							part.Used = true
							continue OUTER
						}
					}
				}
			}
		}
	}
	return result, nil
}

func calculatePart2(data []string, parts map[Coord]Part) (int, error) {
	result := 0
	var gears map[Coord]Gear = make(map[Coord]Gear)
	for coord, part := range parts {
		iY, iX := coord.Y, coord.X
		for houses := 0; houses < part.Length; houses++ {
			for dY := -1; dY <= 1; dY++ {
				for dX := -1; dX <= 1; dX++ {
					coord := Coord{Y: iY + dY, X: iX + houses + dX}
					if coord.Y >= 0 && coord.Y < len(data) && coord.X >= 0 && coord.X < len(data[coord.Y]) {
						if data[coord.Y][coord.X] == '*' {
							if _, ok := gears[coord]; !ok {
								gears[coord] = Gear{Parts: map[Coord]struct{}{}}
							}
							gear := gears[coord]
							gear.Parts[Coord{Y: iY, X: iX}] = struct{}{}
						}
					}
				}
			}
		}
	}
	for _, gear := range gears {
		if len(gear.Parts) == 2 {
			gearRatio := 1
			for partCoord := range gear.Parts {
				gearRatio = gearRatio * parts[partCoord].Number
			}
			result = result + gearRatio
		}
	}
	return result, nil
}

func Challenge1(data []string) (int, error) {
	parts := locateNumbers(data)
	result, err := calculatePart1(data, parts)
	return result, err
}

func Challenge2(data []string) (int, error) {
	parts := locateNumbers(data)
	result, err := calculatePart2(data, parts)
	return result, err
}

func Run() {
	fmt.Println("Day 3 - Gear Ratios")
	path := "gearratios/input.txt"
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
