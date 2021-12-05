package giantsquid

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type House struct {
	Number int
	Marked bool
}

func readLines(path string) ([]int, [][][]House, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var order []int
	var boards [][][]House
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	orderLine := strings.Split(scanner.Text(), ",")
	for _, cur := range orderLine {
		val, err := strconv.Atoi(cur)
		if err != nil {
			return nil, nil, err
		}
		order = append(order, val)
	}

	for scanner.Scan() {
		var board [][]House
		for i := 0; i < 5; i++ {
			var boardLine []House
			scanner.Scan()
			line := strings.Fields(scanner.Text())
			for _, cur := range line {
				val, err := strconv.Atoi(cur)
				if err != nil {
					return nil, nil, err
				}
				boardLine = append(boardLine, House{val, false})
			}
			board = append(board, boardLine)
		}
		boards = append(boards, board)
	}

	return order, boards, nil
}

func winningBoard(board [][]House) bool {
	for i := 0; i < 5; i++ {
		if board[0][i].Marked && board[1][i].Marked && board[2][i].Marked && board[3][i].Marked && board[4][i].Marked {
			return true
		}
		if board[i][0].Marked && board[i][1].Marked && board[i][2].Marked && board[i][3].Marked && board[i][4].Marked {
			return true
		}
	}
	return false
}

func markHouseOnBoard(board [][]House, number int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j].Number == number {
				board[i][j].Marked = true
				return true
			}
		}
	}
	return false
}

func Challenge1(order []int, boards [][][]House) (int, error) {
	score := 0
	for index := 0; index < len(order); index++ {
		curNumber := order[index]
		for _, curBoard := range boards {
			markHouseOnBoard(curBoard, curNumber)
			if winningBoard(curBoard) {
				for i := 0; i < 5; i++ {
					for j := 0; j < 5; j++ {
						if !curBoard[i][j].Marked {
							score = score + curBoard[i][j].Number
						}
					}
				}
				return score * curNumber, nil
			}
		}
	}
	return 0, errors.New("no winning board was found")
}

func Challenge2(order []int, boards [][][]House) (int, error) {
	score, closed := 0, 0
	statuses := make([]bool, len(boards))
	for index := 0; index < len(order); index++ {
		curNumber := order[index]
		for k, curBoard := range boards {
			markHouseOnBoard(curBoard, curNumber)
			if !statuses[k] && winningBoard(curBoard) {
				statuses[k] = true
				closed++
			}
			if closed == len(boards) {
				for i := 0; i < 5; i++ {
					for j := 0; j < 5; j++ {
						if !curBoard[i][j].Marked {
							score = score + curBoard[i][j].Number
						}
					}
				}
				return score * curNumber, nil
			}
		}
	}
	return 0, errors.New("no losing board was found")
}

func Run() {
	fmt.Println("Day 4 - Giant Squid")
	path := "giantsquid/input.txt"
	order, boards, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var cpy [][][]House
	var result int

	cpy = make([][][]House, len(boards))
	copy(cpy, boards)
	result, err = Challenge1(order, cpy)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	cpy = make([][][]House, len(boards))
	copy(cpy, boards)
	result, err = Challenge2(order, cpy)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
