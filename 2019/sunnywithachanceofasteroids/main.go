package sunnywithachanceofasteroids

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const OUTPUT = false

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, v := range strings.Split(line, ",") {
			n, err := strconv.Atoi(v)
			if err != nil {
				return lines, errors.New("something went wrong scanning int")
			}
			lines = append(lines, n)
		}
	}

	return lines, nil
}

func Challenge1(data []int, input []int) (int, error) {
	ip, ii, out := 0, 0, 0
	for true {
		opcode, instruction := data[ip]%100, data[ip]/100
		parammodes := []int{}
		for instruction > 0 {
			parammodes = append(parammodes, instruction%10)
			instruction /= 10
		}
		for len(parammodes) < 3 {
			parammodes = append(parammodes, 0)
		}
		switch opcode {
		case 1:
			val1, val2 := 0, 0
			if parammodes[0] == 0 {
				val1 = data[data[ip+1]]
			} else {
				val1 = data[ip+1]
			}
			if parammodes[1] == 0 {
				val2 = data[data[ip+2]]
			} else {
				val2 = data[ip+2]
			}
			data[data[ip+3]] = val1 + val2
			ip += 4
		case 2:
			val1, val2 := 0, 0
			if parammodes[0] == 0 {
				val1 = data[data[ip+1]]
			} else {
				val1 = data[ip+1]
			}
			if parammodes[1] == 0 {
				val2 = data[data[ip+2]]
			} else {
				val2 = data[ip+2]
			}
			data[data[ip+3]] = val1 * val2
			ip += 4
		case 3:
			data[data[ip+1]] = input[ii]
			ip += 2
			ii++
		case 4:
			if parammodes[0] == 0 {
				out = data[data[ip+1]]
			} else {
				out = data[ip+1]
			}
			if OUTPUT {
				fmt.Println(out)
			}
			ip += 2
		case 99:
			return out, nil
		default:
			return 0, errors.New("invalid opcode")
		}
	}
	return 0, errors.New("something went wrong")
}

func Challenge2(data []int, input []int) (int, error) {
	ip, ii, out := 0, 0, 0
	for true {
		opcode, instruction := data[ip]%100, data[ip]/100
		parammodes := []int{}
		for instruction > 0 {
			parammodes = append(parammodes, instruction%10)
			instruction /= 10
		}
		for len(parammodes) < 3 {
			parammodes = append(parammodes, 0)
		}
		switch opcode {
		case 1:
			val1, val2 := 0, 0
			if parammodes[0] == 0 {
				val1 = data[data[ip+1]]
			} else {
				val1 = data[ip+1]
			}
			if parammodes[1] == 0 {
				val2 = data[data[ip+2]]
			} else {
				val2 = data[ip+2]
			}
			data[data[ip+3]] = val1 + val2
			ip += 4
		case 2:
			val1, val2 := 0, 0
			if parammodes[0] == 0 {
				val1 = data[data[ip+1]]
			} else {
				val1 = data[ip+1]
			}
			if parammodes[1] == 0 {
				val2 = data[data[ip+2]]
			} else {
				val2 = data[ip+2]
			}
			data[data[ip+3]] = val1 * val2
			ip += 4
		case 3:
			data[data[ip+1]] = input[ii]
			ip += 2
			ii++
		case 4:
			if parammodes[0] == 0 {
				out = data[data[ip+1]]
			} else {
				out = data[ip+1]
			}
			if OUTPUT {
				fmt.Println(out)
			}
			ip += 2
		case 5:
			val1, val2 := 0, 0
			if parammodes[0] == 0 {
				val1 = data[data[ip+1]]
			} else {
				val1 = data[ip+1]
			}
			if parammodes[1] == 0 {
				val2 = data[data[ip+2]]
			} else {
				val2 = data[ip+2]
			}
			if val1 != 0 {
				ip = val2
			} else {
				ip += 3
			}
		case 6:
			val1, val2 := 0, 0
			if parammodes[0] == 0 {
				val1 = data[data[ip+1]]
			} else {
				val1 = data[ip+1]
			}
			if parammodes[1] == 0 {
				val2 = data[data[ip+2]]
			} else {
				val2 = data[ip+2]
			}
			if val1 == 0 {
				ip = val2
			} else {
				ip += 3
			}
		case 7:
			val1, val2 := 0, 0
			if parammodes[0] == 0 {
				val1 = data[data[ip+1]]
			} else {
				val1 = data[ip+1]
			}
			if parammodes[1] == 0 {
				val2 = data[data[ip+2]]
			} else {
				val2 = data[ip+2]
			}
			if val1 < val2 {
				data[data[ip+3]] = 1
			} else {
				data[data[ip+3]] = 0
			}
			ip += 4
		case 8:
			val1, val2 := 0, 0
			if parammodes[0] == 0 {
				val1 = data[data[ip+1]]
			} else {
				val1 = data[ip+1]
			}
			if parammodes[1] == 0 {
				val2 = data[data[ip+2]]
			} else {
				val2 = data[ip+2]
			}
			if val1 == val2 {
				data[data[ip+3]] = 1
			} else {
				data[data[ip+3]] = 0
			}
			ip += 4
		case 99:
			return out, nil
		default:
			return 0, errors.New("invalid opcode")
		}
	}
	return 0, errors.New("something went wrong")
}

func Run() {
	fmt.Println("Day 5 - Sunny with a Chance of Asteroids")
	path := "sunnywithachanceofasteroids/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var cpy []int = make([]int, len(data))
	copy(cpy, data)
	var result int
	result, err = Challenge1(cpy, []int{1})
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	copy(cpy, data)
	result, err = Challenge2(cpy, []int{5})
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
