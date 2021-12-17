package spaceimageformat

import (
	"bufio"
	"errors"
	"fmt"
	"math"
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

	var data []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, v := range strings.Split(line, "") {
			n, err := strconv.Atoi(v)
			if err != nil {
				return data, errors.New("something went wrong scanning int")
			}
			data = append(data, n)
		}
	}

	return data, nil
}

func Challenge1(data []int, h int, w int) (int, error) {
	var layers [][][]int
	var values [][]int
	min, b := math.MaxInt32, 0
	for l := 0; l < len(data); {
		var layer [][]int
		var lvalues []int = make([]int, 3)
		for i := 0; i < h; i++ {
			var line []int = make([]int, w)
			for j := 0; j < w; j++ {
				line[j] = data[l]
				if data[l] < 3 {
					lvalues[data[l]]++
				}
				l++
			}
			layer = append(layer, line)
		}
		if lvalues[0] < min {
			min = lvalues[0]
			b = len(values)
		}
		layers = append(layers, layer)
		values = append(values, lvalues)
	}
	return values[b][1] * values[b][2], nil
}

func Challenge2(data []int, h int, w int) {
	var layers [][][]int
	for l := 0; l < len(data); {
		var layer [][]int
		for i := 0; i < h; i++ {
			var line []int = make([]int, w)
			for j := 0; j < w; j++ {
				line[j] = data[l]
				l++
			}
			layer = append(layer, line)
		}
		layers = append(layers, layer)
	}
	var final [][]int = make([][]int, h)
	for i := 0; i < h; i++ {
		final[i] = make([]int, w)
		for j := 0; j < w; j++ {
			final[i][j] = 2
			for k := len(layers) - 1; k >= 0; k-- {
				if x := layers[k][i][j]; x != 2 {
					final[i][j] = x
				}
			}
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if final[i][j] == 0 {
				fmt.Printf("█")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	return
}

func Run() {
	fmt.Println("Day 8 - Space Image Format")
	path := "spaceimageformat/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(data, 6, 25)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	fmt.Println("Challenge 2: ")
	Challenge2(data, 6, 25)
}
