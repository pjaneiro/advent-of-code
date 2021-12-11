package handheldhalting

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
	Rpt bool
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
		r, err := regexp.Compile(`^(acc|jmp|nop) ([\+\-]\d+)$`)
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
		elem.Rpt = false

		lines = append(lines, elem)
	}
	return lines, nil
}

func Challenge1(data []Command) (int, error) {
	result, i := 0, 0
	for true {
		cur := &data[i]
		if cur.Rpt {
			return result, nil
		}
		cur.Rpt = true

		switch cur.Cmd {
		case "nop":
			i++
		case "jmp":
			i = i + cur.Val
		case "acc":
			result = result + cur.Val
			i++
		}
	}
	return 0, errors.New("something went wrong")
}

func challenge2loop(data []Command) (int, int, error) {
	result, i := 0, 0
	for true {
		if i == len(data) {
			return result, i, nil
		}

		cur := &data[i]
		if cur.Rpt {
			return result, i, nil
		}
		cur.Rpt = true

		switch cur.Cmd {
		case "nop":
			i++
		case "jmp":
			i = i + cur.Val
		case "acc":
			result = result + cur.Val
			i++
		}
	}
	return 0, 0, errors.New("something went wrong")
}

func Challenge2(data []Command) (int, error) {
	for i, c := range data {
		cpy := make([]Command, len(data))
		copy(cpy, data)
		if c.Cmd == "jmp" {
			cpy[i].Cmd = "nop"
		} else if c.Cmd == "nop" {
			cpy[i].Cmd = "jmp"
		}

		rr, ri, rerr := challenge2loop(cpy)
		if ri == len(data) && rerr == nil {
			return rr, nil
		}
	}
	return 0, errors.New("something went wrong")
}

func Run() {
	fmt.Println("Day 8 - Handheld Halting")
	path := "handheldhalting/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	var cpy []Command
	cpy = make([]Command, len(data))
	copy(cpy, data)
	result, err = Challenge1(cpy)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	cpy = make([]Command, len(data))
	copy(cpy, data)
	result, err = Challenge2(cpy)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
