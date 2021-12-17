package universalorbitmap

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

func readLines(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make(map[string]string)
	scanner := bufio.NewScanner(file)
	r, err := regexp.Compile(`^(\w+)\)(\w+)$`)
	if err != nil {
		return lines, errors.New("couldn't parse regex expression")
	}
	for scanner.Scan() {
		data := r.FindAllStringSubmatch(scanner.Text(), -1)
		if data == nil {
			return lines, errors.New("string didn't parse")
		}
		lines[data[0][2]] = data[0][1]
	}
	return lines, nil
}

func Challenge1(data map[string]string) (int, error) {
	prev := make(map[string]int)
	count := 0
	missing := true
	for missing {
		missing = false
		for key, val := range data {
			if _, ok := prev[key]; ok {
				continue
			}
			if val == "COM" {
				prev[key] = 1
				count = 1
				continue
			}
			if v, ok := prev[val]; ok {
				prev[key] = 1 + v
				count += prev[key]
				continue
			}
			missing = true
		}
	}
	return count, nil
}

func Challenge2(data map[string]string) (int, error) {
	prev := make(map[string]int)
	for _, starter := range []string{data["YOU"], data["SAN"]} {
		cur := starter
		jumps := 0
		for true {
			if _, ok := prev[cur]; ok {
				return prev[cur] + jumps, nil
			}
			prev[cur] = jumps
			if cur == "COM" {
				break
			}
			cur = data[cur]
			jumps++
		}
	}
	return 0, errors.New("something went wrong")
}

func Run() {
	fmt.Println("Day 6 - Universal Orbit Map")
	path := "universalorbitmap/input.txt"
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
