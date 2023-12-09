package waitforit

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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

func parseData1(data []string) ([][]int, error) {
	var result [][]int = make([][]int, 2)
	times := strings.Fields(strings.Split(data[0], ":")[1])
	for _, t := range times {
		tmp, err := strconv.Atoi(t)
		if err != nil {
			return nil, err
		}
		result[0] = append(result[0], tmp)
	}
	dists := strings.Fields(strings.Split(data[1], ":")[1])
	for _, d := range dists {
		tmp, err := strconv.Atoi(d)
		if err != nil {
			return nil, err
		}
		result[1] = append(result[1], tmp)
	}
	return result, nil
}

func parseData2(data []string) ([]int, error) {
	var result []int = make([]int, 2)
	times := strings.ReplaceAll(strings.Split(data[0], ":")[1], " ", "")
	tmp, err := strconv.Atoi(times)
	if err != nil {
		return nil, err
	}
	result[0] = tmp
	dists := strings.ReplaceAll(strings.Split(data[1], ":")[1], " ", "")
	tmp, err = strconv.Atoi(dists)
	if err != nil {
		return nil, err
	}
	result[1] = tmp
	return result, nil
}

func isIntegral(val float64) bool {
	return val == float64(int(val))
}

func Challenge1(data []string) (int, error) {
	result := 1
	parsedData, err := parseData1(data)
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(parsedData[0]); i++ {
		a, b, c := float64(1), float64(-parsedData[0][i]), float64(parsedData[1][i])
		root1, root2 := (-b-math.Sqrt(b*b-4*a*c))/(2*a), (-b+math.Sqrt(b*b-4*a*c))/(2*a)
		if isIntegral(root1) {
			root1 += 1
		}
		if isIntegral(root2) {
			root2 -= 1
		}
		result *= (int(math.Floor(root2)-math.Ceil(root1)) + 1)
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	parsedData, err := parseData2(data)
	if err != nil {
		return 0, err
	}
	a, b, c := float64(1), float64(-parsedData[0]), float64(parsedData[1])
	root1, root2 := (-b-math.Sqrt(b*b-4*a*c))/(2*a), (-b+math.Sqrt(b*b-4*a*c))/(2*a)
	if isIntegral(root1) {
		root1 += 1
	}
	if isIntegral(root2) {
		root2 -= 1
	}
	return (int(math.Floor(root2)-math.Ceil(root1)) + 1), nil
}

func Run() {
	fmt.Println("Day 6 - Wait For It")
	path := "waitforit/input.txt"
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
