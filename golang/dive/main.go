package dive

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Command struct {
	Cmd string
	Val int
}

func readLines(path string) ([]Command, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Command
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r, err := regexp.Compile(`^(forward|down|up) (\d+)$`)
		if err != nil {
			return lines, errors.New("couldn't parse regex expression")
		}

		var elem Command
		var data []string = r.FindStringSubmatch(scanner.Text())
		if data == nil {
			return lines, nil
		}

		elem.Cmd = data[1]
		elem.Val, err = strconv.Atoi(data[2])

		lines = append(lines, elem)
	}
	return lines, nil
}

func Challenge1(data []Command) (int, error) {
	h, d := 0, 0
	for _, cur := range data {
		switch cur.Cmd {
		case "forward":
			h = h + cur.Val
		case "down":
			d = d + cur.Val
		case "up":
			d = d - cur.Val
		}
	}
	return h * d, nil
}

func Challenge2(data []Command) (int, error) {
	h, d, a := 0, 0, 0
	for _, cur := range data {
		switch cur.Cmd {
		case "forward":
			h = h + cur.Val
			d = d + (a * cur.Val)
		case "down":
			a = a + cur.Val
		case "up":
			a = a - cur.Val
		}
	}
	return h * d, nil
}

func Run() {
	fmt.Println("Day 2 - Dive!")
	path := "dive/input.txt"
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
