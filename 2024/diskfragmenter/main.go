package diskfragmenter

import (
	"bufio"
	"fmt"
	"os"
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
	var disk []int = make([]int, 0)
	for i := 0; i < len(data[0]); i++ {
		if i%2 == 0 {
			id := i / 2
			for j := 0; j < int(data[0][i]-48); j++ {
				disk = append(disk, id)
			}
		} else {
			for j := 0; j < int(data[0][i]-48); j++ {
				disk = append(disk, -1)
			}
		}
	}
	emptyIndex := 0
	for i := len(disk) - 1; i > emptyIndex+1; i-- {
		if disk[i] == -1 {
			continue
		}
		for ; disk[emptyIndex] != -1; emptyIndex++ {
		}
		disk[emptyIndex] = disk[i]
		disk[i] = -1
	}
	result := 0
	for i := 0; i < len(disk) && disk[i] != -1; i++ {
		result += (i * disk[i])
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	var disk []int = make([]int, 0)
	for i := 0; i < len(data[0]); i++ {
		if i%2 == 0 {
			id := i / 2
			for j := 0; j < int(data[0][i]-48); j++ {
				disk = append(disk, id)
			}
		} else {
			for j := 0; j < int(data[0][i]-48); j++ {
				disk = append(disk, -1)
			}
		}
	}
	var ordered map[int]struct{} = map[int]struct{}{}
	for i := len(disk) - 1; i > 0; i-- {
		if disk[i] == -1 {
			continue
		}
		if _, ok := ordered[disk[i]]; ok {
			continue
		}
		id, j, length := disk[i], i-1, 1
		for ; j >= 0 && disk[j] == id; j, length = j-1, length+1 {
		}
		var emptyIndex int
		var foundSpace bool = false
	outer:
		for emptyIndex = 0; emptyIndex < i && !foundSpace; emptyIndex++ {
			for j := 0; j < length && emptyIndex+j < len(disk); j++ {
				if disk[emptyIndex+j] != -1 {
					emptyIndex += j
					continue outer
				}
			}
			foundSpace = true
			break
		}
		if !foundSpace {
			ordered[id] = struct{}{}
			continue
		}
		for j := 0; j < length; j++ {
			ordered[id] = struct{}{}
			disk[emptyIndex+j] = id
			disk[i-j] = -1
		}
	}
	result := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			continue
		}
		result += (i * disk[i])
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 9 - Disk Fragmenter")
	path := "diskfragmenter/input.txt"
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
