package clawcontraption

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type clawmachine struct {
	Ax, Ay, Bx, By, X, Y int
}

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

func parseInput(data []string) ([]clawmachine, error) {
	var result []clawmachine = make([]clawmachine, 0)
	r, err := regexp.Compile(`^[\w\ ]+: X[+=](\d+), Y[+=](\d+)$`)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(data); {
		baLine, bbLine, pLine := data[i], data[i+1], data[i+2]
		i += 4
		var cur clawmachine
		var err error
		baFound := r.FindStringSubmatch(baLine)
		bbFound := r.FindStringSubmatch(bbLine)
		pFound := r.FindStringSubmatch(pLine)
		if cur.Ax, err = strconv.Atoi(baFound[1]); err != nil {
			return nil, err
		}
		if cur.Ay, err = strconv.Atoi(baFound[2]); err != nil {
			return nil, err
		}
		if cur.Bx, err = strconv.Atoi(bbFound[1]); err != nil {
			return nil, err
		}
		if cur.By, err = strconv.Atoi(bbFound[2]); err != nil {
			return nil, err
		}
		if cur.X, err = strconv.Atoi(pFound[1]); err != nil {
			return nil, err
		}
		if cur.Y, err = strconv.Atoi(pFound[2]); err != nil {
			return nil, err
		}
		result = append(result, cur)
	}
	return result, nil
}

func Challenge1(data []string) (int, error) {
	machines, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, machine := range machines {
		a := ((machine.Bx * machine.Y) - (machine.By * machine.X)) / ((machine.Bx * machine.Ay) - (machine.By * machine.Ax))
		b := ((machine.Ax * machine.Y) - (machine.Ay * machine.X)) / ((machine.Ax * machine.By) - (machine.Ay * machine.Bx))
		if a*machine.Ax+b*machine.Bx == machine.X && a*machine.Ay+b*machine.By == machine.Y {
			result += (3*a + b)
		}
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	machines, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, machine := range machines {
		X, Y := machine.X+10000000000000, machine.Y+10000000000000
		a := ((machine.Bx * Y) - (machine.By * X)) / ((machine.Bx * machine.Ay) - (machine.By * machine.Ax))
		b := ((machine.Ax * Y) - (machine.Ay * X)) / ((machine.Ax * machine.By) - (machine.Ay * machine.Bx))
		if a*machine.Ax+b*machine.Bx == X && a*machine.Ay+b*machine.By == Y {
			result += (3*a + b)
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 13 - Claw Contraption")
	path := "clawcontraption/input.txt"
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
