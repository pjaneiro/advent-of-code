package marblemania

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Marble struct {
	Value int
	Left  *Marble
	Right *Marble
}

func readLines(path string) (int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	r, err := regexp.Compile(`^(\d+) players; last marble is worth (\d+) points$`)
	if err != nil {
		return 0, 0, errors.New("couldn't parse regex expression")
	}

	var nPlayers, lastMarble int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	data := r.FindStringSubmatch(line)
	if nPlayers, err = strconv.Atoi(data[1]); err != nil {
		return 0, 0, errors.New("couldn't parse input")
	}
	if lastMarble, err = strconv.Atoi(data[2]); err != nil {
		return 0, 0, errors.New("couldn't parse input")
	}

	return nPlayers, lastMarble, nil
}

func insertRight(current *Marble, value int) *Marble {
	var newMarble Marble
	newMarble.Value = value
	pastRight := current.Right

	newMarble.Right = pastRight
	newMarble.Left = current
	current.Right = &newMarble
	pastRight.Left = &newMarble

	return &newMarble
}

func remove(current *Marble) *Marble {
	current.Left.Right = current.Right
	current.Right.Left = current.Left
	return current.Right
}

func solve(nPlayers int, lastMarble int) int {
	startMarble := Marble{Value: 0}
	startMarble.Right = &startMarble
	startMarble.Left = &startMarble

	current := &startMarble

	var scores []int = make([]int, nPlayers)

	for marble, player := 1, 0; marble <= lastMarble; marble, player = marble+1, (player+1)%nPlayers {
		if marble%23 == 0 {
			scores[player] += marble
			for i := 0; i < 7; i++ {
				current = current.Left
			}
			value := current.Value
			scores[player] += value
			current = remove(current)
		} else {
			current = insertRight(current.Right, marble)
		}
	}

	result := 0
	for _, val := range scores {
		if val > result {
			result = val
		}
	}
	return result
}

func Challenge1(nPlayers int, lastMarble int) (int, error) {
	return solve(nPlayers, lastMarble), nil
}

func Challenge2(nPlayers int, lastMarble int) (int, error) {
	return solve(nPlayers, lastMarble*100), nil
}

func Run() {
	fmt.Println("Day 9 - Marble Mania")
	path := "marblemania/input.txt"
	nPlayers, lastMarble, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(nPlayers, lastMarble)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(nPlayers, lastMarble)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
