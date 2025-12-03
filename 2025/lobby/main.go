package lobby

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

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func parseInput(data []string) ([][]int, error) {
	var result [][]int = make([][]int, 0)
	for _, bank := range data {
		var batteries []int = make([]int, 0)
		for _, battery := range bank {
			batteries = append(batteries, int(battery)-48)
		}
		result = append(result, batteries)
	}
	return result, nil
}

func findHighestSequence(bank []int, digits int) int {
	result, spare := 0, len(bank)-digits
	picked := make([]int, 0, digits)

	for _, battery := range bank {
		for len(picked) > 0 && spare > 0 && battery > picked[len(picked)-1] {
			picked = picked[:len(picked)-1]
			spare--
		}
		if len(picked) < digits {
			picked = append(picked, battery)
		} else if spare > 0 {
			spare--
		}
	}
	for _, cur := range picked {
		result = result*10 + cur
	}
	return result
}

func Challenge1(data []string) (int, error) {
	input, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, bank := range input {
		// best_left, best_pair := bank[0], bank[0]*10+bank[1]
		// for i := 1; i < len(bank); i++ {
		// 	if best_left*10+bank[i] > best_pair {
		// 		best_pair = best_left*10 + bank[i]
		// 	}
		// 	if int(bank[i]) > best_left {
		// 		best_left = bank[i]
		// 	}
		// }
		// result += best_pair
		result += findHighestSequence(bank, 2)
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	input, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, bank := range input {
		result += findHighestSequence(bank, 12)
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 3 - Lobby")
	path := "lobby/input.txt"
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
