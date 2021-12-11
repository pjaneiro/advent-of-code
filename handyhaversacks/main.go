package handyhaversacks

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readLines(path string) (map[string]map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make(map[string]map[string]int)
	scanner := bufio.NewScanner(file)
	r, err := regexp.Compile(`^(\w+ \w+) bags contain ((?:(?:, )?\d+ \w+ \w+ bags?)+)\.$|^(\w+ \w+) bags contain no other bags.$`)
	if err != nil {
		return lines, errors.New("couldn't parse regex expression")
	}
	for scanner.Scan() {
		data := r.FindAllStringSubmatch(scanner.Text(), -1)
		if data == nil {
			return lines, errors.New("string didn't parse")
		}
		if len(data[0][3]) != 0 {
			lines[data[0][3]] = make(map[string]int)
			continue
		}
		lines[data[0][1]] = make(map[string]int)

		c := strings.Split(data[0][2], ", ")
		for _, e := range c {
			s := strings.Split(e, " ")
			lines[data[0][1]][s[1]+" "+s[2]], _ = strconv.Atoi(s[0])
		}
	}
	return lines, nil
}

func challenge1loop(data map[string]map[string]int, key string) bool {
	e := data[key]
	if _, ok := e["shiny gold"]; ok {
		return true
	}

	for k := range e {
		if challenge1loop(data, k) {
			return true
		}
	}
	return false
}

func Challenge1(data map[string]map[string]int) (int, error) {
	result := 0
	for k, e := range data {
		if _, ok := e["shiny gold"]; ok {
			result++
			continue
		}

		if challenge1loop(data, k) {
			result++
			continue
		}
	}
	return result, nil
}

func challenge2loop(data map[string]map[string]int, key string) int {
	result := 0
	for k, v := range data[key] {
		result = result + v
		result = result + v*challenge2loop(data, k)
	}
	return result
}

func Challenge2(data map[string]map[string]int) (int, error) {
	result := 0 + challenge2loop(data, "shiny gold")

	return result, nil
}

func Run() {
	fmt.Println("Day 7 - Handy Haversacks")
	path := "handyhaversacks/input.txt"
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
