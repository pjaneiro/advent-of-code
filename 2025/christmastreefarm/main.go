package christmastreefarm

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type shape struct {
	index    int
	format   []string
	occupied int
}

type region struct {
	width, length int
	quantities    []int
}

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

func parseInput(data []string) ([]shape, []region, error) {
	var shapes []shape = make([]shape, 0)
	var regions []region = make([]region, 0)
	var i int
	for i = 0; i < len(data); i++ {
		if data[i][len(data[i])-1] == ':' {
			cur := shape{-1, []string{}, 0}
			index, err := strconv.Atoi(data[i][:len(data[i])-1])
			if err != nil {
				return shapes, regions, errors.New("error parsing shape")
			}
			cur.index = index

			for j := i + 1; j < i+4; j++ {
				if len(data[j]) > 0 {
					cur.format = append(cur.format, data[j])
					for _, k := range data[j] {
						if k == '#' {
							cur.occupied++
						}
					}
					continue
				}
				break
			}

			shapes = append(shapes, cur)

			i += 4
		} else {
			break
		}
	}

	for ; i < len(data); i++ {
		cur := region{-1, -1, []int{}}

		parts := strings.Split(data[i], ": ")
		areaParts := strings.Split(parts[0], "x")

		if width, err := strconv.Atoi(areaParts[0]); err != nil {
			return shapes, regions, err
		} else {
			cur.width = width
		}

		if length, err := strconv.Atoi(areaParts[1]); err != nil {
			return shapes, regions, err
		} else {
			cur.length = length
		}

		quantities := strings.Split(parts[1], " ")
		for _, quantity := range quantities {
			if q, err := strconv.Atoi(quantity); err != nil {
				return shapes, regions, err
			} else {
				cur.quantities = append(cur.quantities, q)
			}
		}

		regions = append(regions, cur)
	}

	return shapes, regions, nil
}

func Challenge1(data []string) (int, error) {
	shapes, regions, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	result := 0

	for _, r := range regions {
		totalPieces, totalOccupied := 0, 0
		for i, j := range r.quantities {
			totalPieces++
			totalOccupied += j * shapes[i].occupied
		}
		if totalOccupied > r.width*r.length {
			// fmt.Printf("Region %v could never fit all these presents.\n", r)
			continue
		}
		if totalPieces*9 <= r.width*r.length {
			// fmt.Printf("Region %v can for sure fit all these presents.\n", r)
			result++
			continue
		}
		return result, fmt.Errorf("region %v is gonna need to be analyzed in more detail", r)
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 12 - Christmas Tree Farm")
	path := "christmastreefarm/input.txt"
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
}
