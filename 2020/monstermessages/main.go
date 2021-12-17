package monstermessages

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func readLines(path string) (map[string]string, []string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var rules map[string]string = make(map[string]string)
	var messages []string

	r, err := regexp.Compile(`^(\d+): ((?:\d+|\"\w\")(?: (?:\d+|\"\w\"|\|))*)$`)
	if err != nil {
		fmt.Println(err)
		return nil, nil, errors.New("couldn't parse regex expression")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if r.MatchString(line) {
			data := r.FindStringSubmatch(line)
			rules[data[1]] = data[2]
		} else if len(line) == 0 {
			continue
		} else {
			line := scanner.Text()
			messages = append(messages, line)
		}
	}

	return rules, messages, nil
}

func validate(messages []string, rules map[string]string, index string) (bool, []string) {
	rule := rules[index]
	if strings.Contains(rule, "\"") {
		var remainders []string
		for _, cur := range messages {
			if len(cur) != 0 && cur[0:1] == rule[1:2] {
				remainders = append(remainders, cur[1:])
			}
		}
		return len(remainders) > 0, remainders
	} else {
		variants := strings.Split(rule, " | ")
		var rest, rests []string
		var valid bool
		for _, curVariant := range variants {
			newMessages := messages
			steps := strings.Fields(curVariant)
			for _, curStep := range steps {
				valid, rest = validate(newMessages, rules, curStep)
				if !valid {
					break
				}
				newMessages = rest
			}
			if valid {
				rests = append(rests, newMessages...)
			}
		}
		if len(rests) != 0 {
			return true, rests
		}
	}
	return false, messages
}

func Challenge1(rules map[string]string, messages []string) (int, error) {
	count := 0
	for _, cur := range messages {
		if len(cur) == 0 {
			continue
		}
		if valid, remainders := validate([]string{cur}, rules, "0"); valid {
			for _, curRemainder := range remainders {
				if len(curRemainder) == 0 {
					count++
				}
			}
		}
	}
	return count, nil
}

func Challenge2(rules map[string]string, messages []string) (int, error) {
	rules["8"] = "42 | 42 8"
	rules["11"] = "42 31 | 42 11 31"
	count := 0
	for _, cur := range messages {
		if len(cur) == 0 {
			continue
		}
		if valid, remainders := validate([]string{cur}, rules, "0"); valid {
			for _, curRemainder := range remainders {
				if len(curRemainder) == 0 {
					count++
				}
			}
		}
	}
	return count, nil
}

func Run() {
	fmt.Println("Day 19 - Monster Messages")
	path := "monstermessages/input.txt"
	rules, messages, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(rules, messages)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(rules, messages)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
