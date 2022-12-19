package proboscideavolcanium

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Valve struct {
	Name     string
	FlowRate int
	LeadsTo  []string
}

func ReadLines(path string) ([]*Valve, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r, err := regexp.Compile(`^Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (\w+(?:, \w+)*)$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}

	var result []*Valve = make([]*Valve, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := r.FindStringSubmatch(scanner.Text())
		var valve Valve
		valve.Name = line[1]
		if valve.FlowRate, err = strconv.Atoi(line[2]); err != nil {
			return nil, errors.New("error parsing input")
		}
		valve.LeadsTo = strings.Split(line[3], ", ")
		result = append(result, &valve)
	}

	return result, nil
}

func contains[T string | *Valve](slice []T, key T) bool {
	for _, cur := range slice {
		if cur == key {
			return true
		}
	}
	return false
}

func floydwarshall(data []*Valve) map[*Valve]map[*Valve]int {
	var distances map[*Valve]map[*Valve]int = make(map[*Valve]map[*Valve]int)
	for _, valve1 := range data {
		distances[valve1] = make(map[*Valve]int)
		for _, valve2 := range data {
			distances[valve1][valve2] = int(math.MaxInt32)
		}
	}
	for _, valve1 := range data {
		for _, valve2 := range data {
			if valve1 == valve2 {
				distances[valve1][valve2] = 0
			}
			if contains(valve1.LeadsTo, valve2.Name) {
				distances[valve1][valve2] = 1
			}
		}
	}
	for k := 0; k < len(data); k++ {
		for i := 0; i < len(data); i++ {
			for j := 0; j < len(data); j++ {
				if distances[data[i]][data[j]] > distances[data[i]][data[k]]+distances[data[k]][data[j]] {
					distances[data[i]][data[j]] = distances[data[i]][data[k]] + distances[data[k]][data[j]]
				}
			}
		}
	}
	return distances
}

func traversePart1(curNode *Valve, curPressure int, curMinute int, distances map[*Valve]map[*Valve]int, toVisit []*Valve, opened []*Valve) int {
	max := curPressure
	for _, nextNode := range toVisit {
		if nextNode == curNode || nextNode.Name == "AA" || contains(opened, nextNode) {
			continue
		}
		dist := distances[curNode][nextNode] + 1
		if curMinute+dist > 30 {
			continue
		}
		tmp := traversePart1(nextNode, curPressure+(30-curMinute-dist)*nextNode.FlowRate, curMinute+dist, distances, toVisit, append(opened, nextNode))
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func traversePart2(curNode1 *Valve, curNode2 *Valve, curPressure int, curMinute1 int, curMinute2 int, distances map[*Valve]map[*Valve]int, toVisit []*Valve, opened []*Valve) int {
	max := curPressure
	var tmp int
	for _, nextNode := range toVisit {
		if nextNode == curNode1 || nextNode.Name == "AA" || contains(opened, nextNode) {
			continue
		}

		dist1 := distances[curNode1][nextNode] + 1
		if curMinute1+dist1 > 26 {
			continue
		}
		tmp = traversePart2(nextNode, curNode2, curPressure+(26-curMinute1-dist1)*nextNode.FlowRate, curMinute1+dist1, curMinute2, distances, toVisit, append(opened, nextNode))
		if tmp > max {
			max = tmp
		}

		dist2 := distances[curNode2][nextNode] + 1
		if curMinute2+dist2 > 26 {
			continue
		}
		tmp = traversePart2(curNode1, nextNode, curPressure+(26-curMinute2-dist2)*nextNode.FlowRate, curMinute1, curMinute2+dist2, distances, toVisit, append(opened, nextNode))
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func Challenge1(data []*Valve) (int, error) {
	distances := floydwarshall(data)
	var toVisit []*Valve = make([]*Valve, 0)
	var opened []*Valve = make([]*Valve, 0)
	var start *Valve

	for _, valve := range data {
		if valve.FlowRate > 0 || valve.Name == "AA" {
			toVisit = append(toVisit, valve)
		}
		if valve.Name == "AA" {
			start = valve
		}
	}

	return traversePart1(start, 0, 0, distances, toVisit, opened), nil
}

func Challenge2(data []*Valve) (int, error) {
	distances := floydwarshall(data)
	var toVisit []*Valve = make([]*Valve, 0)
	var opened []*Valve = make([]*Valve, 0)
	var start *Valve

	for _, valve := range data {
		if valve.FlowRate > 0 || valve.Name == "AA" {
			toVisit = append(toVisit, valve)
		}
		if valve.Name == "AA" {
			start = valve
		}
	}

	return traversePart2(start, start, 0, 0, 0, distances, toVisit, opened), nil
}

func Run() {
	fmt.Println("Day 16 - Proboscidea Volcanium")
	path := "proboscideavolcanium/input.txt"

	data, err := ReadLines(path)
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
