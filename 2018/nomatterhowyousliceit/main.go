package nomatterhowyousliceit

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Pos struct {
	X, Y int
}

type Claim struct {
	Id   int
	X, Y int
	W, H int
}

func readLines(path string) ([]Claim, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r, err := regexp.Compile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}

	var claims []Claim
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := r.FindStringSubmatch(line)
		var claim Claim
		if claim.Id, err = strconv.Atoi(data[1]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if claim.X, err = strconv.Atoi(data[2]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if claim.Y, err = strconv.Atoi(data[3]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if claim.W, err = strconv.Atoi(data[4]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if claim.H, err = strconv.Atoi(data[5]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		claims = append(claims, claim)
	}

	return claims, nil
}

func Challenge1(data []Claim) (int, error) {
	var hits map[Pos]int = make(map[Pos]int)
	count := 0
	for _, cur := range data {
		for x := cur.X; x < cur.X+cur.W; x++ {
			for y := cur.Y; y < cur.Y+cur.H; y++ {
				pos := Pos{X: x, Y: y}
				if _, ok := hits[pos]; !ok {
					hits[pos] = 0
				}
				hits[pos]++
			}
		}
	}
	for _, cur := range hits {
		if cur > 1 {
			count++
		}
	}
	return count, nil
}

func Challenge2(data []Claim) (int, error) {
	var hits map[Pos]int = make(map[Pos]int)
	for _, cur := range data {
		for x := cur.X; x < cur.X+cur.W; x++ {
			for y := cur.Y; y < cur.Y+cur.H; y++ {
				pos := Pos{X: x, Y: y}
				if _, ok := hits[pos]; !ok {
					hits[pos] = 0
				}
				hits[pos]++
			}
		}
	}
outerloop:
	for _, cur := range data {
		for x := cur.X; x < cur.X+cur.W; x++ {
			for y := cur.Y; y < cur.Y+cur.H; y++ {
				pos := Pos{X: x, Y: y}
				if hits[pos] > 1 {
					continue outerloop
				}
			}
		}
		return cur.Id, nil
	}
	return 0, errors.New("something went wrong")
}

func Run() {
	fmt.Println("Day 3 - No Matter How You Slice It")
	path := "nomatterhowyousliceit/input.txt"
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
