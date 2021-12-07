package thetreacheryofwhales

import (
	"bufio"
	"fmt"
	"math"
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

	var data []int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")
	for _, cur := range line {
		elem, err := strconv.Atoi(cur)
		if err != nil {
			return data, err
		}
		data = append(data, elem)
	}
	return data, nil
}

func abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

func Challenge1(data []int) (int, error) {
	min, max, totalFuel := math.MaxInt32, 0, math.MaxInt32
	weights := make(map[int]int)
	for _, cur := range data {
		if cur < min {
			min = cur
		}
		if cur > max {
			max = cur
		}
		weights[cur]++
	}
	for i := min; i <= max; i++ {
		curFuel := 0
		for k, v := range weights {
			curDistance := abs(k - i)
			curFuel = curFuel + (v * curDistance)
		}
		if curFuel < totalFuel {
			totalFuel = curFuel
		}
	}
	return totalFuel, nil
}

func Challenge2(data []int) (int, error) {
	min, max, totalFuel := math.MaxInt32, 0, math.MaxInt32
	weights := make(map[int]int)
	for _, cur := range data {
		if cur < min {
			min = cur
		}
		if cur > max {
			max = cur
		}
		weights[cur]++
	}
	for i := min; i <= max; i++ {
		curFuel := 0
		for k, v := range weights {
			curDistance := abs(k - i)
			crabFuel := (curDistance * (curDistance + 1)) / 2
			curFuel = curFuel + (v * crabFuel)
		}
		if curFuel < totalFuel {
			totalFuel = curFuel
		}
	}
	return totalFuel, nil
}

func Run() {
	fmt.Println("Day 7 - The Treachery of Whales")
	path := "thetreacheryofwhales/input.txt"
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
