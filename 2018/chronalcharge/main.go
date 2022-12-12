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
	resultX, resultY, resultSize, power := 0, 0, 0, 0
	var powers [][]int = make([][]int, len(data))
	for i := 0; i < len(data); i++ {
		powers[i] = make([]int, len(data[0]))
	}
	for size := 1; size <= len(data); size++ {
		for y := 0; y < len(data)-size; y++ {
			for x := 0; x < len(data[y])-size; x++ {
				powers[y][x] += data[y+size-1][x+size-1]
				for i := 0; i < size-1; i++ {
					powers[y][x] += data[y+size-1][x+i]
					powers[y][x] += data[y+i][x+size-1]
				}
				if powers[y][x] > power {
					resultX, resultY, resultSize = x+1, y+1, size
					power = powers[y][x]
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
