package subterraneansustainability

import (
	"bufio"
	"fmt"
	"os"
)

type Rule struct {
	Plants   [5]bool
	AddPlant bool
}

func readLines(path string) ([]bool, []Rule, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var state []bool
	var rules []Rule

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	stateLine := scanner.Text()
	for i := 15; i < len([]rune(stateLine)); i++ {
		if stateLine[i] == '#' {
			state = append(state, true)
		} else {
			state = append(state, false)
		}
	}

	scanner.Scan()
	for scanner.Scan() {
		ruleLine := []rune(scanner.Text())
		var rule Rule
		for i := 0; i < 5; i++ {
			if ruleLine[i] == '#' {
				rule.Plants[i] = true
			} else {
				rule.Plants[i] = false
			}
		}
		if ruleLine[9] == '#' {
			rule.AddPlant = true
		} else {
			rule.AddPlant = false
		}
		rules = append(rules, rule)
	}

	return state, rules, nil
}

func Challenge1(state []bool, rules []Rule) (int, error) {
	var ruleMap map[[5]bool]bool = make(map[[5]bool]bool)
	for _, rule := range rules {
		ruleMap[rule.Plants] = rule.AddPlant
	}
	var curState []bool = make([]bool, len(state))
	copy(curState, state)

	start := 0
	for gen := 1; gen <= 20; gen++ {
		curState = append([]bool{false, false, false, false}, curState...)
		curState = append(curState, false, false, false, false)
		newState := make([]bool, len(curState))
		start -= 4
		for i := 0; i < len(curState)-4; i++ {
			var tmp [5]bool
			copy(tmp[:], curState[i:i+5])
			if add, ok := ruleMap[tmp]; ok {
				newState[i+2] = add
			}
		}
		curState = newState
	}
	result := 0
	for i := 0; i < len(curState); i++ {
		if curState[i] {
			result += (i + start)
		}
	}
	return result, nil
}

func Challenge2(state []bool, rules []Rule) (int, error) {
	var ruleMap map[[5]bool]bool = make(map[[5]bool]bool)
	for _, rule := range rules {
		ruleMap[rule.Plants] = rule.AddPlant
	}
	var curState []bool = make([]bool, len(state))
	copy(curState, state)
	prevResult, curResult, nAttempts := 0, 0, 150

	start := 0
	for gen := 1; gen <= nAttempts; gen++ {
		curState = append([]bool{false, false, false, false}, curState...)
		curState = append(curState, false, false, false, false)
		newState := make([]bool, len(curState))
		start -= 4
		for i := 0; i < len(curState)-4; i++ {
			var tmp [5]bool
			copy(tmp[:], curState[i:i+5])
			if add, ok := ruleMap[tmp]; ok {
				newState[i+2] = add
			}
		}
		curState = newState
		result := 0
		for i := 0; i < len(curState); i++ {
			if curState[i] {
				result += (i + start)
			}
		}
		prevResult, curResult = curResult, result
	}
	result := curResult + (50000000000-nAttempts)*(curResult-prevResult)
	return result, nil
}

func Run() {
	fmt.Println("Day 12 - Subterranean Sustainability")
	path := "subterraneansustainability/input.txt"
	state, rules, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(state, rules)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %v\n", result)
	}

	result, err = Challenge2(state, rules)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %v\n", result)
	}
}
