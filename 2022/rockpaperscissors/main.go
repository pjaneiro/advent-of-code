package rockpaperscissors

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

func Challenge1(data []string) (int, error) {
	scores := map[string]int{
		"A X": 3 + 1,
		"A Y": 6 + 2,
		"A Z": 0 + 3,
		"B X": 0 + 1,
		"B Y": 3 + 2,
		"B Z": 6 + 3,
		"C X": 6 + 1,
		"C Y": 0 + 2,
		"C Z": 3 + 3,
	}
	result := 0
	for _, play := range data {
		result += scores[play]
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	scores := map[string]int{
		"A X": 0 + 3,
		"A Y": 3 + 1,
		"A Z": 6 + 2,
		"B X": 0 + 1,
		"B Y": 3 + 2,
		"B Z": 6 + 3,
		"C X": 0 + 2,
		"C Y": 3 + 3,
		"C Z": 6 + 1,
	}
	result := 0
	for _, play := range data {
		result += scores[play]
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 2 - Rock Paper Scissors")
	path := "rockpaperscissors/input.txt"
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
