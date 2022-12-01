package combobreaker

import (
	"errors"
	"fmt"
)

func getLoopSize(key int) int {
	value, result := 1, 0
	for value != key {
		result++
		value = (value * 7) % 20201227
	}
	return result
}

func Challenge1(cardPublicKey int, doorPublicKey int) (int, error) {
	cardLoopSize, doorLoopSize := getLoopSize(cardPublicKey), getLoopSize(doorPublicKey)

	cardKey, doorKey := 1, 1
	for i := 0; i < cardLoopSize; i++ {
		cardKey = (cardKey * doorPublicKey) % 20201227
	}
	for i := 0; i < doorLoopSize; i++ {
		doorKey = (doorKey * cardPublicKey) % 20201227
	}

	if cardKey != doorKey {
		return 0, errors.New("keys don't match")
	}

	return cardKey, nil
}

func Run() {
	fmt.Println("Day 25 - Combo Breaker")
	var cardPK, doorPK int = 19241437, 17346587

	var result int
	var err error
	result, err = Challenge1(cardPK, doorPK)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}
}
