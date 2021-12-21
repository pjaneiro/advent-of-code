package diracdice

import (
	"errors"
	"fmt"
)

type State struct {
	player int
	pos1   int
	pos2   int
	score1 int
	score2 int
}

func Challenge1(player1 int, player2 int) (int, error) {
	pos1, pos2, current, score1, score2, rolls, rolled := player1, player2, 1, 0, 0, 0, 0

	for score1 < 1000 && score2 < 1000 {
		rolled = (rolled % 100) + 1
		roll1 := rolled
		rolled = (rolled % 100) + 1
		roll2 := rolled
		rolled = (rolled % 100) + 1
		roll3 := rolled
		rolls = rolls + 3
		if current == 1 {
			pos1 = (pos1+roll1+roll2+roll3-1)%10 + 1
			score1 = score1 + pos1
			if score1 >= 1000 {
				return rolls * score2, nil
			}
			current = 2
		} else {
			pos2 = (pos2+roll1+roll2+roll3-1)%10 + 1
			score2 = score2 + pos2
			if score2 >= 1000 {
				return rolls * score1, nil
			}
			current = 1
		}
	}

	return 0, errors.New("something went wrong")
}

func solve(states map[State][]int64, player int, pos1 int, pos2 int, score1 int, score2 int) []int64 {
	if score1 >= 21 {
		return []int64{1, 0}
	}
	if score2 >= 21 {
		return []int64{0, 1}
	}
	if cur, ok := states[State{player: player, pos1: pos1, pos2: pos2, score1: score1, score2: score2}]; ok {
		return cur
	}

	var result []int64 = make([]int64, 2)
	for die1 := 1; die1 < 4; die1++ {
		for die2 := 1; die2 < 4; die2++ {
			for die3 := 1; die3 < 4; die3++ {
				if player == 1 {
					newPos1 := (pos1+die1+die2+die3-1)%10 + 1
					newScore1 := score1 + newPos1
					subResult := solve(states, 2, newPos1, pos2, newScore1, score2)
					result[0] = result[0] + subResult[0]
					result[1] = result[1] + subResult[1]
				} else {
					newPos2 := (pos2+die1+die2+die3-1)%10 + 1
					newScore2 := score2 + newPos2
					subResult := solve(states, 1, pos1, newPos2, score1, newScore2)
					result[0] = result[0] + subResult[0]
					result[1] = result[1] + subResult[1]
				}
			}
		}
	}
	states[State{player: player, pos1: pos1, pos2: pos2, score1: score1, score2: score2}] = result
	return result
}

func Challenge2(player1 int, player2 int) (int64, error) {
	var states map[State][]int64 = make(map[State][]int64)
	result := solve(states, 1, player1, player2, 0, 0)
	if result[0] > result[1] {
		return result[0], nil
	}
	return result[1], nil
}

func Run() {
	fmt.Println("Day 21 - Dirac Dice")

	result1, err := Challenge1(8, 5)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result1)
	}

	result2, err := Challenge2(8, 5)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result2)
	}
}
