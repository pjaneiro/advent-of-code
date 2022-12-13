package chronalcharge

import (
	"fmt"
)

func getMap(serialNumber int) [][]int {
	var result [][]int = make([][]int, 300)
	for i := 0; i < 300; i++ {
		result[i] = make([]int, 300)
	}

	for y := 0; y < 300; y++ {
		for x := 0; x < 300; x++ {
			rackId := x + 11
			powerLevel := rackId * (y + 1)
			powerLevel += serialNumber
			powerLevel *= rackId
			powerLevel /= 100
			powerLevel %= 10
			powerLevel -= 5

			result[y][x] = powerLevel
		}
	}

	return result
}

func Challenge1(serialNumber int) ([]int, error) {
	data := getMap(serialNumber)
	resultX, resultY, power := 0, 0, 0
	for y := 0; y < len(data)-3; y++ {
		for x := 0; x < len(data[y])-3; x++ {
			p := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					p += data[y+i][x+j]
				}
			}
			if p > power {
				power = p
				resultX, resultY = x+1, y+1
			}
		}
	}
	return []int{resultX, resultY}, nil
}

func Challenge2(serialNumber int) ([]int, error) {
	data := getMap(serialNumber)

	var sums [][]int = make([][]int, len(data))
	for i := 0; i < len(data); i++ {
		sums[i] = make([]int, len(data))
	}

	sums[len(data)-1][len(data)-1] = data[len(data)-1][len(data)-1]

	for y := len(data) - 2; y >= 0; y-- {
		sums[y][0] = data[y][0] + sums[y+1][0]
	}
	for x := len(data) - 2; x >= 0; x-- {
		sums[0][x] = data[0][x] + sums[0][x+1]
	}
	for y := len(data) - 2; y >= 0; y-- {
		for x := len(data) - 2; x >= 0; x-- {
			val := data[y][x] + sums[y+1][x] + sums[y][x+1] - sums[y+1][x+1]
			sums[y][x] = val
		}
	}

	resultX, resultY, resultSize, power := 0, 0, 0, 0
	for size := 1; size <= 300; size++ {
		for y := len(data) - size - 2; y >= 0; y-- {
			for x := len(data) - size - 2; x >= 0; x-- {
				tot := sums[y][x] - sums[y][x+size] - sums[y+size][x] + sums[y+size][x+size]
				if tot > power {
					resultX, resultY = x+1, y+1
					power = tot
					resultSize = size
				}
			}
		}
	}
	return []int{resultX, resultY, resultSize}, nil
}

func Run() {
	fmt.Println("Day 11 - Chronal Charge")
	data := 9435

	var result []int
	var err error
	result, err = Challenge1(data)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %v\n", result)
	}

	result, err = Challenge2(data)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %v\n", result)
	}
}
