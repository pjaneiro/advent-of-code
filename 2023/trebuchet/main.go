package trebuchet

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func valuesInLine(data string) []int {
	var numbers []int
	for _, char := range data {
		if char >= 48 && char <= 57 {
			numbers = append(numbers, int(char-48))
		}
	}
	return numbers
}

func replaceNumbers(data string) string {
	mappings := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	result := data
	for i := 0; i < len(result); i++ {
		for k, v := range mappings {
			if i+len(k) <= len(result) && strings.Index(result[i:i+len(k)], k) == 0 {
				result = result[0:i] + v + result[i+len(k)-1:]
				break
			}
		}
	}
	return result
}

func Challenge1(data []string) (int, error) {
	result := 0
	for _, line := range data {
		numbers := valuesInLine(line)
		result += 10*numbers[0] + numbers[len(numbers)-1]
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	result := 0
	for _, line := range data {
		numbers := valuesInLine(replaceNumbers(line))
		result += 10*numbers[0] + numbers[len(numbers)-1]
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 1 - Trebuchet?!")
	path := "trebuchet/input.txt"
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
