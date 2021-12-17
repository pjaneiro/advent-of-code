package allergenassessment

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Recipe struct {
	Ingredients []string
	Allergens   []string
}

func readLines(path string) ([]Recipe, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r, err := regexp.Compile(`^(\w+(?: \w+)*) \(contains (\w+(?:, \w+)*)\)$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}

	var result []Recipe = make([]Recipe, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var recipe Recipe
		if r.MatchString(line) {
			data := r.FindStringSubmatch(line)
			recipe.Ingredients = strings.Fields(data[1])
			recipe.Allergens = strings.Split(data[2], ", ")
			result = append(result, recipe)
		} else {
			fmt.Printf("Failed parsing: %s\n", line)
		}

	}

	return result, nil
}

func Challenge1(data []Recipe) (int, error) {
	var mapping map[string]map[string]int = make(map[string]map[string]int)
	var assigns map[string]string = make(map[string]string)

	for _, cur := range data {
		for _, alg := range cur.Allergens {
			if _, ok := mapping[alg]; !ok {
				mapping[alg] = make(map[string]int)
			}
			for _, ing := range cur.Ingredients {
				mapping[alg][ing]++
			}
		}
	}
	for kAlg, vAlg := range mapping {
		max := 0
		best := ""
		for kIng, vIng := range vAlg {
			if vIng > max {
				if _, ok := assigns[kIng]; !ok {
					max = vIng
					best = kIng
				}
			}
		}
		assigns[best] = kAlg
	}
	count := 0
	for _, cur := range data {
		for _, ing := range cur.Ingredients {
			if _, ok := assigns[ing]; !ok {
				count++
			}
		}
	}
	return count, nil
}

func Challenge2(data []Recipe) (string, error) {
	var mapping map[string]map[string]int = make(map[string]map[string]int)
	var assigns map[string]string = make(map[string]string)
	var reverseAssigns map[string]string = make(map[string]string)

	for _, cur := range data {
		for _, alg := range cur.Allergens {
			if _, ok := mapping[alg]; !ok {
				mapping[alg] = make(map[string]int)
			}
			for _, ing := range cur.Ingredients {
				mapping[alg][ing]++
			}
		}
	}
	for kAlg, vAlg := range mapping {
		max := 0
		best := ""
		for kIng, vIng := range vAlg {
			if vIng > max {
				if _, ok := assigns[kIng]; !ok {
					max = vIng
					best = kIng
				}
			}
		}
		assigns[best] = kAlg
		reverseAssigns[kAlg] = best
	}
	var keys []string = make([]string, 0)
	var values []string = make([]string, 0)

	for k, _ := range reverseAssigns {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		values = append(values, reverseAssigns[k])
	}
	return strings.Join(values, ","), nil
}

func Run() {
	fmt.Println("Day 21 - Allergen Assessment")
	path := "allergenassessment/input.txt"
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

	var result2 string
	result2, err = Challenge2(data)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %s\n", result2)
	}
}
