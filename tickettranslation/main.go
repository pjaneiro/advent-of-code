package tickettranslation

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Field struct {
	Name string
	Min1 int
	Max1 int
	Min2 int
	Max2 int
}

func readLines(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return [][]int{}, err
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, ",")
		var newline []int = make([]int, len(tmp))
		for i, j := range tmp {
			var err error
			newline[i], err = strconv.Atoi(j)
			if err != nil {
				return [][]int{}, err
			}
		}
		lines = append(lines, newline)
	}
	return lines, nil
}

func Challenge1(data [][]int, myTicket []int, fields []Field) (int, error) {
	result := 0
	for _, ticket := range data {
	ENTRIES:
		for _, entry := range ticket {
			for _, field := range fields {
				if (entry >= field.Min1 && entry <= field.Max1) || (entry >= field.Min2 && entry <= field.Max2) {
					continue ENTRIES
				}
			}
			result += entry
		}
	}
	return result, nil
}

func makeRange(length int) []int {
	a := make([]int, length)
	for i := range a {
		a[i] = i
	}
	return a
}

func Challenge2(data [][]int, myticket []int, fields []Field) (int, error) {
	result := 1
	for i := 0; i < len(data); i++ {
	ENTRIES:
		for _, entry := range data[i] {
			for _, field := range fields {
				if (entry >= field.Min1 && entry <= field.Max1) || (entry >= field.Min2 && entry <= field.Max2) {
					continue ENTRIES
				}
			}
			data = append(data[:i], data[i+1:]...)
			i--
			break
		}
	}

	data = append(data, myticket)

	options := make(map[string][]int, len(fields))
	for i := range fields {
		options[(fields[i]).Name] = makeRange(len(myticket))
	}

	for _, ticket := range data {
		for i, entry := range ticket {
			for _, field := range fields {
				if (entry >= field.Min1 && entry <= field.Max1) || (entry >= field.Min2 && entry <= field.Max2) {
					continue
				}
				n := 0
				for _, x := range options[field.Name] {
					if x != i {
						options[field.Name][n] = x
						n++
					}
				}
				options[field.Name] = options[field.Name][:n]
			}
		}
	}

	fixed := 0
	var indexes map[int]string = make(map[int]string, len(myticket))

	for fixed < len(myticket) {
		for key, value := range options {
			if len(value) == 1 {
				indexes[value[0]] = key
				fixed++
				delete(options, key)
				continue
			}
			n := 0
			for _, j := range value {
				if _, ok := indexes[j]; !ok {
					value[n] = j
					n++
				}
			}
			options[key] = options[key][:n]
		}
	}

	for key, value := range indexes {
		if strings.HasPrefix(value, "departure") {
			result *= myticket[key]
		}
	}

	return result, nil
}

func Run() {
	fmt.Println("Day 16")
	path := "tickettranslation/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	fields := []Field{
		Field{Name: "departure location", Min1: 26, Max1: 404, Min2: 427, Max2: 951},
		Field{Name: "departure station", Min1: 43, Max1: 307, Min2: 325, Max2: 967},
		Field{Name: "departure platform", Min1: 39, Max1: 383, Min2: 399, Max2: 950},
		Field{Name: "departure track", Min1: 31, Max1: 157, Min2: 178, Max2: 969},
		Field{Name: "departure date", Min1: 28, Max1: 109, Min2: 135, Max2: 950},
		Field{Name: "departure time", Min1: 38, Max1: 622, Min2: 631, Max2: 958},
		Field{Name: "arrival location", Min1: 35, Max1: 61, Min2: 69, Max2: 957},
		Field{Name: "arrival station", Min1: 36, Max1: 216, Min2: 241, Max2: 951},
		Field{Name: "arrival platform", Min1: 41, Max1: 586, Min2: 606, Max2: 967},
		Field{Name: "arrival track", Min1: 47, Max1: 573, Min2: 586, Max2: 951},
		Field{Name: "class", Min1: 31, Max1: 439, Min2: 445, Max2: 957},
		Field{Name: "duration", Min1: 35, Max1: 925, Min2: 939, Max2: 965},
		Field{Name: "price", Min1: 41, Max1: 473, Min2: 494, Max2: 952},
		Field{Name: "route", Min1: 45, Max1: 742, Min2: 754, Max2: 963},
		Field{Name: "row", Min1: 41, Max1: 338, Min2: 357, Max2: 952},
		Field{Name: "seat", Min1: 45, Max1: 848, Min2: 873, Max2: 968},
		Field{Name: "train", Min1: 37, Max1: 183, Min2: 197, Max2: 952},
		Field{Name: "type", Min1: 46, Max1: 509, Min2: 522, Max2: 974},
		Field{Name: "wagon", Min1: 32, Max1: 69, Min2: 81, Max2: 967},
		Field{Name: "zone", Min1: 37, Max1: 759, Min2: 780, Max2: 967},
	}
	myTicket := []int{103, 197, 83, 101, 109, 181, 61, 157, 199, 137, 97, 179, 151, 89, 211, 59, 139, 149, 53, 107}

	var result int
	result, err = Challenge1(data, myTicket, fields)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(data, myTicket, fields)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
