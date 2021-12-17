package lanternfish

import (
	"bufio"
	"fmt"
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

func Challenge1(data []int) (int64, error) {
	for generation := 0; generation < 80; generation++ {
		curLength := len(data)
		for i := 0; i < curLength; i++ {
			if data[i] == 0 {
				data[i] = 6
				data = append(data, 8)
			} else {
				data[i] = data[i] - 1
			}
		}
	}
	return int64(len(data)), nil
}

func Challenge2(data []int) (int64, error) {
	buckets := make([]int64, 9)
	for _, cur := range data {
		buckets[cur]++
	}
	for generation := 0; generation < 256; generation++ {
		newBuckets := make([]int64, 9)
		for i := 0; i < 9; i++ {
			count := buckets[i]
			if i == 0 {
				newBuckets[6] = newBuckets[6] + count
				newBuckets[8] = newBuckets[8] + count
			} else {
				newBuckets[i-1] = newBuckets[i-1] + count
			}
		}
		copy(buckets, newBuckets)
	}
	var total int64
	for _, cur := range buckets {
		total = total + cur
	}
	return total, nil
}

func Run() {
	fmt.Println("Day 6 - Lantern Fish")
	path := "lanternfish/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var cpy []int
	var result int64

	cpy = make([]int, len(data))
	copy(cpy, data)
	result, err = Challenge1(cpy)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	cpy = make([]int, len(data))
	copy(cpy, data)
	result, err = Challenge2(cpy)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
