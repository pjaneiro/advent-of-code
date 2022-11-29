package inventorymanagementsystem

import (
	"bufio"
	"errors"
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

func checkDuplicates(data string) (int, int, error) {
	count2, count3 := 0, 0
	values := make(map[rune]int)
	for _, each := range data {
		if _, ok := values[each]; ok {
			values[each]++
		} else {
			values[each] = 1
		}
	}
	for _, v := range values {
		if v == 2 {
			count2++
		}
		if v == 3 {
			count3++
		}
	}
	return count2, count3, nil
}

func Challenge1(data []string) (int, error) {
	total2, total3 := 0, 0
	for _, each := range data {
		count2, count3, err := checkDuplicates(each)
		if err != nil {
			return 0, errors.New("something went wrong")
		}
		if count2 > 0 {
			total2++
		}
		if count3 > 0 {
			total3++
		}
	}
	return total2 * total3, nil
}

func Challenge2(data []string) (string, error) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			index, diff := 0, 0
			for cur := 0; cur < len(data[i]); cur++ {
				if data[i][cur] != data[j][cur] {
					diff++
					index = cur
				}
			}
			if diff == 1 {
				return data[i][0:index] + data[i][index+1:], nil
			}
		}
	}
	return "", errors.New("not implemented yet")
}

func Run() {
	fmt.Println("Day 2 - Inventory Management System")
	path := "inventorymanagementsystem/input.txt"
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

	var result2 string
	result2, err = Challenge2(data)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %s\n", result2)
	}
}
