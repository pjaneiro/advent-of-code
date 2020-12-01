package week1

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, num)
	}
	return lines, scanner.Err()
}

func challenge1(data []int) (int, error) {
	for _, a := range data {
		for _, b := range data {
			if a+b == 2020 {
				return a * b, nil
			}
		}
	}
	return 0, errors.New("couldn't find a satisfying solution")
}

func challenge2(data []int) (int, error) {
	for _, a := range data {
		for _, b := range data {
			for _, c := range data {
				if a+b+c == 2020 {
					return a * b * c, nil
				}
			}
		}
	}
	return 0, errors.New("couldn't find a satisfying solution")
}

func main() {
	path := flag.String("path", "input.txt", "Path to file with challenge input")
	challenge := flag.Int("challenge", 1, "Which challenge to run (1 or 2)")
	flag.Parse()
	data, err := readLines(*path)
	if err != nil {
		fmt.Println(err)
	}
	if *challenge == 1 {
		result, err := challenge1(data)
		if err != nil {
			fmt.Printf("Error running challenge: %v\n", err)
		} else {
			fmt.Println(result)
		}
	} else if *challenge == 2 {
		result, err := challenge2(data)
		if err != nil {
			fmt.Printf("Error running challenge: %v\n", err)
		} else {
			fmt.Println(result)
		}
	} else {
		fmt.Println("Invalid challenge number")
	}
}
