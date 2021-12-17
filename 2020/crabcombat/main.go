package crabcombat

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Deck []int

func readLines(path string) (Deck, Deck, error) {
	file, err := os.Open(path)
	if err != nil {
		return Deck{}, Deck{}, err
	}
	defer file.Close()

	var deck1, deck2 Deck = make(Deck, 0), make(Deck, 0)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		val, err := strconv.Atoi(line)
		if err != nil {
			return deck1, deck2, err
		}
		deck1 = append(deck1, val)
	}
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		val, err := strconv.Atoi(line)
		if err != nil {
			return deck1, deck2, err
		}
		deck2 = append(deck2, val)
	}

	return deck1, deck2, nil
}

func (d1 Deck) equals(d2 Deck) bool {
	if len(d1) != len(d2) {
		return false
	}
	for k, v := range d1 {
		if d2[k] != v {
			return false
		}
	}
	return true
}

func contains(decks []Deck, deck Deck) bool {
	for _, cur := range decks {
		if cur.equals(deck) {
			return true
		}
	}
	return false
}

func playRecursiveGame(deck1 Deck, deck2 Deck) (Deck, int) {
	var cpy1, cpy2 Deck = make(Deck, len(deck1)), make(Deck, len(deck2))
	copy(cpy1, deck1)
	copy(cpy2, deck2)
	hist1, hist2 := []Deck{}, []Deck{}
	var lock bool
	for len(cpy1) != 0 && len(cpy2) != 0 {
		cpy1, cpy2, hist1, hist2, lock = playRecursiveRound(cpy1, cpy2, hist1, hist2)
		if lock {
			return cpy1, 1
		}
	}
	if len(cpy1) == 0 {
		return cpy2, 2
	}
	return cpy1, 1
}

func playRecursiveRound(deck1 Deck, deck2 Deck, hist1 []Deck, hist2 []Deck) (Deck, Deck, []Deck, []Deck, bool) {
	if contains(hist1, deck1) && contains(hist2, deck2) {
		return deck1, deck2, hist1, hist2, true
	}

	hist1, hist2 = append(hist1, deck1), append(hist2, deck2)
	draw1, draw2 := deck1[0], deck2[0]
	deck1, deck2 = deck1[1:], deck2[1:]

	if len(deck1) >= draw1 && len(deck2) >= draw2 {
		_, winner := playRecursiveGame(deck1[0:draw1], deck2[0:draw2])
		if winner == 1 {
			deck1 = append(deck1, draw1, draw2)
		} else {
			deck2 = append(deck2, draw2, draw1)
		}
		return deck1, deck2, hist1, hist2, false
	} else if draw1 > draw2 {
		return append(deck1, draw1, draw2), deck2, hist1, hist2, false
	} else {
		return deck1, append(deck2, draw2, draw1), hist1, hist2, false
	}
}

func Challenge1(deck1 Deck, deck2 Deck) (int, error) {
	for len(deck1) != 0 && len(deck2) != 0 {
		c1, c2 := deck1[0], deck2[0]
		if c1 > c2 {
			deck1 = append(deck1[1:], c1, c2)
			deck2 = deck2[1:]
		} else {
			deck1 = deck1[1:]
			deck2 = append(deck2[1:], c2, c1)
		}
	}
	score := 0
	var winningDeck Deck = deck1
	if len(deck1) == 0 {
		winningDeck = deck2
	}
	deckSize := len(winningDeck)
	for i := 0; i < deckSize; i++ {
		score = score + ((deckSize - i) * winningDeck[i])
	}
	return score, nil
}

func Challenge2(deck1 Deck, deck2 Deck) (int, error) {
	winningDeck, _ := playRecursiveGame(deck1, deck2)
	score := 0
	deckSize := len(winningDeck)
	for i := 0; i < deckSize; i++ {
		score = score + ((deckSize - i) * winningDeck[i])
	}
	return score, nil
}

func Run() {
	fmt.Println("Day 22 - Crab Combat")
	path := "crabcombat/input.txt"
	deck1, deck2, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	var cpy1, cpy2 Deck
	cpy1 = make([]int, len(deck1))
	cpy2 = make([]int, len(deck2))

	copy(cpy1, deck1)
	copy(cpy2, deck2)
	result, err = Challenge1(cpy1, cpy2)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	copy(cpy1, deck1)
	copy(cpy2, deck2)
	result, err = Challenge2(cpy1, cpy2)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
