package memorymaneuver

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	data := strings.Split(line, " ")
	var val int
	for _, cur := range data {
		if val, err = strconv.Atoi(cur); err != nil {
			return nil, err
		}
		result = append(result, val)
	}

	return result, nil
}

func scan(data []int, pos int) (int, int, int) {
	nChildren, nMetadata := data[pos], data[pos+1]
	innerPos, metadata, value := pos+2, 0, 0
	var subValues []int = make([]int, nChildren)
	for i := 0; i < nChildren; i++ {
		childValue, childMetadata, childSize := scan(data, innerPos)
		subValues[i] = childValue
		metadata += childMetadata
		innerPos += childSize
	}
	for i := 0; i < nMetadata; i++ {
		metadata += data[innerPos]
		if nChildren == 0 {
			value += data[innerPos]
		} else if data[innerPos] <= nChildren {
			value += subValues[data[innerPos]-1]
		}
		innerPos++
	}
	return value, metadata, innerPos - pos
}

func Challenge1(data []int) (int, error) {
	_, result, _ := scan(data, 0)
	return result, nil
}

func Challenge2(data []int) (int, error) {
	result, _, _ := scan(data, 0)
	return result, nil
}

func Run() {
	fmt.Println("Day 8 - Memory Maneuver")
	path := "memorymaneuver/input.txt"
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
