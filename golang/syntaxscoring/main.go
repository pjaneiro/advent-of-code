package syntaxscoring

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func Challenge1(data []string) (int, error) {
	score := 0
	for _, curLine := range data {
		var pile []rune
		for _, curChar := range curLine {
			if curChar == '(' || curChar == '[' || curChar == '{' || curChar == '<' {
				pile = append(pile, curChar)
			} else if curChar == ')' && pile[len(pile)-1] != '(' {
				score = score + 3
				break
			} else if curChar == ']' && pile[len(pile)-1] != '[' {
				score = score + 57
				break
			} else if curChar == '}' && pile[len(pile)-1] != '{' {
				score = score + 1197
				break
			} else if curChar == '>' && pile[len(pile)-1] != '<' {
				score = score + 25137
				break
			} else {
				pile = pile[:len(pile)-1]
			}
		}
	}
	return score, nil
}

func Challenge2(data []string) (int, error) {
	var scores []int
outerLoop:
	for _, curLine := range data {
		curScore := 0
		var pile []rune
		for _, curChar := range curLine {
			if curChar == '(' || curChar == '[' || curChar == '{' || curChar == '<' {
				pile = append(pile, curChar)
			} else if curChar == ')' && pile[len(pile)-1] != '(' {
				continue outerLoop
			} else if curChar == ']' && pile[len(pile)-1] != '[' {
				continue outerLoop
			} else if curChar == '}' && pile[len(pile)-1] != '{' {
				continue outerLoop
			} else if curChar == '>' && pile[len(pile)-1] != '<' {
				continue outerLoop
			} else {
				pile = pile[:len(pile)-1]
			}
		}
		for i := len(pile) - 1; i >= 0; i-- {
			curChar := pile[i]
			curScore = curScore * 5
			if curChar == '(' {
				curScore = curScore + 1
			} else if curChar == '[' {
				curScore = curScore + 2
			} else if curChar == '{' {
				curScore = curScore + 3
			} else {
				curScore = curScore + 4
			}
		}
		scores = append(scores, curScore)
	}
	sort.Ints(scores)
	return scores[len(scores)/2], nil
}

func Run() {
	fmt.Println("Day 10 - Syntax Scoring")
	path := "syntaxscoring/input.txt"
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
