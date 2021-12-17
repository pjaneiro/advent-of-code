package seatingsystem

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}
	return lines, scanner.Err()
}

func Challenge1(input [][]rune) (int, error) {
	var data [][]rune = make([][]rune, len(input))
	for i, l := range input {
		data[i] = make([]rune, len(l))
		copy(data[i], l)
	}
	for true {
		var cpy [][]rune = make([][]rune, len(data))
		for i, l := range data {
			cpy[i] = make([]rune, len(l))
			copy(cpy[i], l)
		}
		tainted := false
		for i, l := range data {
			for j, v := range l {
				switch v {
				case 'L':
					cnt := 0
					for w := -1; w <= 1; w++ {
						for h := -1; h <= 1; h++ {
							ti, tj := i+w, j+h
							if ti >= 0 && ti < len(data) && tj >= 0 && tj < len(l) && data[ti][tj] == '#' && (w != 0 || h != 0) {
								cnt++
							}
						}
					}
					if cnt != 0 {
						break
					}
					tainted = true
					cpy[i][j] = '#'
				case '#':
					cnt := 0
					for w := -1; w <= 1; w++ {
						for h := -1; h <= 1; h++ {
							ti, tj := i+w, j+h
							if ti >= 0 && ti < len(data) && tj >= 0 && tj < len(l) && data[ti][tj] == '#' && (w != 0 || h != 0) {
								cnt++
							}
						}
					}
					if cnt < 4 {
						break
					}
					tainted = true
					cpy[i][j] = 'L'
				case '.':
					break
				}
			}
		}

		for i, l := range cpy {
			data[i] = make([]rune, len(l))
			copy(data[i], l)
		}

		if !tainted {
			break
		}
	}
	result := 0
	for _, i := range data {
		for _, j := range i {
			if j == '#' {
				result++
			}
		}
	}

	return result, nil
}

func Challenge2(input [][]rune) (int, error) {
	var data [][]rune = make([][]rune, len(input))
	for i, l := range input {
		data[i] = make([]rune, len(l))
		copy(data[i], l)
	}
	for true {
		var cpy [][]rune = make([][]rune, len(data))
		for i, l := range data {
			cpy[i] = make([]rune, len(l))
			copy(cpy[i], l)
		}
		tainted := false
		for i, l := range data {
			for j, v := range l {
				switch v {
				case 'L':
					cnt := 0
					for w := -1; w <= 1; w++ {
						for h := -1; h <= 1; h++ {
							if w == 0 && h == 0 {
								continue
							}
							ti, tj := i, j
							for true {
								ti, tj = ti+w, tj+h
								if ti >= 0 && ti < len(data) && tj >= 0 && tj < len(l) {
									if data[ti][tj] == '#' {
										cnt++
										break
									} else if data[ti][tj] == 'L' {
										break
									} else {
										continue
									}
								} else {
									break
								}
							}
						}
					}
					if cnt != 0 {
						break
					}
					tainted = true
					cpy[i][j] = '#'
				case '#':
					cnt := 0
					for w := -1; w <= 1; w++ {
						for h := -1; h <= 1; h++ {
							if w == 0 && h == 0 {
								continue
							}
							ti, tj := i, j
							for true {
								ti, tj = ti+w, tj+h
								if ti >= 0 && ti < len(data) && tj >= 0 && tj < len(l) {
									if data[ti][tj] == '#' {
										cnt++
										break
									} else if data[ti][tj] == 'L' {
										break
									} else {
										continue
									}
								} else {
									break
								}
							}
						}
					}
					if cnt < 5 {
						break
					}
					tainted = true
					cpy[i][j] = 'L'
				case '.':
					break
				}
			}
		}

		for i, l := range cpy {
			data[i] = make([]rune, len(l))
			copy(data[i], l)
		}

		if !tainted {
			break
		}
	}
	result := 0
	for _, i := range data {
		for _, j := range i {
			if j == '#' {
				result++
			}
		}
	}

	return result, nil
}

func Run() {
	fmt.Println("Day 11 - Seating System")
	path := "seatingsystem/input.txt"
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
