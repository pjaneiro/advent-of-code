package scratchcards

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Card struct {
	WinningNumbers map[string]struct{}
	NumbersIHave   map[string]struct{}
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

func parseCards(data []string) ([]Card, error) {
	var result []Card = make([]Card, 0)
	for _, line := range data {
		var card Card = Card{WinningNumbers: make(map[string]struct{}), NumbersIHave: make(map[string]struct{})}
		parts := strings.Split(line, ":")
		parts = strings.Split(parts[1], " | ")
		winningPart, havePart := parts[0], parts[1]
		for _, candidate := range strings.Fields(winningPart) {
			card.WinningNumbers[candidate] = struct{}{}
		}
		for _, candidate := range strings.Fields(havePart) {
			card.NumbersIHave[candidate] = struct{}{}
		}
		result = append(result, card)
	}
	return result, nil
}

func Challenge1(data []string) (int, error) {
	result := 0
	cards, err := parseCards(data)
	if err != nil {
		return 0, err
	}
	for _, card := range cards {
		cardPrizedNumbers := 0
		for cur := range card.WinningNumbers {
			if _, ok := card.NumbersIHave[cur]; ok {
				cardPrizedNumbers++
			}
		}
		if cardPrizedNumbers > 0 {
			result = result + (1 << (cardPrizedNumbers - 1))
		}
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	result := 0
	cards, err := parseCards(data)
	if err != nil {
		return 0, err
	}
	var cardCounts map[int]int = make(map[int]int)
	for id, card := range cards {
		if _, ok := cardCounts[id]; !ok {
			cardCounts[id] = 1
		}
		cardPrizedNumbers := 0
		for cur := range card.WinningNumbers {
			if _, ok := card.NumbersIHave[cur]; ok {
				cardPrizedNumbers++
			}
		}
		for i := 0; i < cardPrizedNumbers; i++ {
			if _, ok := cardCounts[id+1+i]; !ok {
				cardCounts[id+1+i] = 1 + cardCounts[id]
			} else {
				cardCounts[id+1+i] = cardCounts[id+1+i] + cardCounts[id]
			}
		}
	}
	for _, cur := range cardCounts {
		result += cur
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 4 - Scratchcards")
	path := "scratchcards/input.txt"
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
