package scratchcards

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Id             int
	WinningNumbers map[int]struct{}
	NumbersIHave   map[int]struct{}
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
		var card Card = Card{Id: 0, WinningNumbers: make(map[int]struct{}), NumbersIHave: make(map[int]struct{})}
		parts := strings.Split(line, ":")
		idPart, rest := parts[0], parts[1]
		id, err := strconv.Atoi(strings.TrimSpace(idPart[4:]))
		if err != nil {
			return nil, err
		}
		card.Id = id
		parts = strings.Split(rest, " | ")
		winningPart, havePart := parts[0], parts[1]
		for _, candidate := range strings.Fields(winningPart) {
			tmp, err := strconv.Atoi(candidate)
			if err != nil {
				return nil, err
			}
			card.WinningNumbers[tmp] = struct{}{}
		}
		for _, candidate := range strings.Fields(havePart) {
			tmp, err := strconv.Atoi(candidate)
			if err != nil {
				return nil, err
			}
			card.NumbersIHave[tmp] = struct{}{}
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
			result = result + int(math.Pow(2, float64(cardPrizedNumbers-1)))
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
	for _, card := range cards {
		if _, ok := cardCounts[card.Id]; !ok {
			cardCounts[card.Id] = 1
		}
		cardPrizedNumbers := 0
		for cur := range card.WinningNumbers {
			if _, ok := card.NumbersIHave[cur]; ok {
				cardPrizedNumbers++
			}
		}
		for i := 0; i < cardPrizedNumbers; i++ {
			if _, ok := cardCounts[card.Id+1+i]; !ok {
				cardCounts[card.Id+1+i] = 1 + cardCounts[card.Id]
			} else {
				cardCounts[card.Id+1+i] = cardCounts[card.Id+1+i] + cardCounts[card.Id]
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
