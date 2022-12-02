package rockpaperscissors

import (
	"bufio"
	"fmt"
	"os"
)

type Play struct {
	Them byte
	You  byte
}

func readLines(path string) ([]Play, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var plays []Play
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := Play{Them: line[0], You: line[2]}
		plays = append(plays, data)
	}

	return plays, nil
}

func Challenge1(data []Play) (int, error) {
	scores := map[Play]int{
		Play{Them: 'A', You: 'X'}: 3 + 1,
		Play{Them: 'A', You: 'Y'}: 6 + 2,
		Play{Them: 'A', You: 'Z'}: 0 + 3,
		Play{Them: 'B', You: 'X'}: 0 + 1,
		Play{Them: 'B', You: 'Y'}: 3 + 2,
		Play{Them: 'B', You: 'Z'}: 6 + 3,
		Play{Them: 'C', You: 'X'}: 6 + 1,
		Play{Them: 'C', You: 'Y'}: 0 + 2,
		Play{Them: 'C', You: 'Z'}: 3 + 3,
	}
	result := 0
	for _, play := range data {
		result += scores[play]
	}
	return result, nil
}

func Challenge2(data []Play) (int, error) {
	scores := map[Play]int{
		Play{Them: 'A', You: 'X'}: 0 + 3,
		Play{Them: 'A', You: 'Y'}: 3 + 1,
		Play{Them: 'A', You: 'Z'}: 6 + 2,
		Play{Them: 'B', You: 'X'}: 0 + 1,
		Play{Them: 'B', You: 'Y'}: 3 + 2,
		Play{Them: 'B', You: 'Z'}: 6 + 3,
		Play{Them: 'C', You: 'X'}: 0 + 2,
		Play{Them: 'C', You: 'Y'}: 3 + 3,
		Play{Them: 'C', You: 'Z'}: 6 + 1,
	}
	result := 0
	for _, play := range data {
		result += scores[play]
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 2 - Rock Paper Scissors")
	path := "rockpaperscissors/input.txt"
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
