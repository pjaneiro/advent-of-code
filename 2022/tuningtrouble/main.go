package tuningtrouble

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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

func findCode(data string, size int) (int, error) {
outerloop:
	for i := 0; i <= len(data)-size; i++ {
		for j := i; j < i+size-1; j++ {
			for k := j + 1; k <= i+size-1; k++ {
				if data[j] == data[k] {
					continue outerloop
				}
			}
		}
		return i + size, nil
	}
	return 0, errors.New("something went wrong")
}

func Challenge1(data string) (int, error) {
	return findCode(data, 4)
}

func Challenge2(data string) (int, error) {
	return findCode(data, 14)
}

func Run() {
	fmt.Println("Day 6 - Tuning Trouble")
	path := "tuningtrouble/input.txt"
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
