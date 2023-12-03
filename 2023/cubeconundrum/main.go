package cubeconundrum

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	result := 0
OUTER1:
	for _, line := range data {
		parsedLine := strings.Split(line, ":")
		id, err := strconv.Atoi(parsedLine[0][5:])
		if err != nil {
			return 0, errors.New("error parsing input")
		}
		games := strings.Split(parsedLine[1], ";")
		for _, game := range games {
			plays := strings.Split(game, ",")
			for _, play := range plays {
				aux := strings.Split(strings.TrimSpace(play), " ")
				count, err := strconv.Atoi(aux[0])
				if err != nil {
					return 0, errors.New("error parsing input")
				}
				switch aux[1] {
				case "red":
					if count > 12 {
						continue OUTER1
					}
				case "green":
					if count > 13 {
						continue OUTER1
					}
				case "blue":
					if count > 14 {
						continue OUTER1
					}
				default:
					return 0, errors.New("error parsing input")
				}
			}
		}
		result += id
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	result := 0
	for _, line := range data {
		red, green, blue := 0, 0, 0
		parsedLine := strings.Split(line, ":")
		games := strings.Split(parsedLine[1], ";")
		for _, game := range games {
			plays := strings.Split(game, ",")
			for _, play := range plays {
				aux := strings.Split(strings.TrimSpace(play), " ")
				count, err := strconv.Atoi(aux[0])
				if err != nil {
					return 0, errors.New("error parsing input")
				}
				switch aux[1] {
				case "red":
					if count > red {
						red = count
					}
				case "green":
					if count > green {
						green = count
					}
				case "blue":
					if count > blue {
						blue = count
					}
				default:
					return 0, errors.New("error parsing input")
				}
			}
		}
		result += (red * green * blue)
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 1 - Cube Conundrum")
	path := "cubeconundrum/input.txt"
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
