package rambunctiousrecitation

import (
	"fmt"
)

func Challenge1(data []int) (int, error) {
	turn, current := 0, 0
	var ages map[int][]int = make(map[int][]int, 0)

	for turn < 2020 {
		if turn < len(data) {
			current = data[turn]
			ages[current] = []int{turn}
		} else {
			age := ages[current]
			if len(age) == 1 {
				current = 0
				ages[0] = append(ages[0], turn)
			} else {
				diff := age[len(age)-1] - age[len(age)-2]
				current = diff
				if mem, ok := ages[current]; ok {
					ages[current] = append(mem, turn)
				} else {
					ages[current] = []int{turn}
				}
			}
		}

		// fmt.Printf("Played a %d\n", current)
		turn++
	}
	return current, nil
}

func Challenge2(data []int) (int, error) {
	turn, current := 0, 0
	var ages map[int][]int = make(map[int][]int, 0)

	for turn < 30000000 {
		if turn < len(data) {
			current = data[turn]
			ages[current] = []int{turn}
		} else {
			age := ages[current]
			if len(age) == 1 {
				current = 0
				ages[0] = append(ages[0], turn)
			} else {
				diff := age[len(age)-1] - age[len(age)-2]
				current = diff
				if mem, ok := ages[current]; ok {
					ages[current] = append(mem, turn)
				} else {
					ages[current] = []int{turn}
				}
			}
		}

		// fmt.Printf("Played a %d\n", current)
		turn++
	}
	return current, nil
}

func Run() {
	fmt.Println("Day 15 - Rambunctious Recitation")
	data := []int{0, 1, 5, 10, 3, 12, 19}

	var result int
	var err error
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
