package historianhysteria

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func readLines(path string) ([]int, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		n1, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, nil, err
		}
		list1 = append(list1, n1)
		n2, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, nil, err
		}
		list2 = append(list2, n2)
	}
	return list1, list2, nil
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func Challenge1(input1 []int, input2 []int) (int, error) {
	diff := 0
	list1, list2 := slices.Clone(input1), slices.Clone(input2)
	sort.Ints(list1)
	sort.Ints(list2)
	for i := range list1 {
		diff += abs(list1[i] - list2[i])
	}
	return diff, nil
}

func Challenge2(list1 []int, list2 []int) (int, error) {
	occurences := make(map[int]int)
	for _, val := range list2 {
		occurences[val]++
	}
	score := 0
	for _, v := range list1 {
		score += v * occurences[v]
	}
	return score, nil
}

func Run() {
	fmt.Println("Day 1 - Historian Hysteria")
	path := "historianhysteria/input.txt"
	list1, list2, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(list1, list2)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(list1, list2)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
