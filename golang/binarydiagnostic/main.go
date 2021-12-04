package binarydiagnostic

import (
	"bufio"
	// "errors"
	"fmt"
	"os"
	"strconv"
)

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

func countTheOnes(data []string, length int) []int {
	ones := make([]int, length)
	for _, cur := range data {
		for i, rune := range cur {
			if rune == '1' {
				ones[i]++
			}
		}
	}
	return ones
}

func mostCommonAtIndex(data []string, index int) byte {
	ones := 0
	for _, cur := range data {
		if cur[index] == '1' {
			ones++
		}
	}
	if ones >= (len(data) - ones) {
		return '1'
	} else {
		return '0'
	}
}

func subData(data []string, target byte, index int) []string {
	var newData []string
	for _, cur := range data {
		if cur[index] == target {
			newData = append(newData, cur)
		}
	}
	return newData
}

func Challenge1(data []string) (int, error) {
	gamma, epsilon, elemLength, count := 0, 0, len(data[0]), len(data)
	ones := countTheOnes(data, elemLength)
	for i, cur := range ones {
		if cur > count/2 {
			gamma = gamma + (1 << ((elemLength - 1) - i))
		} else {
			epsilon = epsilon + (1 << ((elemLength - 1) - i))
		}
	}
	return gamma * epsilon, nil
}

func Challenge2(data []string) (int, error) {
	o2, co2, elemLength := 0, 0, len(data[0])
	o2data, co2data := data, data
	for i := 0; i < elemLength && o2 == 0; i++ {
		mostCommono2 := mostCommonAtIndex(o2data, i)
		o2data = subData(o2data, mostCommono2, i)
		if len(o2data) == 1 {
			tmp, err := strconv.ParseInt(o2data[0], 2, 64)
			if err != nil {
				return 0, err
			}
			o2 = int(tmp)
		}
	}
	for i := 0; i < elemLength && co2 == 0; i++ {
		mostCommonco2 := mostCommonAtIndex(co2data, i)
		leastCommonco2 := mostCommonco2
		if mostCommonco2 == '1' {
			leastCommonco2 = '0'
		} else {
			leastCommonco2 = '1'
		}
		co2data = subData(co2data, leastCommonco2, i)
		if len(co2data) == 1 {
			tmp, err := strconv.ParseInt(co2data[0], 2, 64)
			if err != nil {
				return 0, err
			}
			co2 = int(tmp)
		}
	}
	return o2 * co2, nil
}

func Run() {
	fmt.Println("Day 3 - Binary Diagnostic")
	path := "binarydiagnostic/input.txt"
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
