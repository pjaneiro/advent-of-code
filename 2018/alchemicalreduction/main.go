package alchemicalreduction

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func readLines(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text(), nil
}

func removeDuplicates(data string) (string, int) {
	for i := 0; i < len(data)-1; i++ {
		left, right := data[i], data[i+1]
		if left-right == 32 || right-left == 32 {
			return removeDuplicates(data[:i] + data[i+2:])
		}
	}
	return data, len(data)
}

func Challenge1(data string) (int, error) {
	_, result := removeDuplicates(data)
	return result, nil
}

func Challenge2(data string) (int, error) {
	result := math.MaxInt

	start, _ := removeDuplicates(data)

	for letter := 'A'; letter <= 'Z'; letter++ {
		newString := strings.ReplaceAll(start, string([]rune{letter}), "")
		newString = strings.ReplaceAll(newString, string([]rune{letter + 32}), "")
		_, value := removeDuplicates(newString)
		if value < result {
			result = value
		}
	}

	return result, nil
}

func Run() {
	fmt.Println("Day 5 - Alchemical Reduction")
	path := "alchemicalreduction/input.txt"
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
