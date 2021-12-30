package spacestoichiometry

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Chemical struct {
	Name     string
	Quantity int
}

type Reaction struct {
	Input  []Chemical
	Output Chemical
}

type ReactionList []Reaction

func (d ReactionList) reactionFromOutputName(name string) Reaction {
	for _, cur := range d {
		if cur.Output.Name == name {
			return cur
		}
	}
	return Reaction{}
}

func readLines(path string) (ReactionList, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r, err := regexp.Compile(`(?:(\d+) (\w+))`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}

	var result ReactionList = make(ReactionList, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var reac Reaction
		reac.Input = make([]Chemical, 0)

		data := r.FindAllStringSubmatch(line, -1)
		for i := 0; i < len(data)-1; i++ {
			c := Chemical{}
			c.Name = data[i][2]
			c.Quantity, _ = strconv.Atoi(data[i][1])
			reac.Input = append(reac.Input, c)
		}
		reac.Output.Name = data[len(data)-1][2]
		reac.Output.Quantity, _ = strconv.Atoi(data[len(data)-1][1])
		result = append(result, reac)
	}

	return result, nil
}

func minOre(data ReactionList, fuel int) int {
	var reqs map[string]int = make(map[string]int)
	reqs["FUEL"] = fuel

	var done bool = false

	for !done {
		done = true

		for key, val := range reqs {
			if key == "ORE" || val <= 0 {
				continue
			}

			done = false

			var reaction Reaction = data.reactionFromOutputName(key)

			times := int(math.Ceil(float64(val) / float64(reaction.Output.Quantity)))

			for _, cur := range reaction.Input {
				reqs[cur.Name] = reqs[cur.Name] + (times * cur.Quantity)
			}

			reqs[key] = reqs[key] - (times * reaction.Output.Quantity)
		}
	}

	return reqs["ORE"]
}

func Challenge1(data ReactionList) (int, error) {
	return minOre(data, 1), nil
}

func Challenge2(data ReactionList) (int, error) {
	var lowerBound int = int(math.Floor(float64(1000000000000) / float64(minOre(data, 1))))
	var upperBound int = 2 * lowerBound
	for upperBound-lowerBound > 1 {
		diff := upperBound - lowerBound
		fuel := lowerBound + (diff / 2)
		tmp := minOre(data, fuel)

		if tmp > 1000000000000 {
			upperBound = fuel
			continue
		} else if tmp < 1000000000000 {
			lowerBound = fuel
			continue
		} else {
			return tmp, nil
		}
	}
	return lowerBound, nil
}

func Run() {
	fmt.Println("Day 14 - Space Stoichiometry")
	path := "spacestoichiometry/input.txt"
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
