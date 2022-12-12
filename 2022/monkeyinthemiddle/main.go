package monkeyinthemiddle

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items       []int
	Operator    rune
	Factor      int
	Test        int
	True, False int
	Inspected   int
}

func readLines(path string) ([]Monkey, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rItems, err := regexp.Compile(`^\s*Starting items: (.*)$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}
	rOperation, err := regexp.Compile(`^\s*Operation: new = old ([\+\*]) ((?:old)|\d+)$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}
	rTest, err := regexp.Compile(`^\s*Test: divisible by (\d+)$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}
	rTrue, err := regexp.Compile(`^\s*If true: throw to monkey (\d+)$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}
	rFalse, err := regexp.Compile(`^\s*If false: throw to monkey (\d+)$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}

	scanner := bufio.NewScanner(file)
	var result []Monkey
	for scanner.Scan() {
		var monkey Monkey
		scanner.Text()

		scanner.Scan()
		lineItems := rItems.FindStringSubmatch(scanner.Text())
		for _, item := range strings.Split(lineItems[1], ", ") {
			var val int
			if val, err = strconv.Atoi(item); err != nil {
				return nil, errors.New("couldn't parse item list")
			}
			monkey.Items = append(monkey.Items, val)
		}

		scanner.Scan()
		lineOperation := rOperation.FindStringSubmatch(scanner.Text())
		monkey.Operator = []rune(lineOperation[1])[0]
		if lineOperation[2] == "old" {
			monkey.Factor = -1
		} else {
			if monkey.Factor, err = strconv.Atoi(lineOperation[2]); err != nil {
				return nil, errors.New("couldn't parse operation")
			}
		}

		scanner.Scan()
		lineTest := rTest.FindStringSubmatch(scanner.Text())
		if monkey.Test, err = strconv.Atoi(lineTest[1]); err != nil {
			return nil, errors.New("couldn't parse test")
		}

		scanner.Scan()
		lineTrue := rTrue.FindStringSubmatch(scanner.Text())
		if monkey.True, err = strconv.Atoi(lineTrue[1]); err != nil {
			return nil, errors.New("couldn't parse true")
		}

		scanner.Scan()
		lineFalse := rFalse.FindStringSubmatch(scanner.Text())
		if monkey.False, err = strconv.Atoi(lineFalse[1]); err != nil {
			return nil, errors.New("couldn't parse false")
		}

		result = append(result, monkey)
		scanner.Scan()
	}

	return result, nil
}

func solve(data []Monkey, nRounds int, part int) (int, error) {
	multiple := 1
	for _, monkey := range data {
		multiple *= monkey.Test
	}
	for round := 0; round < nRounds; round++ {
		for monkeyId, monkey := range data {
			for len(data[monkeyId].Items) > 0 {
				item := data[monkeyId].Items[0]
				data[monkeyId].Items = data[monkeyId].Items[1:]
				data[monkeyId].Inspected++
				worryStart := item
				var newWorry int
				switch monkey.Operator {
				case '+':
					switch monkey.Factor {
					case -1:
						newWorry = worryStart + worryStart
					default:
						newWorry = worryStart + monkey.Factor
					}
				case '*':
					switch monkey.Factor {
					case -1:
						newWorry = worryStart * worryStart
					default:
						newWorry = worryStart * monkey.Factor
					}
				default:
					return 0, errors.New("invalid operation")
				}
				var bored int
				switch part {
				case 1:
					bored = newWorry / 3
				case 2:
					bored = newWorry % multiple
				}
				if bored%monkey.Test == 0 {
					data[monkey.True].Items = append(data[monkey.True].Items, bored)
				} else {
					data[monkey.False].Items = append(data[monkey.False].Items, bored)
				}
			}
		}
	}

	var inspections []int = make([]int, 0, len(data))
	for _, monkey := range data {
		inspections = append(inspections, monkey.Inspected)
	}

	sort.Ints(inspections)

	return inspections[len(inspections)-1] * inspections[len(inspections)-2], nil
}

func Challenge1(data []Monkey) (int, error) {
	return solve(data, 20, 1)
}

func Challenge2(data []Monkey) (int, error) {
	return solve(data, 10000, 2)
}

func Run() {
	fmt.Println("Day 11 - Monkey in the Middle")
	path := "monkeyinthemiddle/input.txt"

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

	data, err = readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	result, err = Challenge2(data)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
