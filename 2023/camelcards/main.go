package camelcards

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type HandType int

const (
	HighCard     HandType = iota
	OnePair      HandType = iota
	TwoPair      HandType = iota
	ThreeOfAKind HandType = iota
	FullHouse    HandType = iota
	FourOfAKind  HandType = iota
	FiveOfAKind  HandType = iota
)

type Hand struct {
	Cards []rune
	Bid   int
	Type  HandType
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

func parseHands(data []string) ([]Hand, error) {
	var result []Hand = make([]Hand, 0)
	for _, line := range data {
		var hand Hand
		parts := strings.Fields(line)
		hand.Cards = []rune(parts[0])
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		hand.Bid = bid
		result = append(result, hand)
	}
	return result, nil
}

func (hand *Hand) setHandType1() {
	var cardCounts map[rune]int = make(map[rune]int)
	for _, card := range hand.Cards {
		if _, ok := cardCounts[card]; !ok {
			cardCounts[card] = 0
		}
		cardCounts[card]++
	}
	five, four, three, two := 0, 0, 0, 0
	for _, count := range cardCounts {
		switch count {
		case 5:
			five++
		case 4:
			four++
		case 3:
			three++
		case 2:
			two++
		}
	}
	if five > 0 {
		hand.Type = FiveOfAKind
	} else if four > 0 {
		hand.Type = FourOfAKind
	} else if three > 0 && two > 0 {
		hand.Type = FullHouse
	} else if three > 0 {
		hand.Type = ThreeOfAKind
	} else if two > 1 {
		hand.Type = TwoPair
	} else if two > 0 {
		hand.Type = OnePair
	} else {
		hand.Type = HighCard
	}
}

func (hand *Hand) setHandType2() {
	var cardCounts map[rune]int = make(map[rune]int)
	for _, card := range hand.Cards {
		if _, ok := cardCounts[card]; !ok {
			cardCounts[card] = 0
		}
		cardCounts[card]++
	}
	five, four, three, two := 0, 0, 0, 0
	for _, count := range cardCounts {
		switch count {
		case 5:
			five++
		case 4:
			four++
		case 3:
			three++
		case 2:
			two++
		}
	}

	switch cardCounts['J'] {
	case 4:
		five = 1
		four = 0
	case 3:
		if two == 1 {
			five = 1
			three = 0
			two = 0
		} else {
			four = 1
			three = 0
		}
	case 2:
		if three == 1 {
			five = 1
			three = 0
			two = 0
		} else if two == 2 {
			four = 1
			two = 0
		} else {
			three = 1
			two = 0
		}
	case 1:
		if four == 1 {
			five = 1
			four = 0
		} else if three == 1 {
			four = 1
			three = 0
		} else if two == 2 {
			three = 1
			two = 1
		} else if two == 1 {
			three = 1
			two = 0
		} else {
			two = 1
		}
	}
	if five > 0 {
		hand.Type = FiveOfAKind
	} else if four > 0 {
		hand.Type = FourOfAKind
	} else if three > 0 && two > 0 {
		hand.Type = FullHouse
	} else if three > 0 {
		hand.Type = ThreeOfAKind
	} else if two > 1 {
		hand.Type = TwoPair
	} else if two > 0 {
		hand.Type = OnePair
	} else {
		hand.Type = HighCard
	}
}

func Challenge1(data []string) (int, error) {
	result := 0
	parsedHands, err := parseHands(data)
	if err != nil {
		return 0, err
	}
	for i := range parsedHands {
		parsedHands[i].setHandType1()
	}
	slices.SortFunc(parsedHands, func(hand1, hand2 Hand) int {
		if hand1.Type != hand2.Type {
			return int(hand1.Type - hand2.Type)
		}
		return slices.CompareFunc(hand1.Cards, hand2.Cards, func(card1, card2 rune) int {
			var CardValues = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
			return slices.Index(CardValues, card1) - slices.Index(CardValues, card2)
		})
	})
	for index, hand := range parsedHands {
		result = result + ((index + 1) * hand.Bid)
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	result := 0
	parsedHands, err := parseHands(data)
	if err != nil {
		return 0, err
	}
	for i := range parsedHands {
		parsedHands[i].setHandType2()
	}
	slices.SortFunc(parsedHands, func(hand1, hand2 Hand) int {
		if hand1.Type != hand2.Type {
			return int(hand1.Type - hand2.Type)
		}
		return slices.CompareFunc(hand1.Cards, hand2.Cards, func(card1, card2 rune) int {
			var CardValues = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
			return slices.Index(CardValues, card1) - slices.Index(CardValues, card2)
		})
	})
	for index, hand := range parsedHands {
		result = result + ((index + 1) * hand.Bid)
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 7 - Camel Cards")
	path := "camelcards/input.txt"
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
