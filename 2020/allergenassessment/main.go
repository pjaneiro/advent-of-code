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

func contains(list []string, val string) bool {
	for _, cur := range list {
		if cur == val {
			return true
		}
	}
	return false
}

func intersection(s1 []string, s2 []string) []string {
	var result []string = make([]string, 0)
	for _, cur := range s1 {
		if contains(s2, cur) {
			result = append(result, cur)
		}
	}
	return result
}

func remove(list []string, val string) []string {
	var result []string = make([]string, 0)
	for _, cur := range list {
		if cur != val {
			result = append(result, cur)
		}
	}
	return result
}

func solve(data []Recipe) map[string]string {
	var mapping map[string][]string = make(map[string][]string)
	var assigns map[string]string = make(map[string]string)

	for _, cur := range data {
		for _, alg := range cur.Allergens {
			if _, ok := mapping[alg]; !ok {
				mapping[alg] = make([]string, len(cur.Ingredients))
				copy(mapping[alg], cur.Ingredients)
			} else {
				mapping[alg] = intersection(mapping[alg], cur.Ingredients)
			}
		}
	}
	for len(mapping) > 0 {
		for alg, ing := range mapping {
			if len(ing) == 1 {
				assigns[ing[0]] = alg

				for k, v := range mapping {
					mapping[k] = remove(v, ing[0])
				}

				delete(mapping, alg)
				break
			}
		}
	}
	return assigns
}

func Challenge1(data []Recipe) (int, error) {
	assigns := solve(data)
	count := 0
	for _, v := range data {
		for _, ing := range v.Ingredients {
			if _, ok := assigns[ing]; !ok {
				count++
			}
		}
	}
	return count, nil
}

func Challenge2(data []Recipe) (string, error) {
	assigns := solve(data)
	var reverse map[string]string = make(map[string]string)
	var ingredients []string = make([]string, 0)
	var result []string = make([]string, 0)

	for alg, ing := range assigns {
		reverse[ing] = alg
		ingredients = append(ingredients, ing)
	}
	sort.Strings(ingredients)
	for _, ing := range ingredients {
		result = append(result, reverse[ing])
	}
	return strings.Join(result, ","), nil
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
