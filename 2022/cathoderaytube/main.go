package cathoderaytube

import (
	"bufio"
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

	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

func getCycles(data []string) ([240]int, error) {
	X := 1
	cycle := 0
	var result [240]int
	for i := 0; i < 240; i++ {
		result[i] = X
	}
	for _, inst := range data {
		if inst == "noop" {
			cycle++
			continue
		}

		val, err := strconv.Atoi(inst[5:])
		if err != nil {
			return result, err
		}

		X += val

		for i := cycle + 2; i < 240; i++ {
			result[i] = X
		}
		cycle += 2
	}
	return result, nil
}

func Challenge1(data []string) (int, error) {
	cycles, err := getCycles(data)
	if err != nil {
		return 0, err
	}

	result := 0
	for _, c := range []int{20, 60, 100, 140, 180, 220} {
		result += c * cycles[c-1]
	}

	return result, nil
}

func Challenge2(data []string) ([]string, error) {
	values, err := getCycles(data)
	if err != nil {
		return nil, err
	}

	var result []string

	for line := 0; line < 6; line++ {
		out := ""
		for col := 0; col < 40; col++ {
			X := values[(line*40)+col]
			if col == X || col == X+1 || col == X-1 {
				out += "#"
			} else {
				out += "."
			}
		}
		result = append(result, out)
	}

	return result, nil
}

func Run() {
	fmt.Println("Day 10 - Cathode-Ray Tube")
	path := "cathoderaytube/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result1 int
	result1, err = Challenge1(data)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result1)
	}

	var result2 []string
	result2, err = Challenge2(data)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Println("Challenge 2:")
		for _, line := range result2 {
			fmt.Println(line)
		}
	}

}
