package reposerecord

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Entry struct {
	Year, Month, Day int
	Hour, Minute     int
	Command          string
}

func readLines(path string) ([]Entry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r, err := regexp.Compile(`^\[(\d+)\-(\d+)-(\d+) (\d+):(\d+)\] (.+)$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sort.Strings(lines)

	var result []Entry
	for _, line := range lines {
		data := r.FindStringSubmatch(line)
		var entry Entry
		if entry.Year, err = strconv.Atoi(data[1]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if entry.Month, err = strconv.Atoi(data[2]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if entry.Day, err = strconv.Atoi(data[3]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if entry.Hour, err = strconv.Atoi(data[4]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		if entry.Minute, err = strconv.Atoi(data[5]); err != nil {
			return nil, errors.New("couldn't parse input")
		}
		entry.Command = data[6]

		result = append(result, entry)
	}

	return result, nil
}

func getGuardsSleepingMap(data []Entry) (map[int][]int, error) {
	r, err := regexp.Compile(`^Guard #(\d+) begins shift$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}

	var mapping map[int][]int = make(map[int][]int)

	cur := 0

	for k, v := range data {
		switch v.Command {
		case "falls asleep":
			continue
		case "wakes up":
			prev := data[k-1]
			for i := prev.Minute; i < v.Minute; i++ {
				mapping[cur][i]++
			}
		default:
			parsed := r.FindStringSubmatch(v.Command)
			if cur, err = strconv.Atoi(parsed[1]); err != nil {
				return nil, errors.New("something went wrong")
			}
			if _, ok := mapping[cur]; !ok {
				mapping[cur] = make([]int, 60)
			}
		}
	}

	return mapping, nil
}

func Challenge1(data []Entry) (int, error) {
	mapping, err := getGuardsSleepingMap(data)
	if err != nil {
		return 0, err
	}

	chosen_guard, chosen_minute, max := 0, 0, 0

	for guard_id, v := range mapping {
		guard_total := 0
		guard_max := 0
		guard_minute := 0
		for minute, count := range v {
			guard_total += count
			if count > guard_max {
				guard_max = count
				guard_minute = minute
			}
		}
		if guard_total > max {
			max = guard_total
			chosen_guard = guard_id
			chosen_minute = guard_minute
		}
	}

	return chosen_guard * chosen_minute, nil
}

func Challenge2(data []Entry) (int, error) {
	mapping, err := getGuardsSleepingMap(data)
	if err != nil {
		return 0, err
	}

	chosen_guard, chosen_minute, max := 0, 0, 0

	for guard_id, v := range mapping {
		for minute, count := range v {
			if count > max {
				max = count
				chosen_guard = guard_id
				chosen_minute = minute
			}
		}
	}

	return chosen_guard * chosen_minute, nil
}

func Run() {
	fmt.Println("Day 4 - Repose Record")
	path := "reposerecord/input.txt"
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
