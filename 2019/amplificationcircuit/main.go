package amplificationcircuit

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func intcode(data []int, input []int) (int, error) {
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
			fmt.Println(out)
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
			return 0, fmt.Errorf("invalid opcode: %v", opcode)
		}
	}
	return 0, errors.New("something went wrong")
}

func Challenge1(data []int) (int, error) {
	var cpy []int = make([]int, len(data))
	max := 0
	for a1 := 0; a1 < 5; a1++ {
		copy(cpy, data)
		r1, err := intcode(cpy, []int{a1, 0})
		if err != nil {
			return 0, fmt.Errorf("something went wrong with amplifier 1: %v", err)
		}
		for a2 := 0; a2 < 5; a2++ {
			if a2 == a1 {
				continue
			}
			copy(cpy, data)
			r2, err := intcode(cpy, []int{a2, r1})
			if err != nil {
				return 0, fmt.Errorf("something went wrong with amplifier 2: %v", err)
			}
			for a3 := 0; a3 < 5; a3++ {
				if a3 == a1 || a3 == a2 {
					continue
				}
				copy(cpy, data)
				r3, err := intcode(cpy, []int{a3, r2})
				if err != nil {
					return 0, fmt.Errorf("something went wrong with amplifier 3: %v", err)
				}
				for a4 := 0; a4 < 5; a4++ {
					if a4 == a1 || a4 == a2 || a4 == a3 {
						continue
					}
					copy(cpy, data)
					r4, err := intcode(cpy, []int{a4, r3})
					if err != nil {
						return 0, fmt.Errorf("something went wrong with amplifier 4: %v", err)
					}
					for a5 := 0; a5 < 5; a5++ {
						if a5 == a1 || a5 == a2 || a5 == a3 || a5 == a4 {
							continue
						}
						copy(cpy, data)
						r5, err := intcode(cpy, []int{a5, r4})
						if err != nil {
							return 0, fmt.Errorf("something went wrong with amplifier 5: %v", err)
						}
						if r5 <= max {
							continue
						}
						max = r5
					}
				}
			}
		}
	}
	return max, nil
}

func intcode2(data []int, input []int, iip int, iii int) (int, []int, int, int, bool, error) {
	ip, ii, out := iip, iii, 0
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
			// fmt.Println(out)
			ip += 2
			return out, data, ip, ii, false, nil
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
			return out, data, ip, ii, true, nil
		default:
			return 0, data, ip, ii, false, fmt.Errorf("invalid opcode: %v", opcode)
		}
	}
	return 0, data, ip, ii, false, errors.New("something went wrong")
}

func Challenge2(data []int) (int, error) {
	max := 0
	for a1 := 5; a1 < 10; a1++ {
		for a2 := 5; a2 < 10; a2++ {
			if a2 == a1 {
				continue
			}
			for a3 := 5; a3 < 10; a3++ {
				if a3 == a1 || a3 == a2 {
					continue
				}
				for a4 := 5; a4 < 10; a4++ {
					if a4 == a1 || a4 == a2 || a4 == a3 {
						continue
					}
					for a5 := 5; a5 < 10; a5++ {
						if a5 == a1 || a5 == a2 || a5 == a3 || a5 == a4 {
							continue
						}

						inp := 0
						i1, i2, i3, i4, i5 := []int{a1}, []int{a2}, []int{a3}, []int{a4}, []int{a5}
						ip1, ip2, ip3, ip4, ip5, ii1, ii2, ii3, ii4, ii5 := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
						var cpy1 []int = make([]int, len(data))
						copy(cpy1, data)
						var cpy2 []int = make([]int, len(data))
						copy(cpy2, data)
						var cpy3 []int = make([]int, len(data))
						copy(cpy3, data)
						var cpy4 []int = make([]int, len(data))
						copy(cpy4, data)
						var cpy5 []int = make([]int, len(data))
						copy(cpy5, data)
						var tmp []int = make([]int, len(data))
						var err error
						var r1, r2, r3, r4, r5 int
						var done bool

						for true {
							i1 = append(i1, inp)
							r1, tmp, ip1, ii1, done, err = intcode2(cpy1, i1, ip1, ii1)

							if err != nil {
								return 0, fmt.Errorf("something went wrong with amplifier 1: %v", err)
							}
							if done {
								if r5 > max {
									max = r5
								}
								break
							}
							copy(cpy1, tmp)

							i2 = append(i2, r1)
							r2, tmp, ip2, ii2, done, err = intcode2(cpy2, i2, ip2, ii2)
							if err != nil {
								return 0, fmt.Errorf("something went wrong with amplifier 2: %v", err)
							}
							if done {
								if r5 > max {
									max = r5
								}
								break
							}
							copy(cpy2, tmp)

							i3 = append(i3, r2)
							r3, tmp, ip3, ii3, done, err = intcode2(cpy3, i3, ip3, ii3)
							if err != nil {
								return 0, fmt.Errorf("something went wrong with amplifier 3: %v", err)
							}
							if done {
								if r5 > max {
									max = r5
								}
								break
							}
							copy(cpy3, tmp)

							i4 = append(i4, r3)
							r4, tmp, ip4, ii4, done, err = intcode2(cpy4, i4, ip4, ii4)
							if err != nil {
								return 0, fmt.Errorf("something went wrong with amplifier 4: %v", err)
							}
							if done {
								if r5 > max {
									max = r5
								}
								break
							}
							copy(cpy4, tmp)

							i5 = append(i5, r4)
							r5, tmp, ip5, ii5, done, err = intcode2(cpy5, i5, ip5, ii5)
							if err != nil {
								return 0, fmt.Errorf("something went wrong with amplifier 5: %v", err)
							}
							inp = r5
							if done {
								if r5 > max {
									max = r5
								}
								break
							}
							copy(cpy5, tmp)
						}
					}
				}
			}
		}
	}
	return max, nil
}

func Run() {
	fmt.Println("Day 7 - Amplification Circuit")
	path := "amplificationcircuit/input.txt"
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
